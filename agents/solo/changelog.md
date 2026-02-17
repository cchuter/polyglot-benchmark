# Change Log: polyglot-go-counter

## Summary

The counter exercise is a special exercise where the test suite IS the solution. The existing `counter_test.go` already contains a comprehensive, correct test suite that:
- Passes with COUNTER_IMPL=4 (correct implementation)
- Fails with COUNTER_IMPL=1 (detects incorrect line counting)
- Fails with COUNTER_IMPL=2 (detects ASCII-only letter detection)
- Fails with COUNTER_IMPL=3 (detects byte-level Unicode iteration bug)

## Files

- `counter.go` — Package declaration only (no implementation code needed)
- `counter_test.go` — Pre-existing comprehensive test suite with 8 tests and helper

## Verification

- `COUNTER_IMPL=4 go test ./...` — PASS
- `COUNTER_IMPL=1 go test ./...` — FAIL (4 test failures)
- `COUNTER_IMPL=2 go test ./...` — FAIL (1 test failure)
- `COUNTER_IMPL=3 go test ./...` — FAIL (1 test failure)
- `go vet ./...` — clean
