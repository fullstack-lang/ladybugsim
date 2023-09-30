// insertion point for imports
import { LadybugSimulationDB } from './ladybugsimulation-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LadybugDB {

	static GONGSTRUCT_NAME = "Ladybug"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Id: number = 0
	Position: number = 0
	Speed: number = 0
	LadybugStatus: string = ""

	// insertion point for other declarations
	LadybugSimulation_LadybugsDBID: NullInt64 = new NullInt64
	LadybugSimulation_LadybugsDBID_Index: NullInt64  = new NullInt64 // store the index of the ladybug instance in LadybugSimulation.Ladybugs
	LadybugSimulation_Ladybugs_reverse?: LadybugSimulationDB 

}
