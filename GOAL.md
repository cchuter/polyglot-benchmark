# Goal: Implement Go Beer Song Exercise

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song in Go. The stub file `go/exercises/practice/beer-song/beer_song.go` currently only contains the package declaration. Three functions must be implemented: `Verse(n)`, `Verses(start, stop)`, and `Song()`.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** returns the correct verse for a given number:
   - For n=3..99: standard verse with "N bottles" / "N-1 bottles"
   - For n=2: "2 bottles" / "1 bottle" (singular)
   - For n=1: "1 bottle" / "Take it down" / "no more bottles"
   - For n=0: "No more bottles" / "Go to the store and buy some more, 99 bottles"
   - For n<0 or n>99: returns an error

2. **`Verses(start, stop int) (string, error)`** returns a range of verses from start down to stop:
   - Each verse separated by a blank line (extra `\n` between verses)
   - Returns error for invalid start (>99 or <0), invalid stop (<0), or start < stop

3. **`Song() string`** returns the entire song (equivalent to `Verses(99, 0)`)

4. **All tests pass**: `go test` in the beer-song directory passes with no failures

5. **Code compiles cleanly**: no build errors or warnings

## Key Constraints

- Package must be named `beer`
- Must match the exact output format expected by the test file
- The `Verses` output ends with a trailing newline after the last verse
- Go 1.18 compatibility (per go.mod)
