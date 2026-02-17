# Context: polyglot-go-counter

## Key Decisions

- No code changes needed — existing `counter_test.go` already has a correct, comprehensive test suite
- The counter exercise is special: tests ARE the solution, implementations are pre-provided
- Kept individual test functions rather than rewriting as table-driven (lower risk, no functional benefit)

## Files

- `counter.go` — Package declaration only (`package counter`)
- `counter_test.go` — 8 test functions + `assertCounts()` helper (unchanged, already correct)
- `interface.go` — `Counter` interface (pre-provided, not modified)
- `maker.go` — Factory function using `COUNTER_IMPL` env var (pre-provided, not modified)
- `impl1.go` through `impl4.go` — Test implementations (pre-provided, not modified)

## Test Results

| COUNTER_IMPL | Result | Bugs Detected |
|---|---|---|
| 4 | PASS (9/9) | N/A (correct impl) |
| 1 | FAIL (4 failures) | Wrong line counting |
| 2 | FAIL (1 failure) | ASCII-only letter detection |
| 3 | FAIL (1 failure) | Byte-level iteration |

## Branch

`issue-329` pushed to origin
