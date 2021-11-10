package models

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const BackRepoPerStructTemplateCode = `// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
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

	"{{PkgPathRoot}}"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_{{Structname}}_sql sql.NullBool
var dummy_{{Structname}}_time time.Duration
var dummy_{{Structname}}_sort sort.Float64Slice

// {{Structname}}API is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model {{structname}}API
type {{Structname}}API struct {
	gorm.Model

	models.{{Structname}}

	// encoding of pointers
	{{Structname}}PointersEnconding
}

// {{Structname}}PointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type {{Structname}}PointersEnconding struct {
	// insertion for pointer fields encoding declaration{{` + string(rune(BackRepoPointerEncodingFieldsDeclaration)) + `}}
}

// {{Structname}}DB describes a {{structname}} in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model {{structname}}DB
type {{Structname}}DB struct {
	gorm.Model

	// insertion for basic fields declaration{{` + string(rune(BackRepoBasicFieldsDeclaration)) + `}}
	// encoding of pointers
	{{Structname}}PointersEnconding
}

// {{Structname}}DBs arrays {{structname}}DBs
// swagger:response {{structname}}DBsResponse
type {{Structname}}DBs []{{Structname}}DB

// {{Structname}}DBResponse provides response
// swagger:response {{structname}}DBResponse
type {{Structname}}DBResponse struct {
	{{Structname}}DB
}

// {{Structname}}WOP is a {{Structname}} without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type {{Structname}}WOP struct {
	ID int

	// insertion for WOP basic fields{{` + string(rune(BackRepoBasicAndTimeFieldsWOPDeclaration)) + `}}
	// insertion for WOP pointer fields{{` + string(rune(BackRepoPointerEncodingFieldsWOPDeclaration)) + `}}
}

var {{Structname}}_Fields = []string{
	// insertion for WOP basic fields{{` + string(rune(BackRepoBasicAndTimeFieldsName)) + `}}
}

type BackRepo{{Structname}}Struct struct {
	// stores {{Structname}}DB according to their gorm ID
	Map_{{Structname}}DBID_{{Structname}}DB *map[uint]*{{Structname}}DB

	// stores {{Structname}}DB ID according to {{Structname}} address
	Map_{{Structname}}Ptr_{{Structname}}DBID *map[*models.{{Structname}}]uint

	// stores {{Structname}} according to their gorm ID
	Map_{{Structname}}DBID_{{Structname}}Ptr *map[uint]*models.{{Structname}}

	db *gorm.DB
}

func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) GetDB() *gorm.DB {
	return backRepo{{Structname}}.db
}

// Get{{Structname}}DBFrom{{Structname}}Ptr is a handy function to access the back repo instance from the stage instance
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) Get{{Structname}}DBFrom{{Structname}}Ptr({{structname}} *models.{{Structname}}) ({{structname}}DB *{{Structname}}DB) {
	id := (*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]
	{{structname}}DB = (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[id]
	return
}

// BackRepo{{Structname}}.Init set up the BackRepo of the {{Structname}}
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) Init(db *gorm.DB) (Error error) {

	if backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr != nil {
		err := errors.New("In Init, backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr should be nil")
		return err
	}

	if backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB != nil {
		err := errors.New("In Init, backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB should be nil")
		return err
	}

	if backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID != nil {
		err := errors.New("In Init, backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.{{Structname}}, 0)
	backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr = &tmp

	tmpDB := make(map[uint]*{{Structname}}DB, 0)
	backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB = &tmpDB

	tmpID := make(map[*models.{{Structname}}]uint, 0)
	backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID = &tmpID

	backRepo{{Structname}}.db = db
	return
}

// BackRepo{{Structname}}.CommitPhaseOne commits all staged instances of {{Structname}} to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for {{structname}} := range stage.{{Structname}}s {
		backRepo{{Structname}}.CommitPhaseOneInstance({{structname}})
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, {{structname}} := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr {
		if _, ok := stage.{{Structname}}s[{{structname}}]; !ok {
			backRepo{{Structname}}.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepo{{Structname}}.CommitDeleteInstance commits deletion of {{Structname}} to the BackRepo
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CommitDeleteInstance(id uint) (Error error) {

	{{structname}} := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[id]

	// {{structname}} is not staged anymore, remove {{structname}}DB
	{{structname}}DB := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[id]
	query := backRepo{{Structname}}.db.Unscoped().Delete(&{{structname}}DB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID), {{structname}})
	delete((*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr), id)
	delete((*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB), id)

	return
}

// BackRepo{{Structname}}.CommitPhaseOneInstance commits {{structname}} staged instances of {{Structname}} to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CommitPhaseOneInstance({{structname}} *models.{{Structname}}) (Error error) {

	// check if the {{structname}} is not commited yet
	if _, ok := (*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]; ok {
		return
	}

	// initiate {{structname}}
	var {{structname}}DB {{Structname}}DB
	{{structname}}DB.CopyBasicFieldsFrom{{Structname}}({{structname}})

	query := backRepo{{Structname}}.db.Create(&{{structname}}DB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}] = {{structname}}DB.ID
	(*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[{{structname}}DB.ID] = {{structname}}
	(*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[{{structname}}DB.ID] = &{{structname}}DB

	return
}

// BackRepo{{Structname}}.CommitPhaseTwo commits all staged instances of {{Structname}} to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, {{structname}} := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr {
		backRepo{{Structname}}.CommitPhaseTwoInstance(backRepo, idx, {{structname}})
	}

	return
}

// BackRepo{{Structname}}.CommitPhaseTwoInstance commits {{structname }} of models.{{Structname}} to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, {{structname}} *models.{{Structname}}) (Error error) {

	// fetch matching {{structname}}DB
	if {{structname}}DB, ok := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[idx]; ok {

		{{structname}}DB.CopyBasicFieldsFrom{{Structname}}({{structname}})

		// insertion point for translating pointers encodings into actual pointers{{` + string(rune(BackRepoPointerEncodingFieldsCommit)) + `}}
		query := backRepo{{Structname}}.db.Save(&{{structname}}DB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown {{Structname}} intance %s", {{structname}}.Name))
		return err
	}

	return
}

// BackRepo{{Structname}}.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CheckoutPhaseOne() (Error error) {

	{{structname}}DBArray := make([]{{Structname}}DB, 0)
	query := backRepo{{Structname}}.db.Find(&{{structname}}DBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	{{structname}}InstancesToBeRemovedFromTheStage := make(map[*models.{{Structname}}]struct{})
	for key, value := range models.Stage.{{Structname}}s {
		{{structname}}InstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, {{structname}}DB := range {{structname}}DBArray {
		backRepo{{Structname}}.CheckoutPhaseOneInstance(&{{structname}}DB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		{{structname}}, ok := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[{{structname}}DB.ID]
		if ok {
			delete({{structname}}InstancesToBeRemovedFromTheStage, {{structname}})
		}
	}

	// remove from stage and back repo's 3 maps all {{structname}}s that are not in the checkout
	for {{structname}} := range {{structname}}InstancesToBeRemovedFromTheStage {
		{{structname}}.Unstage()

		// remove instance from the back repo 3 maps
		{{structname}}ID := (*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]
		delete((*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID), {{structname}})
		delete((*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB), {{structname}}ID)
		delete((*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr), {{structname}}ID)
	}

	return
}

// CheckoutPhaseOneInstance takes a {{structname}}DB that has been found in the DB, updates the backRepo and stages the
// models version of the {{structname}}DB
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CheckoutPhaseOneInstance({{structname}}DB *{{Structname}}DB) (Error error) {

	{{structname}}, ok := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[{{structname}}DB.ID]
	if !ok {
		{{structname}} = new(models.{{Structname}})

		(*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[{{structname}}DB.ID] = {{structname}}
		(*backRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}] = {{structname}}DB.ID

		// append model store with the new element
		{{structname}}.Name = {{structname}}DB.Name_Data.String
		{{structname}}.Stage()
	}
	{{structname}}DB.CopyBasicFieldsTo{{Structname}}({{structname}})

	// preserve pointer to {{structname}}DB. Otherwise, pointer will is recycled and the map of pointers
	// Map_{{Structname}}DBID_{{Structname}}DB)[{{structname}}DB hold variable pointers
	{{structname}}DB_Data := *{{structname}}DB
	preservedPtrTo{{Structname}} := &{{structname}}DB_Data
	(*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[{{structname}}DB.ID] = preservedPtrTo{{Structname}}

	return
}

// BackRepo{{Structname}}.CheckoutPhaseTwo Checkouts all staged instances of {{Structname}} to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, {{structname}}DB := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB {
		backRepo{{Structname}}.CheckoutPhaseTwoInstance(backRepo, {{structname}}DB)
	}
	return
}

// BackRepo{{Structname}}.CheckoutPhaseTwoInstance Checkouts staged instances of {{Structname}} to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, {{structname}}DB *{{Structname}}DB) (Error error) {

	{{structname}} := (*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}Ptr)[{{structname}}DB.ID]
	_ = {{structname}} // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding{{` + string(rune(BackRepoPointerEncodingFieldsCheckout)) + `}}
	return
}

// Commit{{Structname}} allows commit of a single {{structname}} (if already staged)
func (backRepo *BackRepoStruct) Commit{{Structname}}({{structname}} *models.{{Structname}}) {
	backRepo.BackRepo{{Structname}}.CommitPhaseOneInstance({{structname}})
	if id, ok := (*backRepo.BackRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]; ok {
		backRepo.BackRepo{{Structname}}.CommitPhaseTwoInstance(backRepo, id, {{structname}})
	}
}

// Commit{{Structname}} allows checkout of a single {{structname}} (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) Checkout{{Structname}}({{structname}} *models.{{Structname}}) {
	// check if the {{structname}} is staged
	if _, ok := (*backRepo.BackRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]; ok {

		if id, ok := (*backRepo.BackRepo{{Structname}}.Map_{{Structname}}Ptr_{{Structname}}DBID)[{{structname}}]; ok {
			var {{structname}}DB {{Structname}}DB
			{{structname}}DB.ID = id

			if err := backRepo.BackRepo{{Structname}}.db.First(&{{structname}}DB, id).Error; err != nil {
				log.Panicln("Checkout{{Structname}} : Problem with getting object with id:", id)
			}
			backRepo.BackRepo{{Structname}}.CheckoutPhaseOneInstance(&{{structname}}DB)
			backRepo.BackRepo{{Structname}}.CheckoutPhaseTwoInstance(backRepo, &{{structname}}DB)
		}
	}
}

// CopyBasicFieldsFrom{{Structname}}
func ({{structname}}DB *{{Structname}}DB) CopyBasicFieldsFrom{{Structname}}({{structname}} *models.{{Structname}}) {
	// insertion point for fields commit{{` + string(rune(BackRepoBasicFieldsCommit)) + `}}
}

// CopyBasicFieldsFrom{{Structname}}WOP
func ({{structname}}DB *{{Structname}}DB) CopyBasicFieldsFrom{{Structname}}WOP({{structname}} *{{Structname}}WOP) {
	// insertion point for fields commit{{` + string(rune(BackRepoBasicFieldsCommit)) + `}}
}

// CopyBasicFieldsTo{{Structname}}
func ({{structname}}DB *{{Structname}}DB) CopyBasicFieldsTo{{Structname}}({{structname}} *models.{{Structname}}) {
	// insertion point for checkout of basic fields (back repo to stage){{` + string(rune(BackRepoBasicFieldsCheckout)) + `}}
}

// CopyBasicFieldsTo{{Structname}}WOP
func ({{structname}}DB *{{Structname}}DB) CopyBasicFieldsTo{{Structname}}WOP({{structname}} *{{Structname}}WOP) {
	{{structname}}.ID = int({{structname}}DB.ID)
	// insertion point for checkout of basic fields (back repo to stage){{` + string(rune(BackRepoBasicFieldsCheckout)) + `}}
}

// Backup generates a json file from a slice of all {{Structname}}DB instances in the backrepo
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "{{Structname}}DB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*{{Structname}}DB, 0)
	for _, {{structname}}DB := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB {
		forBackup = append(forBackup, {{structname}}DB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json {{Structname}} ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json {{Structname}} file", err.Error())
	}
}

// Backup generates a json file from a slice of all {{Structname}}DB instances in the backrepo
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*{{Structname}}DB, 0)
	for _, {{structname}}DB := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB {
		forBackup = append(forBackup, {{structname}}DB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("{{Structname}}")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&{{Structname}}_Fields, -1)
	for _, {{structname}}DB := range forBackup {

		var {{structname}}WOP {{Structname}}WOP
		{{structname}}DB.CopyBasicFieldsTo{{Structname}}WOP(&{{structname}}WOP)

		row := sh.AddRow()
		row.WriteStruct(&{{structname}}WOP, -1)
	}
}

// RestorePhaseOne read the file "{{Structname}}DB.json" in dirPath that stores an array
// of {{Structname}}DB and stores it in the database
// the map BackRepo{{Structname}}id_atBckpTime_newID is updated accordingly
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepo{{Structname}}id_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "{{Structname}}DB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json {{Structname}} file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*{{Structname}}DB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_{{Structname}}DBID_{{Structname}}DB
	for _, {{structname}}DB := range forRestore {

		{{structname}}DB_ID_atBackupTime := {{structname}}DB.ID
		{{structname}}DB.ID = 0
		query := backRepo{{Structname}}.db.Create({{structname}}DB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB)[{{structname}}DB.ID] = {{structname}}DB
		BackRepo{{Structname}}id_atBckpTime_newID[{{structname}}DB_ID_atBackupTime] = {{structname}}DB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json {{Structname}} file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<{{Structname}}>id_atBckpTime_newID
// to compute new index
func (backRepo{{Structname}} *BackRepo{{Structname}}Struct) RestorePhaseTwo() {

	for _, {{structname}}DB := range *backRepo{{Structname}}.Map_{{Structname}}DBID_{{Structname}}DB {

		// next line of code is to avert unused variable compilation error
		_ = {{structname}}DB

		// insertion point for reindexing pointers encoding{{` + string(rune(BackRepoPointerEncodingFieldsReindexing)) + `}}
		// update databse with new index encoding
		query := backRepo{{Structname}}.db.Model({{structname}}DB).Updates(*{{structname}}DB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepo{{Structname}}id_atBckpTime_newID map[uint]uint
`

