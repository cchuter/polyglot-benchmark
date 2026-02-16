# Goal: polyglot-go-beer-song

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" exercise in Go. The solution must generate the lyrics to the classic song, handling special cases for singular/plural bottle forms and the final verse.

The stub file `go/exercises/practice/beer-song/beer_song.go` currently contains only a package declaration and must be completed with three exported functions.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** returns a single verse:
   - For n=3..99: `"{n} bottles of beer on the wall, {n} bottles of beer.\nTake one down and pass it around, {n-1} bottles of beer on the wall.\n"`
   - For n=2: Uses "1 bottle" (singular) on the second line
   - For n=1: Uses "bottle" (singular) and "Take it down" instead of "Take one down", ends with "no more bottles"
   - For n=0: "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
   - For n<0 or n>99: Returns an error

2. **`Verses(start, stop int) (string, error)`** returns a range of verses:
   - Iterates from start down to stop (inclusive), separated by blank lines
   - Returns error if start or stop are outside 0-99, or if start < stop

3. **`Song() string`** returns the entire song (equivalent to `Verses(99, 0)`)

4. All tests in `beer_song_test.go` pass
5. Benchmarks run without error

## Key Constraints

- Package name must be `beer`
- Must use Go 1.18+ compatible syntax
- Solution file: `go/exercises/practice/beer-song/beer_song.go`
- Must not modify the test file
