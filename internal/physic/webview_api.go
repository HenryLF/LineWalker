package physic

import (
	"log"
	"reflect"
)

func (S *State) Set(n int, s string, a ...float64) any {
	if n < 0 || n > len(S.Obj) || len(a) == 0 {
		return nil
	}
	v := new(Vect)
	switch s {
	case "Coord":
		if len(a) > 1 {
			v.X = a[0]
			v.Y = a[1]
			S.Obj[n].Coord = v
		}
		return *S.Obj[n].Coord
	case "Speed":
		if len(a) > 1 {
			v.X = a[0]
			v.Y = a[1]
			S.Obj[n].Speed = v
		}
		return *S.Obj[n].Speed
	case "M":
		S.Obj[n].M = a[0]
		return S.Obj[n].M
	case "R":
		S.Obj[n].R = a[1]
		return S.Obj[n].R
	}
	return nil
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

func (S *State) AddObject(X, Y, M, R float64) bool {
	k := NewObject(X, Y, M, R)
	S.Obj = append(S.Obj, k)
	log.Println("New Object", k, S)

	return true
}

func (C *Constants) Set(s string, a float64) any {
	c := reflect.ValueOf(C).Elem().FieldByName(s)
	if c.CanSet() {
		log.Println("Change ", s, " from ", c, " to ", a)
		c.SetFloat(a)
		return a
	}
	log.Println("Error setting", s)
	return C.Get(s)
}

func (C *Constants) Get(s string) float64 {
	return reflect.ValueOf(C).Elem().FieldByName(s).Float()
}
