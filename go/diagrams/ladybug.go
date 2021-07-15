package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/ladybugsim/go/models"
)

var ladybug uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Ladybug{}),
			Position: &uml.Position{
				X: 660.000000,
				Y: 30.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Ladybug{}.Id,
				},
				{
					Field: models.Ladybug{}.Name,
				},
				{
					Field: models.Ladybug{}.Speed,
				},
			},
		},
		{
			Struct: &(models.LadybugSimulation{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 30.000000,
			},
			Width:  240.000000,
			Heigth: 213.000000,
			Links: []*uml.Link{
				{
					Field: models.LadybugSimulation{}.Ladybugs,
					Middlevertice: &uml.Vertice{
						X: 350.000000,
						Y: 86.500000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.LadybugSimulation{}.AbsoluteSpeed,
				},
				{
					Field: models.LadybugSimulation{}.EventNb,
				},
				{
					Field: models.LadybugSimulation{}.LadybugRadius,
				},
				{
					Field: models.LadybugSimulation{}.LeftRelayInitialPosX,
				},
				{
					Field: models.LadybugSimulation{}.MaxDistanceInOneStep,
				},
				{
					Field: models.LadybugSimulation{}.NbLadybugs,
				},
				{
					Field: models.LadybugSimulation{}.NbLadybugsOnTheGround,
				},
				{
					Field: models.LadybugSimulation{}.NbOfCollision,
				},
				{
					Field: models.LadybugSimulation{}.RightRelayInitialPosX,
				},
				{
					Field: models.LadybugSimulation{}.SimulationStep,
				},
			},
		},
	},
}
