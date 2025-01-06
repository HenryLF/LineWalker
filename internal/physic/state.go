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

func (S State) Colide() map[int][]Object {
	out := make(map[int][]Object)
	return out
}

var CurrentState = State{Time: time.Now(), Obj: []Object{NewObject(500, 0, 5, 20)}}
