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
var dummy_UpdateSpeedEvent_sql sql.NullBool
var dummy_UpdateSpeedEvent_time time.Duration
var dummy_UpdateSpeedEvent_sort sort.Float64Slice

// UpdateSpeedEventAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model updatespeedeventAPI
type UpdateSpeedEventAPI struct {
	gorm.Model

	models.UpdateSpeedEvent

	// encoding of pointers
	UpdateSpeedEventPointersEnconding
}

// UpdateSpeedEventPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type UpdateSpeedEventPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// UpdateSpeedEventDB describes a updatespeedevent in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model updatespeedeventDB
type UpdateSpeedEventDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field updatespeedeventDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field updatespeedeventDB.Duration {{BasicKind}} (to be completed)
	Duration_Data sql.NullInt64
	// encoding of pointers
	UpdateSpeedEventPointersEnconding
}

// UpdateSpeedEventDBs arrays updatespeedeventDBs
// swagger:response updatespeedeventDBsResponse
type UpdateSpeedEventDBs []UpdateSpeedEventDB

// UpdateSpeedEventDBResponse provides response
// swagger:response updatespeedeventDBResponse
type UpdateSpeedEventDBResponse struct {
	UpdateSpeedEventDB
}

// UpdateSpeedEventWOP is a UpdateSpeedEvent without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type UpdateSpeedEventWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Duration time.Duration `xlsx:"2"`
	// insertion for WOP pointer fields
}

var UpdateSpeedEvent_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Duration",
}

type BackRepoUpdateSpeedEventStruct struct {
	// stores UpdateSpeedEventDB according to their gorm ID
	Map_UpdateSpeedEventDBID_UpdateSpeedEventDB *map[uint]*UpdateSpeedEventDB

	// stores UpdateSpeedEventDB ID according to UpdateSpeedEvent address
	Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID *map[*models.UpdateSpeedEvent]uint

	// stores UpdateSpeedEvent according to their gorm ID
	Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr *map[uint]*models.UpdateSpeedEvent

	db *gorm.DB
}

func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) GetDB() *gorm.DB {
	return backRepoUpdateSpeedEvent.db
}

// GetUpdateSpeedEventDBFromUpdateSpeedEventPtr is a handy function to access the back repo instance from the stage instance
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) GetUpdateSpeedEventDBFromUpdateSpeedEventPtr(updatespeedevent *models.UpdateSpeedEvent) (updatespeedeventDB *UpdateSpeedEventDB) {
	id := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]
	updatespeedeventDB = (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[id]
	return
}

// BackRepoUpdateSpeedEvent.Init set up the BackRepo of the UpdateSpeedEvent
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) Init(db *gorm.DB) (Error error) {

	if backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr != nil {
		err := errors.New("In Init, backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr should be nil")
		return err
	}

	if backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB != nil {
		err := errors.New("In Init, backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB should be nil")
		return err
	}

	if backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID != nil {
		err := errors.New("In Init, backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.UpdateSpeedEvent, 0)
	backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr = &tmp

	tmpDB := make(map[uint]*UpdateSpeedEventDB, 0)
	backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB = &tmpDB

	tmpID := make(map[*models.UpdateSpeedEvent]uint, 0)
	backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID = &tmpID

	backRepoUpdateSpeedEvent.db = db
	return
}

// BackRepoUpdateSpeedEvent.CommitPhaseOne commits all staged instances of UpdateSpeedEvent to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for updatespeedevent := range stage.UpdateSpeedEvents {
		backRepoUpdateSpeedEvent.CommitPhaseOneInstance(updatespeedevent)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, updatespeedevent := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr {
		if _, ok := stage.UpdateSpeedEvents[updatespeedevent]; !ok {
			backRepoUpdateSpeedEvent.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoUpdateSpeedEvent.CommitDeleteInstance commits deletion of UpdateSpeedEvent to the BackRepo
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CommitDeleteInstance(id uint) (Error error) {

	updatespeedevent := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[id]

	// updatespeedevent is not staged anymore, remove updatespeedeventDB
	updatespeedeventDB := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[id]
	query := backRepoUpdateSpeedEvent.db.Unscoped().Delete(&updatespeedeventDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID), updatespeedevent)
	delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr), id)
	delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB), id)

	return
}

