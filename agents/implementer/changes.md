# Changes: beer-song implementation

## File Modified
- `go/exercises/practice/beer-song/beer_song.go`

## What Was Done
Implemented three exported functions in the `beer` package:

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n` (0–99). Uses a switch to handle special cases for verses 0, 1, and 2 (singular/plural and unique wording). Returns an error for out-of-range values.

2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop` (inclusive), each followed by an extra newline. Validates both bounds and that start >= stop.

3. **`Song() string`** — Returns the full song by calling `Verses(99, 0)`.

## Test Results
All 12 tests pass. `go vet` reports no issues.
