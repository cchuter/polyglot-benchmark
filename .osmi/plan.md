# Implementation Plan: beer-song

## File to Modify

- `go/exercises/practice/beer-song/beer_song.go` — the only file to change

## Approach

Implement three exported functions in the `beer` package using a switch-based approach for the `Verse` function to handle the four special cases, and a loop-based approach for `Verses`.

### Function: `Verse(n int) (string, error)`

Uses a switch statement to handle:
1. **Invalid input** (n < 0 or n > 99): return empty string and error
2. **n == 0**: "No more bottles..." / "Go to the store..."
3. **n == 1**: "1 bottle..." (singular) / "Take it down..." / "no more bottles..."
4. **n == 2**: "2 bottles..." / "Take one down..." / "1 bottle..." (singular on next line)
5. **default (3-99)**: Standard format with `fmt.Sprintf`

### Function: `Verses(start, stop int) (string, error)`

1. Validate `start` and `stop` are in [0, 99] and `start >= stop`
2. Loop from `start` down to `stop`, calling `Verse(i)` for each
3. Join verses with an extra newline between them (each verse already ends with `\n`, so append `\n` after each)
4. Return the concatenated result

### Function: `Song() string`

Simply calls `Verses(99, 0)` and returns the result, ignoring the error (since 99,0 is always valid).

## Imports

- `fmt` — for Sprintf and Errorf
- `bytes` — for Buffer to efficiently concatenate strings

## Ordering

1. Write the complete `beer_song.go`
2. Run tests to verify
3. Run `go vet` to verify
4. Commit
