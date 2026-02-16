# Scope: paasio Exercise (Issue #227)

## In Scope

- Implementing the solution in `go/exercises/practice/paasio/paasio.go`
- Three constructor functions: `NewReadCounter`, `NewWriteCounter`, `NewReadWriteCounter`
- Backing struct types with thread-safe byte/operation counting
- Ensuring all tests in `paasio_test.go` pass

## Out of Scope

- Modifying `interface.go` (read-only, defines the interfaces)
- Modifying `paasio_test.go` (read-only, defines the tests)
- Modifying `go.mod`
- Any files outside the `paasio` exercise directory
- Adding benchmarks or additional tests
- Performance optimization beyond what's needed to pass tests

## Dependencies

- Go standard library: `io`, `sync`
- No external packages required
- Interfaces defined in `interface.go`: `ReadCounter`, `WriteCounter`, `ReadWriteCounter`
