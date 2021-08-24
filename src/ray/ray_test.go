package ray

import (
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/coremath"
)

func TestCreateRay(t *testing.T) {
	origin := coremath.Point(1, 2, 3)
	direction := coremath.Vector(4, 5, 6)
	r := Ray{origin, direction}
	if !origin.Equal(*r.origin) {
		t.Errorf("Expected to create ray with origin: %v, but found ray.origin: %v", origin, r.origin)
	}
	if !direction.Equal(*r.direction) {
		t.Errorf("Expected to create ray with direction: %v, but found ray.direction: %v", direction, r.direction)
	}
}

func TestPosition(t *testing.T) {
	r := &Ray{coremath.Point(2, 3, 4), coremath.Vector(1, 0, 0)}
	p0 := Position(r, 0)
	expectedP0 := coremath.Point(2, 3, 4)
	if !expectedP0.Equal(p0) {
		t.Errorf("Position for ray %v at t=0 is expected to be %v, but got %v", r, expectedP0, p0)
	}

	p1 := Position(r, 1)
	expectedP1 := coremath.Point(3, 3, 4)
	if !expectedP1.Equal(p1) {
		t.Errorf("Position for ray %v at t=1 is expected to be %v, but got %v", r, expectedP1, p1)
	}

	pneg1 := Position(r, -1)
	expectedPneg1 := coremath.Point(1, 3, 4)
	if !expectedPneg1.Equal(pneg1) {
		t.Errorf("Position for ray %v at t=-1 is expected to be %v, but got %v", r, expectedPneg1, pneg1)
	}

	p2_5 := Position(r, 2.5)
	expectedP2_5 := coremath.Point(4.5, 3, 4)
	if !expectedP2_5.Equal(p2_5) {
		t.Errorf("Position for ray %v at t=2.5 is expected to be %v, but got %v", r, expectedP2_5, p2_5)
	}

}
