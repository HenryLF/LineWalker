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

type ObjectMetaData map[string]any

type Object struct {
	Coord *Vect
	Speed *Vect

	M, R float64

	ScreenCoord *VectInt

	meta *ObjectMetaData
}

func (Obj *Object) SetScreenCoord(X, Y int) {
	*Obj.ScreenCoord = VectInt{X: int(Obj.Coord.X) - X, Y: int(Obj.Coord.Y) - Y}
}
func (Obj *Object) ScreenCoordFromTransform(t func(x, y float64) (int, int)) {
	x, y := t(Obj.Coord.X, Obj.Coord.Y)
	*Obj.ScreenCoord = VectInt{X: x, Y: y}
}

func (Obj *Object) SetMetaData(s string, v any) {
	(*Obj.meta)[s] = v
}

func NewObject(X, Y, M, R float64) Object {
	var out Object
	out.Coord = new(Vect)
	*(out.Coord) = Vect{X: X, Y: Y}
	out.Speed = new(Vect)
	out.ScreenCoord = new(VectInt)
	out.M = M
	out.R = R
	meta := make(ObjectMetaData)
	out.meta = &meta
	(*out.meta)["Created"] = true
	return out
}

func ObjectColide(A, B Object) bool {
	return dist(*A.Coord, *B.Coord) <= (A.R + B.R)
}

func (Obj *Object) PFD(Floor func(float64) float64, Input UserInput, Colision []Object, delay float64) {
	delay = math.Min(delay, Const.MaxTimeDelay)
	Obj.SetMetaData("delay", delay)
	grounded := contact(*Obj, Floor)
	Obj.SetMetaData("wasGrounded", grounded)
	ResultingForce := gravityForce(*Obj)
	ResultingForce = ResultingForce.add(frictionForce(*Obj, grounded))
	ResultingForce = ResultingForce.add(movementForce(Input, grounded))
	ResultingForce = ResultingForce.add(reactiveForce(Floor, *Obj, ResultingForce))
	// col := Vect{}
	if len(Colision) > 0 {
		for _, colider := range Colision {
			k := colisionForce(*Obj, colider)
			ResultingForce = ResultingForce.add(k)
		}
	}

	ResultingForce = ResultingForce.multiply(1 / Obj.M)
	//Cap
	newSpeed := Obj.Speed.apply(ResultingForce, delay)
	newSpeedN := newSpeed.norm()
	if newSpeedN > Const.CapSpeed {
		Obj.SetMetaData("SpeedCaped", true)
		newSpeed = newSpeed.unit().multiply(Const.CapSpeed)
	}
	*(Obj.Speed) = newSpeed
	*(Obj.Coord) = Obj.Coord.apply(*(Obj.Speed), delay)
	grounded = contact(*Obj, Floor)
	if grounded {
		Obj.ground(Floor)
	}
	Obj.SetMetaData("isGrounded", false)
}
