# Goal: Implement Go Beer Song Exercise

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The exercise requires implementing three functions in `go/exercises/practice/beer-song/beer_song.go` that generate the lyrics with correct grammar for singular/plural forms and special verses.

The stub file already exists with the correct function signatures. The test file already exists and defines the expected behavior.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** - Returns a single verse for bottle count `n`:
   - For `n >= 3`: Standard verse with "N bottles" / "N-1 bottles"
   - For `n == 2`: "2 bottles" / "1 bottle" (singular next)
   - For `n == 1`: "1 bottle" / "Take it down" / "no more bottles"
   - For `n == 0`: "No more bottles" / "Go to the store" / "99 bottles"
   - For `n < 0` or `n > 99`: Return an error
   - Each verse ends with `\n` (newline)

2. **`Verses(start, stop int) (string, error)`** - Returns verses from `start` down to `stop`:
   - Verses are separated by a blank line (`\n\n` between verses, so `\n` at end of verse + `\n` separator)
   - Returns error if `start > 99`, `stop < 0`, or `start < stop`
   - Trailing `\n` after the last verse

3. **`Song() string`** - Returns the entire song (equivalent to `Verses(99, 0)`)

4. All existing tests in `beer_song_test.go` pass
5. Code compiles without errors via `go build`

## Key Constraints

- Package name must be `beer`
- Must match the exact function signatures already defined in the stub
- Must use Go module `beer` with `go 1.18`
- Output must match the test expectations exactly (whitespace, capitalization, punctuation)
