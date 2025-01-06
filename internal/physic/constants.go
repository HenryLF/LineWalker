package physic

//Default Value for Const struct,

const G float64 = 1000
const AirFrictionCoeff float64 = .2
const FloorFrictionCoeff float64 = 4
const DX float64 = 10
const LateralAcc float64 = 50000
const LateralAirAcc float64 = 0
const VerticalAcc float64 = 10000
const CapSpeed float64 = 500
const TimeSlow float64 = 1
const MaxTimeDelay float64 = .05
const ElasticColision float64 = 1

type Constants struct {
	G, AirFrictionCoeff, FloorFrictionCoeff, DX, LateralAcc, LateralAirAcc, VerticalAcc, CapSpeed, TimeSlow, MaxTimeDelay, ElasticColision float64
}

var Const = Constants{G: G, AirFrictionCoeff: AirFrictionCoeff, FloorFrictionCoeff: FloorFrictionCoeff, DX: DX, LateralAcc: LateralAcc, LateralAirAcc: LateralAirAcc, VerticalAcc: VerticalAcc, CapSpeed: CapSpeed, TimeSlow: TimeSlow, MaxTimeDelay: MaxTimeDelay}
