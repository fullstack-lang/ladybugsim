// insertion point for imports
import { RectDB } from './rect-db'
import { TextDB } from './text-db'
import { CircleDB } from './circle-db'
import { LineDB } from './line-db'
import { EllipseDB } from './ellipse-db'
import { PolylineDB } from './polyline-db'
import { PolygoneDB } from './polygone-db'
import { PathDB } from './path-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SVGDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Display: boolean = false
	Name: string = ""

	// insertion point for other declarations
	Rects?: Array<RectDB>
	Texts?: Array<TextDB>
	Circles?: Array<CircleDB>
	Lines?: Array<LineDB>
	Ellipses?: Array<EllipseDB>
	Polylines?: Array<PolylineDB>
	Polygones?: Array<PolygoneDB>
	Paths?: Array<PathDB>
}
