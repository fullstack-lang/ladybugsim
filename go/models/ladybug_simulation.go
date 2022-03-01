package models

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

type LadybugSimulation struct {
	Name                  string
	EventNb               int
	NbOfCollision         int
	LadybugRadius         float64
	AbsoluteSpeed         float64
	SimulationStep        time.Duration
	MaxDistanceInOneStep  float64
	NbLadybugs            int
	NbLadybugsOnTheGround int

	// an hint on the theorem
	LeftRelayInitialPosX  float64
	RightRelayInitialPosX float64

	Ladybugs []*Ladybug
}

func (specificEngine *LadybugSimulation) EventFired(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {
	return true
}
func (specificEngine *LadybugSimulation) Reset(engine *gongsim_models.Engine) {

}
func (specificEngine *LadybugSimulation) CommitAgents(engine *gongsim_models.Engine) {
	// commit the simulation agent states
	Stage.Commit()
}
func (specificEngine *LadybugSimulation) CheckoutAgents(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) GetLastCommitNb() (commitNb uint) {
	if Stage.BackRepo != nil {
		commitNb = Stage.BackRepo.GetLastCommitFromBackNb()
	}

	return
}

func (specificEngine *LadybugSimulation) GetLastCommitNbFromFront() (commitNb uint) {
	if Stage.BackRepo != nil {
		commitNb = Stage.BackRepo.GetLastCommitFromBackNb()
	}

	return
}

var LadybugSim *LadybugSimulation

const numericalSimuationAdjustment = 0.9999

func init() {
	LadybugSim = new(LadybugSimulation)

	LadybugSim.NbLadybugs = 32
	LadybugSim.Name = "Simulation of ladybugs"
	LadybugSim.EventNb = 0
	LadybugSim.LadybugRadius = 0.005 / numericalSimuationAdjustment //
	LadybugSim.AbsoluteSpeed = 1.0 / 60.0                           // a ladybug is 1m par minute

	seed := time.Now().UnixNano()
	log.Printf("seed %d ", seed)

	seed = 1625472777467
	rand.Seed(seed)

	// Simulation step must
	//
	// This is an event based simulation that is implemented as fixed step simulation
	// the simulation
	// - be the biggest possible to have the biggest simulation step
	// - have 2 ladybug that are running towards each other detect that they are in a collision
	//
	//         |----*----| ->     <- |----*----|
	//            |----*----|      |----*----|
	//              |----*----|  |----*----|
	//                |----*--|--|--*----|         !!! collision !!!
	//
	//         |----*----| ->   <- |----*----|
	//            |----*----|    |----*----|
	//              |----*----||----*----|       At the step before collision, both ladybugs are almost touching
	//                |----*|----|*----|         !!! collision !!!
	//
	//
	// at each step, a lady walks distance  -- This must be less than half the radius of the ladybug
	//
	//
	// note : the 0.999 is for being robust to numerical inacura
	LadybugSim.MaxDistanceInOneStep = (LadybugSim.LadybugRadius / 2.0) * numericalSimuationAdjustment

	//
	// distance = time * speed
	//
	// =>
	//
	// time = distance / speed
	//
	simStep := LadybugSim.MaxDistanceInOneStep / LadybugSim.AbsoluteSpeed

	// from golang time
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
	// The representation limits the largest representable duration to approximately 290 years.
	LadybugSim.SimulationStep = time.Duration(
		simStep *
			1000.0 * 1000.0 * 1000.0)

	gongsim_models.EngineSingloton.SetStartTime(time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC))
	gongsim_models.EngineSingloton.SetCurrentTime(gongsim_models.EngineSingloton.GetStartTime())
	gongsim_models.EngineSingloton.State = gongsim_models.PAUSED
	gongsim_models.EngineSingloton.Speed = 0.5 // realtime
	// log.Printf("Sim start \t\t\t%s\n", gongsim_models.EngineSingloton.GetStartTime())

	gongsim_models.EngineSingloton.SetEndTime(time.Date(2021, time.July, 1, 0, 30, 0, 0, time.UTC))
	// log.Printf("Sim end  \t\t\t%s\n", gongsim_models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	gongsim_models.EngineSingloton.Simulation = LadybugSim

	// initial positions of ladybugs cannot be close to each others than the radius
	initialXPosition := make(map[float64]*Ladybug)

	sortedInitialXPositions := make([]float64, LadybugSim.NbLadybugs)
	for ladybugId := 0; ladybugId < LadybugSim.NbLadybugs; ladybugId = ladybugId + 1 {
		// set up position
		positionX := rand.Float64()

		minNbOfRadiusBetweenInitialLadybugs := 3.0

		// adjust it on a multiple of the ladybug diameter
		positionX = math.Round(positionX*1.0/
			(minNbOfRadiusBetweenInitialLadybugs*LadybugSim.LadybugRadius)) * (minNbOfRadiusBetweenInitialLadybugs * LadybugSim.LadybugRadius)

		for initialXPosition[positionX] != nil {
			// log.Panic("same initial position")
			positionX = rand.Float64()

			// adjust it on a multiple of the ladybug diameter
			positionX = math.Round(positionX*1.0/
				(minNbOfRadiusBetweenInitialLadybugs*LadybugSim.LadybugRadius)) * (minNbOfRadiusBetweenInitialLadybugs * LadybugSim.LadybugRadius)
		}
		initialXPosition[positionX] = &Ladybug{}
		sortedInitialXPositions[ladybugId] = positionX
	}

	sort.Slice(sortedInitialXPositions[:], func(i, j int) bool {
		return sortedInitialXPositions[i] < sortedInitialXPositions[j]
	})

	// append a ladybug agent to feed the discrete event engine
	for ladybugId := 0; ladybugId < LadybugSim.NbLadybugs; ladybugId = ladybugId + 1 {
		ladybug := new(Ladybug)
		ladybug.Name = fmt.Sprintf("Ladybug #%2d", ladybugId)
		ladybug.Id = ladybugId

		LadybugSim.Ladybugs =
			append(LadybugSim.Ladybugs, ladybug)

		ladybug.Position = sortedInitialXPositions[ladybugId]
		ladybug.LadybugStatus = ON_THE_FENCE

		// decide orientaiton of the speed
		if rand.Float64() > 0.5 {
			ladybug.Speed = LadybugSim.AbsoluteSpeed
		} else {
			ladybug.Speed = -LadybugSim.AbsoluteSpeed
		}

		gongsim_models.EngineSingloton.AppendAgent(ladybug)
		var updatePositionEvent UpdatePositionEvent
		updatePositionEvent.SetFireTime(gongsim_models.EngineSingloton.GetStartTime())
		ladybug.QueueEvent(&updatePositionEvent)
	}

	// compute left & right relay positions
	LadybugSim.LeftRelayInitialPosX = 0.0
	for _, ladybug := range LadybugSim.Ladybugs {
		if ladybug.Speed > 0 && ladybug.LadybugStatus == ON_THE_FENCE {
			LadybugSim.LeftRelayInitialPosX = ladybug.Position
			break
		}
	}

	LadybugSim.RightRelayInitialPosX = 0.0
	for i := len(LadybugSim.Ladybugs) - 1; i >= 0; i-- {

		ladybug := LadybugSim.Ladybugs[i]
		if ladybug.Speed < 0 && ladybug.LadybugStatus == ON_THE_FENCE {
			LadybugSim.RightRelayInitialPosX = ladybug.Position
			break
		}
	}
}
