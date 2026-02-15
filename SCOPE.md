# Scope: bottle-song Exercise (Issue #38)

## In Scope

- Implement `Recite(startBottles, takeDown int) []string` in `bottle_song.go`
- Handle number-to-word conversion for integers 0-10
- Handle singular/plural for "bottle"/"bottles"
- Handle verse separation with empty strings
- Pass all 7 test cases in `cases_test.go`

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `.meta/` or `.docs/` directories
- Adding external dependencies
- Input validation beyond what tests require
- Numbers outside 0-10 range

## Dependencies

- Go 1.18+ toolchain
- No external packages; standard library only (`fmt`)
- `Title` function provided in test file for title-casing
