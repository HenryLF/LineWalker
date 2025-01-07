package physic

import (
	"encoding/json"
	"log"
)

type CustomObject = Object

func (S *State) Set(n int, s string, c ...json.Number) any {
	if n < 0 || n > len(S.Obj) || len(c) == 0 {
		return nil
	}
	Obj, ok := interface{}(S.Obj[n]).(CustomObject)
	if !ok {
		return nil
	}
	a := []float64{}
	for _, C := range c {
		A, err := C.Float64()
		if err != nil {
			log.Printf("Error setting %v to %v.\n", s, C)
			return S.Get(n, s)
		}
		a = append(a, A)
	}
	v := new(Vect)
	switch s {
	case "Coord":
		if len(a) > 1 {
			v.X = a[0]
			v.Y = a[1]
			Obj.Coord = v
		}
	case "Speed":
		if len(a) > 1 {
			v.X = a[0]
			v.Y = a[1]
			Obj.Speed = v
		}
	case "M":
		Obj.M = a[0]
	case "R":
		Obj.R = a[1]
	}
	return S.Get(n, s)
}

func (S *State) SetPlayer(s string, a json.Number) any {
	c, err := a.Float64()
	if err != nil || c == 0 {
		log.Printf("Error setting %v with value %v\n", s, a)
	}
	Obj, ok := interface{}(S.Obj).(CustomObject)
	if !ok {
		return nil
	}
	log.Printf("Setting Player %v from %v to %v\n", s, S.GetPlayer(s), c)
	switch s {
	case "X":
		Obj.Coord.X = c
	case "Y":
		Obj.Coord.X = c
	case "M":
		Obj.M = c
	case "R":
		Obj.R = c
	default:
		log.Println("Error setting", s)
	}
	return S.GetPlayer(s)
}

func (S *State) Get(n int, s string) any {
	if n < 0 || n > len(S.Obj) {
		return nil
	}
	Obj, ok := interface{}(S.Obj[n]).(CustomObject)
	if !ok {
		return nil
	}

	switch s {
	case "Coord":
		return Obj.Coord
	case "Speed":
		return Obj.Speed
	case "M":
		return Obj.M
	case "R":
		return Obj.R
	}
	return nil
}
func (S *State) GetPlayer(s string) any {
	Obj, ok := interface{}(S.Obj[0]).(CustomObject)
	if !ok {
		return nil
	}
	switch s {
	case "X":
		return Obj.Coord.X
	case "Y":
		return Obj.Coord.X
	case "M":
		return Obj.M
	case "R":
		return Obj.R
	}
	log.Println("No property ", s)
	return nil
}

func (S *State) AddObject(X, Y, M, R float64) bool {
	k := NewObject(X, Y, M, R)
	S.Obj = append(S.Obj, k)
	log.Println("New Object", k, S)

	return true
}

func (C *Constants) Set(s string, c json.Number) any {
	a, err := c.Float64()
	if err != nil || a == 0 {
		log.Printf("Error setting %v with value %v\n", s, c)
	}
	log.Printf("Setting %v from %v to %v\n", s, C.Get(s), a)

	switch s {
	case "G":
		C.G = a
	case "AirFrictionCoeff":
		C.AirFrictionCoeff = a
	case "FloorFrictionCoeff":
		C.FloorFrictionCoeff = a
	case "DX":
		C.DX = a
	case "LateralAcc":
		C.LateralAcc = a
	case "LateralAirAcc":
		C.LateralAirAcc = a
	case "VerticalAcc":
		C.VerticalAcc = a
	case "VerticalAccDown":
		C.VerticalAccDown = a
	case "CapSpeed":
		C.CapSpeed = a
	case "TimeSlow":
		C.TimeSlow = a
	case "MaxTimeDelay":
		C.MaxTimeDelay = a
	case "ElasticColision":
		C.ElasticColision = a
	default:
		log.Println("Error setting", s)
	}
	return C.Get(s)
}

func (C *Constants) Get(s string) any {
	switch s {
	case "G":
		return C.G
	case "AirFrictionCoeff":
		return C.AirFrictionCoeff
	case "FloorFrictionCoeff":
		return C.FloorFrictionCoeff
	case "DX":
		return C.DX
	case "LateralAcc":
		return C.LateralAcc
	case "LateralAirAcc":
		return C.LateralAirAcc
	case "VerticalAcc":
		return C.VerticalAcc
	case "VerticalAccDown":
		return C.VerticalAccDown
	case "CapSpeed":
		return C.CapSpeed
	case "TimeSlow":
		return C.TimeSlow
	case "MaxTimeDelay":
		return C.MaxTimeDelay
	case "ElasticColision":
		return C.ElasticColision
	}
	log.Println("No property ", s)
	return nil
}
