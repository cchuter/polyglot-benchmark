# Scope: polyglot-go-counter

## In Scope

- Writing `counter_test.go` with comprehensive test functions
- Ensuring tests correctly validate the Counter interface behavior
- Tests that distinguish the correct implementation (Impl4) from buggy ones (Impl1-3)
- Test coverage for: zero state, empty strings, ASCII text, Unicode text, newline handling, multiple AddString calls

## Out of Scope

- Modifying `interface.go`, `maker.go`, `impl1.go`–`impl4.go`
- Modifying `go.mod` or `.meta/` files
- Adding external dependencies
- Writing benchmarks (optional, not required)
- Modifying `counter.go` beyond what's needed (it may stay as just `package counter`)

## Dependencies

- `maker.go` provides `makeCounter()` factory that reads `COUNTER_IMPL` env var
- `interface.go` defines the `Counter` interface
- `impl1.go`–`impl4.go` provide the four implementations to validate against
- Go 1.18+ toolchain
- Standard library only (`testing`, `unicode` if needed)
