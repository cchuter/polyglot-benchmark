# Scope: polyglot-go-counter

## In Scope

- Writing/verifying the test suite in `go/exercises/practice/counter/counter_test.go`
- Ensuring `counter.go` is a clean package declaration stub
- Verifying all four implementations behave as expected against the test suite
- Ensuring the code compiles and passes `go vet`

## Out of Scope

- Modifying the implementation files (`impl1.go` through `impl4.go`)
- Modifying the interface file (`interface.go`)
- Modifying the factory file (`maker.go`)
- Modifying `go.mod`
- Adding any new dependencies
- Changes to exercises outside the `counter` directory

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Standard library only (`testing`, `unicode`)
- Pre-existing files: `interface.go`, `maker.go`, `impl1.go`–`impl4.go`
