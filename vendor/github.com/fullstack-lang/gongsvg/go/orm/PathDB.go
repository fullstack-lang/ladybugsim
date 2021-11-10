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

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Path_sql sql.NullBool
var dummy_Path_time time.Duration
var dummy_Path_sort sort.Float64Slice

// PathAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model pathAPI
type PathAPI struct {
	gorm.Model

	models.Path

	// encoding of pointers
	PathPointersEnconding
}

// PathPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type PathPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field SVG{}.Paths []*Path
	SVG_PathsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	SVG_PathsDBID_Index sql.NullInt64
}

// PathDB describes a path in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model pathDB
type PathDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field pathDB.Name {{BasicKind}} (to be completed)
	Name_Data sql.NullString

	// Declation for basic field pathDB.Definition {{BasicKind}} (to be completed)
	Definition_Data sql.NullString

	// Declation for basic field pathDB.Color {{BasicKind}} (to be completed)
	Color_Data sql.NullString

	// Declation for basic field pathDB.FillOpacity {{BasicKind}} (to be completed)
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field pathDB.Stroke {{BasicKind}} (to be completed)
	Stroke_Data sql.NullString

	// Declation for basic field pathDB.StrokeWidth {{BasicKind}} (to be completed)
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field pathDB.StrokeDashArray {{BasicKind}} (to be completed)
	StrokeDashArray_Data sql.NullString

	// Declation for basic field pathDB.Transform {{BasicKind}} (to be completed)
	Transform_Data sql.NullString
	// encoding of pointers
	PathPointersEnconding
}

// PathDBs arrays pathDBs
// swagger:response pathDBsResponse
type PathDBs []PathDB

// PathDBResponse provides response
// swagger:response pathDBResponse
type PathDBResponse struct {
	PathDB
}

// PathWOP is a Path without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type PathWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Definition string `xlsx:"2"`

	Color string `xlsx:"3"`

	FillOpacity float64 `xlsx:"4"`

	Stroke string `xlsx:"5"`

	StrokeWidth float64 `xlsx:"6"`

	StrokeDashArray string `xlsx:"7"`

	Transform string `xlsx:"8"`
	// insertion for WOP pointer fields
}

var Path_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Definition",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeWidth",
	"StrokeDashArray",
	"Transform",
}

type BackRepoPathStruct struct {
	// stores PathDB according to their gorm ID
	Map_PathDBID_PathDB *map[uint]*PathDB

	// stores PathDB ID according to Path address
	Map_PathPtr_PathDBID *map[*models.Path]uint

	// stores Path according to their gorm ID
	Map_PathDBID_PathPtr *map[uint]*models.Path

	db *gorm.DB
}

func (backRepoPath *BackRepoPathStruct) GetDB() *gorm.DB {
	return backRepoPath.db
}

// GetPathDBFromPathPtr is a handy function to access the back repo instance from the stage instance
func (backRepoPath *BackRepoPathStruct) GetPathDBFromPathPtr(path *models.Path) (pathDB *PathDB) {
	id := (*backRepoPath.Map_PathPtr_PathDBID)[path]
	pathDB = (*backRepoPath.Map_PathDBID_PathDB)[id]
	return
}