// BackRepoUpdateSpeedEvent.CommitPhaseOneInstance commits updatespeedevent staged instances of UpdateSpeedEvent to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CommitPhaseOneInstance(updatespeedevent *models.UpdateSpeedEvent) (Error error) {

	// check if the updatespeedevent is not commited yet
	if _, ok := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]; ok {
		return
	}

	// initiate updatespeedevent
	var updatespeedeventDB UpdateSpeedEventDB
	updatespeedeventDB.CopyBasicFieldsFromUpdateSpeedEvent(updatespeedevent)

	query := backRepoUpdateSpeedEvent.db.Create(&updatespeedeventDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent] = updatespeedeventDB.ID
	(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[updatespeedeventDB.ID] = updatespeedevent
	(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[updatespeedeventDB.ID] = &updatespeedeventDB

	return
}

// BackRepoUpdateSpeedEvent.CommitPhaseTwo commits all staged instances of UpdateSpeedEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, updatespeedevent := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr {
		backRepoUpdateSpeedEvent.CommitPhaseTwoInstance(backRepo, idx, updatespeedevent)
	}

	return
}

// BackRepoUpdateSpeedEvent.CommitPhaseTwoInstance commits {{structname }} of models.UpdateSpeedEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, updatespeedevent *models.UpdateSpeedEvent) (Error error) {

	// fetch matching updatespeedeventDB
	if updatespeedeventDB, ok := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[idx]; ok {

		updatespeedeventDB.CopyBasicFieldsFromUpdateSpeedEvent(updatespeedevent)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoUpdateSpeedEvent.db.Save(&updatespeedeventDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown UpdateSpeedEvent intance %s", updatespeedevent.Name))
		return err
	}

	return
}

// BackRepoUpdateSpeedEvent.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CheckoutPhaseOne() (Error error) {

	updatespeedeventDBArray := make([]UpdateSpeedEventDB, 0)
	query := backRepoUpdateSpeedEvent.db.Find(&updatespeedeventDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	updatespeedeventInstancesToBeRemovedFromTheStage := make(map[*models.UpdateSpeedEvent]any)
	for key, value := range models.Stage.UpdateSpeedEvents {
		updatespeedeventInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, updatespeedeventDB := range updatespeedeventDBArray {
		backRepoUpdateSpeedEvent.CheckoutPhaseOneInstance(&updatespeedeventDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		updatespeedevent, ok := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[updatespeedeventDB.ID]
		if ok {
			delete(updatespeedeventInstancesToBeRemovedFromTheStage, updatespeedevent)
		}
	}

	// remove from stage and back repo's 3 maps all updatespeedevents that are not in the checkout
	for updatespeedevent := range updatespeedeventInstancesToBeRemovedFromTheStage {
		updatespeedevent.Unstage()

		// remove instance from the back repo 3 maps
		updatespeedeventID := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]
		delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID), updatespeedevent)
		delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB), updatespeedeventID)
		delete((*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr), updatespeedeventID)
	}

	return
}

// CheckoutPhaseOneInstance takes a updatespeedeventDB that has been found in the DB, updates the backRepo and stages the
// models version of the updatespeedeventDB
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CheckoutPhaseOneInstance(updatespeedeventDB *UpdateSpeedEventDB) (Error error) {

	updatespeedevent, ok := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[updatespeedeventDB.ID]
	if !ok {
		updatespeedevent = new(models.UpdateSpeedEvent)

		(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[updatespeedeventDB.ID] = updatespeedevent
		(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent] = updatespeedeventDB.ID

		// append model store with the new element
		updatespeedevent.Name = updatespeedeventDB.Name_Data.String
		updatespeedevent.Stage()
	}
	updatespeedeventDB.CopyBasicFieldsToUpdateSpeedEvent(updatespeedevent)

	// preserve pointer to updatespeedeventDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[updatespeedeventDB hold variable pointers
	updatespeedeventDB_Data := *updatespeedeventDB
	preservedPtrToUpdateSpeedEvent := &updatespeedeventDB_Data
	(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[updatespeedeventDB.ID] = preservedPtrToUpdateSpeedEvent

	return
}

// BackRepoUpdateSpeedEvent.CheckoutPhaseTwo Checkouts all staged instances of UpdateSpeedEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, updatespeedeventDB := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB {
		backRepoUpdateSpeedEvent.CheckoutPhaseTwoInstance(backRepo, updatespeedeventDB)
	}
	return
}

