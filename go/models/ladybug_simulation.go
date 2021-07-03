package models

import (
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

type LadybugSimulation struct {
}

var NbOfCollision = 0

func (specificEngine *LadybugSimulation) EventFired(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {
	return true
}
func (specificEngine *LadybugSimulation) Reset(engine *gongsim_models.Engine)          {}
func (specificEngine *LadybugSimulation) CommitAgents(engine *gongsim_models.Engine)   {}
func (specificEngine *LadybugSimulation) CheckoutAgents(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) GetLastCommitNb() uint                        { return 0 }
