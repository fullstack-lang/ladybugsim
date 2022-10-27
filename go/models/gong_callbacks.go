package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Ladybug:
		if stage.OnAfterLadybugCreateCallback != nil {
			stage.OnAfterLadybugCreateCallback.OnAfterCreate(stage, target)
		}
	case *LadybugSimulation:
		if stage.OnAfterLadybugSimulationCreateCallback != nil {
			stage.OnAfterLadybugSimulationCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Ladybug:
		newTarget := any(new).(*Ladybug)
		if stage.OnAfterLadybugUpdateCallback != nil {
			stage.OnAfterLadybugUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LadybugSimulation:
		newTarget := any(new).(*LadybugSimulation)
		if stage.OnAfterLadybugSimulationUpdateCallback != nil {
			stage.OnAfterLadybugSimulationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Ladybug:
		if stage.OnAfterLadybugDeleteCallback != nil {
			staged := any(staged).(*Ladybug)
			stage.OnAfterLadybugDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LadybugSimulation:
		if stage.OnAfterLadybugSimulationDeleteCallback != nil {
			staged := any(staged).(*LadybugSimulation)
			stage.OnAfterLadybugSimulationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Ladybug:
		if stage.OnAfterLadybugReadCallback != nil {
			stage.OnAfterLadybugReadCallback.OnAfterRead(stage, target)
		}
	case *LadybugSimulation:
		if stage.OnAfterLadybugSimulationReadCallback != nil {
			stage.OnAfterLadybugSimulationReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Ladybug:
		stage.OnAfterLadybugUpdateCallback = any(callback).(OnAfterUpdateInterface[Ladybug])
	
	case *LadybugSimulation:
		stage.OnAfterLadybugSimulationUpdateCallback = any(callback).(OnAfterUpdateInterface[LadybugSimulation])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Ladybug:
		stage.OnAfterLadybugCreateCallback = any(callback).(OnAfterCreateInterface[Ladybug])
	
	case *LadybugSimulation:
		stage.OnAfterLadybugSimulationCreateCallback = any(callback).(OnAfterCreateInterface[LadybugSimulation])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Ladybug:
		stage.OnAfterLadybugDeleteCallback = any(callback).(OnAfterDeleteInterface[Ladybug])
	
	case *LadybugSimulation:
		stage.OnAfterLadybugSimulationDeleteCallback = any(callback).(OnAfterDeleteInterface[LadybugSimulation])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Ladybug:
		stage.OnAfterLadybugReadCallback = any(callback).(OnAfterReadInterface[Ladybug])
	
	case *LadybugSimulation:
		stage.OnAfterLadybugSimulationReadCallback = any(callback).(OnAfterReadInterface[LadybugSimulation])
	
	}
}
