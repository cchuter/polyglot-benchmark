# Goal: Implement Beer Song Exercise in Go

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must provide three functions (`Verse`, `Verses`, `Song`) that generate the lyrics with correct grammar handling for special cases (singular/plural "bottle(s)", verse 0, verse 1, verse 2).

## Acceptance Criteria

1. `Verse(n int) (string, error)` returns a single verse for bottle number `n` (0-99), or an error for invalid input.
2. `Verses(start, stop int) (string, error)` returns verses from `start` down to `stop` (inclusive), separated by blank lines, or an error for invalid input.
3. `Song() string` returns the entire song (verses 99 down to 0).
4. Special verse handling:
   - Verse 0: "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
   - Verse 1: "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
   - Verse 2: "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n" (singular "bottle" for the next verse)
   - Verses 3-99: Standard plural format with "Take one down and pass it around"
5. Error cases: verse number outside 0-99, start < stop, start or stop outside 0-99.
6. All tests in `beer_song_test.go` pass.
7. Code passes `go vet`.

## Key Constraints

- Package name must be `beer`
- File must be `beer_song.go`
- Module is `beer` with Go 1.18
- Must export: `Verse`, `Verses`, `Song`
