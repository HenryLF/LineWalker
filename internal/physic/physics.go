package physic

import (
	"linewalker/internal/physic/vect"
	"math"
)

func gravityForce(Obj ObjectSide) vect.Vect {
	return vect.Vect{X: 0, Y: Obj.M * Const.G}
}

func reactiveForce(Floor func(float64) float64, Obj ObjectSide, R vect.Vect) vect.Vect {
	if contact(Obj, Floor) {
		out := vectOf(Floor, Obj.Coord.X).Rot()
		return out.Multiply(-R.Norm())
	} else {
		return vect.Vect{X: 0, Y: 0}
	}
}

func frictionForce(Obj ObjectSide, contact bool) vect.Vect {
	if contact {
		return (*Obj.Speed).Multiply(-Const.FloorFrictionCoeff)
	}
	return (*Obj.Speed).Multiply(-Const.AirFrictionCoeff)
}

func contact(Obj ObjectSide, Floor func(float64) float64) bool {
	return Obj.Coord.Y+Obj.R >= Floor(Obj.Coord.X)
}

func vectOf(Floor func(float64) float64, x float64) vect.Vect {
	out := vect.Vect{
		X: 2 * Const.DX,
		Y: Floor(x+Const.DX) - Floor(x-Const.DX),
	}
	return out.Unit()
}

func (Obj *ObjectSide) ground(Floor func(float64) float64) {
	Obj.Coord.Y = math.Min(Floor(Obj.Coord.X)-Obj.R, Obj.Coord.Y)
	Obj.Speed.Y = math.Min(Floor(Obj.Speed.Y), 0)
}
func colisionForce(Obj ObjectSide, Col ObjectSide) vect.Vect {
	Energy := Obj.M * math.Pow(Obj.Speed.Norm(), 2) / 2
	Energy += Col.M * math.Pow(Col.Speed.Norm(), 2) / 2
	Energy /= 2
	return (*Col.Coord).To(*Obj.Coord).Unit().Multiply(Energy * Const.ElasticColision)
}

func movementForce(Obj ObjectSide, Input UserInput, Floor func(float64) float64, grounded bool) vect.Vect {
	slope := vectOf(Floor, Obj.Coord.X)
	var out vect.Vect

	if Input.Up && grounded {
		out = (vect.Vect{X: 0, Y: -1}).Multiply(Const.VerticalAcc)
		if Input.Right {
			return out.Add(slope.Multiply(Const.LateralAcc))
		} else if Input.Left {
			return out.Add(slope.Multiply(-Const.LateralAcc))
		} else {
			return out
		}
	}
	if Input.Right {
		if grounded {
			out = out.Add(slope.Multiply(Const.LateralAcc))
		} else {
			out = out.Add(slope.Multiply(Const.LateralAirAcc))
		}
	}
	if Input.Left {
		if grounded {
			out = out.Add(slope.Multiply(-Const.LateralAcc))
		} else {
			out = out.Add(slope.Multiply(-Const.LateralAirAcc))
		}
	}
	if Input.Down {
		out = vect.Vect{X: 0, Y: 1}
		return out.Multiply(Const.VerticalAccDown)
	}
	return out
}
