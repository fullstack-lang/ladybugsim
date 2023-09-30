// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/ladybugsim/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Ladybug":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Ladybug Form",
			OnSave: __gong__New__LadybugFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		ladybug := new(models.Ladybug)
		FillUpForm(ladybug, formGroup, playground)
	case "LadybugSimulation":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LadybugSimulation Form",
			OnSave: __gong__New__LadybugSimulationFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		ladybugsimulation := new(models.LadybugSimulation)
		FillUpForm(ladybugsimulation, formGroup, playground)
	}
	formStage.Commit()
}
