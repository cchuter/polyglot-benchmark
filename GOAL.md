# Goal: polyglot-go-matrix

## Problem Statement

Implement a Go `matrix` package that parses a string representation of a matrix of integers and provides methods to:

1. Retrieve rows (left-to-right, top-to-bottom)
2. Retrieve columns (top-to-bottom, left-to-right)
3. Set the value of an element at a given (row, col) position

The matrix string uses newlines to separate rows and spaces to separate values within a row.

## Acceptance Criteria

1. **`New(s string) (Matrix, error)`** — Constructor that parses a string into a Matrix.
   - Splits on newlines for rows, spaces for columns
   - Returns an error for: uneven row lengths, non-integer values, int64 overflow, empty rows (leading/trailing/middle newlines producing empty lines)
   - Returns a non-nil `*Matrix` (or `Matrix` interface value) on success

2. **`Rows() [][]int`** — Returns all rows as a 2D slice.
   - Must return an independent copy (modifying the returned slice must not affect the matrix)

3. **`Cols() [][]int`** — Returns all columns as a 2D slice.
   - Must return an independent copy (modifying the returned slice must not affect the matrix)

4. **`Set(row, col, val int) bool`** — Sets an element at (row, col) to val.
   - Returns `true` on success
   - Returns `false` for out-of-bounds indices (negative or beyond dimensions)
   - Zero-based indexing

5. All existing tests in `matrix_test.go` must pass, including benchmarks.

## Key Constraints

- The `Matrix` type must be `type Matrix [][]int` (as the test file uses `Matrix` as both a type and via `nil` comparison)
- The solution goes in `go/exercises/practice/matrix/matrix.go`
- The test file `matrix_test.go` is read-only and must not be modified
- Module is `matrix` with Go 1.18
- Rows and Cols must return independent copies of the data
