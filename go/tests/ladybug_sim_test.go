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
	for _, ladybug := range ladybugsim_models.LadybugSim.Ladybugs {
		if ladybug.Speed > 0 {
			leftRelayInitialPosX = ladybug.Position
			log.Printf("Left relay %1.4f", leftRelayInitialPosX)
			break
		}
	}

	rightRelayInitialPosX := 0.0
	for i := len(ladybugsim_models.LadybugSim.Ladybugs) - 1; i >= 0; i-- {

		ladybug := ladybugsim_models.LadybugSim.Ladybugs[i]
		if ladybug.Speed < 0 {
			rightRelayInitialPosX = ladybug.Position
			log.Printf("right relay %1.4f", rightRelayInitialPosX)
			break
		}
	}

	// time for the left relay to get out
	leftRelayDistanceToExit := 1.0 - leftRelayInitialPosX
	log.Printf("left distance to exit %1.4f", leftRelayDistanceToExit)
	log.Printf("left time to exit %1.4f", leftRelayDistanceToExit/ladybugsim_models.LadybugSim.AbsoluteSpeed)

	rightRelayDistanceToExit := rightRelayInitialPosX
	log.Printf("right distance to exit %1.4f", rightRelayDistanceToExit)
	log.Printf("right time to exit %1.4f", rightRelayDistanceToExit/ladybugsim_models.LadybugSim.AbsoluteSpeed)

	maxRelayDistanceFromBorder := math.Max(leftRelayInitialPosX, 1.0-rightRelayInitialPosX)
	log.Printf("max distance relay to border %1.4f", maxRelayDistanceFromBorder)

	_, nextSimTime, _ := gongsim_models.EngineSingloton.GetNextEvent()
	for nextSimTime.Before(gongsim_models.EngineSingloton.GetEndTime()) {
		var agent gongsim_models.AgentInterface
		agent, nextSimTime, _ = gongsim_models.EngineSingloton.FireNextEvent()
		if agent == nil {
			log.Printf("Ladybug sim over")
			return
		}
	}

	log.Printf("Ladybug sim over")
}
