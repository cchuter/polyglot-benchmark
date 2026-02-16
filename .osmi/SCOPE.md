# Scope: polyglot-go-counter

## In Scope

- Writing `counter_test.go` with tests that:
  - Pass on `Impl4` (the correct implementation)
  - Fail on `Impl1` (wrong line counting)
  - Fail on `Impl2` (ASCII-only letter detection)
  - Fail on `Impl3` (byte-level iteration instead of rune-level)
- Verifying all tests pass with `COUNTER_IMPL=4 go test`
- Verifying all tests detect bugs in impls 1-3

## Out of Scope

- Modifying `interface.go`, `maker.go`, `impl1.go`–`impl4.go`, or `go.mod`
- Writing production code in `counter.go`
- Creating new implementation variants
- Modifying any files outside `go/exercises/practice/counter/`

## Dependencies

- `interface.go` — defines `Counter` interface with `AddString(string)`, `Lines() int`, `Letters() int`, `Characters() int`
- `maker.go` — `makeCounter()` factory that reads `COUNTER_IMPL` env var to select implementation
- `impl1.go`–`impl4.go` — four implementations, only impl4 is correct
- `go.mod` — module `counter`, Go 1.18
