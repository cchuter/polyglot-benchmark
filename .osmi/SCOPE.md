# Scope: polyglot-go-matrix

## In Scope

- Implement `Matrix` type in `go/exercises/practice/matrix/matrix.go`
- Implement `New()` constructor with full input validation
- Implement `Rows()` method returning independent copy of row data
- Implement `Cols()` method returning independent copy of column data
- Implement `Set()` method with bounds checking
- Ensure all tests and benchmarks in `matrix_test.go` pass
- Ensure `go vet` passes

## Out of Scope

- Modifying `matrix_test.go`
- Modifying `go.mod`
- Adding additional files or packages
- Performance optimization beyond what's needed to pass benchmarks
- Any changes outside `go/exercises/practice/matrix/`

## Dependencies

- Go standard library only (`errors`, `strconv`, `strings`)
- No external packages required
- Go 1.18+ (per go.mod)
