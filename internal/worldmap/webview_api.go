package worldmap

import (
	"encoding/json"
	"log"
	"math/rand"

	"github.com/aquilax/go-perlin"
)

func (C *WorldMap) Set(s string, a json.Number) any {
	c, err := a.Float64()
	if err != nil || c == 0 {
		log.Printf("Error setting %v with value %v\n", s, a)
	}
	log.Printf("Setting %v from %v to %v\n", s, C.Get(s), c)

	switch s {
	case "A":
		C.A = c
		C.Generator = perlin.NewPerlin(C.A, C.B, dn, rand.Int63())
	case "B":
		C.B = c
		C.Generator = perlin.NewPerlin(C.A, C.B, dn, rand.Int63())
	case "Y0":
		C.Y0 = c
	case "ScaleX":
		C.ScaleX = c
	case "ScaleY":
		C.ScaleY = c
	default:
		log.Println("Error setting", s)
	}
	return C.Get(s)
}

func (C *WorldMap) Get(s string) any {
	switch s {
	case "A":
		return C.A
	case "B":
		return C.B
	case "Y0":
		return C.Y0
	case "ScaleX":
		return C.ScaleX
	case "ScaleY":
		return C.ScaleY
	}
	log.Println("No property ", s)
	return nil
}
