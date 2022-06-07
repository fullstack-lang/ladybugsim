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
var dummy_Ladybug_sql sql.NullBool
var dummy_Ladybug_time time.Duration
var dummy_Ladybug_sort sort.Float64Slice

// LadybugAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model ladybugAPI
type LadybugAPI struct {
	gorm.Model

	models.Ladybug

	// encoding of pointers
	LadybugPointersEnconding
}

// LadybugPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LadybugPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field LadybugSimulation{}.Ladybugs []*Ladybug
	LadybugSimulation_LadybugsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	LadybugSimulation_LadybugsDBID_Index sql.NullInt64
}

// LadybugDB describes a ladybug in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model ladybugDB
type LadybugDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field ladybugDB.TechName {{BasicKind}} (to be completed)
	TechName_Data sql.NullString

	// Declation for basic field ladybugDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field ladybugDB.Id {{BasicKind}} (to be completed)
	Id_Data sql.NullInt64

	// Declation for basic field ladybugDB.Position {{BasicKind}} (to be completed)
	Position_Data sql.NullFloat64

	// Declation for basic field ladybugDB.Speed {{BasicKind}} (to be completed)
	Speed_Data sql.NullFloat64

	// Declation for basic field ladybugDB.LadybugStatus {{BasicKind}} (to be completed)
	LadybugStatus_Data sql.NullString
	// encoding of pointers
	LadybugPointersEnconding
}

// LadybugDBs arrays ladybugDBs
// swagger:response ladybugDBsResponse
type LadybugDBs []LadybugDB

// LadybugDBResponse provides response
// swagger:response ladybugDBResponse
type LadybugDBResponse struct {
	LadybugDB
}

// LadybugWOP is a Ladybug without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LadybugWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	TechName string `xlsx:"1"`

	Name string `xlsx:"2"`

	Id int `xlsx:"3"`

	Position float64 `xlsx:"4"`

	Speed float64 `xlsx:"5"`

	LadybugStatus models.LadybugStatus `xlsx:"6"`
	// insertion for WOP pointer fields
}

var Ladybug_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"TechName",
	"Name",
	"Id",
	"Position",
	"Speed",
	"LadybugStatus",
}

type BackRepoLadybugStruct struct {
	// stores LadybugDB according to their gorm ID
	Map_LadybugDBID_LadybugDB *map[uint]*LadybugDB

	// stores LadybugDB ID according to Ladybug address
	Map_LadybugPtr_LadybugDBID *map[*models.Ladybug]uint

	// stores Ladybug according to their gorm ID
	Map_LadybugDBID_LadybugPtr *map[uint]*models.Ladybug

	db *gorm.DB
}

func (backRepoLadybug *BackRepoLadybugStruct) GetDB() *gorm.DB {
	return backRepoLadybug.db
}

// GetLadybugDBFromLadybugPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLadybug *BackRepoLadybugStruct) GetLadybugDBFromLadybugPtr(ladybug *models.Ladybug) (ladybugDB *LadybugDB) {
	id := (*backRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]
	ladybugDB = (*backRepoLadybug.Map_LadybugDBID_LadybugDB)[id]
	return
}

