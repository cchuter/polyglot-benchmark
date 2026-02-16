# Goal: Implement Beer Song Exercise (Go)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song in Go. The solution must produce lyrics with correct grammar variations depending on the bottle count:

- Verses 99–3: standard plural ("N bottles", "Take one down")
- Verse 2: transitions to singular ("1 bottle of beer on the wall")
- Verse 1: uses singular and different phrasing ("1 bottle", "Take it down", "no more bottles")
- Verse 0: capitalized "No more", different action ("Go to the store and buy some more, 99 bottles")

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n` (0–99). Returns an error for out-of-range values.
2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`, separated by blank lines. Returns errors for invalid ranges (out of bounds, start < stop).
3. **`Song() string`** — Returns the entire song (verses 99 down to 0).
4. All tests in `beer_song_test.go` pass.
5. Code compiles with no errors or warnings.

## Key Constraints

- Package name must be `beer` (matching the existing stub and test file).
- Solution goes in `beer_song.go` only — do not modify `beer_song_test.go` or `go.mod`.
- Must handle all grammatical edge cases (singular/plural "bottle(s)", "Take it down" vs "Take one down", "No more" vs "no more").
