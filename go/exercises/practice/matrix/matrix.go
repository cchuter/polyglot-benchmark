package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a matrix of integers as a slice of rows.
type Matrix [][]int

// New parses a string representation of a matrix.
// Rows are separated by newlines, columns by spaces.
func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	m := make(Matrix, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			return nil, errors.New("empty row")
		}
		if i > 0 && len(fields) != len(m[0]) {
			return nil, errors.New("uneven rows")
		}
		row := make([]int, len(fields))
		for j, f := range fields {
			val, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			row[j] = val
		}
		m[i] = row
	}
	return m, nil
}

// Rows returns a copy of the matrix data as rows.
func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m))
	for i, row := range m {
		result[i] = append([]int{}, row...)
	}
	return result
}

// Cols returns the matrix data as columns.
func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}
	ncols := len(m[0])
	result := make([][]int, ncols)
	for c := 0; c < ncols; c++ {
		col := make([]int, len(m))
		for r := range m {
			col[r] = m[r][c]
		}
		result[c] = col
	}
	return result
}

// Set sets the value at (row, col). Returns false if out of bounds.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
