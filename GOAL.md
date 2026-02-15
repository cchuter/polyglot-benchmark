# Goal: Implement bottle-song Exercise in Go

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates the lyrics to "Ten Green Bottles", a popular children's repetitive song. The function receives a starting bottle count and a number of verses to recite, and returns the lyrics as a slice of strings.

The stub file already exists with a `panic("Please implement the Recite function")` placeholder. The test file and test cases are already provided.

## Acceptance Criteria

1. **All 7 test cases pass** (`go test ./...` in the bottle-song directory):
   - `first generic verse` — Recite(10, 1) returns correct 4-line verse
   - `last generic verse` — Recite(3, 1) returns correct 4-line verse
   - `verse with 2 bottles` — Recite(2, 1) handles singular "bottle" in the result line
   - `verse with 1 bottle` — Recite(1, 1) handles singular "bottle" and "no" for zero
   - `first two verses` — Recite(10, 2) returns two verses separated by blank line
   - `last three verses` — Recite(3, 3) returns three verses with correct transitions
   - `all verses` — Recite(10, 10) returns the complete song

2. **Correct lyrical rules**:
   - Numbers are written as words (Ten, Nine, ..., One, no), not digits
   - First line of each verse uses title case for the number word
   - "bottle" is singular when count is 1, "bottles" is plural otherwise
   - Third line always says "one green bottle" (singular)
   - Last line uses "no green bottles" when count reaches zero
   - Verses are separated by an empty string `""` in the returned slice

3. **Code quality**:
   - Package name is `bottlesong`
   - Function signature: `func Recite(startBottles, takeDown int) []string`
   - Code builds without errors
   - No modifications to test files (`bottle_song_test.go`, `cases_test.go`)

## Key Constraints

- Only modify `bottle_song.go` — test files are read-only
- The `Title()` helper function is available in the test file for use
- Package must remain `bottlesong`
- Go module version is 1.18
- Reference implementation exists at `.meta/example.go` for verification
