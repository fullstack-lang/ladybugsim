// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Polygone__dummysDeclaration__ models.Polygone
var __Polygone_time__dummyDeclaration time.Duration

// An PolygoneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPolygone updatePolygone deletePolygone
type PolygoneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PolygoneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPolygone updatePolygone
type PolygoneInput struct {
	// The Polygone to submit or modify
	// in: body
	Polygone *orm.PolygoneAPI
}

// GetPolygones
//
// swagger:route GET /polygones polygones getPolygones
//
// Get all polygones
//
// Responses:
//    default: genericError
//        200: polygoneDBsResponse
func GetPolygones(c *gin.Context) {
	db := orm.BackRepo.BackRepoPolygone.GetDB()

	// source slice
	var polygoneDBs []orm.PolygoneDB
	query := db.Find(&polygoneDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	polygoneAPIs := make([]orm.PolygoneAPI, 0)

	// for each polygone, update fields from the database nullable fields
	for idx := range polygoneDBs {
		polygoneDB := &polygoneDBs[idx]
		_ = polygoneDB
		var polygoneAPI orm.PolygoneAPI

		// insertion point for updating fields
		polygoneAPI.ID = polygoneDB.ID
		polygoneDB.CopyBasicFieldsToPolygone(&polygoneAPI.Polygone)
		polygoneAPI.PolygonePointersEnconding = polygoneDB.PolygonePointersEnconding
		polygoneAPIs = append(polygoneAPIs, polygoneAPI)
	}

	c.JSON(http.StatusOK, polygoneAPIs)
}

// PostPolygone
//
// swagger:route POST /polygones polygones postPolygone
//
// Creates a polygone
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: polygoneDBResponse
func PostPolygone(c *gin.Context) {
	db := orm.BackRepo.BackRepoPolygone.GetDB()

	// Validate input
	var input orm.PolygoneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create polygone
	polygoneDB := orm.PolygoneDB{}
	polygoneDB.PolygonePointersEnconding = input.PolygonePointersEnconding
	polygoneDB.CopyBasicFieldsFromPolygone(&input.Polygone)

	query := db.Create(&polygoneDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, polygoneDB)
}

// GetPolygone
//
// swagger:route GET /polygones/{ID} polygones getPolygone
//
// Gets the details for a polygone.
//
// Responses:
//    default: genericError
//        200: polygoneDBResponse
func GetPolygone(c *gin.Context) {
	db := orm.BackRepo.BackRepoPolygone.GetDB()

	// Get polygoneDB in DB
	var polygoneDB orm.PolygoneDB
	if err := db.First(&polygoneDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var polygoneAPI orm.PolygoneAPI
	polygoneAPI.ID = polygoneDB.ID
	polygoneAPI.PolygonePointersEnconding = polygoneDB.PolygonePointersEnconding
	polygoneDB.CopyBasicFieldsToPolygone(&polygoneAPI.Polygone)

	c.JSON(http.StatusOK, polygoneAPI)
}

// UpdatePolygone
//
// swagger:route PATCH /polygones/{ID} polygones updatePolygone
//
// Update a polygone
//
// Responses:
//    default: genericError
//        200: polygoneDBResponse
func UpdatePolygone(c *gin.Context) {
	db := orm.BackRepo.BackRepoPolygone.GetDB()

	// Get model if exist
	var polygoneDB orm.PolygoneDB

	// fetch the polygone
	query := db.First(&polygoneDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.PolygoneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	polygoneDB.CopyBasicFieldsFromPolygone(&input.Polygone)
	polygoneDB.PolygonePointersEnconding = input.PolygonePointersEnconding

	query = db.Model(&polygoneDB).Updates(polygoneDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the polygoneDB
	c.JSON(http.StatusOK, polygoneDB)
}

// DeletePolygone
//
// swagger:route DELETE /polygones/{ID} polygones deletePolygone
//
// Delete a polygone
//
// Responses:
//    default: genericError
func DeletePolygone(c *gin.Context) {
	db := orm.BackRepo.BackRepoPolygone.GetDB()

	// Get model if exist
	var polygoneDB orm.PolygoneDB
	if err := db.First(&polygoneDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&polygoneDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
