// insertion point for imports
import { LadybugDB } from './ladybug-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LadybugSimulationDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	EventNb?: number
	NbOfCollision?: number
	LadybugRadius?: number
	AbsoluteSpeed?: number
	SimulationStep?: number
	MaxDistanceInOneStep?: number
	NbLadybugs?: number

	// insertion point for other declarations
	SimulationStep_string?: string
	Ladybugs?: Array<LadybugDB>
}
