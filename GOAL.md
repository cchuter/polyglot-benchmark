# Goal: Implement bottle-song Exercise (Go)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates the lyrics to the children's song "Ten Green Bottles". The function takes two parameters: `startBottles` (which bottle number to start from) and `takeDown` (how many verses to produce), and returns the lyrics as a slice of strings (one string per line).

## Acceptance Criteria

1. `Recite(startBottles, takeDown int) []string` is exported from package `bottlesong`
2. Numbers are spelled out as English words (one, two, ..., ten), title-cased at the start of lines
3. "bottle" is singular when the count is exactly 1; "bottles" is plural otherwise (including 0/"no")
4. The last verse (1 bottle) ends with "There'll be no green bottles hanging on the wall."
5. Multiple verses are separated by an empty string `""` in the returned slice
6. All 7 test cases in `cases_test.go` pass:
   - first generic verse (10, 1)
   - last generic verse (3, 1)
   - verse with 2 bottles (2, 1)
   - verse with 1 bottle (1, 1)
   - first two verses (10, 2)
   - last three verses (3, 3)
   - all verses (10, 10)

## Key Constraints

- Package name must be `bottlesong`
- The `Title` helper function is provided in the test file — it can be used but does not need to be reimplemented
- No external dependencies; only the Go standard library is allowed
- Must target Go 1.18+ (per go.mod)
