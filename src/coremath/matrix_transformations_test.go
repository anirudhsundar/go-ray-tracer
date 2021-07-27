package coremath_test

import (
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
