// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Ladybug:
		ok = stage.IsStagedLadybug(target)

	case *LadybugSimulation:
		ok = stage.IsStagedLadybugSimulation(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedLadybug(ladybug *Ladybug) (ok bool) {

		_, ok = stage.Ladybugs[ladybug]
	
		return
	}

	func (stage *StageStruct) IsStagedLadybugSimulation(ladybugsimulation *LadybugSimulation) (ok bool) {

		_, ok = stage.LadybugSimulations[ladybugsimulation]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Ladybug:
		stage.StageBranchLadybug(target)

	case *LadybugSimulation:
		stage.StageBranchLadybugSimulation(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchLadybug(ladybug *Ladybug) {

	// check if instance is already staged
	if IsStaged(stage, ladybug) {
		return
	}

	ladybug.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLadybugSimulation(ladybugsimulation *LadybugSimulation) {

	// check if instance is already staged
	if IsStaged(stage, ladybugsimulation) {
		return
	}

	ladybugsimulation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ladybug := range ladybugsimulation.Ladybugs {
		StageBranch(stage, _ladybug)
	}

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Ladybug:
		stage.UnstageBranchLadybug(target)

	case *LadybugSimulation:
		stage.UnstageBranchLadybugSimulation(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchLadybug(ladybug *Ladybug) {

	// check if instance is already staged
	if ! IsStaged(stage, ladybug) {
		return
	}

	ladybug.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLadybugSimulation(ladybugsimulation *LadybugSimulation) {

	// check if instance is already staged
	if ! IsStaged(stage, ladybugsimulation) {
		return
	}

	ladybugsimulation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _ladybug := range ladybugsimulation.Ladybugs {
		UnstageBranch(stage, _ladybug)
	}

}

