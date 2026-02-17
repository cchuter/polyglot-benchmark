# Goal: Implement Bottle Song (Ten Green Bottles)

## Problem Statement

Implement the `Recite` function in Go for the "Ten Green Bottles" children's song exercise. The function must generate the lyrics for a specified range of verses, starting from a given number of bottles and producing a given number of consecutive verses counting down.

## Acceptance Criteria

1. **Function signature**: `Recite(startBottles, takeDown int) []string` in package `bottlesong`
2. **Single verse output**: Each verse consists of 4 lines returned as 4 string elements
3. **Multiple verse separation**: When `takeDown > 1`, verses are separated by an empty string element (`""`)
4. **Number words**: Numbers must be spelled out as English words with title case in the first two lines (e.g., "Ten", "Nine") and lowercase in the last line (e.g., "nine", "eight")
5. **Singular/plural**: "bottle" (singular) when count is 1; "bottles" (plural) for all other counts including 0
6. **Zero bottles**: When the count reaches 0, use "no green bottles" (not "zero")
7. **All 7 test cases pass**: first generic verse, last generic verse, verse with 2 bottles, verse with 1 bottle, first two verses, last three verses, all verses

## Key Constraints

- Solution must be in `go/exercises/practice/bottle-song/bottle_song.go`
- Package name must be `bottlesong`
- The `Title` function is provided in the test file and available at package scope
- No external dependencies (module is `go 1.18`, no `require` directives)
- The reference solution exists at `.meta/example.go` for validation
