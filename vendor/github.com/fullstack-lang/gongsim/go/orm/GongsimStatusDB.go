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

	"github.com/fullstack-lang/gongsim/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongsimStatus_sql sql.NullBool
var dummy_GongsimStatus_time time.Duration
var dummy_GongsimStatus_sort sort.Float64Slice

// GongsimStatusAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongsimstatusAPI
type GongsimStatusAPI struct {
	gorm.Model

	models.GongsimStatus

	// encoding of pointers
	GongsimStatusPointersEnconding
}

// GongsimStatusPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongsimStatusPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// GongsimStatusDB describes a gongsimstatus in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongsimstatusDB
type GongsimStatusDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongsimstatusDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field gongsimstatusDB.CurrentCommand {{BasicKind}} (to be completed)
	CurrentCommand_Data sql.NullString

	// Declation for basic field gongsimstatusDB.CompletionDate {{BasicKind}} (to be completed)
	CompletionDate_Data sql.NullString

	// Declation for basic field gongsimstatusDB.CurrentSpeedCommand {{BasicKind}} (to be completed)
	CurrentSpeedCommand_Data sql.NullString

	// Declation for basic field gongsimstatusDB.SpeedCommandCompletionDate {{BasicKind}} (to be completed)
	SpeedCommandCompletionDate_Data sql.NullString
	// encoding of pointers
	GongsimStatusPointersEnconding
}

// GongsimStatusDBs arrays gongsimstatusDBs
// swagger:response gongsimstatusDBsResponse
type GongsimStatusDBs []GongsimStatusDB

// GongsimStatusDBResponse provides response
// swagger:response gongsimstatusDBResponse
type GongsimStatusDBResponse struct {
	GongsimStatusDB
}

// GongsimStatusWOP is a GongsimStatus without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongsimStatusWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	CurrentCommand models.GongsimCommandType `xlsx:"2"`

	CompletionDate string `xlsx:"3"`

	CurrentSpeedCommand models.SpeedCommandType `xlsx:"4"`

	SpeedCommandCompletionDate string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var GongsimStatus_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"CurrentCommand",
	"CompletionDate",
	"CurrentSpeedCommand",
	"SpeedCommandCompletionDate",
}

type BackRepoGongsimStatusStruct struct {
	// stores GongsimStatusDB according to their gorm ID
	Map_GongsimStatusDBID_GongsimStatusDB *map[uint]*GongsimStatusDB

	// stores GongsimStatusDB ID according to GongsimStatus address
	Map_GongsimStatusPtr_GongsimStatusDBID *map[*models.GongsimStatus]uint

	// stores GongsimStatus according to their gorm ID
	Map_GongsimStatusDBID_GongsimStatusPtr *map[uint]*models.GongsimStatus

	db *gorm.DB
}

func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) GetDB() *gorm.DB {
	return backRepoGongsimStatus.db
}

// GetGongsimStatusDBFromGongsimStatusPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) GetGongsimStatusDBFromGongsimStatusPtr(gongsimstatus *models.GongsimStatus) (gongsimstatusDB *GongsimStatusDB) {
	id := (*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]
	gongsimstatusDB = (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[id]
	return
}

