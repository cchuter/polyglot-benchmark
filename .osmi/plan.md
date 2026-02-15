# Implementation Plan: polyglot-go-beer-song (Issue #15)

## Current State

The beer-song exercise already has a complete implementation at `go/exercises/practice/beer-song/` that passes all tests. The implementation includes:

- `beer_song.go` — Three functions: `Song()`, `Verses()`, `Verse()`
- `beer_song_test.go` — Tests covering individual verses, verse ranges, full song, and error cases
- `go.mod` — Module `beer` with Go 1.18

All 11 test cases pass, including benchmarks.

## Plan

### Step 1: Create feature branch

```
git checkout -b issue-15
```

### Step 2: Verify implementation correctness

Run full test suite to confirm all tests pass:
```
cd go/exercises/practice/beer-song && go test -v ./...
```

Verify the implementation handles all edge cases:
- Standard verses (3-99): plural "bottles", "Take one down"
- Verse 2: transition to singular "1 bottle"
- Verse 1: singular "bottle", "Take it down", "no more bottles"
- Verse 0: "No more bottles", "Go to the store"
- Invalid inputs: returns error for out-of-range values

### Step 3: Verify code quality

- Run `go vet` for static analysis
- Run `gofmt` to verify formatting
- Confirm no unnecessary dependencies

### Step 4: Commit and push

Create a commit on the `issue-15` branch that closes the issue.

## Files to Create or Modify

No code changes are expected — the existing implementation is complete and correct. The work is verification and branch/commit management.

- `go/exercises/practice/beer-song/beer_song.go` — Verify (no changes expected)
- `go/exercises/practice/beer-song/beer_song_test.go` — Read-only verification
- `go/exercises/practice/beer-song/go.mod` — Verify (no changes expected)

## Architectural Decisions

1. **No code changes needed** — The existing implementation is correct, idiomatic Go, and passes all tests.
2. **Use feature branch workflow** — Create `issue-15` branch per the standard workflow.
3. **Test file is read-only** — Following exercism conventions, the test file defines the exercise spec.

## Risks

- Low risk: the implementation already exists and passes all tests
- If any issues are found during review, they will be addressed in the implementation phase
