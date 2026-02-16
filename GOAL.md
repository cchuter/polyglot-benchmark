# Goal: Implement Go Beer Song Exercise

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The stub file `go/exercises/practice/beer-song/beer_song.go` currently contains only a package declaration and needs a complete implementation of three functions: `Verse`, `Verses`, and `Song`.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle number `n`.
   - Verses 3-99: Standard format with plural "bottles", "Take one down".
   - Verse 2: Plural "bottles" in first line, singular "1 bottle" in second line.
   - Verse 1: Singular "bottle", "Take it down", "no more bottles".
   - Verse 0: "No more bottles" / "no more bottles", "Go to the store and buy some more, 99 bottles".
   - Invalid verse numbers (outside 0-99): return an error.

2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`, each separated by a blank line.
   - Validates start and stop are in range 0-99.
   - Validates start >= stop.
   - Returns error for invalid inputs.

3. **`Song() string`** — Returns the complete song (verses 99 down to 0).

4. **All tests pass**: `go test` in `go/exercises/practice/beer-song/` passes with zero failures.

5. **Package**: Must be `package beer`.

## Key Constraints

- Must match exact string output expected by tests (punctuation, capitalization, newlines).
- Each verse ends with `\n`, and verses are separated by `\n` (resulting in blank line between verses).
- Error handling is required for out-of-range inputs.
