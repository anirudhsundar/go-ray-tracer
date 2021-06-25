package ray_tracer

import "math"

const epsilon float64 = 0.00001

func floatEq(x, y float64) bool {
	return math.Abs(x-y) <= epsilon
}

type Tuple struct {
	x, y, z, w float64
}

func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func (t1 Tuple) Equal(t2 Tuple) bool {
	return floatEq(t1.x, t2.x) &&
		floatEq(t1.y, t2.y) &&
		floatEq(t1.z, t2.z) &&
		floatEq(t1.w, t2.w)
}

func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t1.x + t2.x,
		t1.y + t2.y,
		t1.z + t2.z,
		t1.w + t2.w}
}

func (t1 Tuple) Sub(t2 Tuple) Tuple {
	return Tuple{t1.x - t2.x,
		t1.y - t2.y,
		t1.z - t2.z,
		t1.w - t2.w}
}

func (t1 Tuple) Negate() Tuple {
	return Tuple{-t1.x,
		-t1.y,
		-t1.z,
		-t1.w}
}

func (t1 Tuple) ScalarMultiply(v float64) Tuple {
	return Tuple{t1.x * v,
		t1.y * v,
		t1.z * v,
		t1.w * v}
}

func (t1 Tuple) ScalarDivide(v float64) Tuple {
	return Tuple{t1.x / v,
		t1.y / v,
		t1.z / v,
		t1.w / v}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

func (t Tuple) Normalize() Tuple {
	return t.ScalarDivide(t.Magnitude())
}
