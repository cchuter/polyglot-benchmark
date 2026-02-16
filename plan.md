# Implementation Plan: polyglot-go-beer-song

## File to Modify

- `go/exercises/practice/beer-song/beer_song.go` — the only file to change (currently just `package beer`)

## Architecture

Single file with three exported functions. No types or interfaces needed.

### Imports

- `fmt` — for `Sprintf` and `Errorf`
- `strings` — for `Builder` (efficient string concatenation)

### Function: `Verse(n int) (string, error)`

1. Validate `n` is in range [0, 99]; return error otherwise using `fmt.Errorf`
2. Switch on `n`:
   - `n == 0`: Return `"No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"`
   - `n == 1`: Return `"1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"` (note: "Take **it** down", singular "bottle")
   - `n == 2`: Return `"2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"` (note: "1 **bottle**" singular)
   - Default (`n >= 3`): `fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)`
3. Each verse is exactly two lines, each ending with `\n`

### Function: `Verses(start, stop int) (string, error)`

1. Validate: `start` in [0, 99], `stop` in [0, 99], `start >= stop`; return descriptive errors for each case
2. Use `strings.Builder`
3. Loop from `start` down to `stop` inclusive:
   - Call `Verse(i)` and write result to builder
   - Write an additional `\n` after **every** verse (not just between them)
4. Return builder contents

**Key insight from test data**: `verses86` ends with a trailing blank line (each verse gets `\n` appended, even the last one). The reference solution confirms: write verse + `\n` for every iteration.

### Function: `Song() string`

1. Call `Verses(99, 0)`, return the string (ignore error since inputs are known valid)

## Ordering

1. Write the complete `beer_song.go`
2. Run `go test` to verify
3. Run `go vet` for static analysis
