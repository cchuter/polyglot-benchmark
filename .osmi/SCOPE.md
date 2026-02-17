# Scope: bottle-song Exercise

## In Scope

- Implement the `Recite(startBottles, takeDown int) []string` function in `bottle_song.go`.
- Handle number-to-word conversion for integers 0-10.
- Handle singular/plural for "bottle"/"bottles".
- Handle capitalization of number words at the start of verse lines.
- Handle "no green bottles" for the zero case.
- Separate verses with empty string elements.
- Pass all existing test cases in `bottle_song_test.go` and `cases_test.go`.

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`).
- Modifying `go.mod`.
- Handling numbers outside 0-10.
- Error handling for invalid inputs (not tested).
- Any changes to other exercises or languages.

## Dependencies

- No external dependencies.
- Standard library only (`strings`, `fmt` if needed).
- Must compile with Go 1.18+.