// insertion points
type BackRepoInsertionPoint int

const (
	BackRepoBasicFieldsDeclaration BackRepoInsertionPoint = iota
	BackRepoBasicAndTimeFieldsName
	BackRepoBasicAndTimeFieldsWOPDeclaration
	BackRepoPointerEncodingFieldsDeclaration
	BackRepoPointerEncodingFieldsWOPDeclaration
	BackRepoBasicFieldsCommit
	BackRepoPointerEncodingFieldsCommit
	BackRepoBasicFieldsCheckout
	BackRepoPointerEncodingFieldsCheckout
	BackRepoPointerEncodingFieldsReindexing
	BackRepoNbInsertionPoints
)

//
// Sub Templates
//
type BackRepoPerStructSubTemplate int

const (
	BackRepoDeclarationBasicField BackRepoPerStructSubTemplate = iota
	BackRepoCommitBasicField
	BackRepoCheckoutBasicField

	BackRepoDeclarationTimeField
	BackRepoCommitTimeField
	BackRepoCheckoutTimeField

	BackRepoCommitBasicFieldEnum
	BackRepoCheckoutBasicFieldEnum

	BackRepoCommitBasicFieldInt
	BackRepoCheckoutBasicFieldInt

	BackRepoDeclarationBasicBooleanField
	BackRepoCommitBasicBooleanField
	BackRepoCheckoutBasicFieldBoolean

	BackRepoDeclarationPointerToStructField
	BackRepoCommitPointerToStructField
	BackRepoCheckoutPointerToStructStageField
	BackRepoReindexingPointerToStruct

	BackRepoDeclarationSliceOfPointerToStructField
	BackRepoCommitSliceOfPointerToStructField
	BackRepoCheckoutSliceOfPointerToStructStageField
	BackRepoReindexingSliceOfPointerToStruct
)

