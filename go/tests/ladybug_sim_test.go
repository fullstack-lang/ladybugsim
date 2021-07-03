package tests

import (
	"log"
	"math"
	"testing"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	ladybugsim_models "github.com/fullstack-lang/ladybugsim/go/models"
)

func TestLadybugSim(t *testing.T) {

	log.SetFlags(0)

	// rand.Seed(time.Now().UnixNano())

	// Kills Engine Simulation goroutine
	gongsim_models.Quit <- true

	// compute left & right relay positions
	leftRelayInitialPosX := 0.0
	for _, ladybug := range ladybugsim_models.LadyBugSimulationSingloton.Ladybugs {
		if ladybug.Speed > 0 {
			leftRelayInitialPosX = ladybug.Position
			log.Printf("Left relay %1.4f", leftRelayInitialPosX)
			break
		}
	}

	rightRelayInitialPosX := 0.0
	for i := len(ladybugsim_models.LadyBugSimulationSingloton.Ladybugs) - 1; i >= 0; i-- {

		ladybug := ladybugsim_models.LadyBugSimulationSingloton.Ladybugs[i]
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
