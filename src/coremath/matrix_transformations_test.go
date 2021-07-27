package coremath_test

import (
	"math"
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/coremath"
)

func TestTranslation1(t *testing.T) {
	tr := coremath.Translation(5, -3, 2)
	p := coremath.Point(-3, 4, 5)

	translated, err := tr.TupleMultiply(*p)
	check(err, t)
	expectedTranslation := coremath.Point(2, 1, 7)
	if !expectedTranslation.Equal(*translated) {
		t.Errorf("Expected Translation of (5, -3, 2) to Point \n%v\n to be \n%v\n, but got \n%v\n", p, expectedTranslation, translated)
	}
}

func TestTranslation2(t *testing.T) {
	tr := coremath.Translation(5, -3, 2)
	inv, err := tr.Inverse()
	check(err, t)
	p := coremath.Point(-3, 4, 5)

	translated, err := inv.TupleMultiply(*p)
	check(err, t)
	expectedTranslation := coremath.Point(-8, 7, 3)
	if !expectedTranslation.Equal(*translated) {
		t.Errorf("Expected Translation of (5, -3, 2) to Point \n%v\n to be \n%v\n, but got \n%v\n", p, expectedTranslation, translated)
	}
}

func TestTranslation3(t *testing.T) {
	tr := coremath.Translation(5, -3, 2)
	p := coremath.Vector(-3, 4, 5)

	translated, err := tr.TupleMultiply(*p)
	check(err, t)
	if !p.Equal(*translated) {
		t.Errorf("Expected Translation of (5, -3, 2) to Vector \n%v\n to be \n%v\n, but got \n%v\n", p, p, translated)
	}
}

func TestScaling1(t *testing.T) {
	tr := coremath.Scaling(2, 3, 4)
	p := coremath.Point(-4, 6, 8)

	scaled, err := tr.TupleMultiply(*p)
	check(err, t)
	expectedScaling := coremath.Point(-8, 18, 32)
	if !expectedScaling.Equal(*scaled) {
		t.Errorf("Expected Translation of (2, 3, 4) to Point \n%v\n to be \n%v\n, but got \n%v\n", p, expectedScaling, scaled)
	}
}

func TestScaling2(t *testing.T) {
	tr := coremath.Scaling(2, 3, 4)
	p := coremath.Vector(-4, 6, 8)

	scaled, err := tr.TupleMultiply(*p)
	check(err, t)
	expectedScaling := coremath.Vector(-8, 18, 32)
	if !expectedScaling.Equal(*scaled) {
		t.Errorf("Expected Translation of (2,3,4) to Vector \n%v\n to be \n%v\n, but got \n%v\n", p, expectedScaling, scaled)
	}
}

func TestScaling3(t *testing.T) {
	tr := coremath.Scaling(2, 3, 4)
	inv, err := tr.Inverse()
	check(err, t)
	p := coremath.Vector(-4, 6, 8)

	translated, err := inv.TupleMultiply(*p)
	check(err, t)
	expectedTranslation := coremath.Vector(-2, 2, 2)
	if !expectedTranslation.Equal(*translated) {
		t.Errorf("Expected Translation of (2,3,4) to Point \n%v\n to be \n%v\n, but got \n%v\n", p, expectedTranslation, translated)
	}
}

func TestReflection(t *testing.T) {
	tr := coremath.Scaling(-1, 1, 1)
	p := coremath.Point(2, 3, 4)

	translated, err := tr.TupleMultiply(*p)
	check(err, t)
	expectedTranslation := coremath.Point(-2, 3, 4)
	if !expectedTranslation.Equal(*translated) {
		t.Errorf("Expected Translation of (-1,1,1) to Point \n%v\n to be \n%v\n, but got \n%v\n", p, expectedTranslation, translated)
	}
}

func TestRotationX(t *testing.T) {
	p := coremath.Point(0, 1, 0)
	halfQuarterValue := math.Pi / 4.0
	halfQuarter := coremath.RotationX(halfQuarterValue)
	halfQuarterRotation, err := halfQuarter.TupleMultiply(*p)
	check(err, t)
	expectedHalfQuarter := coremath.Point(0, math.Sqrt2/2, math.Sqrt2/2)
	if !halfQuarterRotation.Equal(*expectedHalfQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, halfQuarterValue, expectedHalfQuarter, halfQuarterRotation)
	}

	fullQuarterValue := math.Pi / 2.0
	fullQuarter := coremath.RotationX(fullQuarterValue)
	fullQuarterRotation, err := fullQuarter.TupleMultiply(*p)
	check(err, t)
	expectedFullQuarter := coremath.Point(0, 0, 1)
	if !fullQuarterRotation.Equal(*expectedFullQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, fullQuarterValue, expectedFullQuarter, fullQuarterRotation)
	}
}

func TestRotationY(t *testing.T) {
	p := coremath.Point(0, 0, 1)
	halfQuarterValue := math.Pi / 4.0
	halfQuarter := coremath.RotationY(halfQuarterValue)
	halfQuarterRotation, err := halfQuarter.TupleMultiply(*p)
	check(err, t)
	expectedHalfQuarter := coremath.Point(math.Sqrt2/2, 0, math.Sqrt2/2)
	if !halfQuarterRotation.Equal(*expectedHalfQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, halfQuarterValue, expectedHalfQuarter, halfQuarterRotation)
	}

	fullQuarterValue := math.Pi / 2.0
	fullQuarter := coremath.RotationY(fullQuarterValue)
	fullQuarterRotation, err := fullQuarter.TupleMultiply(*p)
	check(err, t)
	expectedFullQuarter := coremath.Point(1, 0, 0)
	if !fullQuarterRotation.Equal(*expectedFullQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, fullQuarterValue, expectedFullQuarter, fullQuarterRotation)
	}
}

func TestRotationZ(t *testing.T) {
	p := coremath.Point(0, 1, 0)
	halfQuarterValue := math.Pi / 4.0
	halfQuarter := coremath.RotationZ(halfQuarterValue)
	halfQuarterRotation, err := halfQuarter.TupleMultiply(*p)
	check(err, t)
	expectedHalfQuarter := coremath.Point(-math.Sqrt2/2, math.Sqrt2/2, 0)
	if !halfQuarterRotation.Equal(*expectedHalfQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, halfQuarterValue, expectedHalfQuarter, halfQuarterRotation)
	}

	fullQuarterValue := math.Pi / 2.0
	fullQuarter := coremath.RotationZ(fullQuarterValue)
	fullQuarterRotation, err := fullQuarter.TupleMultiply(*p)
	check(err, t)
	expectedFullQuarter := coremath.Point(-1, 0, 0)
	if !fullQuarterRotation.Equal(*expectedFullQuarter) {
		t.Errorf("Expected rotation along x axis of point %v by %v to be %v but got %v", p, fullQuarterValue, expectedFullQuarter, fullQuarterRotation)
	}
}
