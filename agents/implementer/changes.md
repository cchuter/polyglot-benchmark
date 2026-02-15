# Changes: beer-song implementation

## Modified Files

- `go/exercises/practice/beer-song/beer_song.go` — Replaced stub implementations with working code

## Summary

Implemented three functions for the beer-song exercise:

- **`Verse(n int) (string, error)`**: Returns a single verse. Uses a switch statement for special cases (0, 1, 2) and `fmt.Sprintf` for verses 3–99. Returns an error for invalid verse numbers.
- **`Verses(start, stop int) (string, error)`**: Returns multiple verses from `start` down to `stop`. Validates inputs and uses `bytes.Buffer` for efficient concatenation.
- **`Song() string`**: Returns the full song by calling `Verses(99, 0)`.

## Test Results

All tests pass: `TestBottlesVerse` (6 subtests), `TestSeveralVerses` (5 subtests), `TestEntireSong`.

## Commit

`fb4a5f3` — feat: implement beer-song exercise for Go
