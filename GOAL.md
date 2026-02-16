# Goal: Implement bottle-song Exercise (Go)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates lyrics to the children's song "Ten Green Bottles." The function takes a starting bottle count and how many verses to take down, returning the lyrics as a slice of strings (one element per line).

## Acceptance Criteria

1. `Recite(startBottles, takeDown int) []string` must be exported from package `bottlesong`
2. Numbers are spelled out as English words with title case at line beginnings (e.g., "Ten", "Nine")
3. Singular "bottle" is used when the count is 1; plural "bottles" for all other counts (including 0)
4. When the remaining count is 0, use "no green bottles"
5. The third line of every verse is always: `And if one green bottle should accidentally fall,`
6. Multiple verses are separated by an empty string `""` in the output slice
7. All 7 test cases in `cases_test.go` must pass: single verses (10, 3, 2, 1), first two verses, last three verses, and all verses

## Key Constraints

- Package name: `bottlesong`
- Go module: `bottlesong` with `go 1.18`
- No external dependencies allowed
- The `Title` helper function is available in the test file for title-casing strings
