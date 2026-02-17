# Goal: Implement Beer Song Exercise (Issue #319)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must provide three functions that generate the lyrics with correct grammar variations for different verse numbers.

## Required Functions

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n` (0-99). Returns error for invalid input.
2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`, separated by blank lines. Returns error for invalid input or if start < stop.
3. **`Song() string`** — Returns the entire song (verses 99 through 0).

## Acceptance Criteria

1. `Verse(n)` returns the correct verse for any n in [0, 99]
2. `Verse(n)` returns an error for n outside [0, 99]
3. Grammar variations are handled correctly:
   - n >= 3: "{n} bottles" / "Take one down" / "{n-1} bottles"
   - n == 2: "2 bottles" / "Take one down" / "1 bottle" (singular)
   - n == 1: "1 bottle" (singular) / "Take it down" (not "one") / "no more bottles"
   - n == 0: "No more bottles" / "Go to the store and buy some more, 99 bottles"
4. `Verses(start, stop)` returns verses separated by "\n" (blank line between verses)
5. `Verses()` returns error when start or stop is outside [0, 99] or start < stop
6. `Song()` returns the same result as `Verses(99, 0)`
7. All tests in `beer_song_test.go` pass
8. `go vet` reports no issues

## Key Constraints

- Package name: `beer`
- Module: `beer` (go 1.18)
- File: `beer_song.go` (edit the existing stub)
- Do not modify `beer_song_test.go`
- Standard library only (no external dependencies)
