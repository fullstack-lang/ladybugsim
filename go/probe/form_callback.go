// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/ladybugsim/go/models"
	"github.com/fullstack-lang/ladybugsim/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__LadybugFormCallback(
	ladybug *models.Ladybug,
	playground *Playground,
) (ladybugFormCallback *LadybugFormCallback) {
	ladybugFormCallback = new(LadybugFormCallback)
	ladybugFormCallback.playground = playground
	ladybugFormCallback.ladybug = ladybug

	ladybugFormCallback.CreationMode = (ladybug == nil)

	return
}

type LadybugFormCallback struct {
	ladybug *models.Ladybug

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (ladybugFormCallback *LadybugFormCallback) OnSave() {

	log.Println("LadybugFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ladybugFormCallback.playground.formStage.Checkout()

	if ladybugFormCallback.ladybug == nil {
		ladybugFormCallback.ladybug = new(models.Ladybug).Stage(ladybugFormCallback.playground.stageOfInterest)
	}
	ladybug_ := ladybugFormCallback.ladybug
	_ = ladybug_

	// get the formGroup
	formGroup := ladybugFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ladybug_.Name), formDiv)
		case "Id":
			FormDivBasicFieldToField(&(ladybug_.Id), formDiv)
		case "Position":
			FormDivBasicFieldToField(&(ladybug_.Position), formDiv)
		case "Speed":
			FormDivBasicFieldToField(&(ladybug_.Speed), formDiv)
		case "LadybugStatus":
			FormDivEnumStringFieldToField(&(ladybug_.LadybugStatus), formDiv)
		case "LadybugSimulation:Ladybugs":
			// we need to retrieve the field owner before the change
			var pastLadybugSimulationOwner *models.LadybugSimulation
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LadybugSimulation"
			rf.Fieldname = "Ladybugs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				ladybugFormCallback.playground.stageOfInterest,
				ladybugFormCallback.playground.backRepoOfInterest,
				ladybug_,
				&rf)

			if reverseFieldOwner != nil {
				pastLadybugSimulationOwner = reverseFieldOwner.(*models.LadybugSimulation)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLadybugSimulationOwner != nil {
					idx := slices.Index(pastLadybugSimulationOwner.Ladybugs, ladybug_)
					pastLadybugSimulationOwner.Ladybugs = slices.Delete(pastLadybugSimulationOwner.Ladybugs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _ladybugsimulation := range *models.GetGongstructInstancesSet[models.LadybugSimulation](ladybugFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _ladybugsimulation.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLadybugSimulationOwner := _ladybugsimulation // we have a match
						if pastLadybugSimulationOwner != nil {
							if newLadybugSimulationOwner != pastLadybugSimulationOwner {
								idx := slices.Index(pastLadybugSimulationOwner.Ladybugs, ladybug_)
								pastLadybugSimulationOwner.Ladybugs = slices.Delete(pastLadybugSimulationOwner.Ladybugs, idx, idx+1)
								newLadybugSimulationOwner.Ladybugs = append(newLadybugSimulationOwner.Ladybugs, ladybug_)
							}
						} else {
							newLadybugSimulationOwner.Ladybugs = append(newLadybugSimulationOwner.Ladybugs, ladybug_)
						}
					}
				}
			}
		}
	}

	ladybugFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Ladybug](
		ladybugFormCallback.playground,
	)
	ladybugFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if ladybugFormCallback.CreationMode {
		ladybugFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LadybugFormCallback(
				nil,
				ladybugFormCallback.playground,
			),
		}).Stage(ladybugFormCallback.playground.formStage)
		ladybug := new(models.Ladybug)
		FillUpForm(ladybug, newFormGroup, ladybugFormCallback.playground)
		ladybugFormCallback.playground.formStage.Commit()
	}

	fillUpTree(ladybugFormCallback.playground)
}
func __gong__New__LadybugSimulationFormCallback(
	ladybugsimulation *models.LadybugSimulation,
	playground *Playground,
) (ladybugsimulationFormCallback *LadybugSimulationFormCallback) {
	ladybugsimulationFormCallback = new(LadybugSimulationFormCallback)
	ladybugsimulationFormCallback.playground = playground
	ladybugsimulationFormCallback.ladybugsimulation = ladybugsimulation

	ladybugsimulationFormCallback.CreationMode = (ladybugsimulation == nil)

	return
}

type LadybugSimulationFormCallback struct {
	ladybugsimulation *models.LadybugSimulation

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (ladybugsimulationFormCallback *LadybugSimulationFormCallback) OnSave() {

	log.Println("LadybugSimulationFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ladybugsimulationFormCallback.playground.formStage.Checkout()

	if ladybugsimulationFormCallback.ladybugsimulation == nil {
		ladybugsimulationFormCallback.ladybugsimulation = new(models.LadybugSimulation).Stage(ladybugsimulationFormCallback.playground.stageOfInterest)
	}
	ladybugsimulation_ := ladybugsimulationFormCallback.ladybugsimulation
	_ = ladybugsimulation_

	// get the formGroup
	formGroup := ladybugsimulationFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ladybugsimulation_.Name), formDiv)
		case "EventNb":
			FormDivBasicFieldToField(&(ladybugsimulation_.EventNb), formDiv)
		case "NbOfCollision":
			FormDivBasicFieldToField(&(ladybugsimulation_.NbOfCollision), formDiv)
		case "LadybugRadius":
			FormDivBasicFieldToField(&(ladybugsimulation_.LadybugRadius), formDiv)
		case "AbsoluteSpeed":
			FormDivBasicFieldToField(&(ladybugsimulation_.AbsoluteSpeed), formDiv)
		case "SimulationStep":
			FormDivBasicFieldToField(&(ladybugsimulation_.SimulationStep), formDiv)
		case "MaxDistanceInOneStep":
			FormDivBasicFieldToField(&(ladybugsimulation_.MaxDistanceInOneStep), formDiv)
		case "NbLadybugs":
			FormDivBasicFieldToField(&(ladybugsimulation_.NbLadybugs), formDiv)
		case "NbLadybugsOnTheGround":
			FormDivBasicFieldToField(&(ladybugsimulation_.NbLadybugsOnTheGround), formDiv)
		case "LeftRelayInitialPosX":
			FormDivBasicFieldToField(&(ladybugsimulation_.LeftRelayInitialPosX), formDiv)
		case "RightRelayInitialPosX":
			FormDivBasicFieldToField(&(ladybugsimulation_.RightRelayInitialPosX), formDiv)
		}
	}

	ladybugsimulationFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LadybugSimulation](
		ladybugsimulationFormCallback.playground,
	)
	ladybugsimulationFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if ladybugsimulationFormCallback.CreationMode {
		ladybugsimulationFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LadybugSimulationFormCallback(
				nil,
				ladybugsimulationFormCallback.playground,
			),
		}).Stage(ladybugsimulationFormCallback.playground.formStage)
		ladybugsimulation := new(models.LadybugSimulation)
		FillUpForm(ladybugsimulation, newFormGroup, ladybugsimulationFormCallback.playground)
		ladybugsimulationFormCallback.playground.formStage.Commit()
	}

	fillUpTree(ladybugsimulationFormCallback.playground)
}
