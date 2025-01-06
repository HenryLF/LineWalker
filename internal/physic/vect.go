package physic

import "math"

type Vect struct {
	X, Y float64
}
type VectInt struct {
	X, Y int
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

func (v Vect) unit() Vect {
	n := v.norm()
	return v.multiply(1 / n)
}

func (v Vect) to(A Vect) Vect {
	return A.add(v.multiply(-1))
}

func (v Vect) rot() Vect {
	return Vect{X: -v.Y, Y: -v.X}
}

func (v Vect) apply(v2 Vect, delay float64) Vect {
	return Vect{X: v.X + v2.X*delay, Y: v.Y + v2.Y*delay}
}

func dist(A, B Vect) float64 {
	return A.to(B).norm()
}
