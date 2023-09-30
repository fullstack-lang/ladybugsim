// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/ladybugsim/go/models"
	"github.com/fullstack-lang/ladybugsim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Ladybug__dummysDeclaration__ models.Ladybug
var __Ladybug_time__dummyDeclaration time.Duration

var mutexLadybug sync.Mutex

// An LadybugID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLadybug updateLadybug deleteLadybug
type LadybugID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LadybugInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLadybug updateLadybug
type LadybugInput struct {
	// The Ladybug to submit or modify
	// in: body
	Ladybug *orm.LadybugAPI
}

// GetLadybugs
//
// swagger:route GET /ladybugs ladybugs getLadybugs
//
// # Get all ladybugs
//
// Responses:
// default: genericError
//
//	200: ladybugDBResponse
func (controller *Controller) GetLadybugs(c *gin.Context) {

	// source slice
	var ladybugDBs []orm.LadybugDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLadybugs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybug.GetDB()

	query := db.Find(&ladybugDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ladybugAPIs := make([]orm.LadybugAPI, 0)

	// for each ladybug, update fields from the database nullable fields
	for idx := range ladybugDBs {
		ladybugDB := &ladybugDBs[idx]
		_ = ladybugDB
		var ladybugAPI orm.LadybugAPI

		// insertion point for updating fields
		ladybugAPI.ID = ladybugDB.ID
		ladybugDB.CopyBasicFieldsToLadybug(&ladybugAPI.Ladybug)
		ladybugAPI.LadybugPointersEnconding = ladybugDB.LadybugPointersEnconding
		ladybugAPIs = append(ladybugAPIs, ladybugAPI)
	}

	c.JSON(http.StatusOK, ladybugAPIs)
}

// PostLadybug
//
// swagger:route POST /ladybugs ladybugs postLadybug
//
// Creates a ladybug
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLadybug(c *gin.Context) {

	mutexLadybug.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLadybugs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybug.GetDB()

	// Validate input
	var input orm.LadybugAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ladybug
	ladybugDB := orm.LadybugDB{}
	ladybugDB.LadybugPointersEnconding = input.LadybugPointersEnconding
	ladybugDB.CopyBasicFieldsFromLadybug(&input.Ladybug)

	query := db.Create(&ladybugDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLadybug.CheckoutPhaseOneInstance(&ladybugDB)
	ladybug := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugPtr[ladybugDB.ID]

	if ladybug != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), ladybug)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, ladybugDB)

	mutexLadybug.Unlock()
}

// GetLadybug
//
// swagger:route GET /ladybugs/{ID} ladybugs getLadybug
//
// Gets the details for a ladybug.
//
// Responses:
// default: genericError
//
//	200: ladybugDBResponse
func (controller *Controller) GetLadybug(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLadybug", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybug.GetDB()

	// Get ladybugDB in DB
	var ladybugDB orm.LadybugDB
	if err := db.First(&ladybugDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var ladybugAPI orm.LadybugAPI
	ladybugAPI.ID = ladybugDB.ID
	ladybugAPI.LadybugPointersEnconding = ladybugDB.LadybugPointersEnconding
	ladybugDB.CopyBasicFieldsToLadybug(&ladybugAPI.Ladybug)

	c.JSON(http.StatusOK, ladybugAPI)
}

// UpdateLadybug
//
// swagger:route PATCH /ladybugs/{ID} ladybugs updateLadybug
//
// # Update a ladybug
//
// Responses:
// default: genericError
//
//	200: ladybugDBResponse
func (controller *Controller) UpdateLadybug(c *gin.Context) {

	mutexLadybug.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLadybug", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybug.GetDB()

	// Validate input
	var input orm.LadybugAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ladybugDB orm.LadybugDB

	// fetch the ladybug
	query := db.First(&ladybugDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ladybugDB.CopyBasicFieldsFromLadybug(&input.Ladybug)
	ladybugDB.LadybugPointersEnconding = input.LadybugPointersEnconding

	query = db.Model(&ladybugDB).Updates(ladybugDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ladybugNew := new(models.Ladybug)
	ladybugDB.CopyBasicFieldsToLadybug(ladybugNew)

	// get stage instance from DB instance, and call callback function
	ladybugOld := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugPtr[ladybugDB.ID]
	if ladybugOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), ladybugOld, ladybugNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ladybugDB
	c.JSON(http.StatusOK, ladybugDB)

	mutexLadybug.Unlock()
}

// DeleteLadybug
//
// swagger:route DELETE /ladybugs/{ID} ladybugs deleteLadybug
//
// # Delete a ladybug
//
// default: genericError
//
//	200: ladybugDBResponse
func (controller *Controller) DeleteLadybug(c *gin.Context) {

	mutexLadybug.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLadybug", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybug.GetDB()

	// Get model if exist
	var ladybugDB orm.LadybugDB
	if err := db.First(&ladybugDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&ladybugDB)

	// get an instance (not staged) from DB instance, and call callback function
	ladybugDeleted := new(models.Ladybug)
	ladybugDB.CopyBasicFieldsToLadybug(ladybugDeleted)

	// get stage instance from DB instance, and call callback function
	ladybugStaged := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugPtr[ladybugDB.ID]
	if ladybugStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), ladybugStaged, ladybugDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLadybug.Unlock()
}
