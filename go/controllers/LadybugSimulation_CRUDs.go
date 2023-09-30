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
var __LadybugSimulation__dummysDeclaration__ models.LadybugSimulation
var __LadybugSimulation_time__dummyDeclaration time.Duration

var mutexLadybugSimulation sync.Mutex

// An LadybugSimulationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLadybugSimulation updateLadybugSimulation deleteLadybugSimulation
type LadybugSimulationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LadybugSimulationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLadybugSimulation updateLadybugSimulation
type LadybugSimulationInput struct {
	// The LadybugSimulation to submit or modify
	// in: body
	LadybugSimulation *orm.LadybugSimulationAPI
}

// GetLadybugSimulations
//
// swagger:route GET /ladybugsimulations ladybugsimulations getLadybugSimulations
//
// # Get all ladybugsimulations
//
// Responses:
// default: genericError
//
//	200: ladybugsimulationDBResponse
func (controller *Controller) GetLadybugSimulations(c *gin.Context) {

	// source slice
	var ladybugsimulationDBs []orm.LadybugSimulationDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLadybugSimulations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybugSimulation.GetDB()

	query := db.Find(&ladybugsimulationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ladybugsimulationAPIs := make([]orm.LadybugSimulationAPI, 0)

	// for each ladybugsimulation, update fields from the database nullable fields
	for idx := range ladybugsimulationDBs {
		ladybugsimulationDB := &ladybugsimulationDBs[idx]
		_ = ladybugsimulationDB
		var ladybugsimulationAPI orm.LadybugSimulationAPI

		// insertion point for updating fields
		ladybugsimulationAPI.ID = ladybugsimulationDB.ID
		ladybugsimulationDB.CopyBasicFieldsToLadybugSimulation(&ladybugsimulationAPI.LadybugSimulation)
		ladybugsimulationAPI.LadybugSimulationPointersEnconding = ladybugsimulationDB.LadybugSimulationPointersEnconding
		ladybugsimulationAPIs = append(ladybugsimulationAPIs, ladybugsimulationAPI)
	}

	c.JSON(http.StatusOK, ladybugsimulationAPIs)
}

// PostLadybugSimulation
//
// swagger:route POST /ladybugsimulations ladybugsimulations postLadybugSimulation
//
// Creates a ladybugsimulation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLadybugSimulation(c *gin.Context) {

	mutexLadybugSimulation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLadybugSimulations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybugSimulation.GetDB()

	// Validate input
	var input orm.LadybugSimulationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ladybugsimulation
	ladybugsimulationDB := orm.LadybugSimulationDB{}
	ladybugsimulationDB.LadybugSimulationPointersEnconding = input.LadybugSimulationPointersEnconding
	ladybugsimulationDB.CopyBasicFieldsFromLadybugSimulation(&input.LadybugSimulation)

	query := db.Create(&ladybugsimulationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLadybugSimulation.CheckoutPhaseOneInstance(&ladybugsimulationDB)
	ladybugsimulation := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]

	if ladybugsimulation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), ladybugsimulation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, ladybugsimulationDB)

	mutexLadybugSimulation.Unlock()
}

// GetLadybugSimulation
//
// swagger:route GET /ladybugsimulations/{ID} ladybugsimulations getLadybugSimulation
//
// Gets the details for a ladybugsimulation.
//
// Responses:
// default: genericError
//
//	200: ladybugsimulationDBResponse
func (controller *Controller) GetLadybugSimulation(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLadybugSimulation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybugSimulation.GetDB()

	// Get ladybugsimulationDB in DB
	var ladybugsimulationDB orm.LadybugSimulationDB
	if err := db.First(&ladybugsimulationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var ladybugsimulationAPI orm.LadybugSimulationAPI
	ladybugsimulationAPI.ID = ladybugsimulationDB.ID
	ladybugsimulationAPI.LadybugSimulationPointersEnconding = ladybugsimulationDB.LadybugSimulationPointersEnconding
	ladybugsimulationDB.CopyBasicFieldsToLadybugSimulation(&ladybugsimulationAPI.LadybugSimulation)

	c.JSON(http.StatusOK, ladybugsimulationAPI)
}

// UpdateLadybugSimulation
//
// swagger:route PATCH /ladybugsimulations/{ID} ladybugsimulations updateLadybugSimulation
//
// # Update a ladybugsimulation
//
// Responses:
// default: genericError
//
//	200: ladybugsimulationDBResponse
func (controller *Controller) UpdateLadybugSimulation(c *gin.Context) {

	mutexLadybugSimulation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLadybugSimulation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybugSimulation.GetDB()

	// Validate input
	var input orm.LadybugSimulationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ladybugsimulationDB orm.LadybugSimulationDB

	// fetch the ladybugsimulation
	query := db.First(&ladybugsimulationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ladybugsimulationDB.CopyBasicFieldsFromLadybugSimulation(&input.LadybugSimulation)
	ladybugsimulationDB.LadybugSimulationPointersEnconding = input.LadybugSimulationPointersEnconding

	query = db.Model(&ladybugsimulationDB).Updates(ladybugsimulationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ladybugsimulationNew := new(models.LadybugSimulation)
	ladybugsimulationDB.CopyBasicFieldsToLadybugSimulation(ladybugsimulationNew)

	// get stage instance from DB instance, and call callback function
	ladybugsimulationOld := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]
	if ladybugsimulationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), ladybugsimulationOld, ladybugsimulationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ladybugsimulationDB
	c.JSON(http.StatusOK, ladybugsimulationDB)

	mutexLadybugSimulation.Unlock()
}

// DeleteLadybugSimulation
//
// swagger:route DELETE /ladybugsimulations/{ID} ladybugsimulations deleteLadybugSimulation
//
// # Delete a ladybugsimulation
//
// default: genericError
//
//	200: ladybugsimulationDBResponse
func (controller *Controller) DeleteLadybugSimulation(c *gin.Context) {

	mutexLadybugSimulation.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLadybugSimulation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/ladybugsim/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLadybugSimulation.GetDB()

	// Get model if exist
	var ladybugsimulationDB orm.LadybugSimulationDB
	if err := db.First(&ladybugsimulationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&ladybugsimulationDB)

	// get an instance (not staged) from DB instance, and call callback function
	ladybugsimulationDeleted := new(models.LadybugSimulation)
	ladybugsimulationDB.CopyBasicFieldsToLadybugSimulation(ladybugsimulationDeleted)

	// get stage instance from DB instance, and call callback function
	ladybugsimulationStaged := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]
	if ladybugsimulationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), ladybugsimulationStaged, ladybugsimulationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLadybugSimulation.Unlock()
}
