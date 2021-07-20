// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/ladybugsim/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_UpdatePositionEvent sql.NullBool
var __UpdatePositionEvent_time__dummyDeclaration time.Duration
var dummy_UpdatePositionEvent_sort sort.Float64Slice

// UpdatePositionEventAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model updatepositioneventAPI
type UpdatePositionEventAPI struct {
	gorm.Model

	models.UpdatePositionEvent

	// encoding of pointers
	UpdatePositionEventPointersEnconding
}

// UpdatePositionEventPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type UpdatePositionEventPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// UpdatePositionEventDB describes a updatepositionevent in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model updatepositioneventDB
type UpdatePositionEventDB struct {
	gorm.Model

	// insertion for basic fields declaration
	// Declation for basic field updatepositioneventDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field updatepositioneventDB.Duration {{BasicKind}} (to be completed)
	Duration_Data sql.NullInt64

	// encoding of pointers
	UpdatePositionEventPointersEnconding
}

// UpdatePositionEventDBs arrays updatepositioneventDBs
// swagger:response updatepositioneventDBsResponse
type UpdatePositionEventDBs []UpdatePositionEventDB

// UpdatePositionEventDBResponse provides response
// swagger:response updatepositioneventDBResponse
type UpdatePositionEventDBResponse struct {
	UpdatePositionEventDB
}

// UpdatePositionEventWOP is a UpdatePositionEvent without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type UpdatePositionEventWOP struct {
	ID int

	// insertion for WOP basic fields

	Name string

	Duration time.Duration
	// insertion for WOP pointer fields
}

var UpdatePositionEvent_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Duration",
}

type BackRepoUpdatePositionEventStruct struct {
	// stores UpdatePositionEventDB according to their gorm ID
	Map_UpdatePositionEventDBID_UpdatePositionEventDB *map[uint]*UpdatePositionEventDB

	// stores UpdatePositionEventDB ID according to UpdatePositionEvent address
	Map_UpdatePositionEventPtr_UpdatePositionEventDBID *map[*models.UpdatePositionEvent]uint

	// stores UpdatePositionEvent according to their gorm ID
	Map_UpdatePositionEventDBID_UpdatePositionEventPtr *map[uint]*models.UpdatePositionEvent

	db *gorm.DB
}

func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) GetDB() *gorm.DB {
	return backRepoUpdatePositionEvent.db
}

// GetUpdatePositionEventDBFromUpdatePositionEventPtr is a handy function to access the back repo instance from the stage instance
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) GetUpdatePositionEventDBFromUpdatePositionEventPtr(updatepositionevent *models.UpdatePositionEvent) (updatepositioneventDB *UpdatePositionEventDB) {
	id := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]
	updatepositioneventDB = (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[id]
	return
}

// BackRepoUpdatePositionEvent.Init set up the BackRepo of the UpdatePositionEvent
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) Init(db *gorm.DB) (Error error) {

	if backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr != nil {
		err := errors.New("In Init, backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr should be nil")
		return err
	}

	if backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB != nil {
		err := errors.New("In Init, backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB should be nil")
		return err
	}

	if backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID != nil {
		err := errors.New("In Init, backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.UpdatePositionEvent, 0)
	backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr = &tmp

	tmpDB := make(map[uint]*UpdatePositionEventDB, 0)
	backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB = &tmpDB

	tmpID := make(map[*models.UpdatePositionEvent]uint, 0)
	backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID = &tmpID

	backRepoUpdatePositionEvent.db = db
	return
}

// BackRepoUpdatePositionEvent.CommitPhaseOne commits all staged instances of UpdatePositionEvent to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for updatepositionevent := range stage.UpdatePositionEvents {
		backRepoUpdatePositionEvent.CommitPhaseOneInstance(updatepositionevent)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, updatepositionevent := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr {
		if _, ok := stage.UpdatePositionEvents[updatepositionevent]; !ok {
			backRepoUpdatePositionEvent.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUpdatePositionEvent.CommitDeleteInstance commits deletion of UpdatePositionEvent to the BackRepo
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CommitDeleteInstance(id uint) (Error error) {

	updatepositionevent := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[id]

	// updatepositionevent is not staged anymore, remove updatepositioneventDB
	updatepositioneventDB := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[id]
	query := backRepoUpdatePositionEvent.db.Unscoped().Delete(&updatepositioneventDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID), updatepositionevent)
	delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr), id)
	delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB), id)

	return
}

// BackRepoUpdatePositionEvent.CommitPhaseOneInstance commits updatepositionevent staged instances of UpdatePositionEvent to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CommitPhaseOneInstance(updatepositionevent *models.UpdatePositionEvent) (Error error) {

	// check if the updatepositionevent is not commited yet
	if _, ok := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]; ok {
		return
	}

	// initiate updatepositionevent
	var updatepositioneventDB UpdatePositionEventDB
	updatepositioneventDB.CopyBasicFieldsFromUpdatePositionEvent(updatepositionevent)

	query := backRepoUpdatePositionEvent.db.Create(&updatepositioneventDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent] = updatepositioneventDB.ID
	(*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[updatepositioneventDB.ID] = updatepositionevent
	(*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[updatepositioneventDB.ID] = &updatepositioneventDB

	return
}

// BackRepoUpdatePositionEvent.CommitPhaseTwo commits all staged instances of UpdatePositionEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, updatepositionevent := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr {
		backRepoUpdatePositionEvent.CommitPhaseTwoInstance(backRepo, idx, updatepositionevent)
	}

	return
}

