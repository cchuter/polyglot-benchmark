# Goal: Implement paasio Exercise (Go)

## Problem Statement

Implement IO statistics wrappers for a PaaS billing system. The `paasio` package needs concrete types that wrap `io.Reader`, `io.Writer`, and `io.ReadWriter` to track the total number of bytes and operations for reads and writes.

The interfaces (`ReadCounter`, `WriteCounter`, `ReadWriteCounter`) are already defined in `interface.go`. The stub file `paasio.go` exists but is empty (only contains the package declaration). The solution must implement three constructor functions and their backing types.

## Acceptance Criteria

1. `NewReadCounter(io.Reader) ReadCounter` — wraps a reader, tracks bytes read and read operations
2. `NewWriteCounter(io.Writer) WriteCounter` — wraps a writer, tracks bytes written and write operations
3. `NewReadWriteCounter(io.ReadWriter) ReadWriteCounter` — wraps a read-writer, tracks both read and write stats
4. All counters must be **thread-safe**: concurrent reads/writes must produce consistent counts
5. `ReadCount()` returns `(totalBytes int64, totalOps int)` — must be consistent even during concurrent `Read()` calls
6. `WriteCount()` returns `(totalBytes int64, totalOps int)` — must be consistent even during concurrent `Write()` calls
7. All existing tests in `paasio_test.go` pass (`go test ./...`)
8. Code passes `go vet ./...`

## Key Constraints

- Must use the `sync.Mutex` (or equivalent synchronization) to ensure thread safety
- The `ReadCount`/`WriteCount` methods must return atomically consistent snapshots (bytes and ops must correspond)
- The package name is `paasio` and module is `paasio` with Go 1.18
- Only `paasio.go` should be modified; `interface.go` and test files are read-only
