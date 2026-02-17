# Goal: Implement bottle-song Exercise (Go)

## Problem Statement

Implement the `Recite` function in the `bottlesong` package that generates lyrics for the "Ten Green Bottles" children's song. The function takes a starting bottle count and the number of verses to produce, returning a slice of strings where each string is one line of the song.

## Acceptance Criteria

1. `Recite(startBottles, takeDown int) []string` is exported and returns the correct lyrics.
2. Each verse has 4 lines:
   - `"{N} green bottle(s) hanging on the wall,"`
   - `"{N} green bottle(s) hanging on the wall,"`
   - `"And if one green bottle should accidentally fall,"`
   - `"There'll be {N-1} green bottle(s) hanging on the wall."`
3. Numbers are written as capitalized English words (e.g., "Ten", "Nine") in lines 1, 2, and 4.
4. Singular "bottle" is used when the count is 1; plural "bottles" otherwise, including 0 ("no green bottles").
5. When the remaining count is 0, the text reads "no green bottles".
6. Multiple verses are separated by a single empty string element (`""`).
7. The third line always says "one green bottle" (lowercase, singular) regardless of current count.
8. All 7 test cases in `cases_test.go` must pass.
9. `go vet` reports no issues.

## Key Constraints

- The solution must be in a single file: `bottle_song.go`.
- Package name must be `bottlesong`.
- No external dependencies (go.mod specifies `go 1.18` only).
- Number words needed: "no", "one" through "ten" (numbers 0-10).
- First letter of the number word is capitalized only in lines 1, 2, and 4 (the verse lines), not in line 3.
