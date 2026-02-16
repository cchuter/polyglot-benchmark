# Context: bottle-song Exercise (Issue #109)

## Status: Complete

## What Was Done
Implemented `Recite(startBottles, takeDown int) []string` in `go/exercises/practice/bottle-song/bottle_song.go` for the "Ten Green Bottles" song exercise.

## Key Files
- `go/exercises/practice/bottle-song/bottle_song.go` — solution (modified)
- `go/exercises/practice/bottle-song/bottle_song_test.go` — tests (unchanged)
- `go/exercises/practice/bottle-song/cases_test.go` — test cases (unchanged)

## Implementation Approach
- `numberToWord` map: 0→"no", 1→"one", ..., 10→"ten"
- `titleCase()`: capitalizes first letter for line starts
- `bottleStr(n)`: returns "bottle" for n==1, "bottles" otherwise
- `verse(n)`: generates 4-line verse using Sprintf
- `Recite()`: loops from startBottles, assembles verses with empty-string separators

## Test Results
All 7 test cases pass (first generic verse, last generic verse, 2 bottles, 1 bottle, first two verses, last three verses, all verses).

## Branch
`issue-109` pushed to origin. Ready for PR.
