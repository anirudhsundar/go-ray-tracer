package coremath

import "math"

func Translation(x, y, z float64) *Matrix {
	m := IdentityMatrix(4)
	m.Set(x, 0, 3)
	m.Set(y, 1, 3)
	m.Set(z, 2, 3)
	return m
}

func Scaling(x, y, z float64) *Matrix {
	m := IdentityMatrix(4)
	m.Set(x, 0, 0)
	m.Set(y, 1, 1)
	m.Set(z, 2, 2)
	return m
}

func RotationX(val float64) *Matrix {
	m := IdentityMatrix(4)
	m.Set(math.Cos(val), 1, 1)
	m.Set(-math.Sin(val), 1, 2)
	m.Set(math.Sin(val), 2, 1)
	m.Set(math.Cos(val), 2, 2)
	return m
}

func RotationY(val float64) *Matrix {
	m := IdentityMatrix(4)
	m.Set(math.Cos(val), 0, 0)
	m.Set(math.Sin(val), 0, 2)
	m.Set(-math.Sin(val), 2, 0)
	m.Set(math.Cos(val), 2, 2)
	return m
}

func RotationZ(val float64) *Matrix {
	m := IdentityMatrix(4)
	m.Set(math.Cos(val), 0, 0)
	m.Set(-math.Sin(val), 0, 1)
	m.Set(math.Sin(val), 1, 0)
	m.Set(math.Cos(val), 1, 1)
	return m
}
