// insertion point for imports
import { AnimateDB } from './animate-db'
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class EllipseDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CX: number = 0
	CY: number = 0
	RX: number = 0
	RY: number = 0
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	SVG_EllipsesDBID: NullInt64 = new NullInt64
	SVG_EllipsesDBID_Index: NullInt64  = new NullInt64 // store the index of the ellipse instance in SVG.Ellipses
	SVG_Ellipses_reverse?: SVGDB 

}
