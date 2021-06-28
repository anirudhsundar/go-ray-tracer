package canvas

import (
	"fmt"

	"github.com/anirudhsundar/go-ray-tracer/src/color"
)

type CanvasType struct {
	Width, Height int
	Pixels        [][]color.ColorType
}

func Canvas(w, h int) *CanvasType {
	var c CanvasType
	c.Height = h
	c.Width = w
	c.Pixels = make([][]color.ColorType, c.Height)
	for i := 0; i < c.Height; i++ {
		c.Pixels[i] = make([]color.ColorType, w)
		for j := 0; j < c.Width; j++ {
			c.Pixels[i][j] = *color.Color(0, 0, 0)
		}
	}
	return &c
}

func (c *CanvasType) WritePixel(x, y int, col color.ColorType) (bool, error) {
	if x < 0 || y < 0 {
		return false, fmt.Errorf("Cannot write to negative index (%v,%v)", x, y)
	}
	c.Pixels[y][x] = col
	return true, nil
}

func (c *CanvasType) PixelAt(x, y int) color.ColorType {
	return c.Pixels[y][x]
}
