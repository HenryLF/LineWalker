package physic

import (
	"math"
)

func gravityForce(Obj Object) Vect {
	return Vect{X: 0, Y: Obj.M * Const.G}
}

func reactiveForce(Floor func(float64) float64, Obj Object, R Vect) Vect {
	if contact(Obj, Floor) {
		out := vectOf(Floor, Obj.Coord.X).rot()
		return out.multiply(-R.norm())
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

func vectOf(Floor func(float64) float64, x float64) Vect {
	out := Vect{
		X: 2 * Const.DX,
		Y: Floor(x+Const.DX) - Floor(x-Const.DX),
	}
	return out.unit()
}

func (Obj *Object) ground(Floor func(float64) float64) {
	Obj.Coord.Y = math.Min(Floor(Obj.Coord.X)-Obj.R, Obj.Coord.Y)
	Obj.Speed.Y = math.Min(Floor(Obj.Speed.Y), 0)
}
func colisionForce(Obj Object, Col Object) Vect {
	Energy := Obj.M * math.Pow(Obj.Speed.norm(), 2) / 2
	Energy += Col.M * math.Pow(Col.Speed.norm(), 2) / 2
	Energy /= 2
	return (*Col.Coord).to(*Obj.Coord).unit().multiply(Energy * Const.ElasticColision)
}

func movementForce(Obj Object, Input UserInput, Floor func(float64) float64, grounded bool) Vect {
	slope := vectOf(Floor, Obj.Coord.X)
	var out Vect

	if Input.Up && grounded {
		out = (Vect{0, -1}).multiply(Const.VerticalAcc)
		if Input.Right {
			return out.add(slope.multiply(Const.LateralAcc))
		} else if Input.Left {
			return out.add(slope.multiply(-Const.LateralAcc))
		} else {
			return out
		}
	}
	if Input.Right {
		if grounded {
			out = out.add(slope.multiply(Const.LateralAcc))
		} else {
			out = out.add(slope.multiply(Const.LateralAirAcc))
		}
	}
	if Input.Left {
		if grounded {
			out = out.add(slope.multiply(-Const.LateralAcc))
		} else {
			out = out.add(slope.multiply(-Const.LateralAirAcc))
		}
	}
	if Input.Down {
		out = Vect{0, 1}
		return out.multiply(Const.VerticalAccDown)
	}
	return out
}