// BackRepoGongsimStatus.Init set up the BackRepo of the GongsimStatus
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) Init(db *gorm.DB) (Error error) {

	if backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr != nil {
		err := errors.New("In Init, backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr should be nil")
		return err
	}

	if backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB != nil {
		err := errors.New("In Init, backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB should be nil")
		return err
	}

	if backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID != nil {
		err := errors.New("In Init, backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.GongsimStatus, 0)
	backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr = &tmp

	tmpDB := make(map[uint]*GongsimStatusDB, 0)
	backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB = &tmpDB

	tmpID := make(map[*models.GongsimStatus]uint, 0)
	backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID = &tmpID

	backRepoGongsimStatus.db = db
	return
}

// BackRepoGongsimStatus.CommitPhaseOne commits all staged instances of GongsimStatus to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongsimstatus := range stage.GongsimStatuss {
		backRepoGongsimStatus.CommitPhaseOneInstance(gongsimstatus)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongsimstatus := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr {
		if _, ok := stage.GongsimStatuss[gongsimstatus]; !ok {
			backRepoGongsimStatus.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongsimStatus.CommitDeleteInstance commits deletion of GongsimStatus to the BackRepo
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CommitDeleteInstance(id uint) (Error error) {

	gongsimstatus := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[id]

	// gongsimstatus is not staged anymore, remove gongsimstatusDB
	gongsimstatusDB := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[id]
	query := backRepoGongsimStatus.db.Unscoped().Delete(&gongsimstatusDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID), gongsimstatus)
	delete((*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr), id)
	delete((*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB), id)

	return
}

// BackRepoGongsimStatus.CommitPhaseOneInstance commits gongsimstatus staged instances of GongsimStatus to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CommitPhaseOneInstance(gongsimstatus *models.GongsimStatus) (Error error) {

	// check if the gongsimstatus is not commited yet
	if _, ok := (*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]; ok {
		return
	}

	// initiate gongsimstatus
	var gongsimstatusDB GongsimStatusDB
	gongsimstatusDB.CopyBasicFieldsFromGongsimStatus(gongsimstatus)

	query := backRepoGongsimStatus.db.Create(&gongsimstatusDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus] = gongsimstatusDB.ID
	(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[gongsimstatusDB.ID] = gongsimstatus
	(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[gongsimstatusDB.ID] = &gongsimstatusDB

	return
}

// BackRepoGongsimStatus.CommitPhaseTwo commits all staged instances of GongsimStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongsimstatus := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr {
		backRepoGongsimStatus.CommitPhaseTwoInstance(backRepo, idx, gongsimstatus)
	}

	return
}

// BackRepoGongsimStatus.CommitPhaseTwoInstance commits {{structname }} of models.GongsimStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongsimstatus *models.GongsimStatus) (Error error) {

	// fetch matching gongsimstatusDB
	if gongsimstatusDB, ok := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[idx]; ok {

		gongsimstatusDB.CopyBasicFieldsFromGongsimStatus(gongsimstatus)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGongsimStatus.db.Save(&gongsimstatusDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongsimStatus intance %s", gongsimstatus.Name))
		return err
	}

	return
}

// BackRepoGongsimStatus.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CheckoutPhaseOne() (Error error) {

	gongsimstatusDBArray := make([]GongsimStatusDB, 0)
	query := backRepoGongsimStatus.db.Find(&gongsimstatusDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongsimstatusInstancesToBeRemovedFromTheStage := make(map[*models.GongsimStatus]struct{})
	for key, value := range models.Stage.GongsimStatuss {
		gongsimstatusInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongsimstatusDB := range gongsimstatusDBArray {
		backRepoGongsimStatus.CheckoutPhaseOneInstance(&gongsimstatusDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongsimstatus, ok := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[gongsimstatusDB.ID]
		if ok {
			delete(gongsimstatusInstancesToBeRemovedFromTheStage, gongsimstatus)
		}
	}

	// remove from stage and back repo's 3 maps all gongsimstatuss that are not in the checkout
	for gongsimstatus := range gongsimstatusInstancesToBeRemovedFromTheStage {
		gongsimstatus.Unstage()

		// remove instance from the back repo 3 maps
		gongsimstatusID := (*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]
		delete((*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID), gongsimstatus)
		delete((*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB), gongsimstatusID)
		delete((*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr), gongsimstatusID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongsimstatusDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongsimstatusDB
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CheckoutPhaseOneInstance(gongsimstatusDB *GongsimStatusDB) (Error error) {

	gongsimstatus, ok := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[gongsimstatusDB.ID]
	if !ok {
		gongsimstatus = new(models.GongsimStatus)

		(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[gongsimstatusDB.ID] = gongsimstatus
		(*backRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus] = gongsimstatusDB.ID

		// append model store with the new element
		gongsimstatus.Name = gongsimstatusDB.Name_Data.String
		gongsimstatus.Stage()
	}
	gongsimstatusDB.CopyBasicFieldsToGongsimStatus(gongsimstatus)

	// preserve pointer to gongsimstatusDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongsimStatusDBID_GongsimStatusDB)[gongsimstatusDB hold variable pointers
	gongsimstatusDB_Data := *gongsimstatusDB
	preservedPtrToGongsimStatus := &gongsimstatusDB_Data
	(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[gongsimstatusDB.ID] = preservedPtrToGongsimStatus

	return
}

// BackRepoGongsimStatus.CheckoutPhaseTwo Checkouts all staged instances of GongsimStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongsimstatusDB := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB {
		backRepoGongsimStatus.CheckoutPhaseTwoInstance(backRepo, gongsimstatusDB)
	}
	return
}

// BackRepoGongsimStatus.CheckoutPhaseTwoInstance Checkouts staged instances of GongsimStatus to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongsimstatusDB *GongsimStatusDB) (Error error) {

	gongsimstatus := (*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusPtr)[gongsimstatusDB.ID]
	_ = gongsimstatus // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongsimStatus allows commit of a single gongsimstatus (if already staged)
func (backRepo *BackRepoStruct) CommitGongsimStatus(gongsimstatus *models.GongsimStatus) {
	backRepo.BackRepoGongsimStatus.CommitPhaseOneInstance(gongsimstatus)
	if id, ok := (*backRepo.BackRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]; ok {
		backRepo.BackRepoGongsimStatus.CommitPhaseTwoInstance(backRepo, id, gongsimstatus)
	}
}

// CommitGongsimStatus allows checkout of a single gongsimstatus (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongsimStatus(gongsimstatus *models.GongsimStatus) {
	// check if the gongsimstatus is staged
	if _, ok := (*backRepo.BackRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]; ok {

		if id, ok := (*backRepo.BackRepoGongsimStatus.Map_GongsimStatusPtr_GongsimStatusDBID)[gongsimstatus]; ok {
			var gongsimstatusDB GongsimStatusDB
			gongsimstatusDB.ID = id

			if err := backRepo.BackRepoGongsimStatus.db.First(&gongsimstatusDB, id).Error; err != nil {
				log.Panicln("CheckoutGongsimStatus : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongsimStatus.CheckoutPhaseOneInstance(&gongsimstatusDB)
			backRepo.BackRepoGongsimStatus.CheckoutPhaseTwoInstance(backRepo, &gongsimstatusDB)
		}
	}
}

// CopyBasicFieldsFromGongsimStatus
func (gongsimstatusDB *GongsimStatusDB) CopyBasicFieldsFromGongsimStatus(gongsimstatus *models.GongsimStatus) {
	// insertion point for fields commit

	gongsimstatusDB.Name_Data.String = gongsimstatus.Name
	gongsimstatusDB.Name_Data.Valid = true

	gongsimstatusDB.CurrentCommand_Data.String = string(gongsimstatus.CurrentCommand)
	gongsimstatusDB.CurrentCommand_Data.Valid = true

	gongsimstatusDB.CompletionDate_Data.String = gongsimstatus.CompletionDate
	gongsimstatusDB.CompletionDate_Data.Valid = true

	gongsimstatusDB.CurrentSpeedCommand_Data.String = string(gongsimstatus.CurrentSpeedCommand)
	gongsimstatusDB.CurrentSpeedCommand_Data.Valid = true

	gongsimstatusDB.SpeedCommandCompletionDate_Data.String = gongsimstatus.SpeedCommandCompletionDate
	gongsimstatusDB.SpeedCommandCompletionDate_Data.Valid = true
}

// CopyBasicFieldsFromGongsimStatusWOP
func (gongsimstatusDB *GongsimStatusDB) CopyBasicFieldsFromGongsimStatusWOP(gongsimstatus *GongsimStatusWOP) {
	// insertion point for fields commit

	gongsimstatusDB.Name_Data.String = gongsimstatus.Name
	gongsimstatusDB.Name_Data.Valid = true

	gongsimstatusDB.CurrentCommand_Data.String = string(gongsimstatus.CurrentCommand)
	gongsimstatusDB.CurrentCommand_Data.Valid = true

	gongsimstatusDB.CompletionDate_Data.String = gongsimstatus.CompletionDate
	gongsimstatusDB.CompletionDate_Data.Valid = true

	gongsimstatusDB.CurrentSpeedCommand_Data.String = string(gongsimstatus.CurrentSpeedCommand)
	gongsimstatusDB.CurrentSpeedCommand_Data.Valid = true

	gongsimstatusDB.SpeedCommandCompletionDate_Data.String = gongsimstatus.SpeedCommandCompletionDate
	gongsimstatusDB.SpeedCommandCompletionDate_Data.Valid = true
}

// CopyBasicFieldsToGongsimStatus
func (gongsimstatusDB *GongsimStatusDB) CopyBasicFieldsToGongsimStatus(gongsimstatus *models.GongsimStatus) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongsimstatus.Name = gongsimstatusDB.Name_Data.String
	gongsimstatus.CurrentCommand = models.GongsimCommandType(gongsimstatusDB.CurrentCommand_Data.String)
	gongsimstatus.CompletionDate = gongsimstatusDB.CompletionDate_Data.String
	gongsimstatus.CurrentSpeedCommand = models.SpeedCommandType(gongsimstatusDB.CurrentSpeedCommand_Data.String)
	gongsimstatus.SpeedCommandCompletionDate = gongsimstatusDB.SpeedCommandCompletionDate_Data.String
}

// CopyBasicFieldsToGongsimStatusWOP
func (gongsimstatusDB *GongsimStatusDB) CopyBasicFieldsToGongsimStatusWOP(gongsimstatus *GongsimStatusWOP) {
	gongsimstatus.ID = int(gongsimstatusDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongsimstatus.Name = gongsimstatusDB.Name_Data.String
	gongsimstatus.CurrentCommand = models.GongsimCommandType(gongsimstatusDB.CurrentCommand_Data.String)
	gongsimstatus.CompletionDate = gongsimstatusDB.CompletionDate_Data.String
	gongsimstatus.CurrentSpeedCommand = models.SpeedCommandType(gongsimstatusDB.CurrentSpeedCommand_Data.String)
	gongsimstatus.SpeedCommandCompletionDate = gongsimstatusDB.SpeedCommandCompletionDate_Data.String
}

// Backup generates a json file from a slice of all GongsimStatusDB instances in the backrepo
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongsimStatusDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongsimStatusDB, 0)
	for _, gongsimstatusDB := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB {
		forBackup = append(forBackup, gongsimstatusDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json GongsimStatus ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json GongsimStatus file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongsimStatusDB instances in the backrepo
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongsimStatusDB, 0)
	for _, gongsimstatusDB := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB {
		forBackup = append(forBackup, gongsimstatusDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongsimStatus")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongsimStatus_Fields, -1)
	for _, gongsimstatusDB := range forBackup {

		var gongsimstatusWOP GongsimStatusWOP
		gongsimstatusDB.CopyBasicFieldsToGongsimStatusWOP(&gongsimstatusWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongsimstatusWOP, -1)
	}
}

// RestoreXL from the "GongsimStatus" sheet all GongsimStatusDB instances
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongsimStatusid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongsimStatus"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongsimStatus.rowVisitorGongsimStatus)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) rowVisitorGongsimStatus(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongsimstatusWOP GongsimStatusWOP
		row.ReadStruct(&gongsimstatusWOP)

		// add the unmarshalled struct to the stage
		gongsimstatusDB := new(GongsimStatusDB)
		gongsimstatusDB.CopyBasicFieldsFromGongsimStatusWOP(&gongsimstatusWOP)

		gongsimstatusDB_ID_atBackupTime := gongsimstatusDB.ID
		gongsimstatusDB.ID = 0
		query := backRepoGongsimStatus.db.Create(gongsimstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[gongsimstatusDB.ID] = gongsimstatusDB
		BackRepoGongsimStatusid_atBckpTime_newID[gongsimstatusDB_ID_atBackupTime] = gongsimstatusDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongsimStatusDB.json" in dirPath that stores an array
// of GongsimStatusDB and stores it in the database
// the map BackRepoGongsimStatusid_atBckpTime_newID is updated accordingly
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongsimStatusid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongsimStatusDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json GongsimStatus file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongsimStatusDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongsimStatusDBID_GongsimStatusDB
	for _, gongsimstatusDB := range forRestore {

		gongsimstatusDB_ID_atBackupTime := gongsimstatusDB.ID
		gongsimstatusDB.ID = 0
		query := backRepoGongsimStatus.db.Create(gongsimstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB)[gongsimstatusDB.ID] = gongsimstatusDB
		BackRepoGongsimStatusid_atBckpTime_newID[gongsimstatusDB_ID_atBackupTime] = gongsimstatusDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json GongsimStatus file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongsimStatus>id_atBckpTime_newID
// to compute new index
func (backRepoGongsimStatus *BackRepoGongsimStatusStruct) RestorePhaseTwo() {

	for _, gongsimstatusDB := range *backRepoGongsimStatus.Map_GongsimStatusDBID_GongsimStatusDB {

		// next line of code is to avert unused variable compilation error
		_ = gongsimstatusDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGongsimStatus.db.Model(gongsimstatusDB).Updates(*gongsimstatusDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongsimStatusid_atBckpTime_newID map[uint]uint
