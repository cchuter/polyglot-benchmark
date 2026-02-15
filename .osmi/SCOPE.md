# Scope: bottle-song Fix (Issue #26)

## In Scope

- Modifying `go/exercises/practice/bottle-song/bottle_song.go` to remove usage of deprecated `strings.Title`
- Replacing `strings.Title` with a manual first-letter capitalization function
- Ensuring all existing tests continue to pass
- Ensuring no linter warnings from `go vet` or `staticcheck`

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Changing the public API (`Recite` function signature)
- Modifying any other exercises in the repository

## Dependencies

- Go 1.18+ (already satisfied, running Go 1.21.6)
- No external package dependencies
- Standard library only (`fmt`, `strings`, `unicode`)
