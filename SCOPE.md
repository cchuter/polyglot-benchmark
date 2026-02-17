# Scope: polyglot-go-paasio

## In Scope

- Implement `NewReadCounter`, `NewWriteCounter`, and `NewReadWriteCounter` in `paasio.go`.
- Implement supporting struct types (`readCounter`, `writeCounter`, `rwCounter`, and a shared `counter` helper).
- Ensure thread-safe counting using `sync.Mutex`.
- Pass all existing tests in `paasio_test.go`.

## Out of Scope

- Modifying `interface.go`, `paasio_test.go`, `go.mod`, or any file outside `paasio.go`.
- Adding new test files or benchmarks.
- Changing the module name or Go version.
- Implementing features beyond what the interfaces and tests require.

## Dependencies

- Go standard library only: `io`, `sync`.
- No external packages required.
- Existing `interface.go` defines `ReadCounter`, `WriteCounter`, `ReadWriteCounter` interfaces.
