// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"net/http"
	"time"

	"github.com/fullstack-lang/ladybugsim/go/models"
	"github.com/fullstack-lang/ladybugsim/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __UpdateSpeedEvent__dummysDeclaration__ models.UpdateSpeedEvent
var __UpdateSpeedEvent_time__dummyDeclaration time.Duration

// An UpdateSpeedEventID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getUpdateSpeedEvent updateUpdateSpeedEvent deleteUpdateSpeedEvent
type UpdateSpeedEventID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// UpdateSpeedEventInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postUpdateSpeedEvent updateUpdateSpeedEvent
type UpdateSpeedEventInput struct {
	// The UpdateSpeedEvent to submit or modify
	// in: body
	UpdateSpeedEvent *orm.UpdateSpeedEventAPI
}

// GetUpdateSpeedEvents
//
// swagger:route GET /updatespeedevents updatespeedevents getUpdateSpeedEvents
//
// Get all updatespeedevents
//
// Responses:
//    default: genericError
//        200: updatespeedeventDBsResponse
func GetUpdateSpeedEvents(c *gin.Context) {
	db := orm.BackRepo.BackRepoUpdateSpeedEvent.GetDB()
	
	// source slice
	var updatespeedeventDBs []orm.UpdateSpeedEventDB
	query := db.Find(&updatespeedeventDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	updatespeedeventAPIs := make([]orm.UpdateSpeedEventAPI, 0)

	// for each updatespeedevent, update fields from the database nullable fields
	for idx := range updatespeedeventDBs {
		updatespeedeventDB := &updatespeedeventDBs[idx]
		_ = updatespeedeventDB
		var updatespeedeventAPI orm.UpdateSpeedEventAPI

		// insertion point for updating fields
		updatespeedeventAPI.ID = updatespeedeventDB.ID
		updatespeedeventDB.CopyBasicFieldsToUpdateSpeedEvent(&updatespeedeventAPI.UpdateSpeedEvent)
		updatespeedeventAPI.UpdateSpeedEventPointersEnconding = updatespeedeventDB.UpdateSpeedEventPointersEnconding
		updatespeedeventAPIs = append(updatespeedeventAPIs, updatespeedeventAPI)
	}

	c.JSON(http.StatusOK, updatespeedeventAPIs)
}

// PostUpdateSpeedEvent
//
// swagger:route POST /updatespeedevents updatespeedevents postUpdateSpeedEvent
//
// Creates a updatespeedevent
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: updatespeedeventDBResponse
func PostUpdateSpeedEvent(c *gin.Context) {
	db := orm.BackRepo.BackRepoUpdateSpeedEvent.GetDB()

	// Validate input
	var input orm.UpdateSpeedEventAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create updatespeedevent
	updatespeedeventDB := orm.UpdateSpeedEventDB{}
	updatespeedeventDB.UpdateSpeedEventPointersEnconding = input.UpdateSpeedEventPointersEnconding
	updatespeedeventDB.CopyBasicFieldsFromUpdateSpeedEvent(&input.UpdateSpeedEvent)

	query := db.Create(&updatespeedeventDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, updatespeedeventDB)
}

// GetUpdateSpeedEvent
//
// swagger:route GET /updatespeedevents/{ID} updatespeedevents getUpdateSpeedEvent
//
// Gets the details for a updatespeedevent.
//
// Responses:
//    default: genericError
//        200: updatespeedeventDBResponse
func GetUpdateSpeedEvent(c *gin.Context) {
	db := orm.BackRepo.BackRepoUpdateSpeedEvent.GetDB()

	// Get updatespeedeventDB in DB
	var updatespeedeventDB orm.UpdateSpeedEventDB
	if err := db.First(&updatespeedeventDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var updatespeedeventAPI orm.UpdateSpeedEventAPI
	updatespeedeventAPI.ID = updatespeedeventDB.ID
	updatespeedeventAPI.UpdateSpeedEventPointersEnconding = updatespeedeventDB.UpdateSpeedEventPointersEnconding
	updatespeedeventDB.CopyBasicFieldsToUpdateSpeedEvent(&updatespeedeventAPI.UpdateSpeedEvent)

	c.JSON(http.StatusOK, updatespeedeventAPI)
}

// UpdateUpdateSpeedEvent
//
// swagger:route PATCH /updatespeedevents/{ID} updatespeedevents updateUpdateSpeedEvent
//
// Update a updatespeedevent
//
// Responses:
//    default: genericError
//        200: updatespeedeventDBResponse
func UpdateUpdateSpeedEvent(c *gin.Context) {
	db := orm.BackRepo.BackRepoUpdateSpeedEvent.GetDB()

	// Get model if exist
	var updatespeedeventDB orm.UpdateSpeedEventDB

	// fetch the updatespeedevent
	query := db.First(&updatespeedeventDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.UpdateSpeedEventAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	updatespeedeventDB.CopyBasicFieldsFromUpdateSpeedEvent(&input.UpdateSpeedEvent)
	updatespeedeventDB.UpdateSpeedEventPointersEnconding = input.UpdateSpeedEventPointersEnconding

	query = db.Model(&updatespeedeventDB).Updates(updatespeedeventDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the updatespeedeventDB
	c.JSON(http.StatusOK, updatespeedeventDB)
}

// DeleteUpdateSpeedEvent
//
// swagger:route DELETE /updatespeedevents/{ID} updatespeedevents deleteUpdateSpeedEvent
//
// Delete a updatespeedevent
//
// Responses:
//    default: genericError
func DeleteUpdateSpeedEvent(c *gin.Context) {
	db := orm.BackRepo.BackRepoUpdateSpeedEvent.GetDB()

	// Get model if exist
	var updatespeedeventDB orm.UpdateSpeedEventDB
	if err := db.First(&updatespeedeventDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&updatespeedeventDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}