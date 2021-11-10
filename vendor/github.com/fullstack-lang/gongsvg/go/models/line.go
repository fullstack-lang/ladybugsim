package models

type Line struct {
	Name           string
	X1, Y1, X2, Y2 float64
	Presentation

	Animates []*Animate
}
