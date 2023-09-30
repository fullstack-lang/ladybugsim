// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/ladybugsim/go/models"
	"github.com/fullstack-lang/ladybugsim/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Ladybug:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Id", instanceWithInferedType.Id, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Position", instanceWithInferedType.Position, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Speed", instanceWithInferedType.Speed, instanceWithInferedType, playground.formStage, formGroup, false)
		EnumTypeStringToForm("LadybugStatus", instanceWithInferedType.LadybugStatus, instanceWithInferedType, playground.formStage, formGroup)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LadybugSimulation"
			rf.Fieldname = "Ladybugs"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.LadybugSimulation),
					"Ladybugs",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.LadybugSimulation, *models.Ladybug](
					nil,
					"Ladybugs",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.LadybugSimulation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("EventNb", instanceWithInferedType.EventNb, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbOfCollision", instanceWithInferedType.NbOfCollision, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("LadybugRadius", instanceWithInferedType.LadybugRadius, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AbsoluteSpeed", instanceWithInferedType.AbsoluteSpeed, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("SimulationStep", instanceWithInferedType.SimulationStep, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("MaxDistanceInOneStep", instanceWithInferedType.MaxDistanceInOneStep, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbLadybugs", instanceWithInferedType.NbLadybugs, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("NbLadybugsOnTheGround", instanceWithInferedType.NbLadybugsOnTheGround, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("LeftRelayInitialPosX", instanceWithInferedType.LeftRelayInitialPosX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("RightRelayInitialPosX", instanceWithInferedType.RightRelayInitialPosX, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Ladybugs", instanceWithInferedType, &instanceWithInferedType.Ladybugs, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
