# Context: Beer Song Implementation

## Key Facts
- Branch: `issue-319`
- Single file modified: `go/exercises/practice/beer-song/beer_song.go`
- Package: `beer`, Module: `beer`, Go 1.18
- All 11 tests pass, go vet clean

## Functions Implemented
1. `Verse(n int) (string, error)` — switch on n for cases 0, 1, 2, default
2. `Verses(start, stop int) (string, error)` — validates inputs, loops with bytes.Buffer
3. `Song() string` — delegates to `Verses(99, 0)`

## Grammar Rules
- n >= 3: "{n} bottles" / "Take one down" / "{n-1} bottles"
- n == 2: "2 bottles" / "Take one down" / "1 bottle"
- n == 1: "1 bottle" / "Take it down" / "no more bottles"
- n == 0: "No more bottles" / "Go to the store" / "99 bottles"

## Verse Separator
Each verse ends with `\n`, and `Verses()` appends an additional `\n` after each verse, creating blank lines between verses and a trailing blank line.
