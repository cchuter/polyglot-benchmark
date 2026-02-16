# Goal: Implement Beer Song Exercise (Go)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song lyrics generator in Go. The solution must handle special cases for verses 0, 1, and 2, which differ from the standard verse format.

## Acceptance Criteria

1. `Verse(n int) (string, error)` returns the correct verse for any valid number (0-99)
   - Verses 3-99: standard format with "bottles" (plural), "Take one down and pass it around"
   - Verse 2: uses "1 bottle" (singular) in the second line
   - Verse 1: uses "bottle" (singular), "Take it down", "no more bottles"
   - Verse 0: starts with "No more bottles", second line is "Go to the store and buy some more, 99 bottles of beer on the wall."
   - Returns error for values outside 0-99

2. `Verses(start, stop int) (string, error)` returns multiple verses from `start` down to `stop`
   - Verses separated by blank lines
   - Returns error if start or stop outside 0-99, or if start < stop

3. `Song() string` returns the entire song (verses 99 down to 0)

4. All tests in `beer_song_test.go` pass
5. `go vet` reports no issues

## Key Constraints

- Package name must be `beer`
- Module name is `beer` with Go 1.18
- Must not modify test file `beer_song_test.go`
- Solution goes in `beer_song.go`
