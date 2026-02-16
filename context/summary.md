# Context: paasio Implementation (Issue #227)

## Key Decisions
- Selected Branch 1 (shared counter struct with embedding) — simplest, DRY, matches reference solution
- Used `sync.Mutex` (not atomics) for thread safety — required for consistent byte+ops pair updates
- `rwCounter` embeds interface types (`WriteCounter`, `ReadCounter`) not concrete types — idiomatic Go

## Architecture
- `counter` struct: shared base with `bytes int64`, `ops int`, `mutex *sync.Mutex`
- `readCounter`: wraps `io.Reader` + embeds `counter`
- `writeCounter`: wraps `io.Writer` + embeds `counter`
- `rwCounter`: embeds both interface types, composes independent read/write counters

## Files Modified
- `go/exercises/practice/paasio/paasio.go` — complete solution (86 lines)

## Test Results
- 13/13 tests pass including concurrency consistency tests with 8000 goroutines
- Build clean, no warnings

## Branch
- Feature branch: `issue-227`
- Pushed to origin
