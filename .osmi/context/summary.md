# Context: polyglot-go-counter

## Key Decisions

- The "solution" for this exercise is the test file (`counter_test.go`), not production code
- Used an `assertCounts` helper to reduce test boilerplate
- Individual test functions (not table-driven) for clarity since each tests a distinct scenario

## Files Modified

- `go/exercises/practice/counter/counter_test.go` — complete test suite (already present from previous PR #116)

## Files Read (not modified)

- `go/exercises/practice/counter/interface.go` — Counter interface
- `go/exercises/practice/counter/maker.go` — factory function selecting impl by env var
- `go/exercises/practice/counter/impl1.go` through `impl4.go` — four implementations
- `go/exercises/practice/counter/.meta/config.json` — exercise config
- `go/exercises/practice/counter/.meta/example.go` — reference solution
- `go/exercises/practice/counter/.docs/instructions.md` — exercise instructions

## Test Results

All acceptance criteria met. Tests pass on Impl4 and detect all three bugs.

## Branch

`issue-158` pushed to origin, ready for PR.
