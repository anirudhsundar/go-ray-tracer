package coremath

import (
	"bytes"
	"fmt"
	"math"
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

func NewEmptyMatrix(rows, cols int) *Matrix {
	val := Matrix{M: rows, N: cols, Data: make([]float64, rows*cols)}
	return &val
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

func IdentityMatrix(n int) *Matrix {
	m := NewEmptyMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				m.Set(1, i, j)
			} else {
				m.Set(0, i, j)
			}
		}
	}
	return m
}

func (m *Matrix) Copy() *Matrix {
	newM := NewEmptyMatrix(m.M, m.N)
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			newM.Set(m.At(i, j), i, j)
		}
	}
	return newM
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
	if m1.M != m2.M || m1.N != m2.N {
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

func (m1 *Matrix) MatrixMultiply(m2 *Matrix) (*Matrix, error) {
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

func (m *Matrix) TupleMultiply(t Tuple) (*Tuple, error) {
	if m.N != 4 {
		return nil, fmt.Errorf("Matrix should have only 4 columns")
	}
	result_vals := [4]float64{0, 0, 0, 0}
	tuple_vals := t.GetArray()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result_vals[i] += m.At(i, j) * tuple_vals[j]
		}
	}
	result := Tuple{result_vals[0], result_vals[1], result_vals[2], result_vals[3]}
	return &result, nil
}

func (m *Matrix) Transpose() *Matrix {
	trans := NewEmptyMatrix(m.N, m.M)
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			trans.Set(m.At(i, j), j, i)
		}
	}
	return trans
}

func (m *Matrix) Determinant2x2() float64 {
	return m.At(0, 0)*m.At(1, 1) - m.At(1, 0)*m.At(0, 1)
}

func (m *Matrix) SubMatrix(row, col int) (*Matrix, error) {
	if m.M != m.N {
		return nil, fmt.Errorf("SubMatrix can only be calculated for square matrices but got matrix with dims %dx%d", row, col)
	}

	outRow, outCol := m.M-1, m.N-1
	ret := NewEmptyMatrix(outRow, outCol)
	for i, x := 0, 0; i < m.M; i++ {
		for j, y := 0, 0; j < m.N; j++ {
			if i != row && j != col {
				ret.Set(m.At(i, j), x, y)
				if y == outRow-1 {
					y = 0
					x++
				} else {
					y++
				}
			}
		}
	}
	return ret, nil
}

func (m *Matrix) Minor(row, col int) (float64, error) {
	sub, err := m.SubMatrix(row, col)
	if err != nil {
		return 0, err
	}
	return sub.Determinant()
}

func (m *Matrix) Cofactor(row, col int) (float64, error) {
	minor, err := m.Minor(row, col)
	if err != nil {
		return 0, err
	}
	if (row+col)%2 == 0 {
		return minor, nil
	} else {
		return -minor, nil
	}
}

func (m *Matrix) Determinant() (float64, error) {
	if m.M == 2 && m.N == 2 {
		return m.Determinant2x2(), nil
	} else {
		det := 0.0
		for i := 0; i < m.M; i++ {
			cofactor, err := m.Cofactor(0, i)
			if err != nil {
				return math.NaN(), err
			}
			det += cofactor * m.At(0, i)
		}
		return det, nil
	}
}

func (m *Matrix) IsInvertible() (bool, error) {
	det, err := m.Determinant()
	if err != nil {
		return false, err
	}
	if det == 0 {
		return false, nil
	}
	return true, nil
}

func (m *Matrix) Inverse() (*Matrix, error) {
	isInvertible, err := m.IsInvertible()
	if err != nil {
		return nil, err
	}
	if !isInvertible {
		return nil, fmt.Errorf("Trying to compute inverse of non-invertible matrix \n%v\n", m)
	}
	inv := NewEmptyMatrix(m.N, m.M)
	det, err := m.Determinant()
	if err != nil {
		return nil, err
	}
	for i := 0; i < m.M; i++ {
		for j := 0; j < m.N; j++ {
			cofactor, err := m.Cofactor(i, j)
			if err != nil {
				return nil, err
			}
			inv.Set(cofactor/det, j, i)
		}
	}
	return inv, nil
}
