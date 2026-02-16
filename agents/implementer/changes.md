# Changes: beer-song exercise implementation

## What was done
- Implemented `beer_song.go` in `go/exercises/practice/beer-song/`
- Added three exported functions: `Verse`, `Verses`, and `Song`

## Details
- **`Verse(n int)`**: Returns a single verse for bottle number `n` (0-99), with special cases for 0, 1, and 2 bottles. Returns an error for out-of-range values.
- **`Verses(start, stop int)`**: Returns a range of verses from `start` down to `stop`, separated by blank lines. Validates both bounds and that start >= stop.
- **`Song()`**: Convenience function returning the full song (verses 99 down to 0).
- Build verified with `go build ./...`
- Committed on branch `issue-191`