var BackRepoFieldSubTemplateCode map[BackRepoPerStructSubTemplate]string = map[BackRepoPerStructSubTemplate]string{

	//
	// Declarations
	//

	BackRepoDeclarationBasicField: `
	// Declation for basic field {{structname}}DB.{{FieldName}} {{BasicKind}} (to be completed)
	{{FieldName}}_Data sql.{{SqlNullType}}
`,

	BackRepoDeclarationTimeField: `
	// Declation for basic field {{structname}}DB.{{FieldName}}
	{{FieldName}}_Data sql.NullTime
`,

	BackRepoDeclarationBasicBooleanField: `
	// Declation for basic field {{structname}}DB.{{FieldName}} {{BasicKind}} (to be completed)
	// provide the sql storage for the boolan
	{{FieldName}}_Data sql.NullBool
`,

	BackRepoDeclarationPointerToStructField: `
	// field {{FieldName}} is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	{{FieldName}}ID sql.NullInt64
`,

	BackRepoDeclarationSliceOfPointerToStructField: `
	// Implementation of a reverse ID for field {{AssociationStructName}}{}.{{FieldName}} []*{{Structname}}
	{{AssociationStructName}}_{{FieldName}}DBID sql.NullInt64

	// implementation of the index of the withing the slice
	{{AssociationStructName}}_{{FieldName}}DBID_Index sql.NullInt64`,

	//
	// Commit sub templates
	//

	BackRepoCommitBasicField: `
	{{structname}}DB.{{FieldName}}_Data.{{SqlNullType}} = {{structname}}.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	BackRepoCommitBasicFieldEnum: `
	{{structname}}DB.{{FieldName}}_Data.String = string({{structname}}.{{FieldName}})
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	BackRepoCommitBasicFieldInt: `
	{{structname}}DB.{{FieldName}}_Data.Int64 = int64({{structname}}.{{FieldName}})
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,
	BackRepoCommitTimeField: `
	{{structname}}DB.{{FieldName}}_Data.Time = {{structname}}.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	BackRepoCommitBasicBooleanField: `
	{{structname}}DB.{{FieldName}}_Data.Bool = {{structname}}.{{FieldName}}
	{{structname}}DB.{{FieldName}}_Data.Valid = true
`,

	BackRepoCommitPointerToStructField: `
		// commit pointer value {{structname}}.{{FieldName}} translates to updating the {{structname}}.{{FieldName}}ID
		{{structname}}DB.{{FieldName}}ID.Valid = true // allow for a 0 value (nil association)
		if {{structname}}.{{FieldName}} != nil {
			if {{FieldName}}Id, ok := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}Ptr_{{AssociationStructName}}DBID)[{{structname}}.{{FieldName}}]; ok {
				{{structname}}DB.{{FieldName}}ID.Int64 = int64({{FieldName}}Id)
				{{structname}}DB.{{FieldName}}ID.Valid = true
			}
		}
`,

	BackRepoCommitSliceOfPointerToStructField: `
		// This loop encodes the slice of pointers {{structname}}.{{FieldName}} into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, {{associationStructName}}AssocEnd := range {{structname}}.{{FieldName}} {

			// get the back repo instance at the association end
			{{associationStructName}}AssocEnd_DB :=
				backRepo.BackRepo{{AssociationStructName}}.Get{{AssociationStructName}}DBFrom{{AssociationStructName}}Ptr({{associationStructName}}AssocEnd)

			// encode reverse pointer in the association end back repo instance
			{{associationStructName}}AssocEnd_DB.{{Structname}}_{{FieldName}}DBID.Int64 = int64({{structname}}DB.ID)
			{{associationStructName}}AssocEnd_DB.{{Structname}}_{{FieldName}}DBID.Valid = true
			{{associationStructName}}AssocEnd_DB.{{Structname}}_{{FieldName}}DBID_Index.Int64 = int64(idx)
			{{associationStructName}}AssocEnd_DB.{{Structname}}_{{FieldName}}DBID_Index.Valid = true
			if q := backRepo{{Structname}}.db.Save({{associationStructName}}AssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}
`,

	//
	// sub template for checkouts
	//

	BackRepoCheckoutBasicField: `
	{{structname}}.{{FieldName}} = {{structname}}DB.{{FieldName}}_Data.{{SqlNullType}}`,

	BackRepoCheckoutTimeField: `
	{{structname}}.{{FieldName}} = {{structname}}DB.{{FieldName}}_Data.Time`,

	BackRepoCheckoutBasicFieldEnum: `
	{{structname}}.{{FieldName}} = models.{{EnumType}}({{structname}}DB.{{FieldName}}_Data.String)`,

	BackRepoCheckoutBasicFieldInt: `
	{{structname}}.{{FieldName}} = {{FieldType}}({{structname}}DB.{{FieldName}}_Data.Int64)`,

	BackRepoCheckoutBasicFieldBoolean: `
	{{structname}}.{{FieldName}} = {{structname}}DB.{{FieldName}}_Data.Bool`,

	BackRepoCheckoutPointerToStructStageField: `
	// {{FieldName}} field
	if {{structname}}DB.{{FieldName}}ID.Int64 != 0 {
		{{structname}}.{{FieldName}} = (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}DBID_{{AssociationStructName}}Ptr)[uint({{structname}}DB.{{FieldName}}ID.Int64)]
	}`,

	BackRepoReindexingPointerToStruct: `
		// reindexing {{FieldName}} field
		if {{structname}}DB.{{FieldName}}ID.Int64 != 0 {
			{{structname}}DB.{{FieldName}}ID.Int64 = int64(BackRepo{{AssociationStructName}}id_atBckpTime_newID[uint({{structname}}DB.{{FieldName}}ID.Int64)])
			{{structname}}DB.{{FieldName}}ID.Valid = true
		}
`,

	BackRepoCheckoutSliceOfPointerToStructStageField: `
	// This loop redeem {{structname}}.{{FieldName}} in the stage from the encode in the back repo
	// It parses all {{AssociationStructName}}DB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	{{structname}}.{{FieldName}} = {{structname}}.{{FieldName}}[:0]
	// 2. loop all instances in the type in the association end
	for _, {{associationStructName}}DB_AssocEnd := range *backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}DBID_{{AssociationStructName}}DB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if {{associationStructName}}DB_AssocEnd.{{Structname}}_{{FieldName}}DBID.Int64 == int64({{structname}}DB.ID) {
			// 4. fetch the associated instance in the stage
			{{associationStructName}}_AssocEnd := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}DBID_{{AssociationStructName}}Ptr)[{{associationStructName}}DB_AssocEnd.ID]
			// 5. append it the association slice
			{{structname}}.{{FieldName}} = append({{structname}}.{{FieldName}}, {{associationStructName}}_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice({{structname}}.{{FieldName}}, func(i, j int) bool {
		{{associationStructName}}DB_i_ID := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}Ptr_{{AssociationStructName}}DBID)[{{structname}}.{{FieldName}}[i]]
		{{associationStructName}}DB_j_ID := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}Ptr_{{AssociationStructName}}DBID)[{{structname}}.{{FieldName}}[j]]

		{{associationStructName}}DB_i := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}DBID_{{AssociationStructName}}DB)[{{associationStructName}}DB_i_ID]
		{{associationStructName}}DB_j := (*backRepo.BackRepo{{AssociationStructName}}.Map_{{AssociationStructName}}DBID_{{AssociationStructName}}DB)[{{associationStructName}}DB_j_ID]

		return {{associationStructName}}DB_i.{{Structname}}_{{FieldName}}DBID_Index.Int64 < {{associationStructName}}DB_j.{{Structname}}_{{FieldName}}DBID_Index.Int64
	})
`,

	BackRepoReindexingSliceOfPointerToStruct: `
		// This reindex {{structname}}.{{FieldName}}
		if {{structname}}DB.{{AssociationStructName}}_{{FieldName}}DBID.Int64 != 0 {
			{{structname}}DB.{{AssociationStructName}}_{{FieldName}}DBID.Int64 =
				int64(BackRepo{{AssociationStructName}}id_atBckpTime_newID[uint({{structname}}DB.{{AssociationStructName}}_{{FieldName}}DBID.Int64)])
		}
`,
}

