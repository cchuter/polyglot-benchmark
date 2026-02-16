package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a matrix of integers stored as rows of int slices.
type Matrix [][]int

// New creates a Matrix from a string where rows are separated by newlines
// and values within a row are separated by whitespace.
func New(s string) (Matrix, error) {
	lines := strings.Split(s, "\n")
	m := make(Matrix, 0, len(lines))
	numCols := -1

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			return nil, errors.New("empty row")
		}
		if numCols == -1 {
			numCols = len(fields)
		} else if len(fields) != numCols {
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
		m = append(m, row)
	}

	return m, nil
}

// Rows returns a deep copy of the matrix rows.
func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m))
	for i, row := range m {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}
	return result
}

// Cols returns the columns of the matrix as a deep copy.
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

// Set sets the value at the given row and column. Returns true if the
// indices are valid, false otherwise.
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
