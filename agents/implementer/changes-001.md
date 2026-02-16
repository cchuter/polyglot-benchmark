# Changes 001: Implement Beer Song Solution

## What changed
- Replaced the stub `beer_song.go` (which only contained `package beer`) with the full implementation.

## Functions implemented
- **`Verse(n int) (string, error)`**: Returns a single verse for bottle number `n` (0-99). Handles special cases for 0, 1, and 2 bottles (singular/plural and unique lyrics). Returns an error for out-of-range values.
- **`Verses(start, stop int) (string, error)`**: Returns a range of verses from `start` down to `stop`, separated by blank lines. Validates both bounds and that start >= stop.
- **`Song() string`**: Returns the complete song (verses 99 down to 0).

## Files modified
- `go/exercises/practice/beer-song/beer_song.go`
