# Goal: Implement Beer Song Exercise in Go

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must produce correct lyrics for individual verses, ranges of verses, and the complete song. Verses are not all identical — special handling is required for verses 2, 1, and 0.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** — Returns a single verse for the given bottle number.
   - For n=3..99: `"{n} bottles of beer on the wall, {n} bottles of beer.\nTake one down and pass it around, {n-1} bottles of beer on the wall.\n"`
   - For n=2: Uses "1 bottle" (singular) in the second line.
   - For n=1: Uses "1 bottle" (singular), "Take it down" (not "Take one down"), and "no more bottles" in the second line.
   - For n=0: "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
   - For n<0 or n>99: Returns an error.

2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`, separated by blank lines.
   - Each verse is followed by a newline (producing a blank line between verses).
   - Returns an error if start or stop are out of range (0-99) or if start < stop.

3. **`Song() string`** — Returns the entire song (`Verses(99, 0)`).

4. All tests in `beer_song_test.go` pass.
5. `go vet` reports no issues.

## Key Constraints

- Package name must be `beer`.
- File must be `beer_song.go`.
- Must use Go 1.18+ (as specified in go.mod).
- Do not modify test files.
