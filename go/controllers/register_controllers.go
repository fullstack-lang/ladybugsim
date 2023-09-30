// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/ladybugsim/go")
	{ // insertion point for registrations
		v1.GET("/v1/ladybugs", GetController().GetLadybugs)
		v1.GET("/v1/ladybugs/:id", GetController().GetLadybug)
		v1.POST("/v1/ladybugs", GetController().PostLadybug)
		v1.PATCH("/v1/ladybugs/:id", GetController().UpdateLadybug)
		v1.PUT("/v1/ladybugs/:id", GetController().UpdateLadybug)
		v1.DELETE("/v1/ladybugs/:id", GetController().DeleteLadybug)

		v1.GET("/v1/ladybugsimulations", GetController().GetLadybugSimulations)
		v1.GET("/v1/ladybugsimulations/:id", GetController().GetLadybugSimulation)
		v1.POST("/v1/ladybugsimulations", GetController().PostLadybugSimulation)
		v1.PATCH("/v1/ladybugsimulations/:id", GetController().UpdateLadybugSimulation)
		v1.PUT("/v1/ladybugsimulations/:id", GetController().UpdateLadybugSimulation)
		v1.DELETE("/v1/ladybugsimulations/:id", GetController().DeleteLadybugSimulation)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
