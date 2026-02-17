# Context Summary: paasio (Issue #355)

## Status: Complete

## Solution
Implemented in `go/exercises/practice/paasio/paasio.go` using a mutex-based counter pattern:
- `counter` struct holds `bytes int64`, `ops int`, and `*sync.Mutex`
- `readCounter` wraps `io.Reader`, delegates reads and tracks stats
- `writeCounter` wraps `io.Writer`, delegates writes and tracks stats
- `rwCounter` composes `WriteCounter` and `ReadCounter` for `ReadWriteCounter`

## Key Decisions
- Used mutex (not atomics) to guarantee consistency between bytes and ops counts
- Solution matches the reference solution in `.meta/example.go`

## Files Modified
- `go/exercises/practice/paasio/paasio.go` — full solution (87 lines)

## Test Results
- 13/13 tests pass with `-race` flag
- `go vet` clean

## Branch
- `issue-355` pushed to origin
