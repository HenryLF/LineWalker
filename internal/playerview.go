package bindings

import "linewalker/internal/physic"

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

const A = 1
const B = 1

func (P PlayerView) ScreenTransform(x, y float64) (int, int) {
	return int(x)*A - P.X, int(y)*B - P.Y
}

func (P *PlayerView) Center(Obj physic.Object) {
	P.X = int(Obj.Coord.X) - P.Width/2
	P.Y = int(Obj.Coord.Y) - P.Height/2
}
