package physic

import (
	"linewalker/internal/physic/vect"
	"math"
)

//Implementation of Object for 2D sidescroller

type ObjectSide struct {
	Coord *vect.Vect
	Speed *vect.Vect

	M, R float64

	ScreenCoord *vect.VectInt

	Meta *ObjectMetaData
}

func (Obj ObjectSide) X() float64 {
	return Obj.Coord.X
}
func (Obj ObjectSide) Y() float64 {
	return Obj.Coord.Y
}
func (Obj ObjectSide) Mass() float64 {
	return Obj.M
}
func (Obj ObjectSide) Radius() float64 {
	return Obj.R
}

func (Obj *ObjectSide) ScreenCoordFromTransform(t func(x, y float64) (int, int)) {
	x, y := t(Obj.Coord.X, Obj.Coord.Y)
	*Obj.ScreenCoord = vect.VectInt{X: x, Y: y}
}

func (Obj *ObjectSide) SetMetaData(s string, v any) {
	(*Obj.Meta)[s] = v
}

func (A ObjectSide) Colide(B ObjectSide) bool {
	return vect.Dist(*A.Coord, *B.Coord) <= (A.R + B.R)
}

func (Obj *ObjectSide) PFD(Floor func(float64) float64, Input UserInput, Colision []ObjectWithColision, delay float64) {
	delay = math.Min(delay, Const.MaxTimeDelay)
	Obj.SetMetaData("delay", delay)
	grounded := contact(*Obj, Floor)
	Obj.SetMetaData("wasGrounded", grounded)
	ResultingForce := gravityForce(*Obj)
	ResultingForce = ResultingForce.Add(frictionForce(*Obj, grounded))
	ResultingForce = ResultingForce.Add(reactiveForce(Floor, *Obj, ResultingForce))
	ResultingForce = ResultingForce.Add(movementForce(*Obj, Input, Floor, grounded))

	if A, ok := interface{}(Obj).(ObjectSide); len(Colision) > 0 && ok {
		for _, colider := range Colision {
			B, ok := interface{}(colider).(ObjectSide)
			if ok {
				k := colisionForce(A, B)
				ResultingForce = ResultingForce.Add(k)
			}
		}
	} else {
		Obj.SetMetaData("Colision", false)
	}
	ResultingForce = ResultingForce.Multiply(1 / Obj.M)
	//Cap
	newSpeed := Obj.Speed.Apply(ResultingForce, delay)
	newSpeedN := newSpeed.Norm()
	if newSpeedN > Const.CapSpeed {
		Obj.SetMetaData("SpeedCaped", true)
		newSpeed = newSpeed.Unit().Multiply(Const.CapSpeed)
	} else {
		Obj.SetMetaData("SpeedCaped", false)
	}
	*(Obj.Speed) = newSpeed
	*(Obj.Coord) = Obj.Coord.Apply(*(Obj.Speed), delay)
	grounded = contact(*Obj, Floor)
	if grounded {
		Obj.ground(Floor)
	}
	Obj.SetMetaData("isGrounded", grounded)
}

func NewObjectSide(X, Y, M, R float64) *ObjectSide {
	var out ObjectSide
	out.Coord = new(vect.Vect)
	*(out.Coord) = vect.Vect{X: X, Y: Y}
	out.Speed = new(vect.Vect)
	out.ScreenCoord = new(vect.VectInt)
	out.M = M
	out.R = R
	meta := make(ObjectMetaData)
	out.Meta = &meta
	(*out.Meta)["Created"] = true
	return &out
}
