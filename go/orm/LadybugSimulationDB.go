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
var dummy_LadybugSimulation_sql sql.NullBool
var dummy_LadybugSimulation_time time.Duration
var dummy_LadybugSimulation_sort sort.Float64Slice

// LadybugSimulationAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model ladybugsimulationAPI
type LadybugSimulationAPI struct {
	gorm.Model

	models.LadybugSimulation

	// encoding of pointers
	LadybugSimulationPointersEnconding
}

// LadybugSimulationPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LadybugSimulationPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// LadybugSimulationDB describes a ladybugsimulation in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model ladybugsimulationDB
type LadybugSimulationDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field ladybugsimulationDB.Name
	Name_Data sql.NullString

	// Declation for basic field ladybugsimulationDB.EventNb
	EventNb_Data sql.NullInt64

	// Declation for basic field ladybugsimulationDB.NbOfCollision
	NbOfCollision_Data sql.NullInt64

	// Declation for basic field ladybugsimulationDB.LadybugRadius
	LadybugRadius_Data sql.NullFloat64

	// Declation for basic field ladybugsimulationDB.AbsoluteSpeed
	AbsoluteSpeed_Data sql.NullFloat64

	// Declation for basic field ladybugsimulationDB.SimulationStep
	SimulationStep_Data sql.NullInt64

	// Declation for basic field ladybugsimulationDB.MaxDistanceInOneStep
	MaxDistanceInOneStep_Data sql.NullFloat64

	// Declation for basic field ladybugsimulationDB.NbLadybugs
	NbLadybugs_Data sql.NullInt64

	// Declation for basic field ladybugsimulationDB.NbLadybugsOnTheGround
	NbLadybugsOnTheGround_Data sql.NullInt64

	// Declation for basic field ladybugsimulationDB.LeftRelayInitialPosX
	LeftRelayInitialPosX_Data sql.NullFloat64

	// Declation for basic field ladybugsimulationDB.RightRelayInitialPosX
	RightRelayInitialPosX_Data sql.NullFloat64
	// encoding of pointers
	LadybugSimulationPointersEnconding
}

// LadybugSimulationDBs arrays ladybugsimulationDBs
// swagger:response ladybugsimulationDBsResponse
type LadybugSimulationDBs []LadybugSimulationDB

// LadybugSimulationDBResponse provides response
// swagger:response ladybugsimulationDBResponse
type LadybugSimulationDBResponse struct {
	LadybugSimulationDB
}

// LadybugSimulationWOP is a LadybugSimulation without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LadybugSimulationWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	EventNb int `xlsx:"2"`

	NbOfCollision int `xlsx:"3"`

	LadybugRadius float64 `xlsx:"4"`

	AbsoluteSpeed float64 `xlsx:"5"`

	SimulationStep time.Duration `xlsx:"6"`

	MaxDistanceInOneStep float64 `xlsx:"7"`

	NbLadybugs int `xlsx:"8"`

	NbLadybugsOnTheGround int `xlsx:"9"`

	LeftRelayInitialPosX float64 `xlsx:"10"`

	RightRelayInitialPosX float64 `xlsx:"11"`
	// insertion for WOP pointer fields
}

var LadybugSimulation_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"EventNb",
	"NbOfCollision",
	"LadybugRadius",
	"AbsoluteSpeed",
	"SimulationStep",
	"MaxDistanceInOneStep",
	"NbLadybugs",
	"NbLadybugsOnTheGround",
	"LeftRelayInitialPosX",
	"RightRelayInitialPosX",
}

type BackRepoLadybugSimulationStruct struct {
	// stores LadybugSimulationDB according to their gorm ID
	Map_LadybugSimulationDBID_LadybugSimulationDB map[uint]*LadybugSimulationDB

	// stores LadybugSimulationDB ID according to LadybugSimulation address
	Map_LadybugSimulationPtr_LadybugSimulationDBID map[*models.LadybugSimulation]uint

	// stores LadybugSimulation according to their gorm ID
	Map_LadybugSimulationDBID_LadybugSimulationPtr map[uint]*models.LadybugSimulation

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoLadybugSimulation.stage
	return
}

func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) GetDB() *gorm.DB {
	return backRepoLadybugSimulation.db
}

