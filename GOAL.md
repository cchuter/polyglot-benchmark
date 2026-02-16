# Goal: Implement paasio Exercise (Issue #227)

## Problem Statement

Implement the `paasio` Go exercise — a PaaS IO statistics wrapper. The solution file `paasio.go` currently contains only a package declaration. We need to implement three constructor functions and their backing types that wrap `io.Reader`, `io.Writer`, and `io.ReadWriter` to track byte counts and operation counts.

## What Needs to Be Built

Implement in `go/exercises/practice/paasio/paasio.go`:

1. **`NewReadCounter(r io.Reader) ReadCounter`** — wraps an `io.Reader`, tracking bytes read and read operations
2. **`NewWriteCounter(w io.Writer) WriteCounter`** — wraps an `io.Writer`, tracking bytes written and write operations
3. **`NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter`** — wraps an `io.ReadWriter`, combining read and write counting

The interfaces `ReadCounter`, `WriteCounter`, and `ReadWriteCounter` are already defined in `interface.go`.

## Acceptance Criteria

1. All tests in `paasio_test.go` pass (`go test -cpu 2 ./...`)
2. `NewReadCounter` correctly delegates reads and tracks total bytes + operation count
3. `NewWriteCounter` correctly delegates writes and tracks total bytes + operation count
4. `NewReadWriteCounter` correctly delegates both reads and writes with independent counters
5. `ReadCount()` and `WriteCount()` return correct values even under concurrent access (thread-safe)
6. Consistency tests pass: when `ReadCount()` or `WriteCount()` is called concurrently with `Read()`/`Write()`, the byte count and operation count are always consistent (bytes = ops * bytesPerOp)
7. Only `paasio.go` is modified (the solution file per `.meta/config.json`)

## Key Constraints

- Must use `sync.Mutex` or equivalent synchronization for thread safety
- The `count()` and `addBytes()` operations must be atomic with respect to each other (both bytes and ops updated under the same lock)
- Must work with Go 1.18 (per `go.mod`)
