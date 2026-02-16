# Goal: Implement Beer Song Exercise (Issue #82)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must produce the correct lyrics for the entire song, handling singular/plural forms and special verses (0, 1, 2).

The stub file at `go/exercises/practice/beer-song/beer_song.go` currently contains only `package beer` and needs a full implementation.

## Functions to Implement

### `Verse(n int) (string, error)`
Returns a single verse for bottle count `n` (0-99). Returns an error for invalid values.

Special cases:
- **n >= 3**: `"{n} bottles of beer on the wall, {n} bottles of beer.\nTake one down and pass it around, {n-1} bottles of beer on the wall.\n"`
- **n == 2**: `"2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"` (singular "bottle" for 1)
- **n == 1**: `"1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"` (singular "bottle", "Take it down", "no more")
- **n == 0**: `"No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"`
- **n < 0 or n > 99**: Return error

### `Verses(start, stop int) (string, error)`
Returns verses from `start` down to `stop`, separated by blank lines. Validates inputs.

Error conditions:
- start < 0 or start > 99
- stop < 0 or stop > 99
- start < stop

### `Song() string`
Returns the entire song (verses 99 down to 0). Equivalent to `Verses(99, 0)`.

## Acceptance Criteria

1. All tests in `beer_song_test.go` pass (`go test` exits with code 0)
2. `Verse(n)` returns correct lyrics for n=0,1,2,3..99 with proper singular/plural handling
3. `Verse(n)` returns an error for n < 0 or n > 99
4. `Verses(start, stop)` returns correctly formatted multi-verse output with blank line separators
5. `Verses(start, stop)` returns errors for invalid inputs
6. `Song()` returns the complete song matching `Verses(99, 0)`
7. Code compiles without errors (`go build`)
8. Benchmarks run without errors

## Key Constraints

- Package must be named `beer`
- Module is `beer` with Go 1.18
- Must match exact string formatting expected by tests (whitespace, newlines, capitalization)
- No external dependencies allowed
