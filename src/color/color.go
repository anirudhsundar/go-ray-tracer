package color

import (
	"fmt"

	"github.com/anirudhsundar/go-ray-tracer/src/tuple"
)

type colorType struct {
	Red, Blue, Green float64
}

func tupleFromColor(c colorType) tuple.Tuple {
	return tuple.Tuple{X: c.Red, Y: c.Blue, Z: c.Green, W: 0}
}

func colorFromTuple(t tuple.Tuple) *colorType {
	return Color(t.X, t.Y, t.Z)
}

func Color(red, blue, green float64) *colorType {
	t := tuple.Tuple{X: red, Y: blue, Z: green, W: 0}
	c := colorType{t.X, t.Y, t.Z}
	return &c
}

func (c colorType) String() string {
	return fmt.Sprintf("RGB: (%v, %v, %v)", c.Red, c.Blue, c.Green)
}

func (c1 colorType) Add(c2 colorType) *colorType {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return colorFromTuple(t1.Add(t2))
}

func (c1 colorType) Sub(c2 colorType) *colorType {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return colorFromTuple(t1.Sub(t2))
}

func (c1 colorType) Multiply(c2 colorType) *colorType {
	return Color(c1.Red*c2.Red,
		c1.Blue*c2.Blue,
		c1.Green*c2.Green)
}

func (c1 colorType) ScalarMultiply(v float64) *colorType {
	t1 := tupleFromColor(c1)
	return colorFromTuple(t1.ScalarMultiply(v))
}

func (c1 colorType) Equal(c2 colorType) bool {
	t1 := tupleFromColor(c1)
	t2 := tupleFromColor(c2)
	return t1.Equal(t2)
}
