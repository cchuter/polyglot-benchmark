# Goal: polyglot-go-beer-song (Issue #93)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song in Go. The solution must generate correct lyrics for individual verses, ranges of verses, and the entire song. Key lyrical variations exist at verses 2, 1, and 0 that must be handled correctly.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n` (0-99). Returns an error for invalid values (n < 0 or n > 99).
   - Verses 3-99: standard format with "bottles" (plural), "Take one down and pass it around"
   - Verse 2: uses "1 bottle" (singular) in the second line
   - Verse 1: uses "bottle" (singular), "Take it down" (not "Take one down"), "no more bottles"
   - Verse 0: "No more bottles…", "Go to the store and buy some more, 99 bottles of beer on the wall."

2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`, separated by blank lines. Returns errors for out-of-range values or when start < stop.

3. **`Song() string`** — Returns the entire song (Verses 99 down to 0).

4. All existing tests in `beer_song_test.go` pass.
5. `go vet` reports no issues.

## Key Constraints

- Package name must be `beer`
- Module defined as `beer` with `go 1.18`
- Must not modify test file `beer_song_test.go`
- Solution goes in `beer_song.go`