// BackRepoLadybug.Init set up the BackRepo of the Ladybug
func (backRepoLadybug *BackRepoLadybugStruct) Init(db *gorm.DB) (Error error) {

	if backRepoLadybug.Map_LadybugDBID_LadybugPtr != nil {
		err := errors.New("In Init, backRepoLadybug.Map_LadybugDBID_LadybugPtr should be nil")
		return err
	}

	if backRepoLadybug.Map_LadybugDBID_LadybugDB != nil {
		err := errors.New("In Init, backRepoLadybug.Map_LadybugDBID_LadybugDB should be nil")
		return err
	}

	if backRepoLadybug.Map_LadybugPtr_LadybugDBID != nil {
		err := errors.New("In Init, backRepoLadybug.Map_LadybugPtr_LadybugDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Ladybug, 0)
	backRepoLadybug.Map_LadybugDBID_LadybugPtr = &tmp

	tmpDB := make(map[uint]*LadybugDB, 0)
	backRepoLadybug.Map_LadybugDBID_LadybugDB = &tmpDB

	tmpID := make(map[*models.Ladybug]uint, 0)
	backRepoLadybug.Map_LadybugPtr_LadybugDBID = &tmpID

	backRepoLadybug.db = db
	return
}

// BackRepoLadybug.CommitPhaseOne commits all staged instances of Ladybug to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLadybug *BackRepoLadybugStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for ladybug := range stage.Ladybugs {
		backRepoLadybug.CommitPhaseOneInstance(ladybug)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, ladybug := range *backRepoLadybug.Map_LadybugDBID_LadybugPtr {
		if _, ok := stage.Ladybugs[ladybug]; !ok {
			backRepoLadybug.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLadybug.CommitDeleteInstance commits deletion of Ladybug to the BackRepo
func (backRepoLadybug *BackRepoLadybugStruct) CommitDeleteInstance(id uint) (Error error) {

	ladybug := (*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[id]

	// ladybug is not staged anymore, remove ladybugDB
	ladybugDB := (*backRepoLadybug.Map_LadybugDBID_LadybugDB)[id]
	query := backRepoLadybug.db.Unscoped().Delete(&ladybugDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoLadybug.Map_LadybugPtr_LadybugDBID), ladybug)
	delete((*backRepoLadybug.Map_LadybugDBID_LadybugPtr), id)
	delete((*backRepoLadybug.Map_LadybugDBID_LadybugDB), id)

	return
}

// BackRepoLadybug.CommitPhaseOneInstance commits ladybug staged instances of Ladybug to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLadybug *BackRepoLadybugStruct) CommitPhaseOneInstance(ladybug *models.Ladybug) (Error error) {

	// check if the ladybug is not commited yet
	if _, ok := (*backRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]; ok {
		return
	}

	// initiate ladybug
	var ladybugDB LadybugDB
	ladybugDB.CopyBasicFieldsFromLadybug(ladybug)

	query := backRepoLadybug.db.Create(&ladybugDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug] = ladybugDB.ID
	(*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[ladybugDB.ID] = ladybug
	(*backRepoLadybug.Map_LadybugDBID_LadybugDB)[ladybugDB.ID] = &ladybugDB

	return
}

// BackRepoLadybug.CommitPhaseTwo commits all staged instances of Ladybug to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybug *BackRepoLadybugStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, ladybug := range *backRepoLadybug.Map_LadybugDBID_LadybugPtr {
		backRepoLadybug.CommitPhaseTwoInstance(backRepo, idx, ladybug)
	}

	return
}

// BackRepoLadybug.CommitPhaseTwoInstance commits {{structname }} of models.Ladybug to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybug *BackRepoLadybugStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, ladybug *models.Ladybug) (Error error) {

	// fetch matching ladybugDB
	if ladybugDB, ok := (*backRepoLadybug.Map_LadybugDBID_LadybugDB)[idx]; ok {

		ladybugDB.CopyBasicFieldsFromLadybug(ladybug)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoLadybug.db.Save(&ladybugDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Ladybug intance %s", ladybug.Name))
		return err
	}

	return
}

// BackRepoLadybug.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoLadybug *BackRepoLadybugStruct) CheckoutPhaseOne() (Error error) {

	ladybugDBArray := make([]LadybugDB, 0)
	query := backRepoLadybug.db.Find(&ladybugDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	ladybugInstancesToBeRemovedFromTheStage := make(map[*models.Ladybug]any)
	for key, value := range models.Stage.Ladybugs {
		ladybugInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, ladybugDB := range ladybugDBArray {
		backRepoLadybug.CheckoutPhaseOneInstance(&ladybugDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		ladybug, ok := (*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[ladybugDB.ID]
		if ok {
			delete(ladybugInstancesToBeRemovedFromTheStage, ladybug)
		}
	}

	// remove from stage and back repo's 3 maps all ladybugs that are not in the checkout
	for ladybug := range ladybugInstancesToBeRemovedFromTheStage {
		ladybug.Unstage()

		// remove instance from the back repo 3 maps
		ladybugID := (*backRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]
		delete((*backRepoLadybug.Map_LadybugPtr_LadybugDBID), ladybug)
		delete((*backRepoLadybug.Map_LadybugDBID_LadybugDB), ladybugID)
		delete((*backRepoLadybug.Map_LadybugDBID_LadybugPtr), ladybugID)
	}

	return
}

// CheckoutPhaseOneInstance takes a ladybugDB that has been found in the DB, updates the backRepo and stages the
// models version of the ladybugDB
func (backRepoLadybug *BackRepoLadybugStruct) CheckoutPhaseOneInstance(ladybugDB *LadybugDB) (Error error) {

	ladybug, ok := (*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[ladybugDB.ID]
	if !ok {
		ladybug = new(models.Ladybug)

		(*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[ladybugDB.ID] = ladybug
		(*backRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug] = ladybugDB.ID

		// append model store with the new element
		ladybug.Name = ladybugDB.Name_Data.String
		ladybug.Stage()
	}
	ladybugDB.CopyBasicFieldsToLadybug(ladybug)

	// preserve pointer to ladybugDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LadybugDBID_LadybugDB)[ladybugDB hold variable pointers
	ladybugDB_Data := *ladybugDB
	preservedPtrToLadybug := &ladybugDB_Data
	(*backRepoLadybug.Map_LadybugDBID_LadybugDB)[ladybugDB.ID] = preservedPtrToLadybug

	return
}

// BackRepoLadybug.CheckoutPhaseTwo Checkouts all staged instances of Ladybug to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybug *BackRepoLadybugStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, ladybugDB := range *backRepoLadybug.Map_LadybugDBID_LadybugDB {
		backRepoLadybug.CheckoutPhaseTwoInstance(backRepo, ladybugDB)
	}
	return
}

// BackRepoLadybug.CheckoutPhaseTwoInstance Checkouts staged instances of Ladybug to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybug *BackRepoLadybugStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, ladybugDB *LadybugDB) (Error error) {

	ladybug := (*backRepoLadybug.Map_LadybugDBID_LadybugPtr)[ladybugDB.ID]
	_ = ladybug // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitLadybug allows commit of a single ladybug (if already staged)
func (backRepo *BackRepoStruct) CommitLadybug(ladybug *models.Ladybug) {
	backRepo.BackRepoLadybug.CommitPhaseOneInstance(ladybug)
	if id, ok := (*backRepo.BackRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]; ok {
		backRepo.BackRepoLadybug.CommitPhaseTwoInstance(backRepo, id, ladybug)
	}
}

// CommitLadybug allows checkout of a single ladybug (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLadybug(ladybug *models.Ladybug) {
	// check if the ladybug is staged
	if _, ok := (*backRepo.BackRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]; ok {

		if id, ok := (*backRepo.BackRepoLadybug.Map_LadybugPtr_LadybugDBID)[ladybug]; ok {
			var ladybugDB LadybugDB
			ladybugDB.ID = id

			if err := backRepo.BackRepoLadybug.db.First(&ladybugDB, id).Error; err != nil {
				log.Panicln("CheckoutLadybug : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLadybug.CheckoutPhaseOneInstance(&ladybugDB)
			backRepo.BackRepoLadybug.CheckoutPhaseTwoInstance(backRepo, &ladybugDB)
		}
	}
}

// CopyBasicFieldsFromLadybug
func (ladybugDB *LadybugDB) CopyBasicFieldsFromLadybug(ladybug *models.Ladybug) {
	// insertion point for fields commit

	ladybugDB.TechName_Data.String = ladybug.TechName
	ladybugDB.TechName_Data.Valid = true

	ladybugDB.Name_Data.String = ladybug.Name
	ladybugDB.Name_Data.Valid = true

	ladybugDB.Id_Data.Int64 = int64(ladybug.Id)
	ladybugDB.Id_Data.Valid = true

	ladybugDB.Position_Data.Float64 = ladybug.Position
	ladybugDB.Position_Data.Valid = true

	ladybugDB.Speed_Data.Float64 = ladybug.Speed
	ladybugDB.Speed_Data.Valid = true

	ladybugDB.LadybugStatus_Data.String = ladybug.LadybugStatus.ToString()
	ladybugDB.LadybugStatus_Data.Valid = true
}

// CopyBasicFieldsFromLadybugWOP
func (ladybugDB *LadybugDB) CopyBasicFieldsFromLadybugWOP(ladybug *LadybugWOP) {
	// insertion point for fields commit

	ladybugDB.TechName_Data.String = ladybug.TechName
	ladybugDB.TechName_Data.Valid = true

	ladybugDB.Name_Data.String = ladybug.Name
	ladybugDB.Name_Data.Valid = true

	ladybugDB.Id_Data.Int64 = int64(ladybug.Id)
	ladybugDB.Id_Data.Valid = true

	ladybugDB.Position_Data.Float64 = ladybug.Position
	ladybugDB.Position_Data.Valid = true

	ladybugDB.Speed_Data.Float64 = ladybug.Speed
	ladybugDB.Speed_Data.Valid = true

	ladybugDB.LadybugStatus_Data.String = ladybug.LadybugStatus.ToString()
	ladybugDB.LadybugStatus_Data.Valid = true
}

// CopyBasicFieldsToLadybug
func (ladybugDB *LadybugDB) CopyBasicFieldsToLadybug(ladybug *models.Ladybug) {
	// insertion point for checkout of basic fields (back repo to stage)
	ladybug.TechName = ladybugDB.TechName_Data.String
	ladybug.Name = ladybugDB.Name_Data.String
	ladybug.Id = int(ladybugDB.Id_Data.Int64)
	ladybug.Position = ladybugDB.Position_Data.Float64
	ladybug.Speed = ladybugDB.Speed_Data.Float64
	ladybug.LadybugStatus.FromString(ladybugDB.LadybugStatus_Data.String)
}

// CopyBasicFieldsToLadybugWOP
func (ladybugDB *LadybugDB) CopyBasicFieldsToLadybugWOP(ladybug *LadybugWOP) {
	ladybug.ID = int(ladybugDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	ladybug.TechName = ladybugDB.TechName_Data.String
	ladybug.Name = ladybugDB.Name_Data.String
	ladybug.Id = int(ladybugDB.Id_Data.Int64)
	ladybug.Position = ladybugDB.Position_Data.Float64
	ladybug.Speed = ladybugDB.Speed_Data.Float64
	ladybug.LadybugStatus.FromString(ladybugDB.LadybugStatus_Data.String)
}

// Backup generates a json file from a slice of all LadybugDB instances in the backrepo
func (backRepoLadybug *BackRepoLadybugStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LadybugDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LadybugDB, 0)
	for _, ladybugDB := range *backRepoLadybug.Map_LadybugDBID_LadybugDB {
		forBackup = append(forBackup, ladybugDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Ladybug ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Ladybug file", err.Error())
	}
}

// Backup generates a json file from a slice of all LadybugDB instances in the backrepo
func (backRepoLadybug *BackRepoLadybugStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LadybugDB, 0)
	for _, ladybugDB := range *backRepoLadybug.Map_LadybugDBID_LadybugDB {
		forBackup = append(forBackup, ladybugDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Ladybug")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Ladybug_Fields, -1)
	for _, ladybugDB := range forBackup {

		var ladybugWOP LadybugWOP
		ladybugDB.CopyBasicFieldsToLadybugWOP(&ladybugWOP)

		row := sh.AddRow()
		row.WriteStruct(&ladybugWOP, -1)
	}
}

// RestoreXL from the "Ladybug" sheet all LadybugDB instances
func (backRepoLadybug *BackRepoLadybugStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLadybugid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Ladybug"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLadybug.rowVisitorLadybug)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLadybug *BackRepoLadybugStruct) rowVisitorLadybug(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var ladybugWOP LadybugWOP
		row.ReadStruct(&ladybugWOP)

		// add the unmarshalled struct to the stage
		ladybugDB := new(LadybugDB)
		ladybugDB.CopyBasicFieldsFromLadybugWOP(&ladybugWOP)

		ladybugDB_ID_atBackupTime := ladybugDB.ID
		ladybugDB.ID = 0
		query := backRepoLadybug.db.Create(ladybugDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLadybug.Map_LadybugDBID_LadybugDB)[ladybugDB.ID] = ladybugDB
		BackRepoLadybugid_atBckpTime_newID[ladybugDB_ID_atBackupTime] = ladybugDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LadybugDB.json" in dirPath that stores an array
// of LadybugDB and stores it in the database
// the map BackRepoLadybugid_atBckpTime_newID is updated accordingly
func (backRepoLadybug *BackRepoLadybugStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLadybugid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LadybugDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Ladybug file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LadybugDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LadybugDBID_LadybugDB
	for _, ladybugDB := range forRestore {

		ladybugDB_ID_atBackupTime := ladybugDB.ID
		ladybugDB.ID = 0
		query := backRepoLadybug.db.Create(ladybugDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLadybug.Map_LadybugDBID_LadybugDB)[ladybugDB.ID] = ladybugDB
		BackRepoLadybugid_atBckpTime_newID[ladybugDB_ID_atBackupTime] = ladybugDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Ladybug file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Ladybug>id_atBckpTime_newID
// to compute new index
func (backRepoLadybug *BackRepoLadybugStruct) RestorePhaseTwo() {

	for _, ladybugDB := range *backRepoLadybug.Map_LadybugDBID_LadybugDB {

		// next line of code is to avert unused variable compilation error
		_ = ladybugDB

		// insertion point for reindexing pointers encoding
		// This reindex ladybug.Ladybugs
		if ladybugDB.LadybugSimulation_LadybugsDBID.Int64 != 0 {
			ladybugDB.LadybugSimulation_LadybugsDBID.Int64 =
				int64(BackRepoLadybugSimulationid_atBckpTime_newID[uint(ladybugDB.LadybugSimulation_LadybugsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLadybug.db.Model(ladybugDB).Updates(*ladybugDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLadybugid_atBckpTime_newID map[uint]uint
