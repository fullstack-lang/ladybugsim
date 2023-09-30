// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for LadybugStatus
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (ladybugstatus LadybugStatus) ToString() (res string) {

	// migration of former implementation of enum
	switch ladybugstatus {
	// insertion code per enum code
	case ON_THE_FENCE:
		res = "ON_THE_FENCE"
	case ON_THE_GROUND:
		res = "ON_THE_GROUND"
	}
	return
}

func (ladybugstatus *LadybugStatus) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ON_THE_FENCE":
		*ladybugstatus = ON_THE_FENCE
	case "ON_THE_GROUND":
		*ladybugstatus = ON_THE_GROUND
	default:
		return errUnkownEnum
	}
	return
}

func (ladybugstatus *LadybugStatus) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ON_THE_FENCE":
		*ladybugstatus = ON_THE_FENCE
	case "ON_THE_GROUND":
		*ladybugstatus = ON_THE_GROUND
	default:
		return errUnkownEnum
	}
	return
}

func (ladybugstatus *LadybugStatus) ToCodeString() (res string) {

	switch *ladybugstatus {
	// insertion code per enum code
	case ON_THE_FENCE:
		res = "ON_THE_FENCE"
	case ON_THE_GROUND:
		res = "ON_THE_GROUND"
	}
	return
}

func (ladybugstatus LadybugStatus) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ON_THE_FENCE")
	res = append(res, "ON_THE_GROUND")

	return
}

func (ladybugstatus LadybugStatus) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ON_THE_FENCE")
	res = append(res, "ON_THE_GROUND")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | LadybugStatus
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*LadybugStatus
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