// GetLadybugSimulationDBFromLadybugSimulationPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) GetLadybugSimulationDBFromLadybugSimulationPtr(ladybugsimulation *models.LadybugSimulation) (ladybugsimulationDB *LadybugSimulationDB) {
	id := backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]
	ladybugsimulationDB = backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[id]
	return
}

// BackRepoLadybugSimulation.CommitPhaseOne commits all staged instances of LadybugSimulation to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for ladybugsimulation := range stage.LadybugSimulations {
		backRepoLadybugSimulation.CommitPhaseOneInstance(ladybugsimulation)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, ladybugsimulation := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr {
		if _, ok := stage.LadybugSimulations[ladybugsimulation]; !ok {
			backRepoLadybugSimulation.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLadybugSimulation.CommitDeleteInstance commits deletion of LadybugSimulation to the BackRepo
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CommitDeleteInstance(id uint) (Error error) {

	ladybugsimulation := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[id]

	// ladybugsimulation is not staged anymore, remove ladybugsimulationDB
	ladybugsimulationDB := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[id]
	query := backRepoLadybugSimulation.db.Unscoped().Delete(&ladybugsimulationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID, ladybugsimulation)
	delete(backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr, id)
	delete(backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB, id)

	return
}

// BackRepoLadybugSimulation.CommitPhaseOneInstance commits ladybugsimulation staged instances of LadybugSimulation to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CommitPhaseOneInstance(ladybugsimulation *models.LadybugSimulation) (Error error) {

	// check if the ladybugsimulation is not commited yet
	if _, ok := backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]; ok {
		return
	}

	// initiate ladybugsimulation
	var ladybugsimulationDB LadybugSimulationDB
	ladybugsimulationDB.CopyBasicFieldsFromLadybugSimulation(ladybugsimulation)

	query := backRepoLadybugSimulation.db.Create(&ladybugsimulationDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation] = ladybugsimulationDB.ID
	backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID] = ladybugsimulation
	backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[ladybugsimulationDB.ID] = &ladybugsimulationDB

	return
}

// BackRepoLadybugSimulation.CommitPhaseTwo commits all staged instances of LadybugSimulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, ladybugsimulation := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr {
		backRepoLadybugSimulation.CommitPhaseTwoInstance(backRepo, idx, ladybugsimulation)
	}

	return
}

// BackRepoLadybugSimulation.CommitPhaseTwoInstance commits {{structname }} of models.LadybugSimulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, ladybugsimulation *models.LadybugSimulation) (Error error) {

	// fetch matching ladybugsimulationDB
	if ladybugsimulationDB, ok := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[idx]; ok {

		ladybugsimulationDB.CopyBasicFieldsFromLadybugSimulation(ladybugsimulation)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers ladybugsimulation.Ladybugs into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, ladybugAssocEnd := range ladybugsimulation.Ladybugs {

			// get the back repo instance at the association end
			ladybugAssocEnd_DB :=
				backRepo.BackRepoLadybug.GetLadybugDBFromLadybugPtr(ladybugAssocEnd)

			// encode reverse pointer in the association end back repo instance
			ladybugAssocEnd_DB.LadybugSimulation_LadybugsDBID.Int64 = int64(ladybugsimulationDB.ID)
			ladybugAssocEnd_DB.LadybugSimulation_LadybugsDBID.Valid = true
			ladybugAssocEnd_DB.LadybugSimulation_LadybugsDBID_Index.Int64 = int64(idx)
			ladybugAssocEnd_DB.LadybugSimulation_LadybugsDBID_Index.Valid = true
			if q := backRepoLadybugSimulation.db.Save(ladybugAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoLadybugSimulation.db.Save(&ladybugsimulationDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown LadybugSimulation intance %s", ladybugsimulation.Name))
		return err
	}

	return
}

// BackRepoLadybugSimulation.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CheckoutPhaseOne() (Error error) {

	ladybugsimulationDBArray := make([]LadybugSimulationDB, 0)
	query := backRepoLadybugSimulation.db.Find(&ladybugsimulationDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	ladybugsimulationInstancesToBeRemovedFromTheStage := make(map[*models.LadybugSimulation]any)
	for key, value := range backRepoLadybugSimulation.stage.LadybugSimulations {
		ladybugsimulationInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, ladybugsimulationDB := range ladybugsimulationDBArray {
		backRepoLadybugSimulation.CheckoutPhaseOneInstance(&ladybugsimulationDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		ladybugsimulation, ok := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]
		if ok {
			delete(ladybugsimulationInstancesToBeRemovedFromTheStage, ladybugsimulation)
		}
	}

	// remove from stage and back repo's 3 maps all ladybugsimulations that are not in the checkout
	for ladybugsimulation := range ladybugsimulationInstancesToBeRemovedFromTheStage {
		ladybugsimulation.Unstage(backRepoLadybugSimulation.GetStage())

		// remove instance from the back repo 3 maps
		ladybugsimulationID := backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]
		delete(backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID, ladybugsimulation)
		delete(backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB, ladybugsimulationID)
		delete(backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr, ladybugsimulationID)
	}

	return
}

// CheckoutPhaseOneInstance takes a ladybugsimulationDB that has been found in the DB, updates the backRepo and stages the
// models version of the ladybugsimulationDB
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CheckoutPhaseOneInstance(ladybugsimulationDB *LadybugSimulationDB) (Error error) {

	ladybugsimulation, ok := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]
	if !ok {
		ladybugsimulation = new(models.LadybugSimulation)

		backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID] = ladybugsimulation
		backRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation] = ladybugsimulationDB.ID

		// append model store with the new element
		ladybugsimulation.Name = ladybugsimulationDB.Name_Data.String
		ladybugsimulation.Stage(backRepoLadybugSimulation.GetStage())
	}
	ladybugsimulationDB.CopyBasicFieldsToLadybugSimulation(ladybugsimulation)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	ladybugsimulation.Stage(backRepoLadybugSimulation.GetStage())

	// preserve pointer to ladybugsimulationDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LadybugSimulationDBID_LadybugSimulationDB)[ladybugsimulationDB hold variable pointers
	ladybugsimulationDB_Data := *ladybugsimulationDB
	preservedPtrToLadybugSimulation := &ladybugsimulationDB_Data
	backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[ladybugsimulationDB.ID] = preservedPtrToLadybugSimulation

	return
}

