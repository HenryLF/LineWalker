package worldmap

import (
	"math"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

const b = 2
const a = 2
const n = 2

const ScaleX float64 = 2000
const ScaleY float64 = 1500

type WorldMap struct {
	A, B, Y0, ScaleX, ScaleY float64
	Generator                *perlin.Perlin
	TransformCoordX          func(x int) float64
	TransformScreenY         func(y float64) int
}

type Response map[int]int

func newPerlin() WorldMap {
	var out = WorldMap{A: a, B: b, Y0: 0, ScaleX: ScaleX, ScaleY: ScaleY}
	out.Generator = perlin.NewPerlin(a, b, n, rand.Int63())
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
