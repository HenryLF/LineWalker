package worldmap

import (
	"log"
	"math/rand"
	"reflect"

	"github.com/aquilax/go-perlin"
)

func (C *WorldMap) Set(s string, a float64) any {
	c := reflect.ValueOf(C).Elem().FieldByName(s)
	if c.CanSet() {
		log.Println("Change ", s, " from ", c, " to ", a)
		c.SetFloat(a)
		if s == "A" || s == "B" {
			C.Generator = perlin.NewPerlin(C.A, C.B, dn, rand.Int63())
		}
		return a
	}
	log.Println("Error setting", s)
	return C.Get(s)
}
func (C *WorldMap) Get(s string) float64 {
	return reflect.ValueOf(C).Elem().FieldByName(s).Float()
}
