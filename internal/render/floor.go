package render

import (
	"image"
	"image/color"
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

func (P PlayerView) rect() image.Rectangle {
	return image.Rect(0, 0, P.Width, P.Height)
}

func NewImage(P PlayerView) *image.RGBA {
	return image.NewRGBA(P.rect())
}

var FloorColor = color.RGBA{0, 0, 0, 255}
var BackColor = color.RGBA{120, 120, 120, 255}

const FloorWidth = 10

func DrawFloor(im *image.RGBA, P PlayerView, Floor func(float64) float64) {
	for x := range P.Width {
		for y := range P.Height {
			im.SetRGBA(x, y, BackColor)
		}
		y := Floor(float64(x+P.X)) + float64(P.Y)
		for k := range FloorWidth / 2 {
			im.SetRGBA(x, int(y)+k, FloorColor)
			im.SetRGBA(x, int(y)-k, FloorColor)
		}
	}
}
