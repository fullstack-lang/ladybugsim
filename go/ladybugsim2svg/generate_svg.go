package ladybugsim2svg

import (
	ladybugsim_models "github.com/fullstack-lang/ladybugsim/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// type of the singloton for interception of ladybugsim commit in order to generate
// the svg
type ladybugsimToSVGTranformer struct {
}

var LadybugsimToSVGTranformerSingloton ladybugsimToSVGTranformer

// callback on the commit function
func (ladybugsimToSVGTranformer *ladybugsimToSVGTranformer) BeforeCommit(stage *ladybugsim_models.StageStruct) {

	// remove all gongsvg stage/repo
	gongsvg_models.Stage.Checkout()
	gongsvg_models.Stage.Reset()
	gongsvg_models.Stage.Commit()
	ladybugsim_models.Stage.Checkout()

	//
	// SVG
	//
	svg := new(gongsvg_models.SVG).Stage()
	svg.Name = "New ladybugsim Chart"
	svg.Display = true

	fence := new(gongsvg_models.Rect).Stage()
	fence.Name = "fence"
	svg.Rects = append(svg.Rects, fence)

	fence.X = 50
	fence.Y = 50
	fence.Height = 30
	fence.Width = 2000

	fence.Color = "blue"
	fence.FillOpacity = 0.1
	fence.Stroke = "blue"
	fence.StrokeWidth = 0.5

	var ladybugSimulation *ladybugsim_models.LadybugSimulation
	for _ladybugsim := range stage.LadybugSimulations {
		ladybugSimulation = _ladybugsim
	}

	for _, ladybug := range ladybugSimulation.Ladybugs {
		circle := new(gongsvg_models.Circle).Stage()
		circle.Name = ladybug.Name
		svg.Circles = append(svg.Circles, circle)
		circle.Color = "red"
		circle.FillOpacity = 0.1
		circle.Stroke = "red"
		circle.StrokeWidth = 0.5
		circle.CX = fence.X + ladybug.Position*fence.Width
		circle.CY = fence.Y
		circle.Radius = 5

		if ladybug.LadybugStatus == ladybugsim_models.ON_THE_GROUND {
			circle.CY = fence.Y + fence.Height
		}

	}

	gongsvg_models.Stage.Commit()

	// log.Printf("Before Commit")
}
