package physic

import (
	"math"
	"time"
)

type UserInput struct {
	Up    bool `json:"Up"`
	Left  bool `json:"Left"`
	Down  bool `json:"Down"`
	Right bool `json:"Right"`
}

type ObjectInterface interface {
	PFD(Floor func(float64) float64, Input UserInput, Colision []ObjectWithColision, delay float64)
	ScreenCoordFromTransform(func(x, y float64) (a, b int))
	SetMetaData(string, any)
	X() float64
	Y() float64
}

type ObjectWithColision interface {
	PFD(Floor func(float64) float64, Input UserInput, Colision []ObjectWithColision, delay float64)
	ScreenCoordFromTransform(func(x, y float64) (a, b int))
	SetMetaData(string, any)
	Colide(ObjectWithColision) bool
}

type State struct {
	Time time.Time
	Obj  []ObjectInterface
}

func (S State) TimeElapsed() float64 {
	out := float64(time.Now().UnixMilli() - S.Time.UnixMilli())
	return (out) / (1000 * Const.TimeSlow)
}

func (S *State) ColisionMap() map[int][]ObjectWithColision {
	if len(S.Obj) < 2 {
		return map[int][]ObjectWithColision{}
	}
	out := make(map[int][]ObjectWithColision)
	meta := map[int][]int{}
	for j, A := range S.Obj[:len(S.Obj)-1] {
		A, ok := interface{}(A).(ObjectWithColision)
		if !ok {
			continue
		}
		for k := 1; k < len(S.Obj)-j; k++ {
			B, ok := interface{}(S.Obj[j+k]).(ObjectWithColision)
			if !ok {
				continue
			}
			if A.Colide(B) {
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

var CurrentState = State{Time: time.Now(), Obj: []ObjectInterface{NewObject(500, 0, 5, 4500)}}
