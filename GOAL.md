# Goal: Implement bottle-song Exercise (Go)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates the lyrics to the "Ten Green Bottles" children's song. The function takes a starting bottle count and the number of verses to recite, returning the lyrics as a slice of strings.

Key details:
- Numbers must be spelled out as words (e.g., "Ten", "Nine") and capitalized at the start of lines
- Singular/plural: "bottle" (1) vs "bottles" (0, 2+)
- The last verse ends with "no green bottles"
- Verses are separated by an empty string element in the slice
- The third line always says "one green bottle" (singular)

## Acceptance Criteria

1. `Recite(startBottles, takeDown)` returns `[]string` with correct lyrics
2. Single verse for any bottle count (10 down to 1) produces correct 4-line output
3. Multiple consecutive verses are separated by an empty string `""`
4. Number words are capitalized at the start of lines 1 and 2 (e.g., "Ten green bottles")
5. Number words are lowercase in lines 3-4 (e.g., "There'll be nine green bottles")
6. Singular "bottle" used when count is 1; plural "bottles" for 0 and 2+
7. "no green bottles" used when the remaining count reaches 0
8. All 7 test cases in `cases_test.go` pass

## Key Constraints

- Package must be `bottlesong`
- Function signature: `func Recite(startBottles, takeDown int) []string`
- Must use the `Title` function from `bottle_song_test.go` for capitalization (or `strings.Title`)
- No external dependencies (go.mod has no requires)
- Go 1.18 compatibility
