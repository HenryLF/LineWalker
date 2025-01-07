package physic

//Default Value for Const struct,

const dG float64 = 50
const dAirFrictionCoeff float64 = .2
const dFloorFrictionCoeff float64 = 1
const dDX float64 = 0.01
const dLateralAcc float64 = 1000
const dLateralAirAcc float64 = 100
const dVerticalAcc float64 = 10000
const dVerticalAccDown float64 = 1000
const dCapSpeed float64 = 5000
const dTimeSlow float64 = .01
const dMaxTimeDelay float64 = 1
const dElasticColision float64 = 1

type Constants struct {
	G, AirFrictionCoeff, FloorFrictionCoeff, DX,
	LateralAcc, LateralAirAcc, VerticalAcc,
	VerticalAccDown, CapSpeed, TimeSlow, MaxTimeDelay,
	ElasticColision float64
}

var Const = Constants{
	G:                  dG,
	AirFrictionCoeff:   dAirFrictionCoeff,
	FloorFrictionCoeff: dFloorFrictionCoeff,
	DX:                 dDX,
	LateralAcc:         dLateralAcc,
	LateralAirAcc:      dLateralAirAcc,
	VerticalAcc:        dVerticalAcc,
	VerticalAccDown:    dVerticalAccDown,
	CapSpeed:           dCapSpeed,
	TimeSlow:           dTimeSlow,
	MaxTimeDelay:       dMaxTimeDelay,
	ElasticColision:    dElasticColision}