// BackRepoPath.Init set up the BackRepo of the Path
func (backRepoPath *BackRepoPathStruct) Init(db *gorm.DB) (Error error) {

	if backRepoPath.Map_PathDBID_PathPtr != nil {
		err := errors.New("In Init, backRepoPath.Map_PathDBID_PathPtr should be nil")
		return err
	}

	if backRepoPath.Map_PathDBID_PathDB != nil {
		err := errors.New("In Init, backRepoPath.Map_PathDBID_PathDB should be nil")
		return err
	}

	if backRepoPath.Map_PathPtr_PathDBID != nil {
		err := errors.New("In Init, backRepoPath.Map_PathPtr_PathDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Path, 0)
	backRepoPath.Map_PathDBID_PathPtr = &tmp

	tmpDB := make(map[uint]*PathDB, 0)
	backRepoPath.Map_PathDBID_PathDB = &tmpDB

	tmpID := make(map[*models.Path]uint, 0)
	backRepoPath.Map_PathPtr_PathDBID = &tmpID

	backRepoPath.db = db
	return
}

// BackRepoPath.CommitPhaseOne commits all staged instances of Path to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPath *BackRepoPathStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for path := range stage.Paths {
		backRepoPath.CommitPhaseOneInstance(path)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, path := range *backRepoPath.Map_PathDBID_PathPtr {
		if _, ok := stage.Paths[path]; !ok {
			backRepoPath.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoPath.CommitDeleteInstance commits deletion of Path to the BackRepo
func (backRepoPath *BackRepoPathStruct) CommitDeleteInstance(id uint) (Error error) {

	path := (*backRepoPath.Map_PathDBID_PathPtr)[id]

	// path is not staged anymore, remove pathDB
	pathDB := (*backRepoPath.Map_PathDBID_PathDB)[id]
	query := backRepoPath.db.Unscoped().Delete(&pathDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoPath.Map_PathPtr_PathDBID), path)
	delete((*backRepoPath.Map_PathDBID_PathPtr), id)
	delete((*backRepoPath.Map_PathDBID_PathDB), id)

	return
}

// BackRepoPath.CommitPhaseOneInstance commits path staged instances of Path to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPath *BackRepoPathStruct) CommitPhaseOneInstance(path *models.Path) (Error error) {

	// check if the path is not commited yet
	if _, ok := (*backRepoPath.Map_PathPtr_PathDBID)[path]; ok {
		return
	}

	// initiate path
	var pathDB PathDB
	pathDB.CopyBasicFieldsFromPath(path)

	query := backRepoPath.db.Create(&pathDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoPath.Map_PathPtr_PathDBID)[path] = pathDB.ID
	(*backRepoPath.Map_PathDBID_PathPtr)[pathDB.ID] = path
	(*backRepoPath.Map_PathDBID_PathDB)[pathDB.ID] = &pathDB

	return
}

// BackRepoPath.CommitPhaseTwo commits all staged instances of Path to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPath *BackRepoPathStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, path := range *backRepoPath.Map_PathDBID_PathPtr {
		backRepoPath.CommitPhaseTwoInstance(backRepo, idx, path)
	}

	return
}

