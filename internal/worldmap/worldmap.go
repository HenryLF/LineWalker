package worldmap

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"reflect"

	"github.com/aquilax/go-perlin"
)

const b = 2
const a = 2
const n = 2

const ScaleX float64 = 200
const ScaleY float64 = 200

type WorldMap struct {
	A, B, Y0, ScaleX, ScaleY float64
	Generator                *perlin.Perlin
}

func (C *WorldMap) Set(s string, a float64) any {
	c := reflect.ValueOf(C).Elem().FieldByName(s)
	if c.CanSet() {
		log.Println("Change ", s, " from ", c, " to ", a)
		c.SetFloat(a)
		if s == "A" || s == "B" {
			C.Generator = perlin.NewPerlin(C.A, C.B, n, rand.Int63())
		}
		return a
	}
	log.Println("Error setting", s)
	return errors.New("trying to set unadressable field")
}
func (C *WorldMap) Get(s string) float64 {
	return reflect.ValueOf(C).Elem().FieldByName(s).Float()
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

func FlatFloor(c float64) func(float64) float64 {
	return func(x float64) float64 { return c }
}

func SlopeFloor(a, b float64) func(float64) float64 {
	return func(x float64) float64 { return a*x + b }
}
