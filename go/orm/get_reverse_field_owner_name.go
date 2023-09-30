// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/ladybugsim/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Ladybug:
		tmp := GetInstanceDBFromInstance[models.Ladybug, LadybugDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Ladybug":
			switch reverseField.Fieldname {
			}
		case "LadybugSimulation":
			switch reverseField.Fieldname {
			case "Ladybugs":
				if tmp != nil && tmp.LadybugSimulation_LadybugsDBID.Int64 != 0 {
					id := uint(tmp.LadybugSimulation_LadybugsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.LadybugSimulation:
		tmp := GetInstanceDBFromInstance[models.LadybugSimulation, LadybugSimulationDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Ladybug":
			switch reverseField.Fieldname {
			}
		case "LadybugSimulation":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Ladybug:
		tmp := GetInstanceDBFromInstance[models.Ladybug, LadybugDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Ladybug":
			switch reverseField.Fieldname {
			}
		case "LadybugSimulation":
			switch reverseField.Fieldname {
			case "Ladybugs":
				if tmp != nil && tmp.LadybugSimulation_LadybugsDBID.Int64 != 0 {
					id := uint(tmp.LadybugSimulation_LadybugsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoLadybugSimulation.Map_LadybugSimulationDBID_LadybugSimulationPtr[id]
					res = reservePointerTarget
				}
			}
		}
	
	case *models.LadybugSimulation:
		tmp := GetInstanceDBFromInstance[models.LadybugSimulation, LadybugSimulationDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Ladybug":
			switch reverseField.Fieldname {
			}
		case "LadybugSimulation":
			switch reverseField.Fieldname {
			}
		}
	
	default:
		_ = inst
	}
	return res
}
