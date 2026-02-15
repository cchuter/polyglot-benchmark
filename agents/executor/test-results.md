# Executor Test Results

**Date**: 2026-02-15
**Package**: `bottlesong` (go/exercises/practice/bottle-song/)

## 1. `go build ./...`

**Result**: PASS (no errors)

## 2. `go test -v`

**Result**: PASS (all 7 tests pass)

```
=== RUN   TestRecite
=== RUN   TestRecite/first_generic_verse
=== RUN   TestRecite/last_generic_verse
=== RUN   TestRecite/verse_with_2_bottles
=== RUN   TestRecite/verse_with_1_bottle
=== RUN   TestRecite/first_two_verses
=== RUN   TestRecite/last_three_verses
=== RUN   TestRecite/all_verses
--- PASS: TestRecite (0.00s)
    --- PASS: TestRecite/first_generic_verse (0.00s)
    --- PASS: TestRecite/last_generic_verse (0.00s)
    --- PASS: TestRecite/verse_with_2_bottles (0.00s)
    --- PASS: TestRecite/verse_with_1_bottle (0.00s)
    --- PASS: TestRecite/first_two_verses (0.00s)
    --- PASS: TestRecite/last_three_verses (0.00s)
    --- PASS: TestRecite/all_verses (0.00s)
PASS
ok  	bottlesong	0.004s
```

## 3. `go vet ./...`

**Result**: PASS (no issues)

## 4. `staticcheck ./...`

**Result**: PASS (no issues)

## Summary

| Check         | Result |
|---------------|--------|
| go build      | PASS   |
| go test       | PASS (7/7) |
| go vet        | PASS   |
| staticcheck   | PASS   |

All checks pass with zero errors or warnings.
