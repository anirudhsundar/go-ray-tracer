package coremath_test

import (
	"testing"

	"github.com/anirudhsundar/go-ray-tracer/src/coremath"
)

func checkMatrixValueAt(m *coremath.Matrix, row, col int, expected float64, t *testing.T) {
	result := m.At(row, col)
	if result != expected {
		t.Errorf("Matrix value at %d,%d is expected to be %v but got %v", row, col, expected, result)
	}
}

func check(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

func getNewMatrix(vals []float64, rows, cols int, t *testing.T) *coremath.Matrix {
	m, err := coremath.NewMatrix(vals, rows, cols)
	check(err, t)
	return m
}

func Test4x4Matrix(t *testing.T) {
	m := getNewMatrix([]float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5}, 4, 4, t)
	checkMatrixValueAt(m, 0, 0, 1, t)
	checkMatrixValueAt(m, 0, 3, 4, t)
	checkMatrixValueAt(m, 1, 0, 5.5, t)
	checkMatrixValueAt(m, 1, 2, 7.5, t)
	checkMatrixValueAt(m, 2, 2, 11, t)
	checkMatrixValueAt(m, 3, 0, 13.5, t)
	checkMatrixValueAt(m, 3, 2, 15.5, t)
}

func Test2x2Matrix(t *testing.T) {
	m := getNewMatrix([]float64{-3, 5, 1, -2}, 2, 2, t)
	checkMatrixValueAt(m, 0, 0, -3, t)
	checkMatrixValueAt(m, 0, 1, 5, t)
	checkMatrixValueAt(m, 1, 0, 1, t)
	checkMatrixValueAt(m, 1, 1, -2, t)
}

func Test3x3Matrix(t *testing.T) {
	m := getNewMatrix([]float64{-3, 5, 0, 1, -2, 7, 0, 1, 1}, 3, 3, t)
	checkMatrixValueAt(m, 0, 0, -3, t)
	checkMatrixValueAt(m, 1, 1, -2, t)
	checkMatrixValueAt(m, 2, 2, 1, t)
}

func Test4x4MatrixEqual(t *testing.T) {
	m1 := getNewMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}, 4, 4, t)
	m2 := getNewMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}, 4, 4, t)
	ok, err := m1.CheckEqual(m2)
	if !ok {
		t.Error(err)
	}
}

func Test4x4MatrixNotEqual(t *testing.T) {
	m1 := getNewMatrix([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2}, 4, 4, t)
	m2 := getNewMatrix([]float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 4, 4, t)
	if m1.Equal(m2) {
		t.Error("Matrices are not supposed to be equal")
	}
}

func Test4x4MatrixMultiply(t *testing.T) {
	m1 := getNewMatrix(
		[]float64{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 8, 7, 6,
			5, 4, 3, 2},
		4, 4, t)
	m2 := getNewMatrix(
		[]float64{
			-2, 1, 2, 3,
			3, 2, 1, -1,
			4, 3, 6, 5,
			1, 2, 7, 8},
		4, 4, t)
	expected := getNewMatrix(
		[]float64{
			20, 22, 50, 48,
			44, 54, 114, 108,
			40, 58, 110, 102,
			16, 26, 46, 42},
		4, 4, t)
	result, err := m1.MatrixMultiply(m2)
	check(err, t)
	if !result.Equal(expected) {
		t.Errorf("Matrix product of \n%v and \n%v is expected to be \n%v, but got \n%v", m1, m2, expected, result)
	}
}

func TestMatrixTupl(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			1, 2, 3, 4,
			2, 4, 4, 2,
			8, 6, 4, 1,
			0, 0, 0, 1},
		4, 4, t)
	b := coremath.Tuple{1, 2, 3, 1}
	expected := coremath.Tuple{18, 24, 33, 1}
	result, err := A.TupleMultiply(b)
	check(err, t)
	if !result.Equal(expected) {
		t.Errorf("Product of \n%v and \n%v\n is expected to be \n%v\n but got \n%v", A, b, expected, result)
	}
}

