// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/ladybugsim/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | LadybugDB | LadybugSimulationDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Ladybug:
		ladybugInstance := any(concreteInstance).(*models.Ladybug)
		ret2 := backRepo.BackRepoLadybug.GetLadybugDBFromLadybugPtr(ladybugInstance)
		ret = any(ret2).(*T2)
	case *models.LadybugSimulation:
		ladybugsimulationInstance := any(concreteInstance).(*models.LadybugSimulation)
		ret2 := backRepo.BackRepoLadybugSimulation.GetLadybugSimulationDBFromLadybugSimulationPtr(ladybugsimulationInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Ladybug:
		tmp := GetInstanceDBFromInstance[models.Ladybug, LadybugDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LadybugSimulation:
		tmp := GetInstanceDBFromInstance[models.LadybugSimulation, LadybugSimulationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Ladybug:
		tmp := GetInstanceDBFromInstance[models.Ladybug, LadybugDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.LadybugSimulation:
		tmp := GetInstanceDBFromInstance[models.LadybugSimulation, LadybugSimulationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
