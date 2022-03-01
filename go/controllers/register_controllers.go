package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/ladybugsim/go/orm"
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

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/ladybugsim/go")
	{ // insertion point for registrations
		v1.GET("/v1/ladybugs", GetLadybugs)
		v1.GET("/v1/ladybugs/:id", GetLadybug)
		v1.POST("/v1/ladybugs", PostLadybug)
		v1.PATCH("/v1/ladybugs/:id", UpdateLadybug)
		v1.PUT("/v1/ladybugs/:id", UpdateLadybug)
		v1.DELETE("/v1/ladybugs/:id", DeleteLadybug)

		v1.GET("/v1/ladybugsimulations", GetLadybugSimulations)
		v1.GET("/v1/ladybugsimulations/:id", GetLadybugSimulation)
		v1.POST("/v1/ladybugsimulations", PostLadybugSimulation)
		v1.PATCH("/v1/ladybugsimulations/:id", UpdateLadybugSimulation)
		v1.PUT("/v1/ladybugsimulations/:id", UpdateLadybugSimulation)
		v1.DELETE("/v1/ladybugsimulations/:id", DeleteLadybugSimulation)

		v1.GET("/v1/updatepositionevents", GetUpdatePositionEvents)
		v1.GET("/v1/updatepositionevents/:id", GetUpdatePositionEvent)
		v1.POST("/v1/updatepositionevents", PostUpdatePositionEvent)
		v1.PATCH("/v1/updatepositionevents/:id", UpdateUpdatePositionEvent)
		v1.PUT("/v1/updatepositionevents/:id", UpdateUpdatePositionEvent)
		v1.DELETE("/v1/updatepositionevents/:id", DeleteUpdatePositionEvent)

		v1.GET("/v1/updatespeedevents", GetUpdateSpeedEvents)
		v1.GET("/v1/updatespeedevents/:id", GetUpdateSpeedEvent)
		v1.POST("/v1/updatespeedevents", PostUpdateSpeedEvent)
		v1.PATCH("/v1/updatespeedevents/:id", UpdateUpdateSpeedEvent)
		v1.PUT("/v1/updatespeedevents/:id", UpdateUpdateSpeedEvent)
		v1.DELETE("/v1/updatespeedevents/:id", DeleteUpdateSpeedEvent)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
