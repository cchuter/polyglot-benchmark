# Goal: Implement paasio Exercise (Issue #355)

## Problem Statement

Implement the `paasio` Go exercise from Exercism. The exercise requires creating wrappers for network connections and files that report IO statistics — specifically, the total number of bytes read/written and the total number of read/write operations. These wrappers must be thread-safe.

## Acceptance Criteria

1. **`NewReadCounter(r io.Reader) ReadCounter`** — returns a wrapper around an `io.Reader` that tracks bytes read and number of read operations.
2. **`NewWriteCounter(w io.Writer) WriteCounter`** — returns a wrapper around an `io.Writer` that tracks bytes written and number of write operations.
3. **`NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter`** — returns a wrapper around an `io.ReadWriter` that tracks both read and write statistics.
4. **`ReadCount() (n int64, nops int)`** — returns correct cumulative bytes read and operation count, even under concurrent access.
5. **`WriteCount() (n int64, nops int)`** — returns correct cumulative bytes written and operation count, even under concurrent access.
6. **Thread safety** — `ReadCount()` and `WriteCount()` must return consistent results even when `Read()` and `Write()` are called concurrently from multiple goroutines.
7. All tests in `paasio_test.go` pass (`go test ./...`).
8. Code passes `go vet ./...`.

## Key Constraints

- The solution must be written in a single file: `paasio.go`.
- The interfaces `ReadCounter`, `WriteCounter`, and `ReadWriteCounter` are defined in `interface.go` and must not be modified.
- The solution must handle concurrent access correctly (tests use 8000 goroutines).
- The consistency tests require that bytes and ops are updated atomically together (reading `ReadCount`/`WriteCount` while concurrent operations are in-flight must see `nops == n / numBytes`).
