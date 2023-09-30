// generated code - do not edit
package models

import (
	"errors"
	"fmt"
	"time"
)

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Ladybugs           map[*Ladybug]any
	Ladybugs_mapString map[string]*Ladybug

	OnAfterLadybugCreateCallback OnAfterCreateInterface[Ladybug]
	OnAfterLadybugUpdateCallback OnAfterUpdateInterface[Ladybug]
	OnAfterLadybugDeleteCallback OnAfterDeleteInterface[Ladybug]
	OnAfterLadybugReadCallback   OnAfterReadInterface[Ladybug]

	LadybugSimulations           map[*LadybugSimulation]any
	LadybugSimulations_mapString map[string]*LadybugSimulation

	OnAfterLadybugSimulationCreateCallback OnAfterCreateInterface[LadybugSimulation]
	OnAfterLadybugSimulationUpdateCallback OnAfterUpdateInterface[LadybugSimulation]
	OnAfterLadybugSimulationDeleteCallback OnAfterDeleteInterface[LadybugSimulation]
	OnAfterLadybugSimulationReadCallback   OnAfterReadInterface[LadybugSimulation]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitLadybug(ladybug *Ladybug)
	CheckoutLadybug(ladybug *Ladybug)
	CommitLadybugSimulation(ladybugsimulation *LadybugSimulation)
	CheckoutLadybugSimulation(ladybugsimulation *LadybugSimulation)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Ladybugs:           make(map[*Ladybug]any),
		Ladybugs_mapString: make(map[string]*Ladybug),

		LadybugSimulations:           make(map[*LadybugSimulation]any),
		LadybugSimulations_mapString: make(map[string]*LadybugSimulation),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Ladybug"] = len(stage.Ladybugs)
	stage.Map_GongStructName_InstancesNb["LadybugSimulation"] = len(stage.LadybugSimulations)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Ladybug"] = len(stage.Ladybugs)
	stage.Map_GongStructName_InstancesNb["LadybugSimulation"] = len(stage.LadybugSimulations)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts ladybug to the model stage
func (ladybug *Ladybug) Stage(stage *StageStruct) *Ladybug {
	stage.Ladybugs[ladybug] = __member
	stage.Ladybugs_mapString[ladybug.Name] = ladybug

	return ladybug
}

// Unstage removes ladybug off the model stage
func (ladybug *Ladybug) Unstage(stage *StageStruct) *Ladybug {
	delete(stage.Ladybugs, ladybug)
	delete(stage.Ladybugs_mapString, ladybug.Name)
	return ladybug
}

// UnstageVoid removes ladybug off the model stage
func (ladybug *Ladybug) UnstageVoid(stage *StageStruct) {
	delete(stage.Ladybugs, ladybug)
	delete(stage.Ladybugs_mapString, ladybug.Name)
}

// commit ladybug to the back repo (if it is already staged)
func (ladybug *Ladybug) Commit(stage *StageStruct) *Ladybug {
	if _, ok := stage.Ladybugs[ladybug]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLadybug(ladybug)
		}
	}
	return ladybug
}

func (ladybug *Ladybug) CommitVoid(stage *StageStruct) {
	ladybug.Commit(stage)
}

// Checkout ladybug to the back repo (if it is already staged)
func (ladybug *Ladybug) Checkout(stage *StageStruct) *Ladybug {
	if _, ok := stage.Ladybugs[ladybug]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLadybug(ladybug)
		}
	}
	return ladybug
}

// for satisfaction of GongStruct interface
func (ladybug *Ladybug) GetName() (res string) {
	return ladybug.Name
}

// Stage puts ladybugsimulation to the model stage
func (ladybugsimulation *LadybugSimulation) Stage(stage *StageStruct) *LadybugSimulation {
	stage.LadybugSimulations[ladybugsimulation] = __member
	stage.LadybugSimulations_mapString[ladybugsimulation.Name] = ladybugsimulation

	return ladybugsimulation
}

// Unstage removes ladybugsimulation off the model stage
func (ladybugsimulation *LadybugSimulation) Unstage(stage *StageStruct) *LadybugSimulation {
	delete(stage.LadybugSimulations, ladybugsimulation)
	delete(stage.LadybugSimulations_mapString, ladybugsimulation.Name)
	return ladybugsimulation
}