// BackRepoUpdatePositionEvent.CommitPhaseTwoInstance commits {{structname }} of models.UpdatePositionEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, updatepositionevent *models.UpdatePositionEvent) (Error error) {

	// fetch matching updatepositioneventDB
	if updatepositioneventDB, ok := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[idx]; ok {

		updatepositioneventDB.CopyBasicFieldsFromUpdatePositionEvent(updatepositionevent)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoUpdatePositionEvent.db.Save(&updatepositioneventDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown UpdatePositionEvent intance %s", updatepositionevent.Name))
		return err
	}

	return
}

// BackRepoUpdatePositionEvent.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CheckoutPhaseOne() (Error error) {

	updatepositioneventDBArray := make([]UpdatePositionEventDB, 0)
	query := backRepoUpdatePositionEvent.db.Find(&updatepositioneventDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	updatepositioneventInstancesToBeRemovedFromTheStage := make(map[*models.UpdatePositionEvent]struct{})
	for key, value := range models.Stage.UpdatePositionEvents {
		updatepositioneventInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, updatepositioneventDB := range updatepositioneventDBArray {
		backRepoUpdatePositionEvent.CheckoutPhaseOneInstance(&updatepositioneventDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		updatepositionevent, ok := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[updatepositioneventDB.ID]
		if ok {
			delete(updatepositioneventInstancesToBeRemovedFromTheStage, updatepositionevent)
		}
	}

	// remove from stage and back repo's 3 maps all updatepositionevents that are not in the checkout
	for updatepositionevent := range updatepositioneventInstancesToBeRemovedFromTheStage {
		updatepositionevent.Unstage()

		// remove instance from the back repo 3 maps
		updatepositioneventID := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]
		delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID), updatepositionevent)
		delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB), updatepositioneventID)
		delete((*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr), updatepositioneventID)
	}

	return
}

// CheckoutPhaseOneInstance takes a updatepositioneventDB that has been found in the DB, updates the backRepo and stages the
// models version of the updatepositioneventDB
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CheckoutPhaseOneInstance(updatepositioneventDB *UpdatePositionEventDB) (Error error) {

	updatepositionevent, ok := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[updatepositioneventDB.ID]
	if !ok {
		updatepositionevent = new(models.UpdatePositionEvent)

		(*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[updatepositioneventDB.ID] = updatepositionevent
		(*backRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent] = updatepositioneventDB.ID

		// append model store with the new element
		updatepositionevent.Name = updatepositioneventDB.Name_Data.String
		updatepositionevent.Stage()
	}
	updatepositioneventDB.CopyBasicFieldsToUpdatePositionEvent(updatepositionevent)

	// preserve pointer to updatepositioneventDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_UpdatePositionEventDBID_UpdatePositionEventDB)[updatepositioneventDB hold variable pointers
	updatepositioneventDB_Data := *updatepositioneventDB
	preservedPtrToUpdatePositionEvent := &updatepositioneventDB_Data
	(*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[updatepositioneventDB.ID] = preservedPtrToUpdatePositionEvent

	return
}

// BackRepoUpdatePositionEvent.CheckoutPhaseTwo Checkouts all staged instances of UpdatePositionEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, updatepositioneventDB := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB {
		backRepoUpdatePositionEvent.CheckoutPhaseTwoInstance(backRepo, updatepositioneventDB)
	}
	return
}

// BackRepoUpdatePositionEvent.CheckoutPhaseTwoInstance Checkouts staged instances of UpdatePositionEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, updatepositioneventDB *UpdatePositionEventDB) (Error error) {

	updatepositionevent := (*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventPtr)[updatepositioneventDB.ID]
	_ = updatepositionevent // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitUpdatePositionEvent allows commit of a single updatepositionevent (if already staged)
func (backRepo *BackRepoStruct) CommitUpdatePositionEvent(updatepositionevent *models.UpdatePositionEvent) {
	backRepo.BackRepoUpdatePositionEvent.CommitPhaseOneInstance(updatepositionevent)
	if id, ok := (*backRepo.BackRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]; ok {
		backRepo.BackRepoUpdatePositionEvent.CommitPhaseTwoInstance(backRepo, id, updatepositionevent)
	}
}

