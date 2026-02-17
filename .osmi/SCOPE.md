# Scope: polyglot-go-counter

## In Scope

- Writing/verifying `counter_test.go` with comprehensive test cases
- Ensuring `counter.go` has the correct package declaration
- Tests that pass with Impl4 and fail with Impl1, Impl2, Impl3
- Verifying `go test` and `go vet` pass cleanly

## Out of Scope

- Modifying `interface.go`, `maker.go`, or any `impl*.go` files
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Adding any new dependencies
- Writing implementation code (the exercise is about writing tests)

## Dependencies

- `interface.go` — defines the `Counter` interface
- `maker.go` — provides `makeCounter()` factory using `COUNTER_IMPL` env var
- `impl1.go` through `impl4.go` — pre-provided implementations to test against
- Go 1.18+ (as specified in go.mod)
