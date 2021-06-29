package projectile

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/canvas"
	"github.com/anirudhsundar/go-ray-tracer/src/color"
	"github.com/anirudhsundar/go-ray-tracer/src/tuple"
)

func TestProjectileTick(t *testing.T) {
	// Projectile starts 1 unit above the origin
	// Velocity is normalized to 1 unit/tick
	start := *tuple.Point(0, 1, 0)
	velocity := tuple.Vector(1, 1.8, 0).Normalize().ScalarMultiply(11.25)
	p := Projectile{start, velocity}

	gravity := *tuple.Vector(0, -0.1, 0)
	wind := *tuple.Vector(-0.01, 0, 0)
	// Gravity -0.1 unit/tick and Wind is -0.01 unit/tick
	e := Environment{gravity, wind}
	canvasWidth := 900
	canvasHeight := 550
	c := canvas.Canvas(canvasWidth, canvasHeight)

	fmt.Printf("Projectile position starts at %v\n", p.Position)
	for p.Position.Y > 0 {
		p = Tick(e, p)
		fmt.Printf("Projectile position is %v\n", p.Position)
		canvasY := canvasHeight - int(p.Position.Y)
		if canvasY < 0 {
			canvasY = 0
		}
		if canvasY > canvasHeight {
			canvasY = canvasHeight - 1
		}
		canvasX := int(p.Position.X)
		if canvasX < 0 {
			canvasX = 0
		}
		if canvasX > canvasWidth {
			canvasX = canvasWidth - 1
		}
		c.WritePixel(canvasX, canvasY, *color.Color(1, 0, 0))
	}
	ppm := c.CanvasToPPM()
	err := ioutil.WriteFile("projectile.ppm", []byte(ppm), 0644)
	if err != nil {
		t.Error(err)
	}

}
