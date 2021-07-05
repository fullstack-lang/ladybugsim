package models

import (
	"log"
	"math"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

type Ladybug struct {
	gongsim_models.Agent

	Name string

	Id int

	Position float64 // between 0.0 and 1.0

	Speed float64 // either AbsoluteSpeed or -AbsoluteSpeed

	LadybugStatus LadybugStatus
}

func (ladybug *Ladybug) FireNextEvent() {
	event, eventTime := ladybug.GetNextEventAndRemoveIt()

	sim := gongsim_models.EngineSingloton.Simulation.(*LadybugSimulation)

	if ladybug.LadybugStatus == ON_THE_GROUND {
		return
	}

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		sim.EventNb = sim.EventNb + 1

		//
		// update ladybug position
		//
		ladybug.Position = ladybug.Position + sim.SimulationStep.Seconds()*ladybug.Speed

		// check for colisions (and compute)
		for _, otherLadybug := range sim.Ladybugs {

			// do not compute collision of a ladybug with itslef
			if otherLadybug.Id == ladybug.Id {
				continue
			}

			// do not compute collision if the other ladybug is allready on the ground
			if otherLadybug.LadybugStatus == ON_THE_GROUND {
				continue
			}

			// get the delta between X positions
			deltaX := otherLadybug.Position - ladybug.Position

			// there is a collision if both are within a Ladybug diameter
			if math.Abs(deltaX) < 2*sim.LadybugRadius {

				// if ladybug.Id == 0 {
				log.Printf("Event %10d Time : %s Pos %10f dist %10f ladybug %2d / %2d",
					sim.EventNb, eventTime.Format("15:04:05.000000"),
					otherLadybug.Position, deltaX, ladybug.Id, otherLadybug.Id)
				// }

				if deltaX > 0 && ladybug.Speed > 0 {
					// return
					ladybug.Speed = -ladybug.Speed
				}
				if deltaX < 0 && ladybug.Speed < 0 {
					// return
					ladybug.Speed = -ladybug.Speed
				}

				sim.NbOfCollision = sim.NbOfCollision + 1
			}
		}

		if ladybug.Position < 0 || ladybug.Position > 1.0 {
			ladybug.LadybugStatus = ON_THE_GROUND
		}

		// stop simu if all ladybugs are on the ground
		allLadybugsOnTheGround := true
		nbLadybugsOnTheGround := 0
		for _, _ladybug := range sim.Ladybugs {
			if _ladybug.LadybugStatus == ON_THE_FENCE {
				allLadybugsOnTheGround = false
			} else {
				nbLadybugsOnTheGround = nbLadybugsOnTheGround + 1
			}
		}
		sim.NbLadybugsOnTheGround = nbLadybugsOnTheGround
		if allLadybugsOnTheGround {
			log.Printf("Event %10d Time : %s, nbOfCollisions %d simulation over",
				sim.EventNb, eventTime.Format("15:04:05.000000"), sim.NbOfCollision/2)

			gongsim_models.EngineSingloton.State = gongsim_models.OVER

			return
		}

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		ladybug.QueueEvent(checkStateEvent)
	}
}