func Test4x4MatrixMultiplyIdentity(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			0, 1, 2, 4,
			1, 2, 4, 8,
			2, 4, 8, 16,
			4, 8, 16, 32},
		4, 4, t)
	I := coremath.IdentityMatrix(4)
	result, err := A.MatrixMultiply(I)
	check(err, t)
	if !result.Equal(A) {
		t.Errorf("Multiplying \n%v\n by Identitiy matrix returns \n%v \n", A, I)
	}
}

func TestMatrixTranspose(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			0, 9, 3, 0,
			9, 8, 0, 8,
			1, 8, 5, 3,
			0, 0, 5, 8},
		4, 4, t)
	expected := getNewMatrix(
		[]float64{
			0, 9, 1, 0,
			9, 8, 8, 0,
			3, 0, 5, 5,
			0, 8, 3, 8},
		4, 4, t)
	result := A.Transpose()
	if !result.Equal(expected) {
		t.Errorf("Transpose of \n%v\n is expected to be \n%v\n but got \n%v\n", A, expected, result)
	}
}

func TestMatrixTransposeIdentity(t *testing.T) {
	I := coremath.IdentityMatrix(4)
	if !I.Transpose().Equal(I) {
		t.Error("Transpose of Identity should be Identity")
	}

}

func Test2x2Determinant(t *testing.T) {
	m := getNewMatrix(
		[]float64{
			1, 5,
			-3, 2},
		2, 2, t)

	var expected float64 = 17
	result := m.Determinant2x2()
	if !coremath.FloatEq(result, expected) {
		t.Errorf("Expected determinant of \n%v\n to be %v but got %v", m, expected, result)
	}
}

func Test3x3SubMatrix(t *testing.T) {
	m := getNewMatrix(
		[]float64{
			1, 5, 0,
			-3, 2, 7,
			0, 6, -3},
		3, 3, t)

	row, col := 0, 2
	expected := getNewMatrix(
		[]float64{
			-3, 2,
			0, 6},
		2, 2, t)
	result, err := m.SubMatrix(row, col)
	check(err, t)
	if !result.Equal(expected) {
		t.Errorf("A sub-matrix of row %v and col %v for matrix \n%v\n is expected to be \n%v\n but got \n%v\n", row, col, m, expected, result)
	}
}

func Test4x4SubMatrix(t *testing.T) {
	m := getNewMatrix(
		[]float64{
			-6, 1, 1, 6,
			-8, 5, 8, 6,
			-1, 0, 8, 2,
			-7, 1, -1, 1},
		4, 4, t)
	expected := getNewMatrix(
		[]float64{
			-6, 1, 6,
			-8, 8, 6,
			-7, -1, 1},
		3, 3, t)

	row, col := 2, 1
	result, err := m.SubMatrix(row, col)
	check(err, t)
	if !result.Equal(expected) {
		t.Errorf("A sub-matrix of row %v and col %v for matrix \n%v\n is expected to be \n%v\n but got \n%v\n", row, col, m, expected, result)
	}
}

func Test3x3Minor(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5},
		3, 3, t)
	expected := float64(25)
	result, err := A.Minor(1, 0)
	check(err, t)
	if !coremath.FloatEq(result, expected) {
		t.Errorf("Expected Minro of \n%v\n to be %v but got %v", A, expected, result)
	}
}

func Test3x3Cofactor(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			3, 5, 0,
			2, -1, -7,
			6, -1, 5},
		3, 3, t)
	minor00, err := A.Minor(0, 0)
	check(err, t)
	expectedMinor00 := -12.0
	if !coremath.FloatEq(minor00, expectedMinor00) {
		t.Errorf("Expected minor of \n%v\n at 0,0 to be %f but got %f", A, expectedMinor00, minor00)
	}

	cofactor00, err := A.Cofactor(0, 0)
	check(err, t)
	expectedCofactor00 := -12.0
	if !coremath.FloatEq(cofactor00, expectedCofactor00) {
		t.Errorf("Expected cofactor of \n%v\n at 0,0 to be %f but got %f", A, expectedCofactor00, cofactor00)
	}

	minor10, err := A.Minor(1, 0)
	check(err, t)
	expectedMinor10 := 25.0
	if !coremath.FloatEq(minor10, expectedMinor10) {
		t.Errorf("Expected minor of \n%v\n at 1,0 to be %f but got %f", A, expectedMinor10, minor10)
	}

	cofactor10, err := A.Cofactor(1, 0)
	check(err, t)
	expectedCofactor10 := -25.0
	if !coremath.FloatEq(cofactor10, expectedCofactor10) {
		t.Errorf("Expected cofactor of \n%v\n at 1,0 to be %f but got %f", A, expectedCofactor10, cofactor10)
	}
}

