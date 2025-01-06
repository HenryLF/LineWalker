package physic

import (
	"testing"
)

var ExObject = NewObject(50, 50, 20, 1)

func FlatFloor(i float64) float64 {
	return 500
}
func SlopFloor(i float64) float64 {
	return 500 - i*2
}
func VerySlopFloor(i float64) float64 {
	return 500 - i*100
}
func (Obj Object) log(t *testing.T) {
	t.Logf("Content of Object %v\nCoord : %v\nSpeed : %v\nMass : %v\nRadius : %v", &Obj, *Obj.Coord, *Obj.Speed, Obj.M, Obj.R)
}
func TestObject(t *testing.T) {
	ExObject.log(t)
}

func TestObjectMultiply(t *testing.T) {
	k := ExObject.Coord.multiply(2)
	t.Log("50,50 times 2", k)
	if !(k.X == 100 && k.Y == 100) {
		t.Error("error multiplying", k)
	}
	*(ExObject.Coord) = ExObject.Coord.multiply(2)
	ExObject.log(t)
}

func TestObjectPermanence(t *testing.T) {
	ExObject.log(t)
	if ExObject.Coord.X != 100 && ExObject.Coord.Y != 100 {
		t.Error("not permanent")
	}
}

func TestVectAdd(t *testing.T) {
	v1 := vect{1, 2}
	v2 := vect{1, -2}
	res := v1.add(v2)
	exp := vect{2, 0}
	if res != exp {
		t.Error("error adding", v1, "+", v2, "=", res)
	}
}

func TestContact(t *testing.T) {
	A := NewObject(0, 450.1, 50, 1)
	r := contact(A, FlatFloor)
	t.Log("expected contact", r)
	if !r {
		t.Fail()
	}
	A = NewObject(0, 0, 50, 1)
	r = contact(A, FlatFloor)
	t.Log("expected no contact", !r)
	if r {
		t.Fail()
	}

}

func TestDetectFloor(t *testing.T) {
	A := NewObject(0, 0, 50, 1)
	for {
		A.Coord.Y += 5.3
		A.log(t)
		if contact(A, FlatFloor) {
			break
		}
	}
}

func TestReactiveForce(t *testing.T) {
	A := NewObject(0, 450, 50, 1)
	t.Log("contact flat", contact(A, FlatFloor), angleOf(FlatFloor, 0))
	t.Log(reactiveForce(FlatFloor, A, gravityForce(A)))

	t.Log("contact slop", contact(A, SlopFloor), angleOf(SlopFloor, 0))
	t.Log(reactiveForce(SlopFloor, A, gravityForce(A)))

	t.Log("contact big slop", contact(A, VerySlopFloor), angleOf(VerySlopFloor, 0))
	t.Log(reactiveForce(SlopFloor, A, gravityForce(A)))

}

func TestSetConst(t *testing.T) {
	t.Log(Const)
	Const.Set("G", 0)
	t.Log(Const)
}

func TestGetConst(t *testing.T) {
	t.Log(Const)
	k := Const.Get("G")
	t.Log(k)
}
