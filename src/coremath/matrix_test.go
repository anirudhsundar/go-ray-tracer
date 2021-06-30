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
