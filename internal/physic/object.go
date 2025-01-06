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
	Coord *Vect
	Speed *Vect

	M float64
	R float64
}

func NewObject(X, Y, M, R float64) Object {
	var out Object
	out.Coord = new(Vect)
	*(out.Coord) = Vect{X: X, Y: Y}
	out.Speed = new(Vect)
	out.M = M
	out.R = R
	return out
}

type Vect struct {
	X, Y float64
}

func (v Vect) multiply(k float64) Vect {
	return Vect{X: v.X * k, Y: v.Y * k}
}
func (v Vect) add(A ...Vect) Vect {
	for _, i := range A {
		v.X += i.X
		v.Y += i.Y
	}
	return v
}
func (v Vect) norm() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vect) rot() Vect {
	return Vect{X: -v.Y, Y: -v.X}
}

func (v Vect) apply(v2 Vect, delay float64) Vect {
	return Vect{X: v.X + v2.X*delay, Y: v.Y + v2.Y*delay}
}
