# Implementation Plan: Beer Song

## File to Modify

- `go/exercises/practice/beer-song/beer_song.go` — the only file to change

## Approach

Implement three exported functions in the `beer` package:

### 1. `Verse(n int) (string, error)`

Use a switch statement to handle four cases:
- `n < 0 || n > 99`: return error
- `n == 0`: "No more bottles..." / "Go to the store..."
- `n == 1`: "1 bottle..." (singular) / "Take it down..." / "no more bottles"
- `n == 2`: "2 bottles..." / "Take one down..." / "1 bottle" (singular on second line)
- `default` (3-99): standard format with `fmt.Sprintf` using `n` and `n-1`

Each verse is two lines ending with `\n`.

### 2. `Verses(start, stop int) (string, error)`

- Validate: start and stop must be 0-99, start >= stop
- Loop from start down to stop, calling `Verse(i)` for each
- Join verses with an extra `\n` between them (each verse already ends with `\n`, so append `\n` after each)
- Use `strings.Builder` or `bytes.Buffer` for efficiency

### 3. `Song() string`

- Simply call `Verses(99, 0)` and return the result (ignoring error since inputs are known-valid)

## Imports

- `fmt` for Sprintf and Errorf
- `bytes` for Buffer (or `strings` for Builder)

## Order of Changes

1. Write the complete `beer_song.go` with all three functions
2. Run tests to verify
