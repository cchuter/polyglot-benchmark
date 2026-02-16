# Scope: polyglot-go-matrix

## In Scope

- Implement `Matrix` type in `go/exercises/practice/matrix/matrix.go`
- Implement `New()` constructor, `Rows()`, `Cols()`, and `Set()` methods
- Pass all tests in `matrix_test.go` (unit tests and benchmarks)
- Pass `go vet`

## Out of Scope

- Modifying `matrix_test.go` or `go.mod`
- Adding additional methods beyond what the tests require
- Implementing matrix operations (addition, multiplication, etc.)
- Supporting non-integer matrix elements
- Any files outside `go/exercises/practice/matrix/`

## Dependencies

- Go standard library only (`strconv`, `strings`, `errors`)
- No external packages
- `go 1.18` as specified in `go.mod`
