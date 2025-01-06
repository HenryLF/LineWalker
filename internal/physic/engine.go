package physic

import (
	"log"
	"math"
	"time"
)

var CurrentState = State{Time: time.Now(), Obj: []Object{NewObject(500, 0, 5, 50)}}

func gravityForce(Obj Object) Vect {
	return Vect{X: 0, Y: Obj.M * Const.G}
}

func reactiveForce(Floor func(float64) float64, Obj Object, R Vect) Vect {
	if contact(Obj, Floor) {
		ang := -angleOf(Floor, Obj.Coord.X)
		out := Vect{X: R.norm() * math.Cos(ang), Y: R.norm() * math.Sin(ang)}
		// log.Println(out, out.rot())
		return out.rot()
	} else {
		return Vect{X: 0, Y: 0}
	}
}

func frictionForce(Obj Object, contact bool) Vect {
	if contact {
		return (*Obj.Speed).multiply(-Const.FloorFrictionCoeff)
	}
	return (*Obj.Speed).multiply(-Const.AirFrictionCoeff)
}

func contact(Obj Object, Floor func(float64) float64) bool {
	return Obj.Coord.Y+Obj.R >= Floor(Obj.Coord.X)
}

func angleOf(Floor func(float64) float64, x float64) float64 {
	return math.Atan((Floor(x+Const.DX) - Floor(x-Const.DX)) / (2 * Const.DX))
}

func (Obj *Object) ground(Floor func(float64) float64) {
	Obj.Coord.Y = math.Min(Floor(Obj.Coord.X)-Obj.R, Obj.Coord.Y)
	Obj.Speed.Y = math.Min(Floor(Obj.Speed.Y), 0)
}

func movementForce(Input UserInput, grounded bool) Vect {
	out := Vect{}
	if Input.Up && grounded {
		out = out.add(Vect{X: 0, Y: -Const.VerticalAcc})
		return out
	}
	if Input.Right {
		if grounded {
			out = out.add(Vect{X: Const.LateralAcc, Y: 0})
		} else {
			out = out.add(Vect{X: Const.LateralAirAcc, Y: 0})
		}
	}
	if Input.Left && grounded {
		if grounded {
			out = out.add(Vect{X: -Const.LateralAcc, Y: 0})
		} else {
			out = out.add(Vect{X: -Const.LateralAirAcc, Y: 0})
		}
	}
	if Input.Down && !grounded {
		out = out.add(Vect{X: 0, Y: Const.VerticalAcc})
	}
	return out
}

func PFD(Obj *Object, Floor func(float64) float64, Input UserInput, delay float64) {
	grounded := contact(*Obj, Floor)
	ResultingForce := gravityForce(*Obj)
	ResultingForce = ResultingForce.add(frictionForce(*Obj, grounded))
	ResultingForce = ResultingForce.add(movementForce(Input, grounded))
	ResultingForce = ResultingForce.add(reactiveForce(Floor, *Obj, ResultingForce))
	ResultingForce = ResultingForce.multiply(1 / Obj.M)
	//Cap
	newSpeed := Obj.Speed.apply(ResultingForce, delay)
	newSpeedN := newSpeed.norm()
	if newSpeedN > Const.CapSpeed {
		log.Println("Speed Cap", newSpeed, newSpeed.multiply(Const.CapSpeed/newSpeedN))
		newSpeed = newSpeed.multiply(Const.CapSpeed / newSpeedN)
	}
	*(Obj.Speed) = newSpeed
	*(Obj.Coord) = Obj.Coord.apply(*(Obj.Speed), delay)
	if contact(*Obj, Floor) {
		Obj.ground(Floor)
	}
	// log.Println(*Obj.Coord, contact(*Obj, Floor))
}
