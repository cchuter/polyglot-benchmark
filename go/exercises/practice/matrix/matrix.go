package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

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

func (m Matrix) Rows() [][]int {
	result := make([][]int, len(m))
	for i, row := range m {
		result[i] = make([]int, len(row))
		copy(result[i], row)
	}
	return result
}

func (m Matrix) Cols() [][]int {
	if len(m) == 0 {
		return nil
	}
	nCols := len(m[0])
	result := make([][]int, nCols)
	for c := 0; c < nCols; c++ {
		col := make([]int, len(m))
		for r := 0; r < len(m); r++ {
			col[r] = m[r][c]
		}
		result[c] = col
	}
	return result
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
