// insertion point for imports
import { AnimateDB } from './animate-db'
import { SVGDB } from './svg-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PolylineDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Points: string = ""
	Color: string = ""
	FillOpacity: number = 0
	Stroke: string = ""
	StrokeWidth: number = 0
	StrokeDashArray: string = ""
	Transform: string = ""

	// insertion point for other declarations
	Animates?: Array<AnimateDB>
	SVG_PolylinesDBID: NullInt64 = new NullInt64
	SVG_PolylinesDBID_Index: NullInt64  = new NullInt64 // store the index of the polyline instance in SVG.Polylines
	SVG_Polylines_reverse?: SVGDB 

}
