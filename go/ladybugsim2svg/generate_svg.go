package ladybugsim2svg

import (
	"fmt"

	ladybugsim_models "github.com/fullstack-lang/ladybugsim/go/models"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// type of the singloton for interception of ladybugsim commit in order to generate
// the svg
type ladybugsimToSVGTranformer struct {
}

var LadybugsimToSVGTranformerSingloton ladybugsimToSVGTranformer

var svgSingloton *gongsvg_models.SVG
var fenceSingloton *gongsvg_models.Rect

var mapLadybug_Circle = make(map[*ladybugsim_models.Ladybug]*gongsvg_models.Circle)
var mapLadybug_Animate = make(map[*ladybugsim_models.Ladybug]*gongsvg_models.Animate)

// callback on the commit function
func (ladybugsimToSVGTranformer *ladybugsimToSVGTranformer) BeforeCommit(stage *ladybugsim_models.StageStruct) {

	// remove all gongsvg stage/repo
	gongsvg_models.Stage.Checkout()
	ladybugsim_models.Stage.Checkout()

	//
	// SVG
	//
	if svgSingloton == nil {
		svgSingloton = new(gongsvg_models.SVG).Stage()
		svgSingloton.Name = "New ladybugsim Chart"
		svgSingloton.Display = true
	}

	if fenceSingloton == nil {
		fenceSingloton = new(gongsvg_models.Rect).Stage()
		fenceSingloton.Name = "fence"
		svgSingloton.Rects = append(svgSingloton.Rects, fenceSingloton)
		fenceSingloton.X = 50
		fenceSingloton.Y = 50
		fenceSingloton.Height = 30
		fenceSingloton.Width = 2000

		fenceSingloton.Color = "blue"
		fenceSingloton.FillOpacity = 0.1
		fenceSingloton.Stroke = "blue"
		fenceSingloton.StrokeWidth = 0.5
	}

	simSpeed := gongsim_models.EngineSingloton.Speed

	var sim *ladybugsim_models.LadybugSimulation
	for _ladybugsim := range stage.LadybugSimulations {
		sim = _ladybugsim
	}

	for _, ladybug := range sim.Ladybugs {

		circle := mapLadybug_Circle[ladybug]
		if circle == nil {
			circle = new(gongsvg_models.Circle).Stage()
			circle.Name = ladybug.Name
			circle.FillOpacity = 0.1
			circle.StrokeWidth = 0.5
			circle.Radius = 5

			svgSingloton.Circles = append(svgSingloton.Circles, circle)
			mapLadybug_Circle[ladybug] = circle
		}

		if ladybug.Speed > 0 {
			circle.Color = "red"
			circle.Stroke = "red"
		} else {
			circle.Color = "green"
			circle.Stroke = "green"
		}
		circle.CX = fenceSingloton.X + ladybug.Position*fenceSingloton.Width
		circle.CY = fenceSingloton.Y

		if ladybug.LadybugStatus == ladybugsim_models.ON_THE_GROUND {
			circle.CY = fenceSingloton.Y + fenceSingloton.Height
		}

		animate := mapLadybug_Animate[ladybug]
		if animate == nil {
			animate = new(gongsvg_models.Animate).Stage()
			animate.Name = ladybug.Name
			circle.Animations = append(circle.Animations, animate)
			animate.AttributeName = "cx"
			mapLadybug_Animate[ladybug] = animate
		}

		//  gongsim_models.EngineSingloton.State == gongsim_models.RUNNING
		if ladybug.LadybugStatus == ladybugsim_models.ON_THE_FENCE {
			animate.Stage()

			if gongsim_models.EngineSingloton.State == gongsim_models.RUNNING {
				animate.Values = fmt.Sprintf("%d;%d",
					int64(circle.CX),
					int64(circle.CX+
						simSpeed*
							ladybug.Speed*
							fenceSingloton.Width))
			} else {
				animate.Values = fmt.Sprintf("%d;%d",
					int64(circle.CX),
					int64(circle.CX))
			}
			animate.Dur = "1s"
			animate.RepeatCount = "undefinite"
		} else {
			animate.Unstage()
			circle.Animations = nil
		}

	}

	// {
	// 	circle := new(gongsvg_models.Circle).Stage()
	// 	circle.Name = "laft relay"
	// 	svgSingloton.Circles = append(svgSingloton.Circles, circle)

	// 	circle.Color = "black"
	// 	circle.Stroke = "black"

	// 	circle.FillOpacity = 0.1
	// 	circle.StrokeWidth = 0.5
	// 	circle.CX = fenceSingloton.X + sim.LeftRelayInitialPosX*fenceSingloton.Width
	// 	circle.CY = fenceSingloton.Y
	// 	circle.Radius = 10

	// 	// add animation
	// 	animate := new(gongsvg_models.Animate).Stage()
	// 	animate.Name = circle.Name
	// 	circle.Animations = append(circle.Animations, animate)
	// 	animate.AttributeName = "cx"
	// 	animate.Values = fmt.Sprintf("%d;%d", int64(circle.CX),
	// 		int64(circle.CX+simSpeed*sim.AbsoluteSpeed*fenceSingloton.Width))
	// 	animate.Dur = "1s"
	// 	animate.RepeatCount = "undefinite"
	// }

	// {
	// 	circle := new(gongsvg_models.Circle).Stage()
	// 	circle.Name = "laft relay"
	// 	svgSingloton.Circles = append(svgSingloton.Circles, circle)

	// 	circle.Color = "black"
	// 	circle.Stroke = "black"

	// 	circle.FillOpacity = 0.1
	// 	circle.StrokeWidth = 0.5
	// 	circle.CX = fenceSingloton.X + sim.RightRelayInitialPosX*fenceSingloton.Width
	// 	circle.CY = fenceSingloton.Y
	// 	circle.Radius = 10

	// 	// add animation
	// 	animate := new(gongsvg_models.Animate).Stage()
	// 	animate.Name = circle.Name
	// 	circle.Animations = append(circle.Animations, animate)
	// 	animate.AttributeName = "cx"
	// 	animate.Values = fmt.Sprintf("%d;%d", int64(circle.CX), int64(circle.CX-simSpeed*sim.AbsoluteSpeed*fenceSingloton.Width))
	// 	animate.Dur = "1s"
	// 	animate.RepeatCount = "undefinite"
	// }

	gongsvg_models.Stage.Commit()

	// log.Printf("Before Commit")
}
