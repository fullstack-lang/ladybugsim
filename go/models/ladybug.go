package models

import (
	"fmt"
	"log"
	"math"
	"os"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

type Ladybug struct {
	gongsim_models.Agent

	Name string

	Id int

	Position float64 // between 0.0 and 1.0

	Speed float64 // either AbsoluteSpeed or -AbsoluteSpeed
}

func (ladybug *Ladybug) FireNextEvent() {
	event, eventTime := ladybug.GetNextEventAndRemoveIt()

	ladybugSimulation := gongsim_models.EngineSingloton.Simulation.(*LadybugSimulation)

	if ladybugSimulation.EventNb%32 == 0 && ladybug.Speed != 0.0 {

		var positions string
		for _, _ladybug := range ladybugSimulation.Ladybugs {
			direction := "D"
			if _ladybug.Speed < 0 {
				direction = "G"
			}
			positions = positions + fmt.Sprintf("%0.2f %s", _ladybug.Position, direction)
		}

		// log.Printf("%s", positions)
	}

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// if eventNb%5000 == 0 && ladybug.Id == 0 {
		// 	log.Printf("Event %10d Time : %s Ladybug %s Position %10f Speed %10f", eventNb, eventTime.Format("15:04:05.000000"), ladybug.Name, ladybug.Position, ladybug.Speed)
		// }
		// if eventNb%3200 == 0 && ladybug.Speed != 0.0 {
		// 	log.Printf("Event %10d Time : %s Ladybug %s Position %10f Speed %10f", eventNb, eventTime.Format("15:04:05.000000"), ladybug.Name, ladybug.Position, ladybug.Speed)
		// }
		ladybugSimulation.EventNb = ladybugSimulation.EventNb + 1

		ladybug.Position = ladybug.Position + ladybugSimulation.SimulationStep.Seconds()*ladybug.Speed

		// stop simu if sum of speeds is 0
		sumOfSpeeds := 0.0
		for _, otherLadybug := range ladybugSimulation.Ladybugs {
			// sum speeds
			sumOfSpeeds = sumOfSpeeds + math.Abs(otherLadybug.Speed)
		}
		if sumOfSpeeds == 0 {
			log.Printf("Event %10d Time : %s, nbOfCollisions %d simulation over",
				ladybugSimulation.EventNb, eventTime.Format("15:04:05.000000"), ladybugSimulation.NbOfCollision/2)
			os.Exit(0)
		}

		// check for colisions (and compute)
		for _, otherLadybug := range ladybugSimulation.Ladybugs {
			// sum speeds
			sumOfSpeeds = sumOfSpeeds + math.Abs(otherLadybug.Speed)

			// do not compute collision of a ladybug with itslef
			if otherLadybug.Id == ladybug.Id {
				continue
			}

			// do not compute collision if the other ladybug is allready out
			if otherLadybug.Speed == 0.0 {
				continue
			}

			// get the between X positions
			deltaX := otherLadybug.Position - ladybug.Position

			// there is a collision if both are within a Ladybug diameter
			if math.Abs(deltaX) < 2*ladybugSimulation.LadybugRadius {

				if ladybug.Id == 0 {
					log.Printf("Event %10d Time : %s Pos %10f dist %10f ladybug %2d / %2d",
						ladybugSimulation.EventNb, eventTime.Format("15:04:05.000000"), otherLadybug.Position, deltaX, ladybug.Id, otherLadybug.Id)
				}

				if deltaX > 0 && ladybug.Speed > 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.Id)*1.0
				}
				if deltaX < 0 && ladybug.Speed < 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.Id)*1.0
				}

				ladybugSimulation.NbOfCollision = ladybugSimulation.NbOfCollision + 1
			}
		}

		if ladybug.Position < 0 || ladybug.Position > 1.0 {
			ladybug.Speed = 0.0
		}

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		ladybug.QueueEvent(checkStateEvent)
	}
}
