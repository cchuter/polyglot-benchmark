# Context Summary: bottle-song Exercise

## Issue
#195 — polyglot-go-bottle-song: Implement the "Ten Green Bottles" children's song reciter.

## Solution
Implemented `Recite(startBottles, takeDown int) []string` in `go/exercises/practice/bottle-song/bottle_song.go`.

### Architecture
- `numberToWord` map: 0→"no", 1→"one", ..., 10→"ten"
- `verse(n int) []string`: switch on n==1 (singular+no), n==2 (singular result), default (fmt.Sprintf)
- `Recite`: loop from startBottles down for takeDown iterations, join with "" separators

### Key Details
- Uses `strings.Title` for capitalization (deprecated but functional in Go 1.18)
- n==1 and n==2 hardcoded for correct singular/plural handling
- Default case handles n>=3 with fmt.Sprintf

## Branch
`issue-195` — pushed to origin

## Commit
`3c62e39` — feat: implement Recite function for bottle-song exercise

## Status
Complete — all 7 tests pass, all acceptance criteria verified.
