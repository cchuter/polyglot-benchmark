# Implementation Plan: Beer Song

## Overview

Replace the stub implementations in `go/exercises/practice/beer-song/beer_song.go` with working code. The reference solution in `.meta/example.go` provides a known-correct approach.

## File to Modify

- `go/exercises/practice/beer-song/beer_song.go` (single file change)

## Architectural Decisions

1. **Use `fmt.Sprintf` for standard verses** - Format strings with bottle count interpolation
2. **Use `switch` statement for special cases** - Verse 0, 1, and 2 have unique wording
3. **Use `bytes.Buffer` for `Verses()`** - Efficient string concatenation when joining multiple verses
4. **Imports: `bytes` and `fmt`** - Only standard library packages needed

## Implementation Details

### `Verse(n int) (string, error)`

Handle 5 cases via switch:
- `n < 0 || n > 99`: Return `("", error)`
- `n == 0`: Hard-coded "No more bottles..." verse
- `n == 1`: Hard-coded "1 bottle..." verse with "Take it down"
- `n == 2`: Hard-coded "2 bottles..." verse with "1 bottle" (singular) on next line
- `default` (3-99): `fmt.Sprintf` with `n` and `n-1`

### `Verses(start, stop int) (string, error)`

1. Validate: `start` in [0,99], `stop` in [0,99], `start >= stop`
2. Loop from `start` down to `stop`, calling `Verse(i)` for each
3. Append each verse + `"\n"` separator to a `bytes.Buffer`
4. Return the buffer contents

### `Song() string`

Simply call `Verses(99, 0)` and return the result, ignoring the error (it's always valid).

## Ordering

1. Write the implementation in `beer_song.go`
2. Run `go test` in the exercise directory to verify
3. Commit

## Risks

- Whitespace sensitivity: test expectations are exact string matches. Must ensure each verse ends with `\n` and verses are separated by `\n` (blank line).
