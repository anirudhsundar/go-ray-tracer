package ray_tracer

import (
	"math"
	"testing"
)

const epsilon float64 = 0.00001

func floatEq(x, y float64) bool {
	return math.Abs(x-y) <= epsilon
}

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
