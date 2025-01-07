package physic

import (
	"math"
	"time"
)

type State struct {
	Time time.Time
	Obj  []Object
}

func (S State) TimeElapsed() float64 {
	out := float64(time.Now().UnixMilli() - S.Time.UnixMilli())
	return (out) / (1000 * Const.TimeSlow)
}

func (S *State) ColisionMap() map[int][]Object {
	if len(S.Obj) < 2 {
		return map[int][]Object{}
	}
	out := make(map[int][]Object)
	meta := map[int][]int{}
	for j, A := range S.Obj[:len(S.Obj)-1] {
		for k := 1; k < len(S.Obj)-j; k++ {
			B := S.Obj[j+k]
			if ObjectColide(A, B) {
				out[j] = append(out[j], B)
				out[j+k] = append(out[j+k], A)
				meta[j] = append(meta[j], j+k)
				meta[j+k] = append(meta[j+k], j)
			}
		}
	}
	for k, it := range meta {
		S.Obj[k].SetMetaData("Colide with:", it)
	}

	return out
}

func (S *State) UpdateState(Input UserInput, Floor func(x float64) float64, ScreenTransform func(x, y float64) (int, int)) {
	delay := math.Min(S.TimeElapsed(), Const.MaxTimeDelay)
	Colision := S.ColisionMap()
	for k, obj := range S.Obj {
		if k == 0 {
			obj.PFD(Floor, Input, Colision[k], delay)
		} else {
			obj.PFD(Floor, UserInput{}, Colision[k], delay)
		}
	}
}
func (S *State) ScreenCoordFromTransform(ScreenTransform func(float64, float64) (int, int)) {
	for _, obj := range S.Obj {
		obj.ScreenCoordFromTransform(ScreenTransform)
	}
}

var CurrentState = State{Time: time.Now(), Obj: []Object{NewObject(500, 0, 5, 4500)}}
