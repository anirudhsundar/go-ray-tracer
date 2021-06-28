package canvas_test

import (
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/canvas"
	"github.com/anirudhsundar/go-ray-tracer/src/color"
)

func TestCanvas(t *testing.T) {
	c := canvas.Canvas(10, 20)
	if c.Width != 10 || c.Height != 20 {
		t.Errorf("height and width of Canvas is expected to be 10 and 20 but got %v and %v", c.Width, c.Height)
	}
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			zero := color.Color(0, 0, 0)
			if !c.Pixels[i][j].Equal(*zero) {
				t.Errorf("Found a non-zero color while canvas initialization at index c.Pixels[%v][%v] = %v", i, j, c.Pixels[i][j])
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := canvas.Canvas(10, 20)
	red := color.Color(1, 0, 0)
	ok, err := c.WritePixel(2, 3, *red)
	if !ok || err != nil {
		t.Errorf("Could not write to pixel due to Error:\n\t%v", err)
	}
	if !c.PixelAt(2, 3).Equal(*red) {
		t.Errorf("Writing %v to pixel at (%v,%v) should set that pixel", red, 2, 3)
	}
}
