package tuple

import "math"

const epsilon float64 = 0.00001

func FloatEq(x, y float64) bool {
	return math.Abs(x-y) <= epsilon
}

type Tuple struct {
	X, Y, Z, W float64
}

func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t1 Tuple) Equal(t2 Tuple) bool {
	return FloatEq(t1.X, t2.X) &&
		FloatEq(t1.Y, t2.Y) &&
		FloatEq(t1.Z, t2.Z) &&
		FloatEq(t1.W, t2.W)
}

func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t1.X + t2.X,
		t1.Y + t2.Y,
		t1.Z + t2.Z,
		t1.W + t2.W}
}

func (t1 Tuple) Sub(t2 Tuple) Tuple {
	return Tuple{t1.X - t2.X,
		t1.Y - t2.Y,
		t1.Z - t2.Z,
		t1.W - t2.W}
}

func (t1 Tuple) Negate() Tuple {
	return Tuple{-t1.X,
		-t1.Y,
		-t1.Z,
		-t1.W}
}

func (t1 Tuple) ScalarMultiply(v float64) Tuple {
	return Tuple{t1.X * v,
		t1.Y * v,
		t1.Z * v,
		t1.W * v}
}

func (t1 Tuple) ScalarDivide(v float64) Tuple {
	return Tuple{t1.X / v,
		t1.Y / v,
		t1.Z / v,
		t1.W / v}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Tuple) Normalize() Tuple {
	return t.ScalarDivide(t.Magnitude())
}

func (a Tuple) Dot(b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func (a Tuple) Cross(b Tuple) Tuple {
	return *Vector(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X)
}
