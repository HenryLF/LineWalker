package physic

import (
	"math"
)

type UserInput struct {
	Up    bool `json:"Up"`
	Left  bool `json:"Left"`
	Down  bool `json:"Down"`
	Right bool `json:"Right"`
}

type Object struct {
	Coord *vect
	Speed *vect

	M float64
	R float64
}

func NewObject(X, Y, R, M float64) Object {
	var out Object
	out.Coord = new(vect)
	*(out.Coord) = vect{X: X, Y: Y}
	out.Speed = new(vect)
	out.M = M
	out.R = R
	return out
}

type vect struct {
	X, Y float64
}

func (v vect) multiply(k float64) vect {
	return vect{X: v.X * k, Y: v.Y * k}
}
func (v vect) add(A ...vect) vect {
	for _, i := range A {
		v.X += i.X
		v.Y += i.Y
	}
	return v
}
func (v vect) norm() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v vect) rot() vect {
	return vect{X: -v.Y, Y: -v.X}
}

func (v vect) apply(v2 vect, delay float64) vect {
	return vect{X: v.X + v2.X*delay, Y: v.Y + v2.Y*delay}
}
