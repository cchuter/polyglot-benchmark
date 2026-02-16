# Goal: polyglot-go-matrix

## Problem Statement

Implement the `matrix` Go exercise: given a string representing a matrix of numbers (rows separated by newlines, values separated by spaces), provide functionality to:

1. Parse the string into a Matrix type
2. Return rows (left-to-right, top-to-bottom)
3. Return columns (top-to-bottom, left-to-right)
4. Set the value of an element given its zero-based row and column indices

## Acceptance Criteria

1. **`New(s string) (Matrix, error)`** parses a string into a Matrix:
   - Splits rows by `\n`, values by whitespace
   - Returns error for: uneven rows, empty rows, non-integer values, integer overflow
   - Returns non-nil Matrix for valid input

2. **`Rows() [][]int`** returns all rows as independent copies:
   - Mutating the returned slice must NOT affect the matrix
   - Values read left-to-right, rows top-to-bottom

3. **`Cols() [][]int`** returns all columns as independent copies:
   - Mutating the returned slice must NOT affect the matrix
   - Values read top-to-bottom, columns left-to-right

4. **`Set(row, col, val int) bool`** sets a matrix element:
   - Returns `true` and updates the value for valid in-bounds indices
   - Returns `false` for out-of-bounds indices (negative or >= dimension)
   - After Set, both Rows() and Cols() reflect the change

5. All existing tests in `matrix_test.go` pass (`go test ./...`)
6. Code passes `go vet ./...`

## Key Constraints

- `Matrix` type must be nil-comparable (the benchmark declares `var matrix Matrix` and checks `matrix == nil`)
- Rows and columns use zero-based indexing
- `Rows()` and `Cols()` must return deep copies (tests verify mutation independence)
- Must handle leading/trailing whitespace in rows (test uses `" 8 7 6"`)
- Module is `matrix` with `go 1.18`
