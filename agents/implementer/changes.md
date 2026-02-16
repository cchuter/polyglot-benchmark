# Changes: beer-song implementation

## Modified Files

- `go/exercises/practice/beer-song/beer_song.go` — Full implementation replacing the stub

## What Was Done

Implemented three exported functions in package `beer`:

1. **`Verse(n int) (string, error)`** — Returns a single verse using a switch statement with cases for:
   - Invalid input (n < 0 or n > 99): returns error
   - n == 0: "No more bottles" / "Go to the store" verse
   - n == 1: singular "bottle", "Take it down", ends with "no more bottles"
   - n == 2: plural first line, singular "1 bottle" second line
   - n >= 3 (default): standard plural verse via fmt.Sprintf

2. **`Verses(start, stop int) (string, error)`** — Returns a range of verses from start down to stop (inclusive). Validates start/stop in 0-99 range and start >= stop. Uses bytes.Buffer; appends extra `\n` after every verse (including last) to match test expectations.

3. **`Song() string`** — Calls `Verses(99, 0)`, discards error, returns the full song.

## Test Results

All 11 tests pass (6 Verse tests, 5 Verses tests, 1 Song test + 2 benchmarks).
