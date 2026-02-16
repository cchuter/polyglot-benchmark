# Scope: bottle-song Exercise

## In Scope

- Implement `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Handle number-to-word mapping for 0-10
- Handle singular/plural "bottle"/"bottles"
- Handle capitalization of number words at line starts
- Handle verse separation with empty strings
- Pass all 7 test cases in `cases_test.go`

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` directories
- Supporting numbers beyond 10
- Error handling for invalid inputs (not tested)
- Any other exercises or languages

## Dependencies

- Go standard library only (`fmt` package)
- The `Title` helper function is provided in `bottle_song_test.go` for test use
- No external packages
