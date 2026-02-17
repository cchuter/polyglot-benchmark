# Scope: paasio Exercise (Issue #355)

## In Scope

- Implementing the solution in `go/exercises/practice/paasio/paasio.go`
- Implementing `NewReadCounter`, `NewWriteCounter`, and `NewReadWriteCounter` constructors
- Implementing thread-safe byte/operation counting via `ReadCount()` and `WriteCount()`
- Ensuring all tests in `paasio_test.go` pass
- Ensuring `go vet` passes

## Out of Scope

- Modifying `interface.go` (defines the required interfaces, read-only)
- Modifying `paasio_test.go` (test file, read-only)
- Modifying `go.mod`
- Adding any new files beyond `paasio.go`
- Changes to exercises other than `paasio`

## Dependencies

- Go standard library: `io`, `sync`
- No external packages required
- The `interface.go` file defines `ReadCounter`, `WriteCounter`, and `ReadWriteCounter` interfaces
