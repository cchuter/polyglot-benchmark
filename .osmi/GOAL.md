# Goal: polyglot-go-matrix

## Problem Statement

Implement the `matrix` exercise in Go for the polyglot benchmark suite. Given a string representing a matrix of numbers (rows separated by newlines, columns separated by spaces), provide functionality to:

1. Parse the string into a matrix representation
2. Return the rows of the matrix
3. Return the columns of the matrix
4. Set the value of an element at a given row and column

## Acceptance Criteria

1. **`New(s string) (Matrix, error)`** — Parse a string into a Matrix:
   - Split input by newlines into rows, split each row by spaces into integer elements
   - Handle leading/trailing whitespace in rows (e.g., `" 8 7 6"`)
   - Return an error for: uneven rows, empty rows, non-integer values, int64 overflow
   - Return a non-nil Matrix on success

2. **`Matrix.Rows() [][]int`** — Return all rows as `[][]int`:
   - Must return an independent deep copy (modifying returned data must not affect the matrix)

3. **`Matrix.Cols() [][]int`** — Return all columns as `[][]int`:
   - Must return an independent deep copy (modifying returned data must not affect the matrix)

4. **`Matrix.Set(row, col, val int) bool`** — Set a value at position (row, col):
   - Zero-based indexing
   - Return `true` if the position is valid, `false` if out of bounds
   - Must be reflected in subsequent `Rows()` and `Cols()` calls

5. All existing tests in `matrix_test.go` must pass (`go test ./...`)
6. Code must pass `go vet ./...`

## Key Constraints

- `Matrix` type must be nilable (used as `var matrix Matrix` and compared to `nil` in benchmarks)
- The `Matrix` type is used as both the return type of `New` and in `var` declarations
- Rows and columns are zero-based
- Module is `matrix` with `go 1.18`
