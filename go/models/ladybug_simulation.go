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
	Name                 string
	EventNb              int
	NbOfCollision        int
	LadybugRadius        float64
	AbsoluteSpeed        float64
	SimulationStep       time.Duration
	MaxDistanceInOneStep float64
	NbLadybugs           int
	Ladybugs             []*Ladybug
}

func (specificEngine *LadybugSimulation) EventFired(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {
	return true
}
func (specificEngine *LadybugSimulation) Reset(engine *gongsim_models.Engine)          {}
func (specificEngine *LadybugSimulation) CommitAgents(engine *gongsim_models.Engine)   {}
func (specificEngine *LadybugSimulation) CheckoutAgents(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) GetLastCommitNb() uint                        { return 0 }

var LadybugSimulationSingloton *LadybugSimulation

func init() {
	LadybugSimulationSingloton = new(LadybugSimulation)
	LadybugSimulationSingloton.Name = "Simulation of ladybugs"
	LadybugSimulationSingloton.EventNb = 0
	LadybugSimulationSingloton.LadybugRadius = 0.00005    // a ladybug is 1mm wide
	LadybugSimulationSingloton.AbsoluteSpeed = 1.0 / 60.0 // a ladybug is 1m par minute
	LadybugSimulationSingloton.MaxDistanceInOneStep = float64(0.99 * (LadybugSimulationSingloton.LadybugRadius / 4.0) * 60.0)
	LadybugSimulationSingloton.SimulationStep = time.Microsecond * time.Duration(LadybugSimulationSingloton.MaxDistanceInOneStep*1000000.0)

	LadybugSimulationSingloton.NbLadybugs = 32

	// seven days of simulation
	gongsim_models.EngineSingloton.SetStartTime(time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC))
	gongsim_models.EngineSingloton.SetCurrentTime(gongsim_models.EngineSingloton.GetStartTime())
	gongsim_models.EngineSingloton.State = gongsim_models.PAUSED
	gongsim_models.EngineSingloton.Speed = 1.0 // realtime
	// log.Printf("Sim start \t\t\t%s\n", gongsim_models.EngineSingloton.GetStartTime())

	gongsim_models.EngineSingloton.SetEndTime(time.Date(2021, time.July, 1, 0, 30, 0, 0, time.UTC))
	// log.Printf("Sim end  \t\t\t%s\n", gongsim_models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	gongsim_models.EngineSingloton.Simulation = LadybugSimulationSingloton

	// initial positions of ladybugs cannot be close to each others than the radius
	initialXPosition := make(map[float64]*Ladybug)

	sortedInitialXPositions := make([]float64, LadybugSimulationSingloton.NbLadybugs)
	for ladybugId := 0; ladybugId < LadybugSimulationSingloton.NbLadybugs; ladybugId = ladybugId + 1 {
		// set up position
		positionX := rand.Float64()

		// adjust it on a multiple of the ladybug diameter
		positionX = math.Round(positionX*1.0/LadybugSimulationSingloton.LadybugRadius) * LadybugSimulationSingloton.LadybugRadius

		sortedInitialXPositions[ladybugId] = positionX
		if initialXPosition[positionX] != nil {
			log.Panic("same initial position")
		}
	}

	sort.Slice(sortedInitialXPositions[:], func(i, j int) bool {
		return sortedInitialXPositions[i] < sortedInitialXPositions[j]
	})

	// append a ladybug agent to feed the discrete event engine
	for ladybugId := 0; ladybugId < LadybugSimulationSingloton.NbLadybugs; ladybugId = ladybugId + 1 {
		ladyBug := new(Ladybug)
		ladyBug.Name = fmt.Sprintf("Ladybug #%2d", ladybugId)
		ladyBug.Id = ladybugId

		LadybugSimulationSingloton.Ladybugs =
			append(LadybugSimulationSingloton.Ladybugs, ladyBug)

		ladyBug.Position = sortedInitialXPositions[ladybugId]

		// decide orientaiton of the speed
		if rand.Float64() > 0.5 {
			ladyBug.Speed = LadybugSimulationSingloton.AbsoluteSpeed
		} else {
			ladyBug.Speed = -LadybugSimulationSingloton.AbsoluteSpeed
		}

		gongsim_models.EngineSingloton.AppendAgent(ladyBug)
		var step gongsim_models.UpdateState
		step.SetFireTime(gongsim_models.EngineSingloton.GetStartTime())
		step.Period = LadybugSimulationSingloton.SimulationStep //
		step.Name = "update of laybug motion"
		ladyBug.QueueEvent(&step)
	}
}
