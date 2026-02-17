# Goal: polyglot-go-paasio

## Problem Statement

Implement the `paasio` Go exercise: create IO wrappers that track network/filesystem usage statistics. The stub file `paasio.go` currently contains only the package declaration and needs a complete implementation.

The interfaces `ReadCounter`, `WriteCounter`, and `ReadWriteCounter` are already defined in `interface.go`. We must provide concrete types and constructor functions that satisfy these interfaces.

## Acceptance Criteria

1. **`NewReadCounter(r io.Reader) ReadCounter`** — wraps an `io.Reader`, delegates `Read()` calls, and tracks total bytes read and number of read operations.
2. **`NewWriteCounter(w io.Writer) WriteCounter`** — wraps an `io.Writer`, delegates `Write()` calls, and tracks total bytes written and number of write operations.
3. **`NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter`** — wraps an `io.ReadWriter`, providing both read and write counting.
4. **Thread safety** — `ReadCount()` and `WriteCount()` must return correct results even when `Read()`/`Write()` are being called concurrently from multiple goroutines.
5. **All tests pass** — `go test -cpu 2 -race ./...` in the `paasio` directory must pass with zero failures.
6. **Only `paasio.go` is modified** — the solution must not alter `interface.go`, `paasio_test.go`, or `go.mod`.

## Key Constraints

- Go 1.18 module (`go.mod` specifies `go 1.18`).
- Must use synchronization primitives (e.g., `sync.Mutex`) to ensure concurrent correctness.
- The consistency tests verify that the bytes/ops ratio is always exact (bytes = ops * bytesPerOp), so bytes and ops must be updated atomically together under a single lock.