// BackRepoLadybugSimulation.CheckoutPhaseTwo Checkouts all staged instances of LadybugSimulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, ladybugsimulationDB := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB {
		backRepoLadybugSimulation.CheckoutPhaseTwoInstance(backRepo, ladybugsimulationDB)
	}
	return
}

// BackRepoLadybugSimulation.CheckoutPhaseTwoInstance Checkouts staged instances of LadybugSimulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, ladybugsimulationDB *LadybugSimulationDB) (Error error) {

	ladybugsimulation := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[ladybugsimulationDB.ID]
	_ = ladybugsimulation // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem ladybugsimulation.Ladybugs in the stage from the encode in the back repo
	// It parses all LadybugDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	ladybugsimulation.Ladybugs = ladybugsimulation.Ladybugs[:0]
	// 2. loop all instances in the type in the association end
	for _, ladybugDB_AssocEnd := range backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if ladybugDB_AssocEnd.LadybugSimulation_LadybugsDBID.Int64 == int64(ladybugsimulationDB.ID) {
			// 4. fetch the associated instance in the stage
			ladybug_AssocEnd := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugPtr[ladybugDB_AssocEnd.ID]
			// 5. append it the association slice
			ladybugsimulation.Ladybugs = append(ladybugsimulation.Ladybugs, ladybug_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(ladybugsimulation.Ladybugs, func(i, j int) bool {
		ladybugDB_i_ID := backRepo.BackRepoLadybug.Map_LadybugPtr_LadybugDBID[ladybugsimulation.Ladybugs[i]]
		ladybugDB_j_ID := backRepo.BackRepoLadybug.Map_LadybugPtr_LadybugDBID[ladybugsimulation.Ladybugs[j]]

		ladybugDB_i := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugDB[ladybugDB_i_ID]
		ladybugDB_j := backRepo.BackRepoLadybug.Map_LadybugDBID_LadybugDB[ladybugDB_j_ID]

		return ladybugDB_i.LadybugSimulation_LadybugsDBID_Index.Int64 < ladybugDB_j.LadybugSimulation_LadybugsDBID_Index.Int64
	})

	return
}

// CommitLadybugSimulation allows commit of a single ladybugsimulation (if already staged)
func (backRepo *BackRepoStruct) CommitLadybugSimulation(ladybugsimulation *models.LadybugSimulation) {
	backRepo.BackRepoLadybugSimulation.CommitPhaseOneInstance(ladybugsimulation)
	if id, ok := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]; ok {
		backRepo.BackRepoLadybugSimulation.CommitPhaseTwoInstance(backRepo, id, ladybugsimulation)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLadybugSimulation allows checkout of a single ladybugsimulation (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLadybugSimulation(ladybugsimulation *models.LadybugSimulation) {
	// check if the ladybugsimulation is staged
	if _, ok := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]; ok {

		if id, ok := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationPtr_LadybugSimulationDBID[ladybugsimulation]; ok {
			var ladybugsimulationDB LadybugSimulationDB
			ladybugsimulationDB.ID = id

			if err := backRepo.BackRepoLadybugSimulation.db.First(&ladybugsimulationDB, id).Error; err != nil {
				log.Panicln("CheckoutLadybugSimulation : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLadybugSimulation.CheckoutPhaseOneInstance(&ladybugsimulationDB)
			backRepo.BackRepoLadybugSimulation.CheckoutPhaseTwoInstance(backRepo, &ladybugsimulationDB)
		}
	}
}

// CopyBasicFieldsFromLadybugSimulation
func (ladybugsimulationDB *LadybugSimulationDB) CopyBasicFieldsFromLadybugSimulation(ladybugsimulation *models.LadybugSimulation) {
	// insertion point for fields commit

	ladybugsimulationDB.Name_Data.String = ladybugsimulation.Name
	ladybugsimulationDB.Name_Data.Valid = true

	ladybugsimulationDB.EventNb_Data.Int64 = int64(ladybugsimulation.EventNb)
	ladybugsimulationDB.EventNb_Data.Valid = true

	ladybugsimulationDB.NbOfCollision_Data.Int64 = int64(ladybugsimulation.NbOfCollision)
	ladybugsimulationDB.NbOfCollision_Data.Valid = true

	ladybugsimulationDB.LadybugRadius_Data.Float64 = ladybugsimulation.LadybugRadius
	ladybugsimulationDB.LadybugRadius_Data.Valid = true

	ladybugsimulationDB.AbsoluteSpeed_Data.Float64 = ladybugsimulation.AbsoluteSpeed
	ladybugsimulationDB.AbsoluteSpeed_Data.Valid = true

	ladybugsimulationDB.SimulationStep_Data.Int64 = int64(ladybugsimulation.SimulationStep)
	ladybugsimulationDB.SimulationStep_Data.Valid = true

	ladybugsimulationDB.MaxDistanceInOneStep_Data.Float64 = ladybugsimulation.MaxDistanceInOneStep
	ladybugsimulationDB.MaxDistanceInOneStep_Data.Valid = true

	ladybugsimulationDB.NbLadybugs_Data.Int64 = int64(ladybugsimulation.NbLadybugs)
	ladybugsimulationDB.NbLadybugs_Data.Valid = true

	ladybugsimulationDB.NbLadybugsOnTheGround_Data.Int64 = int64(ladybugsimulation.NbLadybugsOnTheGround)
	ladybugsimulationDB.NbLadybugsOnTheGround_Data.Valid = true

	ladybugsimulationDB.LeftRelayInitialPosX_Data.Float64 = ladybugsimulation.LeftRelayInitialPosX
	ladybugsimulationDB.LeftRelayInitialPosX_Data.Valid = true

	ladybugsimulationDB.RightRelayInitialPosX_Data.Float64 = ladybugsimulation.RightRelayInitialPosX
	ladybugsimulationDB.RightRelayInitialPosX_Data.Valid = true
}

// CopyBasicFieldsFromLadybugSimulationWOP
func (ladybugsimulationDB *LadybugSimulationDB) CopyBasicFieldsFromLadybugSimulationWOP(ladybugsimulation *LadybugSimulationWOP) {
	// insertion point for fields commit

	ladybugsimulationDB.Name_Data.String = ladybugsimulation.Name
	ladybugsimulationDB.Name_Data.Valid = true

	ladybugsimulationDB.EventNb_Data.Int64 = int64(ladybugsimulation.EventNb)
	ladybugsimulationDB.EventNb_Data.Valid = true

	ladybugsimulationDB.NbOfCollision_Data.Int64 = int64(ladybugsimulation.NbOfCollision)
	ladybugsimulationDB.NbOfCollision_Data.Valid = true

	ladybugsimulationDB.LadybugRadius_Data.Float64 = ladybugsimulation.LadybugRadius
	ladybugsimulationDB.LadybugRadius_Data.Valid = true

	ladybugsimulationDB.AbsoluteSpeed_Data.Float64 = ladybugsimulation.AbsoluteSpeed
	ladybugsimulationDB.AbsoluteSpeed_Data.Valid = true

	ladybugsimulationDB.SimulationStep_Data.Int64 = int64(ladybugsimulation.SimulationStep)
	ladybugsimulationDB.SimulationStep_Data.Valid = true

	ladybugsimulationDB.MaxDistanceInOneStep_Data.Float64 = ladybugsimulation.MaxDistanceInOneStep
	ladybugsimulationDB.MaxDistanceInOneStep_Data.Valid = true

	ladybugsimulationDB.NbLadybugs_Data.Int64 = int64(ladybugsimulation.NbLadybugs)
	ladybugsimulationDB.NbLadybugs_Data.Valid = true

	ladybugsimulationDB.NbLadybugsOnTheGround_Data.Int64 = int64(ladybugsimulation.NbLadybugsOnTheGround)
	ladybugsimulationDB.NbLadybugsOnTheGround_Data.Valid = true

	ladybugsimulationDB.LeftRelayInitialPosX_Data.Float64 = ladybugsimulation.LeftRelayInitialPosX
	ladybugsimulationDB.LeftRelayInitialPosX_Data.Valid = true

	ladybugsimulationDB.RightRelayInitialPosX_Data.Float64 = ladybugsimulation.RightRelayInitialPosX
	ladybugsimulationDB.RightRelayInitialPosX_Data.Valid = true
}

// CopyBasicFieldsToLadybugSimulation
func (ladybugsimulationDB *LadybugSimulationDB) CopyBasicFieldsToLadybugSimulation(ladybugsimulation *models.LadybugSimulation) {
	// insertion point for checkout of basic fields (back repo to stage)
	ladybugsimulation.Name = ladybugsimulationDB.Name_Data.String
	ladybugsimulation.EventNb = int(ladybugsimulationDB.EventNb_Data.Int64)
	ladybugsimulation.NbOfCollision = int(ladybugsimulationDB.NbOfCollision_Data.Int64)
	ladybugsimulation.LadybugRadius = ladybugsimulationDB.LadybugRadius_Data.Float64
	ladybugsimulation.AbsoluteSpeed = ladybugsimulationDB.AbsoluteSpeed_Data.Float64
	ladybugsimulation.SimulationStep = time.Duration(ladybugsimulationDB.SimulationStep_Data.Int64)
	ladybugsimulation.MaxDistanceInOneStep = ladybugsimulationDB.MaxDistanceInOneStep_Data.Float64
	ladybugsimulation.NbLadybugs = int(ladybugsimulationDB.NbLadybugs_Data.Int64)
	ladybugsimulation.NbLadybugsOnTheGround = int(ladybugsimulationDB.NbLadybugsOnTheGround_Data.Int64)
	ladybugsimulation.LeftRelayInitialPosX = ladybugsimulationDB.LeftRelayInitialPosX_Data.Float64
	ladybugsimulation.RightRelayInitialPosX = ladybugsimulationDB.RightRelayInitialPosX_Data.Float64
}

// CopyBasicFieldsToLadybugSimulationWOP
func (ladybugsimulationDB *LadybugSimulationDB) CopyBasicFieldsToLadybugSimulationWOP(ladybugsimulation *LadybugSimulationWOP) {
	ladybugsimulation.ID = int(ladybugsimulationDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	ladybugsimulation.Name = ladybugsimulationDB.Name_Data.String
	ladybugsimulation.EventNb = int(ladybugsimulationDB.EventNb_Data.Int64)
	ladybugsimulation.NbOfCollision = int(ladybugsimulationDB.NbOfCollision_Data.Int64)
	ladybugsimulation.LadybugRadius = ladybugsimulationDB.LadybugRadius_Data.Float64
	ladybugsimulation.AbsoluteSpeed = ladybugsimulationDB.AbsoluteSpeed_Data.Float64
	ladybugsimulation.SimulationStep = time.Duration(ladybugsimulationDB.SimulationStep_Data.Int64)
	ladybugsimulation.MaxDistanceInOneStep = ladybugsimulationDB.MaxDistanceInOneStep_Data.Float64
	ladybugsimulation.NbLadybugs = int(ladybugsimulationDB.NbLadybugs_Data.Int64)
	ladybugsimulation.NbLadybugsOnTheGround = int(ladybugsimulationDB.NbLadybugsOnTheGround_Data.Int64)
	ladybugsimulation.LeftRelayInitialPosX = ladybugsimulationDB.LeftRelayInitialPosX_Data.Float64
	ladybugsimulation.RightRelayInitialPosX = ladybugsimulationDB.RightRelayInitialPosX_Data.Float64
}

// Backup generates a json file from a slice of all LadybugSimulationDB instances in the backrepo
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LadybugSimulationDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LadybugSimulationDB, 0)
	for _, ladybugsimulationDB := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB {
		forBackup = append(forBackup, ladybugsimulationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json LadybugSimulation ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json LadybugSimulation file", err.Error())
	}
}

// Backup generates a json file from a slice of all LadybugSimulationDB instances in the backrepo
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LadybugSimulationDB, 0)
	for _, ladybugsimulationDB := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB {
		forBackup = append(forBackup, ladybugsimulationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("LadybugSimulation")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&LadybugSimulation_Fields, -1)
	for _, ladybugsimulationDB := range forBackup {

		var ladybugsimulationWOP LadybugSimulationWOP
		ladybugsimulationDB.CopyBasicFieldsToLadybugSimulationWOP(&ladybugsimulationWOP)

		row := sh.AddRow()
		row.WriteStruct(&ladybugsimulationWOP, -1)
	}
}

// RestoreXL from the "LadybugSimulation" sheet all LadybugSimulationDB instances
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLadybugSimulationid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["LadybugSimulation"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLadybugSimulation.rowVisitorLadybugSimulation)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) rowVisitorLadybugSimulation(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var ladybugsimulationWOP LadybugSimulationWOP
		row.ReadStruct(&ladybugsimulationWOP)

		// add the unmarshalled struct to the stage
		ladybugsimulationDB := new(LadybugSimulationDB)
		ladybugsimulationDB.CopyBasicFieldsFromLadybugSimulationWOP(&ladybugsimulationWOP)

		ladybugsimulationDB_ID_atBackupTime := ladybugsimulationDB.ID
		ladybugsimulationDB.ID = 0
		query := backRepoLadybugSimulation.db.Create(ladybugsimulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[ladybugsimulationDB.ID] = ladybugsimulationDB
		BackRepoLadybugSimulationid_atBckpTime_newID[ladybugsimulationDB_ID_atBackupTime] = ladybugsimulationDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LadybugSimulationDB.json" in dirPath that stores an array
// of LadybugSimulationDB and stores it in the database
// the map BackRepoLadybugSimulationid_atBckpTime_newID is updated accordingly
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLadybugSimulationid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LadybugSimulationDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json LadybugSimulation file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LadybugSimulationDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LadybugSimulationDBID_LadybugSimulationDB
	for _, ladybugsimulationDB := range forRestore {

		ladybugsimulationDB_ID_atBackupTime := ladybugsimulationDB.ID
		ladybugsimulationDB.ID = 0
		query := backRepoLadybugSimulation.db.Create(ladybugsimulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[ladybugsimulationDB.ID] = ladybugsimulationDB
		BackRepoLadybugSimulationid_atBckpTime_newID[ladybugsimulationDB_ID_atBackupTime] = ladybugsimulationDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json LadybugSimulation file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<LadybugSimulation>id_atBckpTime_newID
// to compute new index
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) RestorePhaseTwo() {

	for _, ladybugsimulationDB := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB {

		// next line of code is to avert unused variable compilation error
		_ = ladybugsimulationDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoLadybugSimulation.db.Model(ladybugsimulationDB).Updates(*ladybugsimulationDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoLadybugSimulation.ResetReversePointers commits all staged instances of LadybugSimulation to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, ladybugsimulation := range backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr {
		backRepoLadybugSimulation.ResetReversePointersInstance(backRepo, idx, ladybugsimulation)
	}

	return
}

func (backRepoLadybugSimulation *BackRepoLadybugSimulationStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.LadybugSimulation) (Error error) {

	// fetch matching ladybugsimulationDB
	if ladybugsimulationDB, ok := backRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationDB[idx]; ok {
		_ = ladybugsimulationDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLadybugSimulationid_atBckpTime_newID map[uint]uint
