// insertion point for imports
import { AnimateDB } from './animate-db'
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LineDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	X1: number = 0
	Y1: number = 0
	X2: number = 0
	Y2: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	SVG_LinesDBID: NullInt64 = new NullInt64
	SVG_LinesDBID_Index: NullInt64  = new NullInt64 // store the index of the line instance in SVG.Lines
	SVG_Lines_reverse?: SVGDB 

}
