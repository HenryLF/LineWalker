package vect

import "math"

type Vect struct {
	X, Y float64
}
type VectInt struct {
	X, Y int
}

func (v Vect) Multiply(k float64) Vect {
	return Vect{X: v.X * k, Y: v.Y * k}
}
func (v Vect) Add(A ...Vect) Vect {
	for _, i := range A {
		v.X += i.X
		v.Y += i.Y
	}
	return v
}
func (v Vect) Norm() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vect) Unit() Vect {
	n := v.Norm()
	return v.Multiply(1 / n)
}

func (v Vect) To(A Vect) Vect {
	return A.Add(v.Multiply(-1))
}

func (v Vect) Rot() Vect {
	return Vect{X: -v.Y, Y: -v.X}
}

func (v Vect) Apply(v2 Vect, delay float64) Vect {
	return Vect{X: v.X + v2.X*delay, Y: v.Y + v2.Y*delay}
}

func Dist(A, B Vect) float64 {
	return A.To(B).Norm()
}
