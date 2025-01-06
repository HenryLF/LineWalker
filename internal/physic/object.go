package physic

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