// CommitUpdatePositionEvent allows checkout of a single updatepositionevent (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUpdatePositionEvent(updatepositionevent *models.UpdatePositionEvent) {
	// check if the updatepositionevent is staged
	if _, ok := (*backRepo.BackRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]; ok {

		if id, ok := (*backRepo.BackRepoUpdatePositionEvent.Map_UpdatePositionEventPtr_UpdatePositionEventDBID)[updatepositionevent]; ok {
			var updatepositioneventDB UpdatePositionEventDB
			updatepositioneventDB.ID = id

			if err := backRepo.BackRepoUpdatePositionEvent.db.First(&updatepositioneventDB, id).Error; err != nil {
				log.Panicln("CheckoutUpdatePositionEvent : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUpdatePositionEvent.CheckoutPhaseOneInstance(&updatepositioneventDB)
			backRepo.BackRepoUpdatePositionEvent.CheckoutPhaseTwoInstance(backRepo, &updatepositioneventDB)
		}
	}
}

// CopyBasicFieldsFromUpdatePositionEvent
func (updatepositioneventDB *UpdatePositionEventDB) CopyBasicFieldsFromUpdatePositionEvent(updatepositionevent *models.UpdatePositionEvent) {
	// insertion point for fields commit
	updatepositioneventDB.Name_Data.String = updatepositionevent.Name
	updatepositioneventDB.Name_Data.Valid = true

	updatepositioneventDB.Duration_Data.Int64 = int64(updatepositionevent.Duration)
	updatepositioneventDB.Duration_Data.Valid = true

}

// CopyBasicFieldsFromUpdatePositionEventWOP
func (updatepositioneventDB *UpdatePositionEventDB) CopyBasicFieldsFromUpdatePositionEventWOP(updatepositionevent *UpdatePositionEventWOP) {
	// insertion point for fields commit
	updatepositioneventDB.Name_Data.String = updatepositionevent.Name
	updatepositioneventDB.Name_Data.Valid = true

	updatepositioneventDB.Duration_Data.Int64 = int64(updatepositionevent.Duration)
	updatepositioneventDB.Duration_Data.Valid = true

}

// CopyBasicFieldsToUpdatePositionEvent
func (updatepositioneventDB *UpdatePositionEventDB) CopyBasicFieldsToUpdatePositionEvent(updatepositionevent *models.UpdatePositionEvent) {
	// insertion point for checkout of basic fields (back repo to stage)
	updatepositionevent.Name = updatepositioneventDB.Name_Data.String
	updatepositionevent.Duration = time.Duration(updatepositioneventDB.Duration_Data.Int64)
}

// CopyBasicFieldsToUpdatePositionEventWOP
func (updatepositioneventDB *UpdatePositionEventDB) CopyBasicFieldsToUpdatePositionEventWOP(updatepositionevent *UpdatePositionEventWOP) {
	updatepositionevent.ID = int(updatepositioneventDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	updatepositionevent.Name = updatepositioneventDB.Name_Data.String
	updatepositionevent.Duration = time.Duration(updatepositioneventDB.Duration_Data.Int64)
}

// Backup generates a json file from a slice of all UpdatePositionEventDB instances in the backrepo
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "UpdatePositionEventDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UpdatePositionEventDB, 0)
	for _, updatepositioneventDB := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB {
		forBackup = append(forBackup, updatepositioneventDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json UpdatePositionEvent ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json UpdatePositionEvent file", err.Error())
	}
}

// Backup generates a json file from a slice of all UpdatePositionEventDB instances in the backrepo
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UpdatePositionEventDB, 0)
	for _, updatepositioneventDB := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB {
		forBackup = append(forBackup, updatepositioneventDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("UpdatePositionEvent")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&UpdatePositionEvent_Fields, -1)
	for _, updatepositioneventDB := range forBackup {

		var updatepositioneventWOP UpdatePositionEventWOP
		updatepositioneventDB.CopyBasicFieldsToUpdatePositionEventWOP(&updatepositioneventWOP)

		row := sh.AddRow()
		row.WriteStruct(&updatepositioneventWOP, -1)
	}
}

// RestorePhaseOne read the file "UpdatePositionEventDB.json" in dirPath that stores an array
// of UpdatePositionEventDB and stores it in the database
// the map BackRepoUpdatePositionEventid_atBckpTime_newID is updated accordingly
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoUpdatePositionEventid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "UpdatePositionEventDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json UpdatePositionEvent file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*UpdatePositionEventDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_UpdatePositionEventDBID_UpdatePositionEventDB
	for _, updatepositioneventDB := range forRestore {

		updatepositioneventDB_ID_atBackupTime := updatepositioneventDB.ID
		updatepositioneventDB.ID = 0
		query := backRepoUpdatePositionEvent.db.Create(updatepositioneventDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB)[updatepositioneventDB.ID] = updatepositioneventDB
		BackRepoUpdatePositionEventid_atBckpTime_newID[updatepositioneventDB_ID_atBackupTime] = updatepositioneventDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json UpdatePositionEvent file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<UpdatePositionEvent>id_atBckpTime_newID
// to compute new index
func (backRepoUpdatePositionEvent *BackRepoUpdatePositionEventStruct) RestorePhaseTwo() {

	for _, updatepositioneventDB := range *backRepoUpdatePositionEvent.Map_UpdatePositionEventDBID_UpdatePositionEventDB {

		// next line of code is to avert unused variable compilation error
		_ = updatepositioneventDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoUpdatePositionEvent.db.Model(updatepositioneventDB).Updates(*updatepositioneventDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoUpdatePositionEventid_atBckpTime_newID map[uint]uint