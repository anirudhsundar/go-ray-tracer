package coremath

import (
	"bytes"
	"fmt"
)

type Matrix struct {
	M, N int
	Data []float64
}

func (m *Matrix) At(row, col int) float64 {
	return m.Data[row*m.M+col]
}

func (m *Matrix) Set(val float64, row, col int) {
	m.Data[row*m.M+col] = val
}

func NewMatrix(vals []float64, rows, cols int) (*Matrix, error) {
	m := Matrix{M: rows, N: cols}
	if len(vals) != m.M*m.N {
		return nil, fmt.Errorf("Expected slice of size %v as row major input for matrix of dimensions %vx%v", m.M*m.N, m.M, m.N)
	}
	m.Data = make([]float64, len(vals))
	for i := 0; i < len(vals); i++ {
		m.Data[i] = vals[i]
	}
	return &m, nil
}

func (m *Matrix) String() string {
	var b bytes.Buffer
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			b.WriteString(fmt.Sprintf("%f", m.At(i, j)))
			if j != m.N-1 {
				b.WriteString("\t")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m1 *Matrix) CheckEqual(m2 *Matrix) (bool, error) {
	if m1.M != m2.M || m1.N != m2.M {
		return false, fmt.Errorf("Incompatible dimensions for comparison (%d, %d) != (%d, %d)", m1.M, m1.N, m2.M, m2.N)
	}
	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			if !FloatEq(m1.At(i, j), m2.At(i, j)) {
				return false, fmt.Errorf("Elements at index (%d, %d) not matching: %f != %f", i, j, m1.At(i, j), m2.At(i, j))
			}
		}
	}
	return true, nil
}

func (m1 *Matrix) Equal(m2 *Matrix) bool {
	if m1.M != m2.M || m1.N != m2.M {
		return false
	}
	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			if !FloatEq(m1.At(i, j), m2.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (m1 *Matrix) Multiply(m2 *Matrix) (*Matrix, error) {
	if m1.N != m2.M {
		return nil, fmt.Errorf("Incompatible dimensions for product %d != %d for matrices with dimensions (%d, %d) & (%d, %d)", m1.N, m2.M, m1.M, m1.N, m2.M, m2.N)
	}
	result := Matrix{M: m1.M, N: m2.N}
	result.Data = make([]float64, m1.M*m2.N)
	for i := 0; i < m1.M; i++ {
		for j := 0; j < m1.N; j++ {
			for k := 0; k < m2.M; k++ {
				result.Data[i*m1.M+j] += m1.At(i, k) * m2.At(k, j)
			}
		}
	}
	return &result, nil
}
