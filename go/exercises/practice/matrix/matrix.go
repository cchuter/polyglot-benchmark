package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix of integers.
type Matrix [][]int

// New parses a string representation into a Matrix.
func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	m := make(Matrix, len(lines))
	var numCols int
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			return nil, fmt.Errorf("row %d is empty", i)
		}
		if i == 0 {
			numCols = len(fields)
		} else if len(fields) != numCols {
			return nil, fmt.Errorf("row %d has %d columns, expected %d", i, len(fields), numCols)
		}
		row := make([]int, len(fields))
		for j, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("invalid value at row %d, col %d: %w", i, j, err)
			}
			row[j] = val
		}
		m[i] = row
	}
	return m, nil
}

// Rows returns a deep copy of all rows.
func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m))
	for i, row := range m {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}
	return result
}

// Cols returns a deep copy of all columns.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}
	numCols := len(m[0])
	result := make([][]int, numCols)
	for j := 0; j < numCols; j++ {
		col := make([]int, len(m))
		for i := range m {
			col[i] = m[i][j]
		}
		result[j] = col
	}
	return result
}

// Set sets the value at (row, col). Returns false if out of bounds.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
		return false
	}
	if col < 0 || col >= len(m[row]) {
		return false
	}
	m[row][col] = val
	return true
}
