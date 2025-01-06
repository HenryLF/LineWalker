package physic

import (
	"errors"
	"log"
	"reflect"
)

var G float64 = 1000

var AirFrictionCoeff float64 = .1
var FloorFrictionCoeff float64 = .9999

var DX float64 = 10

var LateralAcc float64 = 1000
var LateralAirAcc float64 = 0
var VerticalAcc float64 = 10000

var CapSpeed float64 = 500

var TimeSlow float64 = 1

type Constants struct {
	G, AirFrictionCoeff, FloorFrictionCoeff, DX, LateralAcc, LateralAirAcc, VerticalAcc, CapSpeed, TimeSlow float64
}

func (C *Constants) Set(s string, a float64) any {
	c := reflect.ValueOf(C).Elem().FieldByName(s)
	if c.CanSet() {
		log.Println("Change ", s, " from ", c, " to ", a)
		c.SetFloat(a)
		return a
	}
	log.Println("Error setting", s)
	return errors.New("trying to set unadressable field")
}

func (C *Constants) Get(s string) float64 {
	return reflect.ValueOf(C).Elem().FieldByName(s).Float()
}

var Const = Constants{G: G, AirFrictionCoeff: AirFrictionCoeff, FloorFrictionCoeff: FloorFrictionCoeff, DX: DX, LateralAcc: LateralAcc, LateralAirAcc: LateralAirAcc, VerticalAcc: VerticalAcc, CapSpeed: CapSpeed, TimeSlow: TimeSlow}
