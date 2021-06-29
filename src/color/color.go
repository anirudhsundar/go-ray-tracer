package color

import (
	"fmt"

	"github.com/anirudhsundar/go-ray-tracer/src/tuple"
)

type ColorType struct {
	Red, Green, Blue float64
}

func tupleFromColor(c ColorType) tuple.Tuple {
	return tuple.Tuple{X: c.Red, Y: c.Green, Z: c.Blue, W: 0}
}

func colorFromTuple(t tuple.Tuple) *ColorType {
	return Color(t.X, t.Y, t.Z)
}

func Color(red, green, blue float64) *ColorType {
	t := tuple.Tuple{X: red, Y: green, Z: blue, W: 0}
	c := ColorType{t.X, t.Y, t.Z}
	return &c
}

func (c ColorType) String() string {
	return fmt.Sprintf("RGB: (%v, %v, %v)", c.Red, c.Green, c.Blue)
}

func (c1 ColorType) Add(c2 ColorType) *ColorType {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return colorFromTuple(t1.Add(t2))
}

func (c1 ColorType) Sub(c2 ColorType) *ColorType {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return colorFromTuple(t1.Sub(t2))
}

func (c1 ColorType) Multiply(c2 ColorType) *ColorType {
	return Color(c1.Red*c2.Red,
		c1.Green*c2.Green,
		c1.Blue*c2.Blue)
}

func (c1 ColorType) ScalarMultiply(v float64) *ColorType {
	t1 := tupleFromColor(c1)
	return colorFromTuple(t1.ScalarMultiply(v))
}

func (c1 ColorType) Equal(c2 ColorType) bool {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return t1.Equal(t2)
}