// BackRepoPath.CommitPhaseTwoInstance commits {{structname }} of models.Path to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPath *BackRepoPathStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, path *models.Path) (Error error) {

	// fetch matching pathDB
	if pathDB, ok := (*backRepoPath.Map_PathDBID_PathDB)[idx]; ok {

		pathDB.CopyBasicFieldsFromPath(path)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers path.Animates into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, animateAssocEnd := range path.Animates {

			// get the back repo instance at the association end
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)

			// encode reverse pointer in the association end back repo instance
			animateAssocEnd_DB.Path_AnimatesDBID.Int64 = int64(pathDB.ID)
			animateAssocEnd_DB.Path_AnimatesDBID.Valid = true
			animateAssocEnd_DB.Path_AnimatesDBID_Index.Int64 = int64(idx)
			animateAssocEnd_DB.Path_AnimatesDBID_Index.Valid = true
			if q := backRepoPath.db.Save(animateAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoPath.db.Save(&pathDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Path intance %s", path.Name))
		return err
	}

	return
}

// BackRepoPath.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoPath *BackRepoPathStruct) CheckoutPhaseOne() (Error error) {

	pathDBArray := make([]PathDB, 0)
	query := backRepoPath.db.Find(&pathDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	pathInstancesToBeRemovedFromTheStage := make(map[*models.Path]struct{})
	for key, value := range models.Stage.Paths {
		pathInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, pathDB := range pathDBArray {
		backRepoPath.CheckoutPhaseOneInstance(&pathDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		path, ok := (*backRepoPath.Map_PathDBID_PathPtr)[pathDB.ID]
		if ok {
			delete(pathInstancesToBeRemovedFromTheStage, path)
		}
	}

	// remove from stage and back repo's 3 maps all paths that are not in the checkout
	for path := range pathInstancesToBeRemovedFromTheStage {
		path.Unstage()

		// remove instance from the back repo 3 maps
		pathID := (*backRepoPath.Map_PathPtr_PathDBID)[path]
		delete((*backRepoPath.Map_PathPtr_PathDBID), path)
		delete((*backRepoPath.Map_PathDBID_PathDB), pathID)
		delete((*backRepoPath.Map_PathDBID_PathPtr), pathID)
	}

	return
}

// CheckoutPhaseOneInstance takes a pathDB that has been found in the DB, updates the backRepo and stages the
// models version of the pathDB
func (backRepoPath *BackRepoPathStruct) CheckoutPhaseOneInstance(pathDB *PathDB) (Error error) {

	path, ok := (*backRepoPath.Map_PathDBID_PathPtr)[pathDB.ID]
	if !ok {
		path = new(models.Path)

		(*backRepoPath.Map_PathDBID_PathPtr)[pathDB.ID] = path
		(*backRepoPath.Map_PathPtr_PathDBID)[path] = pathDB.ID

		// append model store with the new element
		path.Name = pathDB.Name_Data.String
		path.Stage()
	}
	pathDB.CopyBasicFieldsToPath(path)

	// preserve pointer to pathDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_PathDBID_PathDB)[pathDB hold variable pointers
	pathDB_Data := *pathDB
	preservedPtrToPath := &pathDB_Data
	(*backRepoPath.Map_PathDBID_PathDB)[pathDB.ID] = preservedPtrToPath

	return
}

// BackRepoPath.CheckoutPhaseTwo Checkouts all staged instances of Path to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPath *BackRepoPathStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, pathDB := range *backRepoPath.Map_PathDBID_PathDB {
		backRepoPath.CheckoutPhaseTwoInstance(backRepo, pathDB)
	}
	return
}

// BackRepoPath.CheckoutPhaseTwoInstance Checkouts staged instances of Path to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPath *BackRepoPathStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, pathDB *PathDB) (Error error) {

	path := (*backRepoPath.Map_PathDBID_PathPtr)[pathDB.ID]
	_ = path // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem path.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	path.Animates = path.Animates[:0]
	// 2. loop all instances in the type in the association end
	for _, animateDB_AssocEnd := range *backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if animateDB_AssocEnd.Path_AnimatesDBID.Int64 == int64(pathDB.ID) {
			// 4. fetch the associated instance in the stage
			animate_AssocEnd := (*backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr)[animateDB_AssocEnd.ID]
			// 5. append it the association slice
			path.Animates = append(path.Animates, animate_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(path.Animates, func(i, j int) bool {
		animateDB_i_ID := (*backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID)[path.Animates[i]]
		animateDB_j_ID := (*backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID)[path.Animates[j]]

		animateDB_i := (*backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB)[animateDB_i_ID]
		animateDB_j := (*backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB)[animateDB_j_ID]

		return animateDB_i.Path_AnimatesDBID_Index.Int64 < animateDB_j.Path_AnimatesDBID_Index.Int64
	})

	return
}

// CommitPath allows commit of a single path (if already staged)
func (backRepo *BackRepoStruct) CommitPath(path *models.Path) {
	backRepo.BackRepoPath.CommitPhaseOneInstance(path)
	if id, ok := (*backRepo.BackRepoPath.Map_PathPtr_PathDBID)[path]; ok {
		backRepo.BackRepoPath.CommitPhaseTwoInstance(backRepo, id, path)
	}
}

// CommitPath allows checkout of a single path (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutPath(path *models.Path) {
	// check if the path is staged
	if _, ok := (*backRepo.BackRepoPath.Map_PathPtr_PathDBID)[path]; ok {

		if id, ok := (*backRepo.BackRepoPath.Map_PathPtr_PathDBID)[path]; ok {
			var pathDB PathDB
			pathDB.ID = id

			if err := backRepo.BackRepoPath.db.First(&pathDB, id).Error; err != nil {
				log.Panicln("CheckoutPath : Problem with getting object with id:", id)
			}
			backRepo.BackRepoPath.CheckoutPhaseOneInstance(&pathDB)
			backRepo.BackRepoPath.CheckoutPhaseTwoInstance(backRepo, &pathDB)
		}
	}
}

// CopyBasicFieldsFromPath
func (pathDB *PathDB) CopyBasicFieldsFromPath(path *models.Path) {
	// insertion point for fields commit

	pathDB.Name_Data.String = path.Name
	pathDB.Name_Data.Valid = true

	pathDB.Definition_Data.String = path.Definition
	pathDB.Definition_Data.Valid = true

	pathDB.Color_Data.String = path.Color
	pathDB.Color_Data.Valid = true

	pathDB.FillOpacity_Data.Float64 = path.FillOpacity
	pathDB.FillOpacity_Data.Valid = true

	pathDB.Stroke_Data.String = path.Stroke
	pathDB.Stroke_Data.Valid = true

	pathDB.StrokeWidth_Data.Float64 = path.StrokeWidth
	pathDB.StrokeWidth_Data.Valid = true

	pathDB.StrokeDashArray_Data.String = path.StrokeDashArray
	pathDB.StrokeDashArray_Data.Valid = true

	pathDB.Transform_Data.String = path.Transform
	pathDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromPathWOP
func (pathDB *PathDB) CopyBasicFieldsFromPathWOP(path *PathWOP) {
	// insertion point for fields commit

	pathDB.Name_Data.String = path.Name
	pathDB.Name_Data.Valid = true

	pathDB.Definition_Data.String = path.Definition
	pathDB.Definition_Data.Valid = true

	pathDB.Color_Data.String = path.Color
	pathDB.Color_Data.Valid = true

	pathDB.FillOpacity_Data.Float64 = path.FillOpacity
	pathDB.FillOpacity_Data.Valid = true

	pathDB.Stroke_Data.String = path.Stroke
	pathDB.Stroke_Data.Valid = true

	pathDB.StrokeWidth_Data.Float64 = path.StrokeWidth
	pathDB.StrokeWidth_Data.Valid = true

	pathDB.StrokeDashArray_Data.String = path.StrokeDashArray
	pathDB.StrokeDashArray_Data.Valid = true

	pathDB.Transform_Data.String = path.Transform
	pathDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToPath
func (pathDB *PathDB) CopyBasicFieldsToPath(path *models.Path) {
	// insertion point for checkout of basic fields (back repo to stage)
	path.Name = pathDB.Name_Data.String
	path.Definition = pathDB.Definition_Data.String
	path.Color = pathDB.Color_Data.String
	path.FillOpacity = pathDB.FillOpacity_Data.Float64
	path.Stroke = pathDB.Stroke_Data.String
	path.StrokeWidth = pathDB.StrokeWidth_Data.Float64
	path.StrokeDashArray = pathDB.StrokeDashArray_Data.String
	path.Transform = pathDB.Transform_Data.String
}

// CopyBasicFieldsToPathWOP
func (pathDB *PathDB) CopyBasicFieldsToPathWOP(path *PathWOP) {
	path.ID = int(pathDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	path.Name = pathDB.Name_Data.String
	path.Definition = pathDB.Definition_Data.String
	path.Color = pathDB.Color_Data.String
	path.FillOpacity = pathDB.FillOpacity_Data.Float64
	path.Stroke = pathDB.Stroke_Data.String
	path.StrokeWidth = pathDB.StrokeWidth_Data.Float64
	path.StrokeDashArray = pathDB.StrokeDashArray_Data.String
	path.Transform = pathDB.Transform_Data.String
}

// Backup generates a json file from a slice of all PathDB instances in the backrepo
func (backRepoPath *BackRepoPathStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "PathDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PathDB, 0)
	for _, pathDB := range *backRepoPath.Map_PathDBID_PathDB {
		forBackup = append(forBackup, pathDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Path ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Path file", err.Error())
	}
}

// Backup generates a json file from a slice of all PathDB instances in the backrepo
func (backRepoPath *BackRepoPathStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PathDB, 0)
	for _, pathDB := range *backRepoPath.Map_PathDBID_PathDB {
		forBackup = append(forBackup, pathDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Path")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Path_Fields, -1)
	for _, pathDB := range forBackup {

		var pathWOP PathWOP
		pathDB.CopyBasicFieldsToPathWOP(&pathWOP)

		row := sh.AddRow()
		row.WriteStruct(&pathWOP, -1)
	}
}

// RestoreXL from the "Path" sheet all PathDB instances
func (backRepoPath *BackRepoPathStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoPathid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Path"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoPath.rowVisitorPath)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoPath *BackRepoPathStruct) rowVisitorPath(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var pathWOP PathWOP
		row.ReadStruct(&pathWOP)

		// add the unmarshalled struct to the stage
		pathDB := new(PathDB)
		pathDB.CopyBasicFieldsFromPathWOP(&pathWOP)

		pathDB_ID_atBackupTime := pathDB.ID
		pathDB.ID = 0
		query := backRepoPath.db.Create(pathDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoPath.Map_PathDBID_PathDB)[pathDB.ID] = pathDB
		BackRepoPathid_atBckpTime_newID[pathDB_ID_atBackupTime] = pathDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "PathDB.json" in dirPath that stores an array
// of PathDB and stores it in the database
// the map BackRepoPathid_atBckpTime_newID is updated accordingly
func (backRepoPath *BackRepoPathStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoPathid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "PathDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Path file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*PathDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_PathDBID_PathDB
	for _, pathDB := range forRestore {

		pathDB_ID_atBackupTime := pathDB.ID
		pathDB.ID = 0
		query := backRepoPath.db.Create(pathDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoPath.Map_PathDBID_PathDB)[pathDB.ID] = pathDB
		BackRepoPathid_atBckpTime_newID[pathDB_ID_atBackupTime] = pathDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Path file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Path>id_atBckpTime_newID
// to compute new index
func (backRepoPath *BackRepoPathStruct) RestorePhaseTwo() {

	for _, pathDB := range *backRepoPath.Map_PathDBID_PathDB {

		// next line of code is to avert unused variable compilation error
		_ = pathDB

		// insertion point for reindexing pointers encoding
		// This reindex path.Paths
		if pathDB.SVG_PathsDBID.Int64 != 0 {
			pathDB.SVG_PathsDBID.Int64 =
				int64(BackRepoSVGid_atBckpTime_newID[uint(pathDB.SVG_PathsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoPath.db.Model(pathDB).Updates(*pathDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoPathid_atBckpTime_newID map[uint]uint
