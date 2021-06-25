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

func TestPointEqual(t *testing.T) {
	p1 := Point(1, 2, 3)
	p2 := Point(1, 2, 3)
	result := p1.Equal(*p2)
	if !result {
		t.Errorf("Point equality failed")
	}
}

func TestVectorEqual(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(1, 2, 3)
	result := v1.Equal(*v2)
	if !result {
		t.Errorf("Vector equality failed")
	}
}

func AddSubTestUtil(f func(t1, t2 Tuple) Tuple,
	t1, t2, expected Tuple,
	t *testing.T) {
	result := f(t1, t2)
	if !result.Equal(expected) {
		t.Errorf("%v + %v is expected to return %v", t1, t2, expected)
	}
}

func TestTupleAdd(t *testing.T) {
	p := Point(3, -2, 5)
	v := Vector(-2, 3, 1)
	AddSubTestUtil(Tuple.Add, *p, *v, Tuple{1, 1, 6, 1}, t)
}

func TestSubPointFromPoint(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	AddSubTestUtil(Tuple.Sub, *p1, *p2, *Vector(-2, -4, -6), t)
}

func TestSubVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	AddSubTestUtil(Tuple.Sub, *p, *v, *Point(-2, -4, -6), t)
}

func TestSubVectorFromVector(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	AddSubTestUtil(Tuple.Sub, *v1, *v2, *Vector(-2, -4, -6), t)
}

func TestSubVectorFromZero(t *testing.T) {
	zero := Vector(0, 0, 0)
	v := Vector(1, -2, 3)
	AddSubTestUtil(Tuple.Sub, *zero, *v, *Vector(-1, 2, -3), t)
}

func TestNegateTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{-1, 2, -3, 4}
	result := a.Negate()
	if !result.Equal(expected) {
		t.Errorf("Negation of %v is expected to return %v but got %v", a, expected, result)
	}
}

func TestScalarMultiply(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	expected := Tuple{3.5, -7, 10.5, -14}
	result := a.ScalarMultiply(3.5)
	if !result.Equal(expected) {
		t.Errorf("Negation of %v is expected to return %v but got %v", a, expected, result)
	}
}
