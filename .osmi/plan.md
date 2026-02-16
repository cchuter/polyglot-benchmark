# Implementation Plan: polyglot-go-matrix

## File to Modify

- `go/exercises/practice/matrix/matrix.go` — the only file that needs changes

## Type Design

Define `Matrix` as a named slice type:

```go
type Matrix [][]int
```

**Rationale**: The benchmark test declares `var matrix Matrix` and then checks `matrix == nil`. This requires `Matrix` to be a nil-comparable type. A named slice type satisfies this — its zero value is nil, it supports method definitions, and `New()` can return `(Matrix, error)` directly. A struct type wouldn't work because structs can't be compared to nil.

## Functions and Methods

### `New(s string) (Matrix, error)`

1. Split input string by `"\n"` to get row strings
2. For each row string, use `strings.Fields()` to split by whitespace (handles leading/trailing spaces)
3. Validate: if any row has zero fields, return error (catches empty rows)
4. Validate: all rows must have the same number of columns as the first row
5. Parse each field with `strconv.Atoi()` — this catches non-integers, floats, overflow
6. Build and return the `Matrix` (a `[][]int`)

### `Rows() [][]int`

Return a deep copy of the matrix data. For each row, allocate a new slice and `copy()` the values. This ensures the test's mutation-independence check passes.

### `Cols() [][]int`

Build column slices by iterating column-first. For each column index j, create a new slice containing `m[i][j]` for all rows i.

### `Set(row, col, val int) bool`

1. Bounds-check: row must be in `[0, len(m))` and col in `[0, len(m[row]))`
2. If out of bounds, return `false`
3. Set `m[row][col] = val` and return `true`

Since `Matrix` is a slice type, the value receiver `(m Matrix)` still allows modification of elements (slices are reference types — the receiver holds a reference to the underlying array).

## Imports

Only standard library: `errors`, `strconv`, `strings`

## Order of Implementation

1. Write the complete `matrix.go` file
2. Run `go test ./...` from the matrix directory
3. Run `go vet ./...`
4. Fix any issues and repeat
