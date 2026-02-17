# Goal: polyglot-go-matrix

## Problem Statement

Implement the `matrix` Go exercise: given a string representing a matrix of numbers (rows separated by newlines, columns separated by spaces), provide functions to extract rows, columns, and set individual element values.

The solution file `go/exercises/practice/matrix/matrix.go` currently contains only the package declaration. It must be populated with working code that passes all tests in `matrix_test.go`.

## Acceptance Criteria

1. **`New(s string) (Matrix, error)`** — Parse a string into a Matrix. Return an error for:
   - Uneven row lengths
   - Empty rows (leading/trailing/middle newlines producing empty lines)
   - Non-integer values (floats, text)
   - Integer overflow (values exceeding int64 range via `strconv.Atoi`)

2. **`Matrix.Rows() [][]int`** — Return all rows as `[][]int`. Must return an independent copy (modifying the returned slice must not affect the matrix).

3. **`Matrix.Cols() [][]int`** — Return all columns as `[][]int`. Must return an independent copy.

4. **`Matrix.Set(row, col, val int) bool`** — Set element at (row, col) to val. Return `true` on success, `false` if row/col is out of bounds (including negative indices).

5. **`Matrix` type** — Must be usable as an interface type in the benchmark (`var matrix Matrix` and `matrix == nil` check), meaning `Matrix` should be defined as a named slice type or a pointer-receivable type that supports nil comparison.

6. All tests in `matrix_test.go` pass: `go test ./...` from the exercise directory.
7. `go vet ./...` reports no issues.

## Key Constraints

- Rows and columns are zero-indexed.
- The `Matrix` type must be a type that can be compared to `nil` (the benchmarks do `matrix == nil`).
- `Rows()` and `Cols()` must return deep copies — tests verify mutation independence.
- The solution must only modify `matrix.go`; test file is read-only.
- Module is `matrix` with Go 1.18.
