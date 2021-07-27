package coremath

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
