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
