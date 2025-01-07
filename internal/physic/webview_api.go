package physic

import (
	"encoding/json"
	"log"
)

func (S *State) Set(n int, s string, c ...json.Number) any {
	if n < 0 || n > len(S.Obj) || len(c) == 0 {
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
			S.Obj[n].Coord = v
		}
	case "Speed":
		if len(a) > 1 {
			v.X = a[0]
			v.Y = a[1]
			S.Obj[n].Speed = v
		}
	case "M":
		S.Obj[n].M = a[0]
	case "R":
		S.Obj[n].R = a[1]
	}
	return S.Get(n, s)
}

func (S *State) SetPlayer(s string, a json.Number) any {
	c, err := a.Float64()
	if err != nil || c == 0 {
		log.Printf("Error setting %v with value %v\n", s, a)
	}
	log.Printf("Setting Player %v from %v to %v\n", s, S.GetPlayer(s), c)
	switch s {
	case "X":
		S.Obj[0].Coord.X = c
	case "Y":
		S.Obj[0].Coord.X = c
	case "M":
		S.Obj[0].M = c
	case "R":
		S.Obj[0].R = c
	default:
		log.Println("Error setting", s)
	}
	return S.GetPlayer(s)
}

func (S *State) Get(n int, s string) any {
	if n < 0 || n > len(S.Obj) {
		return nil
	}
	k := S.Obj[n]
	switch s {
	case "Coord":
		return *k.Coord
	case "Speed":
		return *k.Speed
	case "M":
		return k.M
	case "R":
		return k.R
	}
	return nil
}
func (S *State) GetPlayer(s string) any {
	switch s {
	case "X":
		return S.Obj[0].Coord.X
	case "Y":
		return S.Obj[0].Coord.X
	case "M":
		return S.Obj[0].M
	case "R":
		return S.Obj[0].R
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