func Test3x3Determinant(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			1, 2, 6,
			-5, 8, -4,
			2, 6, 4},
		3, 3, t)

	cofactor00, err := A.Cofactor(0, 0)
	check(err, t)
	var expectedCofactor00 float64 = 56
	if !coremath.FloatEq(cofactor00, expectedCofactor00) {
		t.Errorf("Expected cofactor of \n%v\n at 0,0 to be %f but got %f", A, expectedCofactor00, cofactor00)
	}

	cofactor10, err := A.Cofactor(0, 1)
	check(err, t)
	var expectedCofactor10 float64 = 12
	if !coremath.FloatEq(cofactor10, expectedCofactor10) {
		t.Errorf("Expected cofactor of \n%v\n at 1,0 to be %f but got %f", A, expectedCofactor10, cofactor10)
	}

	cofactor20, err := A.Cofactor(0, 2)
	check(err, t)
	var expectedCofactor20 float64 = -46
	if !coremath.FloatEq(cofactor20, expectedCofactor20) {
		t.Errorf("Expected cofactor of \n%v\n at 2,0 to be %f but got %f", A, expectedCofactor20, cofactor20)
	}

	var expectedDeterminant float64 = -196
	determinant, err := A.Determinant()
	check(err, t)
	if !coremath.FloatEq(determinant, expectedDeterminant) {
		t.Errorf("Expected determinant of \n%v\n to be %v but got %v", A, expectedDeterminant, determinant)
	}
}

func Test4x4Determinant(t *testing.T) {
	A := getNewMatrix(
		[]float64{
			-2, -8, 3, 5,
			-3, 1, 7, 3,
			1, 2, -9, 6,
			-6, 7, 7, -9},
		4, 4, t)

	cofactor00, err := A.Cofactor(0, 0)
	check(err, t)
	var expectedCofactor00 float64 = 690
	if !coremath.FloatEq(cofactor00, expectedCofactor00) {
		t.Errorf("Expected cofactor of \n%v\n at 0,0 to be %f but got %f", A, expectedCofactor00, cofactor00)
	}

	cofactor10, err := A.Cofactor(0, 1)
	check(err, t)
	var expectedCofactor10 float64 = 447
	if !coremath.FloatEq(cofactor10, expectedCofactor10) {
		t.Errorf("Expected cofactor of \n%v\n at 1,0 to be %f but got %f", A, expectedCofactor10, cofactor10)
	}

	cofactor20, err := A.Cofactor(0, 2)
	check(err, t)
	var expectedCofactor20 float64 = 210
	if !coremath.FloatEq(cofactor20, expectedCofactor20) {
		t.Errorf("Expected cofactor of \n%v\n at 2,0 to be %f but got %f", A, expectedCofactor20, cofactor20)
	}

	cofactor30, err := A.Cofactor(0, 3)
	check(err, t)
	var expectedCofactor30 float64 = 51
	if !coremath.FloatEq(cofactor30, expectedCofactor30) {
		t.Errorf("Expected cofactor of \n%v\n at 3,0 to be %f but got %f", A, expectedCofactor30, cofactor30)
	}

	var expectedDeterminant float64 = -4071
	determinant, err := A.Determinant()
	check(err, t)
	if !coremath.FloatEq(determinant, expectedDeterminant) {
		t.Errorf("Expected determinant of \n%v\n to be %v but got %v", A, expectedDeterminant, determinant)
	}
}
