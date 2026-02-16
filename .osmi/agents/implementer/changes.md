# Changes: bottle-song implementation

## File Modified
- `go/exercises/practice/bottle-song/bottle_song.go`

## What Changed
Implemented the `Recite(startBottles, takeDown int) []string` function for the bottle-song exercise.

### Added:
- `numberToWord` map (0-10) for converting integers to English words
- `titleCase(word string) string` helper to capitalize the first letter
- `bottleStr(n int) string` helper returning singular "bottle" for n=1, plural "bottles" otherwise
- `verse(n int) []string` helper generating the 4 lines of a single verse
- `Recite(startBottles, takeDown int) []string` main function that loops from startBottles downward for takeDown iterations, separating verses with empty strings

## Test Results
All 7 tests pass.
