package physic

import (
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

func (S State) ColisionMap() map[int][]Object {
	if len(S.Obj) < 2 {
		return map[int][]Object{}
	}
	out := make(map[int][]Object)
	for j, A := range S.Obj[:len(S.Obj)-1] {
		for k := 1; k < len(S.Obj)-j; k++ {
			B := S.Obj[j+k]
			if ObjectColide(A, B) {
				out[j] = append(out[j], B)
				out[j+k] = append(out[j+k], A)
			}
		}
	}
	return out
}

var CurrentState = State{Time: time.Now(), Obj: []Object{NewObject(500, 0, 5, 20)}}
