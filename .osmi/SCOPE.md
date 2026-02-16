# Scope: paasio Exercise

## In Scope

- Implement `NewReadCounter`, `NewWriteCounter`, `NewReadWriteCounter` constructors in `paasio.go`
- Implement backing struct types with thread-safe byte/operation counting
- Implement `Read`, `Write`, `ReadCount`, `WriteCount` methods
- Ensure all tests pass and `go vet` is clean

## Out of Scope

- Modifying `interface.go` (defines the interfaces, read-only)
- Modifying `paasio_test.go` (test file, read-only)
- Modifying `go.mod`
- Adding external dependencies
- Error handling beyond what the underlying reader/writer provides
- Any files outside `go/exercises/practice/paasio/`

## Dependencies

- Go standard library: `io`, `sync`
- No external packages required
