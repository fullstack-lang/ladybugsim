package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	ladybugsim_controllers "github.com/fullstack-lang/ladybugsim/go/controllers"
	ladybugsim_models "github.com/fullstack-lang/ladybugsim/go/models"
	ladybugsim_orm "github.com/fullstack-lang/ladybugsim/go/orm"

	gongsim_controllers "github.com/fullstack-lang/gongsim/go/controllers"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_orm "github.com/fullstack-lang/gongsim/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

func main() {

	log.SetPrefix("ladybugsim: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup GORM
	db := ladybugsim_orm.SetupModels(*logDBFlag, ":memory:")
	// add gongsim database
	gongsim_orm.AutoMigrate(db)

	dbDB, err := db.DB()

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	// init back repos
	ladybugsim_orm.BackRepo.Init(db)
	gongsim_orm.BackRepo.Init(db)

	// stage the simulation and the ladybugs
	for _, ladybug := range ladybugsim_models.LadybugSim.Ladybugs {
		ladybug.Stage().Commit()
	}
	ladybugsim_models.LadybugSim.Stage().Commit()

	// Gongsim (if autonomous, sim will run at the start)
	if *clientControlFlag {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.CLIENT_CONTROL
	} else {
		gongsim_models.EngineSingloton.ControlMode = gongsim_models.AUTONOMOUS
	}

	ladybugsim_models.Stage.Commit()
	gongsim_models.Stage.Commit()

	ladybugsim_controllers.RegisterControllers(r)
	gongsim_controllers.RegisterControllers(r)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

//go:embed ng/dist/ng
var ng embed.FS

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