// BackRepoUpdateSpeedEvent.CheckoutPhaseTwoInstance Checkouts staged instances of UpdateSpeedEvent to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, updatespeedeventDB *UpdateSpeedEventDB) (Error error) {

	updatespeedevent := (*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventPtr)[updatespeedeventDB.ID]
	_ = updatespeedevent // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitUpdateSpeedEvent allows commit of a single updatespeedevent (if already staged)
func (backRepo *BackRepoStruct) CommitUpdateSpeedEvent(updatespeedevent *models.UpdateSpeedEvent) {
	backRepo.BackRepoUpdateSpeedEvent.CommitPhaseOneInstance(updatespeedevent)
	if id, ok := (*backRepo.BackRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]; ok {
		backRepo.BackRepoUpdateSpeedEvent.CommitPhaseTwoInstance(backRepo, id, updatespeedevent)
	}
}

// CommitUpdateSpeedEvent allows checkout of a single updatespeedevent (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutUpdateSpeedEvent(updatespeedevent *models.UpdateSpeedEvent) {
	// check if the updatespeedevent is staged
	if _, ok := (*backRepo.BackRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]; ok {

		if id, ok := (*backRepo.BackRepoUpdateSpeedEvent.Map_UpdateSpeedEventPtr_UpdateSpeedEventDBID)[updatespeedevent]; ok {
			var updatespeedeventDB UpdateSpeedEventDB
			updatespeedeventDB.ID = id

			if err := backRepo.BackRepoUpdateSpeedEvent.db.First(&updatespeedeventDB, id).Error; err != nil {
				log.Panicln("CheckoutUpdateSpeedEvent : Problem with getting object with id:", id)
			}
			backRepo.BackRepoUpdateSpeedEvent.CheckoutPhaseOneInstance(&updatespeedeventDB)
			backRepo.BackRepoUpdateSpeedEvent.CheckoutPhaseTwoInstance(backRepo, &updatespeedeventDB)
		}
	}
}

// CopyBasicFieldsFromUpdateSpeedEvent
func (updatespeedeventDB *UpdateSpeedEventDB) CopyBasicFieldsFromUpdateSpeedEvent(updatespeedevent *models.UpdateSpeedEvent) {
	// insertion point for fields commit

	updatespeedeventDB.Name_Data.String = updatespeedevent.Name
	updatespeedeventDB.Name_Data.Valid = true

	updatespeedeventDB.Duration_Data.Int64 = int64(updatespeedevent.Duration)
	updatespeedeventDB.Duration_Data.Valid = true
}

// CopyBasicFieldsFromUpdateSpeedEventWOP
func (updatespeedeventDB *UpdateSpeedEventDB) CopyBasicFieldsFromUpdateSpeedEventWOP(updatespeedevent *UpdateSpeedEventWOP) {
	// insertion point for fields commit

	updatespeedeventDB.Name_Data.String = updatespeedevent.Name
	updatespeedeventDB.Name_Data.Valid = true

	updatespeedeventDB.Duration_Data.Int64 = int64(updatespeedevent.Duration)
	updatespeedeventDB.Duration_Data.Valid = true
}

// CopyBasicFieldsToUpdateSpeedEvent
func (updatespeedeventDB *UpdateSpeedEventDB) CopyBasicFieldsToUpdateSpeedEvent(updatespeedevent *models.UpdateSpeedEvent) {
	// insertion point for checkout of basic fields (back repo to stage)
	updatespeedevent.Name = updatespeedeventDB.Name_Data.String
	updatespeedevent.Duration = time.Duration(updatespeedeventDB.Duration_Data.Int64)
}

