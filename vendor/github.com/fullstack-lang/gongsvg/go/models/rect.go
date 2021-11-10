package models

type Rect struct {
	Name                    string
	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate
}
