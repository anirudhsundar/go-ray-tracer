package canvas

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

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

func scaleMinMax(origMin, origMax, newMin, newMax, x float64) float64 {
	if x < origMin {
		return newMin
	}
	if x > origMax {
		return newMax
	}
	return ((newMax-newMin)*(x-origMin))/(origMax-origMin) + newMin
}

func scaleRGB255(x float64) int {
	// return int(math.Ceil(255 * x))
	return int(math.Ceil(scaleMinMax(0, 1, 0, 255, x)))
}

func writePixelString(val string, lineLength *int, bufLine *bytes.Buffer, b *bytes.Buffer) {
	if *lineLength > 69 {
		b.WriteString(strings.TrimSpace(bufLine.String()))
		bufLine.Reset()
		b.WriteString("\n")
		*lineLength = 0
	}
	bufLine.WriteString(val)
}

func (c *CanvasType) CanvasToPPM() string {
	var b bytes.Buffer
	// First 3 lines are going to be headers
	// First line of PPM format should always be "P3"
	b.WriteString("P3\n")
	// Second line should be "Width Height", eg: "5 3"
	b.WriteString(fmt.Sprintf("%v %v\n", c.Width, c.Height))
	// Third line specifies max color value of RGB. Eg: 255
	b.WriteString("255\n")
	lineLength := 0
	var bufLine bytes.Buffer
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			if j != 0 {
				bufLine.WriteString(" ")
			}
			strPixelRed := strconv.Itoa(scaleRGB255(c.Pixels[i][j].Red))
			lineLength += len(strPixelRed)
			writePixelString(strPixelRed, &lineLength, &bufLine, &b)
			bufLine.WriteString(" ")
			lineLength++

			strPixelGreen := strconv.Itoa(scaleRGB255(c.Pixels[i][j].Green))
			lineLength += len(strPixelGreen)
			writePixelString(strPixelGreen, &lineLength, &bufLine, &b)
			bufLine.WriteString(" ")
			lineLength++

			strPixelBlue := strconv.Itoa(scaleRGB255(c.Pixels[i][j].Blue))
			lineLength += len(strPixelBlue)
			writePixelString(strPixelBlue, &lineLength, &bufLine, &b)
			lineLength++
		}
		b.WriteString(strings.TrimSpace(bufLine.String()))
		bufLine.Reset()
		b.WriteString("\n")
		lineLength = 0
	}
	return b.String()
}

func (c *CanvasType) WriteAllPixels(col color.ColorType) {
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			c.Pixels[i][j] = col
		}
	}
}
