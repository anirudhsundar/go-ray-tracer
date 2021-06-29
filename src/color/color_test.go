package color_test

import (
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/color"
)

func TestColor(t *testing.T) {
	c := color.Color(-0.5, 0.4, 1.7)
	if c.Red != -0.5 || c.Green != 0.4 || c.Blue != 1.7 {
		t.Errorf("The red, green and blue colors are expected to be %v, %v, %v but got %v, %v %v", -0.5, 0.4, 1.7, c.Red, c.Green, c.Blue)
	}
}

func TestColorAdd(t *testing.T) {
	c1 := color.Color(0.9, 0.6, 0.75)
	c2 := color.Color(0.7, 0.1, 0.25)
	result := c1.Add(*c2)
	expected := *color.Color(1.6, 0.7, 1.0)
	if !result.Equal(expected) {
		t.Errorf("%v + %v is expected to be %v, but got %v", c1, c2, expected, result)
	}
}

func TestColorSubtract(t *testing.T) {
	c1 := color.Color(0.9, 0.6, 0.75)
	c2 := color.Color(0.7, 0.1, 0.25)
	result := c1.Sub(*c2)
	expected := *color.Color(0.2, 0.5, 0.5)
	if !result.Equal(expected) {
		t.Errorf("%v - %v is expected to be %v, but got %v", c1, c2, expected, result)
	}
}

func TestColorScalarMultiply(t *testing.T) {
	c1 := color.Color(0.2, 0.3, 0.4)
	result := c1.ScalarMultiply(2)
	expected := *color.Color(0.4, 0.6, 0.8)
	if !result.Equal(expected) {
		t.Errorf("%v * %v is expected to be %v, but got %v", c1, 2, expected, result)
	}
}

func TestColorMultiply(t *testing.T) {
	c1 := color.Color(1, 0.2, 0.4)
	c2 := color.Color(0.9, 1, 0.1)
	result := c1.Multiply(*c2)
	expected := *color.Color(0.9, 0.2, 0.04)
	if !result.Equal(expected) {
		t.Errorf("%v * %v is expected to be %v, but got %v", c1, c2, expected, result)
	}
}
