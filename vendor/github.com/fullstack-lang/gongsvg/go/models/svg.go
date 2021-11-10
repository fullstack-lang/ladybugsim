package models

type SVG struct {
	Display bool

	Name string

	Rects     []*Rect
	Texts     []*Text
	Circles   []*Circle
	Lines     []*Line
	Ellipses  []*Ellipse
	Polylines []*Polyline
	Polygones []*Polygone
	Paths     []*Path
}
