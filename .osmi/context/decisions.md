# Key Decisions

## Type Design: `type Matrix [][]int`

Chose a named slice type over an interface or struct wrapper because:
- Slices are nilable (required by benchmark tests)
- Value receiver `Set` works because slice elements share underlying storage
- Minimal code, idiomatic Go
- Matches Exercism's canonical reference solution

## Parsing: `strings.Fields` over `strings.Split`

Used `strings.Fields` to split each row, which handles leading/trailing whitespace and multiple spaces between values. This correctly handles the test case `" 8 7 6"` (leading space).

## Deep Copy Strategy

Both `Rows()` and `Cols()` allocate fresh slices with `make` + `copy` to ensure returned data is independent of the matrix's internal storage. Tests explicitly verify this by mutating returned slices and checking the matrix is unchanged.

## Files Modified

- `go/exercises/practice/matrix/matrix.go` — complete implementation (was empty stub)
