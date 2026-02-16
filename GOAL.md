# Goal: Implement bottle-song Exercise (Issue #65)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates the lyrics to the children's song "Ten Green Bottles". The function takes two parameters: `startBottles` (the number to start counting from) and `takeDown` (how many verses to generate).

## Acceptance Criteria

1. `Recite(startBottles, takeDown int) []string` is implemented in `bottle_song.go`
2. All 7 test cases in `cases_test.go` pass:
   - Single verse from 10 bottles (first generic verse)
   - Single verse from 3 bottles (last generic verse)
   - Single verse with 2 bottles (singular "bottle" in result line)
   - Single verse with 1 bottle (singular "bottle", "no green bottles" result)
   - Two consecutive verses (separated by empty string)
   - Three consecutive verses
   - All 10 verses
3. Number words are capitalized at the start of lines ("Ten", "Nine", etc.)
4. Singular "bottle" is used when count is 1; plural "bottles" otherwise
5. "no green bottles" is used when count reaches 0
6. Verses are separated by an empty string (`""`) in the returned slice
7. `go test` passes with exit code 0

## Key Constraints

- Package name must be `bottlesong`
- Must use Go 1.18 compatible syntax (per go.mod)
- The solution file is `bottle_song.go`
- The `Title` helper function is available in `bottle_song_test.go` for use
