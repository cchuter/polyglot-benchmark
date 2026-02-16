# Goal: polyglot-go-beer-song (Issue #105)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must produce the correct lyrics including special-case verses for 2, 1, and 0 bottles.

## Functions to Implement

In `go/exercises/practice/beer-song/beer_song.go` (package `beer`):

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n` (0-99). Returns an error for out-of-range values.
2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop` (inclusive), separated by blank lines. Returns an error if parameters are out of range or `start < stop`.
3. **`Song() string`** — Returns the entire song (verses 99 down to 0).

## Acceptance Criteria

1. `Verse(n)` returns the correct lyric for any valid verse (0-99)
2. `Verse(n)` returns an error for `n < 0` or `n > 99`
3. Special cases handled correctly:
   - Verse 2: "...1 bottle of beer on the wall." (singular)
   - Verse 1: "1 bottle..." / "Take it down..." / "no more bottles..."
   - Verse 0: "No more bottles..." / "Go to the store..."
4. `Verses(start, stop)` returns multiple verses separated by `\n\n`
5. `Verses` returns error for invalid ranges (out of bounds, start < stop)
6. `Song()` returns `Verses(99, 0)`
7. All tests in `beer_song_test.go` pass
8. Code compiles without errors

## Key Constraints

- Package name must be `beer`
- Module is `beer` with `go 1.18`
- Must not modify the test file
- Verse output must include a trailing newline
- Multi-verse output: each verse followed by `\n`, with `\n` between verses (resulting in blank line separators)
