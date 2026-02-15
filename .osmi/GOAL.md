# Goal: Implement bottle-song Exercise in Go (Issue #38)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates lyrics for the "Ten Green Bottles" children's song. The function takes two parameters:
- `startBottles`: the number of bottles to start from (1-10)
- `takeDown`: how many verses to recite

## Acceptance Criteria

1. `Recite(startBottles, takeDown int) []string` returns the correct lyrics as a slice of strings
2. Each verse consists of 4 lines:
   - Line 1: `"{N} green bottle(s) hanging on the wall,"`
   - Line 2: same as line 1
   - Line 3: `"And if one green bottle should accidentally fall,"`
   - Line 4: `"There'll be {N-1} green bottle(s) hanging on the wall."`
3. Numbers are spelled out as words with the first letter capitalized in lines 1-2 (e.g., "Ten", "Nine")
4. Numbers are lowercase in line 4 (e.g., "nine", "eight")
5. Singular "bottle" is used when count is 1; plural "bottles" otherwise
6. When count reaches 0, use "no green bottles"
7. Multiple verses are separated by an empty string `""` in the slice
8. All 7 test cases in `cases_test.go` pass: single verses (10, 3, 2, 1), first two verses, last three verses, and all verses
9. `go vet` produces no warnings
10. No external dependencies added

## Key Constraints

- Package must be `bottlesong`
- Must use Go 1.18+ (as specified in go.mod)
- The `Title` function is provided in the test file for title-casing strings
- No external dependencies allowed
