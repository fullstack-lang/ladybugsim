// insertion point for imports
import { LadybugDB } from './ladybug-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LadybugSimulationDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EventNb: number = 0
	NbOfCollision: number = 0
	LadybugRadius: number = 0
	AbsoluteSpeed: number = 0
	SimulationStep: number = 0
	MaxDistanceInOneStep: number = 0
	NbLadybugs: number = 0
	NbLadybugsOnTheGround: number = 0
	LeftRelayInitialPosX: number = 0
	RightRelayInitialPosX: number = 0

	// insertion point for other declarations
	SimulationStep_string?: string
	Ladybugs?: Array<LadybugDB>
}
