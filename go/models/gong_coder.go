package models

import "time"

// GongfieldCoder return an instance of Type where each field 
// encodes the index of the field
//
// This allows for refactorable field names
// 
func GongfieldCoder[Type Gongstruct]() Type {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case Ladybug:
		fieldCoder := Ladybug{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.Id = 1
		fieldCoder.Position = 2.000000
		fieldCoder.Speed = 3.000000
		fieldCoder.LadybugStatus = "4"
		return (any)(fieldCoder).(Type)
	case LadybugSimulation:
		fieldCoder := LadybugSimulation{}
		// insertion point for field dependant code
		fieldCoder.Name = "0"
		fieldCoder.EventNb = 1
		fieldCoder.NbOfCollision = 2
		fieldCoder.LadybugRadius = 3.000000
		fieldCoder.AbsoluteSpeed = 4.000000
		fieldCoder.SimulationStep = 5
		fieldCoder.MaxDistanceInOneStep = 6.000000
		fieldCoder.NbLadybugs = 7
		fieldCoder.NbLadybugsOnTheGround = 8
		fieldCoder.LeftRelayInitialPosX = 9.000000
		fieldCoder.RightRelayInitialPosX = 10.000000
		return (any)(fieldCoder).(Type)
	default:
		return t
	}
}

type Gongfield interface {
	string | bool | int | float64 | time.Time | time.Duration | *Ladybug | []*Ladybug | *LadybugSimulation | []*LadybugSimulation
}

// GongfieldName provides the name of the field by passing the instance of the coder to
// the fonction.
//
// This allows for refactorable field name
//
// fieldCoder := models.GongfieldCoder[models.Astruct]()
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Name))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Booleanfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Intfield))
// log.Println( models.GongfieldName[*models.Astruct](fieldCoder.Floatfield))
// 
// limitations:
// 1. cannot encode boolean fields
// 2. for associations (pointer to gongstruct or slice of pointer to gongstruct, uses GetAssociationName)
func GongfieldName[Type PointerToGongstruct, FieldType Gongfield](field FieldType) string {
	var t Type

	switch any(t).(type) {
	// insertion point for cases
	case *Ladybug:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
			if field == "4" {
				return "LadybugStatus"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "Id"
			}
		case float64:
			// insertion point for field dependant name
			if field == 2.000000 {
				return "Position"
			}
			if field == 3.000000 {
				return "Speed"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	case *LadybugSimulation:
		switch field := any(field).(type) {
		case string:
			// insertion point for field dependant name
			if field == "0" {
				return "Name"
			}
		case int, int64:
			// insertion point for field dependant name
			if field == 1 {
				return "EventNb"
			}
			if field == 2 {
				return "NbOfCollision"
			}
			if field == 5 {
				return "SimulationStep"
			}
			if field == 7 {
				return "NbLadybugs"
			}
			if field == 8 {
				return "NbLadybugsOnTheGround"
			}
		case float64:
			// insertion point for field dependant name
			if field == 3.000000 {
				return "LadybugRadius"
			}
			if field == 4.000000 {
				return "AbsoluteSpeed"
			}
			if field == 6.000000 {
				return "MaxDistanceInOneStep"
			}
			if field == 9.000000 {
				return "LeftRelayInitialPosX"
			}
			if field == 10.000000 {
				return "RightRelayInitialPosX"
			}
		case time.Time:
			// insertion point for field dependant name
		case bool:
			// insertion point for field dependant name
		}
	default:
		return ""
	}
	_ = field
	return ""
}
