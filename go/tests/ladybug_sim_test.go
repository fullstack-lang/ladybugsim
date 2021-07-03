package tests

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	ladybugsim_models "github.com/fullstack-lang/ladybugsim/go/models"
)

func TestLadybugSim(t *testing.T) {

	log.SetFlags(0)

	// rand.Seed(time.Now().UnixNano())

	// Kills Engine Simulation goroutine
	gongsim_models.Quit <- true

	// seven days of simulation
	gongsim_models.EngineSingloton.SetStartTime(time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC))
	gongsim_models.EngineSingloton.SetCurrentTime(gongsim_models.EngineSingloton.GetStartTime())
	gongsim_models.EngineSingloton.State = gongsim_models.PAUSED
	gongsim_models.EngineSingloton.Speed = 1.0 // realtime
	// log.Printf("Sim start \t\t\t%s\n", gongsim_models.EngineSingloton.GetStartTime())

	// Three years
	gongsim_models.EngineSingloton.SetEndTime(time.Date(2021, time.July, 1, 0, 30, 0, 0, time.UTC))
	// log.Printf("Sim end  \t\t\t%s\n", gongsim_models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var ladyBugSimulation ladybugsim_models.LadybugSimulation
	gongsim_models.EngineSingloton.Simulation = &ladyBugSimulation

	// initial positions of ladybugs cannot be close to each others than the radius
	initialXPosition := make(map[float64]*ladybugsim_models.Ladybug)

	// append a ladybug agent to feed the discrete event engine
	for ladybugId := 0; ladybugId < ladybugsim_models.NbLadybugs; ladybugId = ladybugId + 1 {
		ladyBug := new(ladybugsim_models.Ladybug)
		ladyBug.Name = fmt.Sprintf("Ladybug #%2d", ladybugId)
		ladyBug.Id = ladybugId

		ladybugsim_models.LadybugSlice = append(ladybugsim_models.LadybugSlice, ladyBug)

		// set up position
		positionX := rand.Float64()

		// adjust it on a multiple of the ladybug diameter
		positionX = math.Round(positionX*1.0/ladybugsim_models.LadybugRadius) * ladybugsim_models.LadybugRadius

		ladyBug.Position = positionX
		if initialXPosition[positionX] != nil {
			log.Panic("same initial position")
		}

		// decide orientaiton of the speed
		if rand.Float64() > 0.5 {
			ladyBug.Speed = ladybugsim_models.AbsoluteSpeed
		} else {
			ladyBug.Speed = -ladybugsim_models.AbsoluteSpeed
		}

		gongsim_models.EngineSingloton.AppendAgent(ladyBug)
		var step gongsim_models.UpdateState
		step.SetFireTime(gongsim_models.EngineSingloton.GetStartTime())
		step.Period = ladybugsim_models.SimulationStep //
		step.Name = "update of laybug motion"
		ladyBug.QueueEvent(&step)
	}

	sort.Slice(ladybugsim_models.LadybugSlice[:], func(i, j int) bool {
		return ladybugsim_models.LadybugSlice[i].Position < ladybugsim_models.LadybugSlice[j].Position
	})

	// compute left & right relay positions
	leftRelayInitialPosX := 0.0
	for _, ladybug := range ladybugsim_models.LadybugSlice {
		if ladybug.Speed > 0 {
			leftRelayInitialPosX = ladybug.Position
			log.Printf("Left relay %1.4f", leftRelayInitialPosX)
			break
		}
	}

	rightRelayInitialPosX := 0.0
	for i := len(ladybugsim_models.LadybugSlice) - 1; i >= 0; i-- {

		ladybug := ladybugsim_models.LadybugSlice[i]
		if ladybug.Speed > 0 {
			rightRelayInitialPosX = ladybug.Position
			log.Printf("right relay %1.4f", rightRelayInitialPosX)
			break
		}
	}

	middleBetweenRelayX := (leftRelayInitialPosX + rightRelayInitialPosX) / 2.0
	log.Printf("middle between relays %1.4f", middleBetweenRelayX)

	// time for the left relay to get out
	leftRelayDistanceToExit := (middleBetweenRelayX - leftRelayInitialPosX) + middleBetweenRelayX
	log.Printf("left distance to exit %1.4f", leftRelayDistanceToExit)
	log.Printf("left time to exit %1.4f", leftRelayDistanceToExit*60.0)

	rightRelayDistanceToExit := (rightRelayInitialPosX - middleBetweenRelayX) + (1.0 - middleBetweenRelayX)
	log.Printf("right distance to exit %1.4f", rightRelayDistanceToExit)
	log.Printf("right time to exit %1.4f", rightRelayDistanceToExit*60.0)

	maxRelayDistanceFromBorder := math.Max(leftRelayInitialPosX, 1.0-rightRelayInitialPosX)
	log.Printf("max distance relay to border %1.4f", maxRelayDistanceFromBorder)

	_, nextSimTime, _ := gongsim_models.EngineSingloton.GetNextEvent()
	for nextSimTime.Before(gongsim_models.EngineSingloton.GetEndTime()) {
		_, nextSimTime, _ = gongsim_models.EngineSingloton.FireNextEvent()
	}

	log.Printf("Ladybug sim over")
}
