package bindings

import (
	"linewalker/internal/physic"
)

type PlayerView struct {
	X, Y, Width, Height int
}

func (P *PlayerView) SetSize(w, h int) {
	P.Width = w
	P.Height = h
}

func (P *PlayerView) SetCoord(x, y int) {
	P.X = x
	P.Y = y
}

const A float64 = 2
const B float64 = 2

func (P PlayerView) ScreenTransform(x, y float64) (int, int) {
	return (int(x) - P.X) / int(A), (int(y) - P.Y) / int(B)
}
func (P PlayerView) CoordTransformX(x int) float64 {
	return float64(x)*A + float64(P.X)
}
func (P PlayerView) ScreenTransformY(y float64) int {
	return (int(y) - P.Y) / int(B)
}

func (P *PlayerView) Center(Obj physic.Object) {
	P.X = int(Obj.Coord.X) - int(A)*P.Width/2
	P.Y = int(Obj.Coord.Y) - int(B)*P.Height/2
}
