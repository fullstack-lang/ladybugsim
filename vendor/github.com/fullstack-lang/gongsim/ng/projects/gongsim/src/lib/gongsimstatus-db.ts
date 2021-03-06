// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongsimStatusDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CurrentCommand: string = ""
	CompletionDate: string = ""
	CurrentSpeedCommand: string = ""
	SpeedCommandCompletionDate: string = ""

	// insertion point for other declarations
}
