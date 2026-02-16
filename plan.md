# Implementation Plan: bottle-song Exercise

## File to Modify

- `go/exercises/practice/bottle-song/bottle_song.go` — the only file that needs changes

## Approach

The reference solution in `.meta/example.go` provides the canonical approach. The implementation follows this pattern:

### 1. Number-to-Word Map

Define a `map[int]string` mapping integers 1-10 to their English word equivalents (lowercase). This is used for generating verse text.

### 2. `verse(n int) []string` Helper

Generates a single verse for `n` bottles. Three cases:
- **n == 1**: Singular "bottle" in both lines, "no green bottles" in the result line
- **n == 2**: Plural "bottles" in the first two lines, singular "one green bottle" in the result line
- **n >= 3**: Plural "bottles" throughout, uses `Title()` for capitalizing number words at line start, lowercase for the "There'll be" line

### 3. `Recite(startBottles, takeDown int) []string` Main Function

Iterates from `startBottles` down for `takeDown` verses:
- Appends each verse's lines to the result slice
- Inserts an empty string `""` between verses (not after the last verse)

## Key Details

- The `Title` function is defined in `bottle_song_test.go` and is accessible within the same package
- Number words must be capitalized at the start of lines (e.g., "Ten green bottles")
- Number words are lowercase in "There'll be" lines (e.g., "There'll be nine green bottles")
- Special case: "no" instead of "zero" for 0 bottles

## Testing

Run `go test` in `go/exercises/practice/bottle-song/` to verify all 7 test cases pass.
