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

func getNewMatrix(vals []float64, rows, cols int, t *testing.T) *coremath.Matrix {
	m, err := coremath.NewMatrix(vals, rows, cols)
	if err != nil {
		t.Error(err)
	}
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
	if err != nil {
		t.Error(err)
	}
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
	if err != nil {
		t.Error(err)
	}
	if !result.Equal(expected) {
		t.Errorf("Product of \n%v and \n%v\n is expected to be \n%v\n but got \n%v", A, b, expected, result)
	}
}