// CopyBasicFieldsToUpdateSpeedEventWOP
func (updatespeedeventDB *UpdateSpeedEventDB) CopyBasicFieldsToUpdateSpeedEventWOP(updatespeedevent *UpdateSpeedEventWOP) {
	updatespeedevent.ID = int(updatespeedeventDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	updatespeedevent.Name = updatespeedeventDB.Name_Data.String
	updatespeedevent.Duration = time.Duration(updatespeedeventDB.Duration_Data.Int64)
}

// Backup generates a json file from a slice of all UpdateSpeedEventDB instances in the backrepo
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "UpdateSpeedEventDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UpdateSpeedEventDB, 0)
	for _, updatespeedeventDB := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB {
		forBackup = append(forBackup, updatespeedeventDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json UpdateSpeedEvent ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json UpdateSpeedEvent file", err.Error())
	}
}

// Backup generates a json file from a slice of all UpdateSpeedEventDB instances in the backrepo
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*UpdateSpeedEventDB, 0)
	for _, updatespeedeventDB := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB {
		forBackup = append(forBackup, updatespeedeventDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("UpdateSpeedEvent")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&UpdateSpeedEvent_Fields, -1)
	for _, updatespeedeventDB := range forBackup {

		var updatespeedeventWOP UpdateSpeedEventWOP
		updatespeedeventDB.CopyBasicFieldsToUpdateSpeedEventWOP(&updatespeedeventWOP)

		row := sh.AddRow()
		row.WriteStruct(&updatespeedeventWOP, -1)
	}
}

// RestoreXL from the "UpdateSpeedEvent" sheet all UpdateSpeedEventDB instances
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoUpdateSpeedEventid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["UpdateSpeedEvent"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoUpdateSpeedEvent.rowVisitorUpdateSpeedEvent)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) rowVisitorUpdateSpeedEvent(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var updatespeedeventWOP UpdateSpeedEventWOP
		row.ReadStruct(&updatespeedeventWOP)

		// add the unmarshalled struct to the stage
		updatespeedeventDB := new(UpdateSpeedEventDB)
		updatespeedeventDB.CopyBasicFieldsFromUpdateSpeedEventWOP(&updatespeedeventWOP)

		updatespeedeventDB_ID_atBackupTime := updatespeedeventDB.ID
		updatespeedeventDB.ID = 0
		query := backRepoUpdateSpeedEvent.db.Create(updatespeedeventDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[updatespeedeventDB.ID] = updatespeedeventDB
		BackRepoUpdateSpeedEventid_atBckpTime_newID[updatespeedeventDB_ID_atBackupTime] = updatespeedeventDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "UpdateSpeedEventDB.json" in dirPath that stores an array
// of UpdateSpeedEventDB and stores it in the database
// the map BackRepoUpdateSpeedEventid_atBckpTime_newID is updated accordingly
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoUpdateSpeedEventid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "UpdateSpeedEventDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json UpdateSpeedEvent file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*UpdateSpeedEventDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_UpdateSpeedEventDBID_UpdateSpeedEventDB
	for _, updatespeedeventDB := range forRestore {

		updatespeedeventDB_ID_atBackupTime := updatespeedeventDB.ID
		updatespeedeventDB.ID = 0
		query := backRepoUpdateSpeedEvent.db.Create(updatespeedeventDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB)[updatespeedeventDB.ID] = updatespeedeventDB
		BackRepoUpdateSpeedEventid_atBckpTime_newID[updatespeedeventDB_ID_atBackupTime] = updatespeedeventDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json UpdateSpeedEvent file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<UpdateSpeedEvent>id_atBckpTime_newID
// to compute new index
func (backRepoUpdateSpeedEvent *BackRepoUpdateSpeedEventStruct) RestorePhaseTwo() {

	for _, updatespeedeventDB := range *backRepoUpdateSpeedEvent.Map_UpdateSpeedEventDBID_UpdateSpeedEventDB {

		// next line of code is to avert unused variable compilation error
		_ = updatespeedeventDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoUpdateSpeedEvent.db.Model(updatespeedeventDB).Updates(*updatespeedeventDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoUpdateSpeedEventid_atBckpTime_newID map[uint]uint
