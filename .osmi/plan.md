# Implementation Plan: polyglot-go-alphametics (Issue #21)

## Current State

The alphametics exercise has a complete implementation at `go/exercises/practice/alphametics/` that passes all 10 test cases in ~1.2 seconds. The implementation includes:

- `alphametics.go` — `Solve` function with permutation-based brute-force solver, leading-zero validation, and carry propagation
- `alphametics_test.go` — Test runner and benchmark
- `cases_test.go` — 10 test cases covering 3-10 letter puzzles, error cases, and a 199-addend stress test
- `go.mod` — Module `alphametics` with Go 1.18

The implementation was previously fixed in commit `80330aa` (PR #14) to add comprehensive leading-zero checks for all multi-digit words (not just the answer) and carry overflow validation.

## Plan

### Step 1: Create feature branch

```
git checkout -b issue-21
```

### Step 2: Verify implementation correctness

Run full test suite:
```
cd go/exercises/practice/alphametics && go test -v -timeout 300s
```

Verify:
- All 10 test cases pass
- Benchmarks run without error
- `go vet` reports no issues
- Code is properly formatted with `gofmt`

### Step 3: Code review

Review `alphametics.go` against the exercise specification:
- Function signature matches: `func Solve(puzzle string) (map[string]int, error)`
- Leading-zero constraint enforced for all multi-digit words
- Each letter maps to a unique digit (ensured by permutation-based approach)
- Error returned when no solution exists

### Step 4: Commit on feature branch

Commit the verified solution on `issue-21` branch with a descriptive message referencing issue #21.

## Files to Create or Modify

No code changes are expected — the existing implementation is complete and correct. The work is verification and branch/commit management.

- `go/exercises/practice/alphametics/alphametics.go` — Verify (no changes expected)
- `go/exercises/practice/alphametics/alphametics_test.go` — Read-only verification
- `go/exercises/practice/alphametics/cases_test.go` — Read-only verification
- `go/exercises/practice/alphametics/go.mod` — Verify (no changes expected)

## Architectural Decisions

1. **No code changes needed** — The existing implementation is correct, handles all edge cases, and passes all tests.
2. **Use feature branch workflow** — Create `issue-21` branch per the standard workflow.
3. **Test files are read-only** — Following exercism conventions, test files define the exercise spec.

## Approach Rationale

The implementation uses a permutation-based brute-force approach:
- Parse puzzle into structured `problem` type with column-indexed digits
- Generate all P(10, n) permutations where n = number of unique letters
- For each permutation, check column-by-column addition with carry propagation
- Validate no leading zeros on any multi-digit word
- Return first valid solution or error if exhausted

This is the standard approach for this exercise and handles all test cases within acceptable time limits.
