package bindings

import (
	"encoding/json"
	"linewalker/internal/physic"
	"log"
)

// Default value
const dScaleX int = 200
const dScaleY int = 200

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
func (P *PlayerView) Set(s string, a json.Number) any {
	c, err := a.Int64()
	if err != nil || c == 0 {
		log.Printf("Error setting %v with value %v\n", s, a)
	}
	log.Printf("Setting %v from %v to %v\n", s, P.Get(s), c)
	switch s {
	case "X":
		P.X = int(c)
	case "Y":
		P.Y = int(c)
	case "Width":
		P.Width = int(c)
	case "Height":
		P.Height = int(c)
	case "ScaleX":
		P.ScaleX = int(c)
	case "ScaleY":
		P.ScaleY = int(c)
	}
	log.Println("Error setting", s)
	return P.Get(s)
}
func (P *PlayerView) Get(s string) any {
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
	return nil
}

var CurrentView = PlayerView{
	X:      0,
	Y:      0,
	Width:  300,
	Height: 300,
	ScaleX: dScaleX,
	ScaleY: dScaleY}
