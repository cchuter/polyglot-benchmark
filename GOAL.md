# Goal: Implement Beer Song Exercise in Go

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song in Go. The solution must generate correct lyrics with proper grammar handling for singular/plural forms and special verse cases.

The stub file `go/exercises/practice/beer-song/beer_song.go` currently contains only `package beer` and needs three exported functions implemented.

## Required Functions

1. **`Verse(n int) (string, error)`** — Returns lyrics for a single verse (0-99), error for invalid input
2. **`Verses(start, stop int) (string, error)`** — Returns consecutive verses from `start` down to `stop`, error for invalid input
3. **`Song() string`** — Returns the entire song (all verses from 99 to 0)

## Acceptance Criteria

1. `Verse(n)` returns correct lyrics for any valid verse number (0-99)
2. `Verse(n)` returns an error for invalid verse numbers (< 0 or > 99)
3. Verse 0 uses "No more bottles" / "no more bottles" and "Go to the store and buy some more, 99 bottles"
4. Verse 1 uses singular "bottle" and "Take it down" (not "Take one down")
5. Verse 2 uses plural "bottles" but the next line uses singular "1 bottle"
6. Verses 3-99 use plural "bottles" throughout
7. `Verses(start, stop)` returns multiple verses separated by blank lines (extra `\n` between verses)
8. `Verses(start, stop)` returns error when start or stop is outside 0-99, or start < stop
9. `Song()` returns the same output as `Verses(99, 0)`
10. All existing tests in `beer_song_test.go` pass
11. Benchmarks run without error

## Key Constraints

- Package name must be `beer`
- Module is `beer` with Go 1.18
- Must not modify the test file
- Each verse ends with `\n`, and verses in multi-verse output are separated by an additional `\n`
