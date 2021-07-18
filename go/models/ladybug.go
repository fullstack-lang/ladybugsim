package models

import (
	"log"
	"math"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Ladybug struct {
	gongsim_models.Agent

	Name string

	Id int

	Position float64 // between 0.0 and 1.0

	Speed float64 // either AbsoluteSpeed or -AbsoluteSpeed

	LadybugStatus LadybugStatus
}

// A ladybug process 2 kinds of events:
// - update position
// - update speed (change speed direction if there is a collision)
//
func (ladybug *Ladybug) FireNextEvent() {
	event, eventTime := ladybug.GetNextEventAndRemoveIt()

	sim := gongsim_models.EngineSingloton.Simulation.(*LadybugSimulation)

	if ladybug.LadybugStatus == ON_THE_GROUND {
		return
	}

	switch event.(type) {
	case *UpdateSpeedEvent:
		updateSpeedEvent := event.(*UpdateSpeedEvent)
		_ = updateSpeedEvent

		// check for colisions with ladybug on the left
		if ladybug.Id > 0 {

		}
		ladybugOnTheLeftId := max(0, ladybug.Id-1)
		ladybugOnTheRigthId := min(len(sim.Ladybugs)-1, ladybug.Id+1)
		for _, otherLadybug := range sim.Ladybugs[ladybugOnTheLeftId : ladybugOnTheRigthId+1] {

			// do not compute collision of a ladybug with itslef
			if otherLadybug.Id == ladybug.Id {
				continue
			}

			// do not compute collision if the other ladybug is allready on the ground
			if otherLadybug.LadybugStatus == ON_THE_GROUND {
				continue
			}

			// get the delta between X positions
			vectorToOtherLadybug := otherLadybug.Position - ladybug.Position

			// there is a collision if both are within a Ladybug diameter
			if math.Abs(vectorToOtherLadybug) < 2*sim.LadybugRadius {

				// if ladybug.Id == 0 {
				// log.Printf("Event %10d Time : %s Pos %10f dist %10f ladybug %2d / %2d",
				// 	sim.EventNb, eventTime.Format("15:04:05.000000"),
				// 	otherLadybug.Position, vectorToOtherLadybug, ladybug.Id, otherLadybug.Id)
				// }

				// check for strange situation
				// collision with a ladybug on the rigth with a speed to the left
				if vectorToOtherLadybug > 0 && ladybug.Speed < 0 {
					log.Panic("In collision while going away")
				}
				if vectorToOtherLadybug < 0 && ladybug.Speed > 0 {
					log.Panic("In collision while going away")
				}

				ladybug.Speed = -ladybug.Speed

				sim.NbOfCollision = sim.NbOfCollision + 1
			}
		}

		updatePositionEvent := new(UpdatePositionEvent)
		updatePositionEvent.SetFireTime(updateSpeedEvent.GetFireTime().Add(sim.SimulationStep / 2.0))
		ladybug.QueueEvent(updatePositionEvent)

	case *UpdatePositionEvent:
		updatePositionEvent := event.(*UpdatePositionEvent)

		sim.EventNb = sim.EventNb + 1

		//
		// update ladybug position
		//
		ladybug.Position = ladybug.Position + sim.SimulationStep.Seconds()*ladybug.Speed

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
			gongsim_models.Stage.Commit()

			return
		}

		// compute left & right relay positions
		// sim.LeftRelayInitialPosX = 0.0
		// for _, ladybug := range sim.Ladybugs {
		// 	if ladybug.Speed > 0 && ladybug.LadybugStatus == ON_THE_FENCE {
		// 		sim.LeftRelayInitialPosX = ladybug.Position
		// 		break
		// 	}
		// }

		// sim.RightRelayInitialPosX = 0.0
		// for i := len(sim.Ladybugs) - 1; i >= 0; i-- {

		// 	ladybug := sim.Ladybugs[i]
		// 	if ladybug.Speed < 0 && ladybug.LadybugStatus == ON_THE_FENCE {
		// 		sim.RightRelayInitialPosX = ladybug.Position
		// 		break
		// 	}
		// }

		// post next event which is a update speed event
		updateSpeedEvent := new(UpdateSpeedEvent)
		updateSpeedEvent.SetFireTime(updatePositionEvent.GetFireTime().Add(sim.SimulationStep / 2.0))
		ladybug.QueueEvent(updateSpeedEvent)
	}
}
