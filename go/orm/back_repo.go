// generated by genORMTranslation.go
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/fullstack-lang/ladybugsim/go/models"

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoLadybug BackRepoLadybugStruct

	BackRepoLadybugSimulation BackRepoLadybugSimulationStruct

	BackRepoUpdatePositionEvent BackRepoUpdatePositionEventStruct

	BackRepoUpdateSpeedEvent BackRepoUpdateSpeedEventStruct

	CommitNb uint // this ng is updated at the BackRepo level but also at the BackRepo<GongStruct> level

	PushFromFrontNb uint // records increments from push from front
}

func (backRepo *BackRepoStruct) GetLastCommitNb() uint {
	return backRepo.CommitNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitNb() uint {
	if models.Stage.OnInitCommitCallback != nil {
		models.Stage.OnInitCommitCallback.BeforeCommit(&models.Stage)
	}
	backRepo.CommitNb = backRepo.CommitNb + 1
	return backRepo.CommitNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitNb
}

// Init the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) init(db *gorm.DB) {
	// insertion point for per struct back repo declarations
	backRepo.BackRepoLadybug.Init(db)
	backRepo.BackRepoLadybugSimulation.Init(db)
	backRepo.BackRepoUpdatePositionEvent.Init(db)
	backRepo.BackRepoUpdateSpeedEvent.Init(db)

	models.Stage.BackRepo = backRepo
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoLadybug.CommitPhaseOne(stage)
	backRepo.BackRepoLadybugSimulation.CommitPhaseOne(stage)
	backRepo.BackRepoUpdatePositionEvent.CommitPhaseOne(stage)
	backRepo.BackRepoUpdateSpeedEvent.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoLadybug.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLadybugSimulation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUpdatePositionEvent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUpdateSpeedEvent.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoLadybug.CheckoutPhaseOne()
	backRepo.BackRepoLadybugSimulation.CheckoutPhaseOne()
	backRepo.BackRepoUpdatePositionEvent.CheckoutPhaseOne()
	backRepo.BackRepoUpdateSpeedEvent.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoLadybug.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLadybugSimulation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUpdatePositionEvent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUpdateSpeedEvent.CheckoutPhaseTwo(backRepo)
}

var BackRepo BackRepoStruct

func GetLastCommitNb() uint {
	return BackRepo.GetLastCommitNb()
}

func GetLastPushFromFrontNb() uint {
	return BackRepo.GetLastPushFromFrontNb()
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.Mkdir(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoLadybug.Backup(dirPath)
	backRepo.BackRepoLadybugSimulation.Backup(dirPath)
	backRepo.BackRepoUpdatePositionEvent.Backup(dirPath)
	backRepo.BackRepoUpdateSpeedEvent.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.Mkdir(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoLadybug.BackupXL(file)
	backRepo.BackRepoLadybugSimulation.BackupXL(file)
	backRepo.BackRepoUpdatePositionEvent.BackupXL(file)
	backRepo.BackRepoUpdateSpeedEvent.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	models.Stage.Commit()
	models.Stage.Reset()
	models.Stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoLadybug.RestorePhaseOne(dirPath)
	backRepo.BackRepoLadybugSimulation.RestorePhaseOne(dirPath)
	backRepo.BackRepoUpdatePositionEvent.RestorePhaseOne(dirPath)
	backRepo.BackRepoUpdateSpeedEvent.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoLadybug.RestorePhaseTwo()
	backRepo.BackRepoLadybugSimulation.RestorePhaseTwo()
	backRepo.BackRepoUpdatePositionEvent.RestorePhaseTwo()
	backRepo.BackRepoUpdateSpeedEvent.RestorePhaseTwo()

	models.Stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {
}
