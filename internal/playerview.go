package bindings

import (
	"linewalker/internal/physic"
	"log"
)

type PlayerView struct {
	X, Y, Width, Height int
	ScaleX, ScaleY      int
}

func (P *PlayerView) SetSize(w, h int) {
	P.Width = w
	P.Height = h
}

func (P *PlayerView) SetCoord(x, y int) {
	P.X = x
	P.Y = y
}

const dScaleX int = 1
const dScaleY int = 1

func (P PlayerView) ScreenTransform(x, y float64) (int, int) {
	return (int(x) - P.X) / P.ScaleX, (int(y) - P.Y) / P.ScaleY
}
func (P PlayerView) CoordTransformX(x int) float64 {
	return float64(x*P.ScaleX) + float64(P.X)
}
func (P PlayerView) ScreenTransformY(y float64) int {
	return (int(y) - P.Y) / P.ScaleY
}

func (P *PlayerView) Center(Obj physic.Object) {
	P.X = int(Obj.Coord.X) - P.ScaleX*P.Width/2
	P.Y = int(Obj.Coord.Y) - P.ScaleY*P.Height/2
}
func (P *PlayerView) Set(s string, a int) int {
	switch s {
	case "X":
		P.X = a
	case "Y":
		P.Y = a
	case "Width":
		P.Width = a
	case "Height":
		P.Height = a
	case "ScaleX":
		P.ScaleX = a
	case "ScaleY":
		P.ScaleY = a
	}
	log.Println("Error setting", s)
	return P.Get(s)
}
func (P *PlayerView) Get(s string) int {
	switch s {
	case "Width":
		return P.Width
	case "Height":
		return P.Height
	case "ScaleX":
		return P.ScaleX
	case "ScaleY":
		return P.ScaleY
	}
	log.Println("Error getting", s)
	return 0
}

var CurrentView = PlayerView{
	X:      0,
	Y:      0,
	Width:  300,
	Height: 300,
	ScaleX: dScaleX,
	ScaleY: dScaleY}