// UnstageVoid removes ladybugsimulation off the model stage
func (ladybugsimulation *LadybugSimulation) UnstageVoid(stage *StageStruct) {
	delete(stage.LadybugSimulations, ladybugsimulation)
	delete(stage.LadybugSimulations_mapString, ladybugsimulation.Name)
}

// commit ladybugsimulation to the back repo (if it is already staged)
func (ladybugsimulation *LadybugSimulation) Commit(stage *StageStruct) *LadybugSimulation {
	if _, ok := stage.LadybugSimulations[ladybugsimulation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLadybugSimulation(ladybugsimulation)
		}
	}
	return ladybugsimulation
}

func (ladybugsimulation *LadybugSimulation) CommitVoid(stage *StageStruct) {
	ladybugsimulation.Commit(stage)
}

// Checkout ladybugsimulation to the back repo (if it is already staged)
func (ladybugsimulation *LadybugSimulation) Checkout(stage *StageStruct) *LadybugSimulation {
	if _, ok := stage.LadybugSimulations[ladybugsimulation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLadybugSimulation(ladybugsimulation)
		}
	}
	return ladybugsimulation
}

// for satisfaction of GongStruct interface
func (ladybugsimulation *LadybugSimulation) GetName() (res string) {
	return ladybugsimulation.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMLadybug(Ladybug *Ladybug)
	CreateORMLadybugSimulation(LadybugSimulation *LadybugSimulation)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMLadybug(Ladybug *Ladybug)
	DeleteORMLadybugSimulation(LadybugSimulation *LadybugSimulation)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Ladybugs = make(map[*Ladybug]any)
	stage.Ladybugs_mapString = make(map[string]*Ladybug)

	stage.LadybugSimulations = make(map[*LadybugSimulation]any)
	stage.LadybugSimulations_mapString = make(map[string]*LadybugSimulation)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Ladybugs = nil
	stage.Ladybugs_mapString = nil

	stage.LadybugSimulations = nil
	stage.LadybugSimulations_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for ladybug := range stage.Ladybugs {
		ladybug.Unstage(stage)
	}

	for ladybugsimulation := range stage.LadybugSimulations {
		ladybugsimulation.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Ladybug | LadybugSimulation
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Ladybug | *LadybugSimulation
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Ladybug]any |
		map[*LadybugSimulation]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Ladybug |
		map[string]*LadybugSimulation |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Ladybug]any:
		return any(&stage.Ladybugs).(*Type)
	case map[*LadybugSimulation]any:
		return any(&stage.LadybugSimulations).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Ladybug:
		return any(&stage.Ladybugs_mapString).(*Type)
	case map[string]*LadybugSimulation:
		return any(&stage.LadybugSimulations_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Ladybug:
		return any(&stage.Ladybugs).(*map[*Type]any)
	case LadybugSimulation:
		return any(&stage.LadybugSimulations).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Ladybug:
		return any(&stage.Ladybugs).(*map[Type]any)
	case *LadybugSimulation:
		return any(&stage.LadybugSimulations).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Ladybug:
		return any(&stage.Ladybugs_mapString).(*map[string]*Type)
	case LadybugSimulation:
		return any(&stage.LadybugSimulations_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Ladybug:
		return any(&Ladybug{
			// Initialisation of associations
		}).(*Type)
	case LadybugSimulation:
		return any(&LadybugSimulation{
			// Initialisation of associations
			// field is initialized with an instance of Ladybug with the name of the field
			Ladybugs: []*Ladybug{{Name: "Ladybugs"}},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Ladybug
	case Ladybug:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LadybugSimulation
	case LadybugSimulation:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Ladybug
	case Ladybug:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of LadybugSimulation
	case LadybugSimulation:
		switch fieldname {
		// insertion point for per direct association field
		case "Ladybugs":
			res := make(map[*Ladybug]*LadybugSimulation)
			for ladybugsimulation := range stage.LadybugSimulations {
				for _, ladybug_ := range ladybugsimulation.Ladybugs {
					res[ladybug_] = ladybugsimulation
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Ladybug:
		res = "Ladybug"
	case LadybugSimulation:
		res = "LadybugSimulation"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Ladybug:
		res = "Ladybug"
	case *LadybugSimulation:
		res = "LadybugSimulation"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Ladybug:
		res = []string{"Name", "Id", "Position", "Speed", "LadybugStatus"}
	case LadybugSimulation:
		res = []string{"Name", "EventNb", "NbOfCollision", "LadybugRadius", "AbsoluteSpeed", "SimulationStep", "MaxDistanceInOneStep", "NbLadybugs", "NbLadybugsOnTheGround", "LeftRelayInitialPosX", "RightRelayInitialPosX", "Ladybugs"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Ladybug:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "LadybugSimulation"
		rf.Fieldname = "Ladybugs"
		res = append(res, rf)
	case LadybugSimulation:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Ladybug:
		res = []string{"Name", "Id", "Position", "Speed", "LadybugStatus"}
	case *LadybugSimulation:
		res = []string{"Name", "EventNb", "NbOfCollision", "LadybugRadius", "AbsoluteSpeed", "SimulationStep", "MaxDistanceInOneStep", "NbLadybugs", "NbLadybugsOnTheGround", "LeftRelayInitialPosX", "RightRelayInitialPosX", "Ladybugs"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Ladybug:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Id":
			res = fmt.Sprintf("%d", inferedInstance.Id)
		case "Position":
			res = fmt.Sprintf("%f", inferedInstance.Position)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "LadybugStatus":
			enum := inferedInstance.LadybugStatus
			res = enum.ToCodeString()
		}
	case *LadybugSimulation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EventNb":
			res = fmt.Sprintf("%d", inferedInstance.EventNb)
		case "NbOfCollision":
			res = fmt.Sprintf("%d", inferedInstance.NbOfCollision)
		case "LadybugRadius":
			res = fmt.Sprintf("%f", inferedInstance.LadybugRadius)
		case "AbsoluteSpeed":
			res = fmt.Sprintf("%f", inferedInstance.AbsoluteSpeed)
		case "SimulationStep":
			res = fmt.Sprintf("%d", inferedInstance.SimulationStep)
		case "MaxDistanceInOneStep":
			res = fmt.Sprintf("%f", inferedInstance.MaxDistanceInOneStep)
		case "NbLadybugs":
			res = fmt.Sprintf("%d", inferedInstance.NbLadybugs)
		case "NbLadybugsOnTheGround":
			res = fmt.Sprintf("%d", inferedInstance.NbLadybugsOnTheGround)
		case "LeftRelayInitialPosX":
			res = fmt.Sprintf("%f", inferedInstance.LeftRelayInitialPosX)
		case "RightRelayInitialPosX":
			res = fmt.Sprintf("%f", inferedInstance.RightRelayInitialPosX)
		case "Ladybugs":
			for idx, __instance__ := range inferedInstance.Ladybugs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Ladybug:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Id":
			res = fmt.Sprintf("%d", inferedInstance.Id)
		case "Position":
			res = fmt.Sprintf("%f", inferedInstance.Position)
		case "Speed":
			res = fmt.Sprintf("%f", inferedInstance.Speed)
		case "LadybugStatus":
			enum := inferedInstance.LadybugStatus
			res = enum.ToCodeString()
		}
	case LadybugSimulation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EventNb":
			res = fmt.Sprintf("%d", inferedInstance.EventNb)
		case "NbOfCollision":
			res = fmt.Sprintf("%d", inferedInstance.NbOfCollision)
		case "LadybugRadius":
			res = fmt.Sprintf("%f", inferedInstance.LadybugRadius)
		case "AbsoluteSpeed":
			res = fmt.Sprintf("%f", inferedInstance.AbsoluteSpeed)
		case "SimulationStep":
			res = fmt.Sprintf("%d", inferedInstance.SimulationStep)
		case "MaxDistanceInOneStep":
			res = fmt.Sprintf("%f", inferedInstance.MaxDistanceInOneStep)
		case "NbLadybugs":
			res = fmt.Sprintf("%d", inferedInstance.NbLadybugs)
		case "NbLadybugsOnTheGround":
			res = fmt.Sprintf("%d", inferedInstance.NbLadybugsOnTheGround)
		case "LeftRelayInitialPosX":
			res = fmt.Sprintf("%f", inferedInstance.LeftRelayInitialPosX)
		case "RightRelayInitialPosX":
			res = fmt.Sprintf("%f", inferedInstance.RightRelayInitialPosX)
		case "Ladybugs":
			for idx, __instance__ := range inferedInstance.Ladybugs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
