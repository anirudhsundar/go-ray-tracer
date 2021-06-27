package projectile

import (
	"fmt"
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/tuple"
)

func TestProjectileTick(t *testing.T) {
	// Projectile starts 1 unit above the origin
	// Velocity is normalized to 1 unit/tick
	p := Projectile{*tuple.Point(0, 1, 0), *tuple.Vector(1, 10, 0)}

	// Gravity -0.1 unit/tick and Wind is -0.01 unit/tick
	e := Environment{*tuple.Vector(0, -0.1, 0), *tuple.Vector(0, -0.01, 0)}

	fmt.Printf("Projectile position starts at %v\n", p.Position)
	for p.Position.Y > 0 {
		p = Tick(e, p)
		fmt.Printf("Projectile position is %v\n", p.Position)
	}

}
