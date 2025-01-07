package worldmap

import (
	"math"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

const db = 2
const da = 2
const dn = 2

const dScaleX float64 = 2000
const dScaleY float64 = 1500

type WorldMap struct {
	A, B, Y0, ScaleX, ScaleY float64
	Generator                *perlin.Perlin
	TransformCoordX          func(x int) float64
	TransformScreenY         func(y float64) int
}

type Response map[int]int

func newPerlin() WorldMap {
	var out = WorldMap{A: da, B: db, Y0: 0, ScaleX: dScaleX, ScaleY: dScaleY}
	out.Generator = perlin.NewPerlin(da, db, dn, rand.Int63())
	return out
}

var CurrentMap = newPerlin()

func (M WorldMap) Generate(x float64) float64 {
	return M.ScaleY*math.Max(M.Generator.Noise1D(x/M.ScaleX)+1, 0) + M.Y0
}

func (M WorldMap) GenerateFromTransform(x int, CoordTransformX func(x int) float64, ScreenTransformY func(y float64) int) int {
	return ScreenTransformY(M.Generate(CoordTransformX(x)))

}

func FlatFloor(c float64) func(float64) float64 {
	return func(x float64) float64 { return c }
}

func SlopeFloor(a, b float64) func(float64) float64 {
	return func(x float64) float64 { return a*x + b }
}
