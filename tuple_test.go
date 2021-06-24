package ray_tracer

import (
	"testing"
)

func TestTuple(t *testing.T) {
	var testTuple = func(act, ref float64) {
		if !floatEq(act, ref) {
			t.Errorf("Tuple expected to be %v but got %v", ref, act)
		}
	}
	x, y, z, w := 4.3, -4.2, 3.1, 1.0
	vector := Tuple{x, y, z, w}
	testTuple(vector.x, x)
	testTuple(vector.y, y)
	testTuple(vector.z, z)
	testTuple(vector.w, w)
	x, y, z, w = 4.3, -4.2, 3.1, 0.0
	point := Tuple{x, y, z, w}
	testTuple(point.x, x)
	testTuple(point.y, y)
	testTuple(point.z, z)
	testTuple(point.w, w)
}

func TestPoint(t *testing.T) {
	p := Point(4, -3, 4)
	if p == nil {
		t.Errorf("Point(x, y, z) is expected to return a valid Tuple")
	} else if !floatEq(p.w, 1.0) {
		t.Errorf("Point is expected to have w=%v but returned %v", 1.0, p.w)
	}
}
func TestVector(t *testing.T) {
	v := Vector(4, -3, 4)
	if v == nil {
		t.Errorf("Vector(x, y, z) is expected to return a valid Tuple")
	} else if !floatEq(v.w, 0.0) {
		t.Errorf("Vector is expected to have w=%v but returned %v", 0.0, v.w)
	}
}

func TestTupleEq(t *testing.T) {
	p1 := Point(1, 2, 3)
	p2 := Point(1, 2, 3)
	v1 := Vector(1, 2, 3)
	v2 := Vector(1, 2, 3)
	if !p1.Equal(*p2) {
		t.Errorf("Point equality failed")
	}
	if !v1.Equal(*v2) {
		t.Errorf("Vector equality failed")
	}
}

func TestTupleAdd(t *testing.T) {
	p := Point(3, -2, 5)
	v := Vector(-2, 3, 1)
	expectedPoint := Tuple{1, 1, 6, 1}
	if !p.Add(*v).Equal(expectedPoint) {
		t.Errorf("%v + %v is expected to return %v", p, v, expectedPoint)
	}
}

func TestTupleSub(t *testing.T) {
	p := Point(3, 2, 1)
	v := Point(5, 6, 7)
	expectedPoint := Tuple{-2, -4, -6, 0}
	if !p.Sub(*v).Equal(expectedPoint) {
		t.Errorf("%v - %v is expected to return %v", p, v, expectedPoint)
	}
}
