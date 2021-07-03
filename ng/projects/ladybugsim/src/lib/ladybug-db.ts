// insertion point for imports
import { LadybugSimulationDB } from './ladybugsimulation-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LadybugDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	TechName?: string
	Name?: string
	Id?: number
	Position?: number
	Speed?: number
	LadybugStatus?: string

	// insertion point for other declarations
	LadybugSimulation_LadybugsDBID?: NullInt64
	LadybugSimulation_LadybugsDBID_Index?: NullInt64 // store the index of the ladybug instance in LadybugSimulation.Ladybugs
	LadybugSimulation_Ladybugs_reverse?: LadybugSimulationDB

}
