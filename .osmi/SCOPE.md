# Scope: polyglot-go-counter (Issue #201)

## In Scope

- Verifying and ensuring the test suite in `counter_test.go` correctly detects bugs in all three incorrect implementations (Impl1, Impl2, Impl3)
- Ensuring the test suite passes for the correct implementation (Impl4)
- The test file `go/exercises/practice/counter/counter_test.go` is the primary deliverable
- The stub file `go/exercises/practice/counter/counter.go` (package declaration only)

## Out of Scope

- Modifying the provided fixture files: `interface.go`, `maker.go`, `impl1.go`, `impl2.go`, `impl3.go`, `impl4.go`, `go.mod`
- Adding new implementation files
- Changing the Counter interface or makeCounter factory
- Performance optimization
- Adding benchmarks or fuzz tests

## Dependencies

- Go toolchain (go 1.18+, per go.mod)
- No external packages required (standard library only: `testing`, `unicode`)
- Fixture files: `interface.go` (Counter interface), `maker.go` (factory function), `impl1-4.go` (implementations under test)
