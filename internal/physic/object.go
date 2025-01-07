package physic

import (
	"math"
)

type ObjectMetaData map[string]any

type Object struct {
	Coord *Vect
	Speed *Vect

	M, R float64

	ScreenCoord *VectInt

	Meta *ObjectMetaData
}

func (Obj Object) X() float64 {
	return Obj.Coord.X
}

func (Obj Object) Y() float64 {
	return Obj.Coord.Y
}

func (Obj *Object) ScreenCoordFromTransform(t func(x, y float64) (int, int)) {
	x, y := t(Obj.Coord.X, Obj.Coord.Y)
	*Obj.ScreenCoord = VectInt{X: x, Y: y}
}

func (Obj *Object) SetMetaData(s string, v any) {
	(*Obj.Meta)[s] = v
}

func NewObject(X, Y, M, R float64) *Object {
	var out Object
	out.Coord = new(Vect)
	*(out.Coord) = Vect{X: X, Y: Y}
	out.Speed = new(Vect)
	out.ScreenCoord = new(VectInt)
	out.M = M
	out.R = R
	meta := make(ObjectMetaData)
	out.Meta = &meta
	(*out.Meta)["Created"] = true
	return &out
}

func (A Object) Colide(B Object) bool {
	return dist(*A.Coord, *B.Coord) <= (A.R + B.R)
}

func (Obj *Object) PFD(Floor func(float64) float64, Input UserInput, Colision []ObjectWithColision, delay float64) {
	delay = math.Min(delay, Const.MaxTimeDelay)
	Obj.SetMetaData("delay", delay)
	grounded := contact(*Obj, Floor)
	Obj.SetMetaData("wasGrounded", grounded)
	ResultingForce := gravityForce(*Obj)
	ResultingForce = ResultingForce.add(frictionForce(*Obj, grounded))
	ResultingForce = ResultingForce.add(reactiveForce(Floor, *Obj, ResultingForce))
	ResultingForce = ResultingForce.add(movementForce(*Obj, Input, Floor, grounded))

	if A, ok := interface{}(Obj).(Object); len(Colision) > 0 && ok {
		for _, colider := range Colision {
			B, ok := interface{}(colider).(Object)
			if ok {
				k := colisionForce(A, B)
				ResultingForce = ResultingForce.add(k)
			}
		}
	} else {
		Obj.SetMetaData("Colision", false)
	}
	ResultingForce = ResultingForce.multiply(1 / Obj.M)
	//Cap
	newSpeed := Obj.Speed.apply(ResultingForce, delay)
	newSpeedN := newSpeed.norm()
	if newSpeedN > Const.CapSpeed {
		Obj.SetMetaData("SpeedCaped", true)
		newSpeed = newSpeed.unit().multiply(Const.CapSpeed)
	} else {
		Obj.SetMetaData("SpeedCaped", false)
	}
	*(Obj.Speed) = newSpeed
	*(Obj.Coord) = Obj.Coord.apply(*(Obj.Speed), delay)
	grounded = contact(*Obj, Floor)
	if grounded {
		Obj.ground(Floor)
	}
	Obj.SetMetaData("isGrounded", grounded)
}
