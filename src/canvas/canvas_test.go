package canvas_test

import (
	"bufio"
	"strings"
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

func CompareStringLines(low, high int, str []string, compare string, t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(compare))
	i := 0
	for scanner.Scan() {
		if i < low {
			i++
			continue
		}
		if i >= high {
			break
		}
		actual_text := scanner.Text()
		if actual_text != str[i-low] {
			t.Errorf("Line no. %v was expected to be \"%v\" but got \"%v\"", i+1, str[i-low], actual_text)
		}
		i++
	}
}

func TestCanvasToPPMHeader(t *testing.T) {
	c := canvas.Canvas(5, 3)
	ppm := c.GetPPMString()
	str := []string{"P3", "5 3", "255"}
	CompareStringLines(0, 3, str, ppm, t)
}

func TestCanvasToPPMData(t *testing.T) {
	c := canvas.Canvas(5, 3)
	c1 := color.Color(1.5, 0, 0)
	c2 := color.Color(0, 0.5, 0)
	c3 := color.Color(-0.5, 0, 1)
	c.WritePixel(0, 0, *c1)
	c.WritePixel(2, 1, *c2)
	c.WritePixel(4, 2, *c3)
	ppm := c.GetPPMString()
	str := []string{"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"}
	CompareStringLines(3, 6, str, ppm, t)
}

func TestCanvasToPPMDataLongLines(t *testing.T) {
	c := canvas.Canvas(10, 2)
	c1 := color.Color(1, 0.8, 0.6)
	c.WriteAllPixels(*c1)
	ppm := c.GetPPMString()
	str := []string{"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153"}
	CompareStringLines(3, 7, str, ppm, t)
}

func TestCanvasToPPMNewLine(t *testing.T) {
	c := canvas.Canvas(5, 3)
	ppm := c.GetPPMString()
	if ppm[len(ppm)-1] != '\n' {
		t.Errorf("Last character in PPM file is expected to be -1")
	}
}