// MultiCodeGeneratorBackRepo parses mdlPkg and generates the code for the
// back repository code
func MultiCodeGeneratorBackRepo(
	mdlPkg *ModelPkg,
	pkgName string,
	pkgGoPath string,
	dirPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		if _struct.HasNameField() {
			structList = append(structList, _struct)
		}
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		codeGO := BackRepoPerStructTemplateCode

		insertions := make(map[BackRepoInsertionPoint]string)
		for insertion := BackRepoInsertionPoint(0); insertion < BackRepoNbInsertionPoints; insertion++ {
			insertions[insertion] = ""
		}

		insertions[BackRepoBasicAndTimeFieldsName] += "\n\t\"ID\","

		for _, field := range _struct.Fields {

			switch field := field.(type) {
			case *GongBasicField:

				insertions[BackRepoBasicAndTimeFieldsWOPDeclaration] +=
					"\n\n\t" + field.Name + " " +
						strings.ReplaceAll(field.DeclaredType, pkgGoPath+".", "models.")

				insertions[BackRepoBasicAndTimeFieldsName] += "\n\t\"" + field.Name + "\","

				if field.basicKind == types.Bool {

					insertions[BackRepoBasicFieldsDeclaration] += Replace2(
						BackRepoFieldSubTemplateCode[BackRepoDeclarationBasicBooleanField],
						"{{FieldName}}", field.Name,
						"{{BasicKind}}", field.Type.Underlying().String())

					insertions[BackRepoBasicFieldsCommit] += Replace1(
						BackRepoFieldSubTemplateCode[BackRepoCommitBasicBooleanField],
						"{{FieldName}}", field.Name)

					insertions[BackRepoBasicFieldsCheckout] += Replace1(
						BackRepoFieldSubTemplateCode[BackRepoCheckoutBasicFieldBoolean],
						"{{FieldName}}", field.Name)

				} else {
					switch field.basicKind {
					case types.String:
						insertions[BackRepoBasicFieldsDeclaration] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoDeclarationBasicField],
							"{{FieldName}}", field.Name,
							"{{SqlNullType}}", "NullString")

						if field.GongEnum != nil {
							insertions[BackRepoBasicFieldsCommit] += Replace1(
								BackRepoFieldSubTemplateCode[BackRepoCommitBasicFieldEnum],
								"{{FieldName}}", field.Name)

							insertions[BackRepoBasicFieldsCheckout] += Replace2(
								BackRepoFieldSubTemplateCode[BackRepoCheckoutBasicFieldEnum],
								"{{FieldName}}", field.Name,
								"{{EnumType}}", field.GongEnum.Name)

						} else {
							insertions[BackRepoBasicFieldsCommit] += Replace2(
								BackRepoFieldSubTemplateCode[BackRepoCommitBasicField],
								"{{FieldName}}", field.Name,
								"{{SqlNullType}}", "String")

							insertions[BackRepoBasicFieldsCheckout] += Replace2(
								BackRepoFieldSubTemplateCode[BackRepoCheckoutBasicField],
								"{{FieldName}}", field.Name,
								"{{SqlNullType}}", "String")

						}
					case types.Float64:
						insertions[BackRepoBasicFieldsDeclaration] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoDeclarationBasicField],
							"{{FieldName}}", field.Name,
							"{{SqlNullType}}", "NullFloat64")

						insertions[BackRepoBasicFieldsCommit] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoCommitBasicField],
							"{{FieldName}}", field.Name,
							"{{SqlNullType}}", "Float64")

						insertions[BackRepoBasicFieldsCheckout] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoCheckoutBasicField],
							"{{FieldName}}", field.Name,
							"{{SqlNullType}}", "Float64")
					case types.Int, types.Int64:
						insertions[BackRepoBasicFieldsDeclaration] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoDeclarationBasicField],
							"{{FieldName}}", field.Name,
							"{{SqlNullType}}", "NullInt64")

						insertions[BackRepoBasicFieldsCommit] += Replace1(
							BackRepoFieldSubTemplateCode[BackRepoCommitBasicFieldInt],
							"{{FieldName}}", field.Name)

						insertions[BackRepoBasicFieldsCheckout] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoCheckoutBasicFieldInt],
							"{{FieldName}}", field.Name,
							"{{FieldType}}", field.DeclaredType)
					default:
					}
				}

			case *GongTimeField:

				insertions[BackRepoBasicAndTimeFieldsWOPDeclaration] +=
					"\n\n\t" + field.Name + " " + "time.Time"

				insertions[BackRepoBasicAndTimeFieldsName] += "\n\t\"" + field.Name + "\","

				insertions[BackRepoBasicFieldsDeclaration] += Replace1(
					BackRepoFieldSubTemplateCode[BackRepoDeclarationTimeField],
					"{{FieldName}}", field.Name)

				insertions[BackRepoBasicFieldsCheckout] += Replace1(
					BackRepoFieldSubTemplateCode[BackRepoCheckoutTimeField],
					"{{FieldName}}", field.Name)

				insertions[BackRepoBasicFieldsCommit] += Replace1(
					BackRepoFieldSubTemplateCode[BackRepoCommitTimeField],
					"{{FieldName}}", field.Name)

			case *PointerToGongStructField:

				insertions[BackRepoPointerEncodingFieldsDeclaration] += Replace1(
					BackRepoFieldSubTemplateCode[BackRepoDeclarationPointerToStructField],
					"{{FieldName}}", field.Name)

				insertions[BackRepoPointerEncodingFieldsCommit] += Replace3(
					BackRepoFieldSubTemplateCode[BackRepoCommitPointerToStructField],
					"{{AssociationStructName}}", field.GongStruct.Name,
					"{{associationStructName}}", strings.ToLower(field.GongStruct.Name),
					"{{FieldName}}", field.Name)

				insertions[BackRepoPointerEncodingFieldsCheckout] += Replace3(
					BackRepoFieldSubTemplateCode[BackRepoCheckoutPointerToStructStageField],
					"{{AssociationStructName}}", field.GongStruct.Name,
					"{{associationStructName}}", strings.ToLower(field.GongStruct.Name),
					"{{FieldName}}", field.Name)

				insertions[BackRepoPointerEncodingFieldsReindexing] += Replace3(
					BackRepoFieldSubTemplateCode[BackRepoReindexingPointerToStruct],
					"{{AssociationStructName}}", field.GongStruct.Name,
					"{{associationStructName}}", strings.ToLower(field.GongStruct.Name),
					"{{FieldName}}", field.Name)

			case *SliceOfPointerToGongStructField:

				insertions[BackRepoPointerEncodingFieldsCommit] += Replace3(
					BackRepoFieldSubTemplateCode[BackRepoCommitSliceOfPointerToStructField],
					"{{AssociationStructName}}", field.GongStruct.Name,
					"{{associationStructName}}", strings.ToLower(field.GongStruct.Name),
					"{{FieldName}}", field.Name)

				insertions[BackRepoPointerEncodingFieldsCheckout] += Replace3(
					BackRepoFieldSubTemplateCode[BackRepoCheckoutSliceOfPointerToStructStageField],
					"{{AssociationStructName}}", field.GongStruct.Name,
					"{{associationStructName}}", strings.ToLower(field.GongStruct.Name),
					"{{FieldName}}", field.Name)

			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {

						insertions[BackRepoPointerEncodingFieldsDeclaration] += Replace2(
							BackRepoFieldSubTemplateCode[BackRepoDeclarationSliceOfPointerToStructField],
							"{{FieldName}}", field.Name,
							"{{AssociationStructName}}", __struct.Name)

						insertions[BackRepoPointerEncodingFieldsReindexing] += Replace3(
							BackRepoFieldSubTemplateCode[BackRepoReindexingSliceOfPointerToStruct],
							"{{AssociationStructName}}", __struct.Name,
							"{{associationStructName}}", strings.ToLower(__struct.Name),
							"{{FieldName}}", field.Name)
					}
				}
			}
		}

		// substitutes {{<<insertion points>>}} stuff with generated code
		for insertion := BackRepoInsertionPoint(0); insertion < BackRepoNbInsertionPoints; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeGO = strings.ReplaceAll(codeGO, toReplace, insertions[insertion])
		}

		// substitutes struct level {{<...>}} stuff
		codeGO = Replace6(codeGO,
			"{{PkgName}}", pkgName,
			"{{TitlePkgName}}", strings.Title(pkgName),
			"{{pkgname}}", strings.ToLower(pkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(pkgGoPath, "/orm", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))

		file, err := os.Create(filepath.Join(dirPath, _struct.Name+"DB.go"))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()
		fmt.Fprint(file, codeGO)

	}
}
