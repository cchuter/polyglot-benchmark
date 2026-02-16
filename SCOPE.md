# Scope: bottle-song Exercise

## In Scope

- Implement `Recite(startBottles, takeDown int) []string` in `bottle_song.go`
- Number-to-word mapping for integers 0-10
- Correct singular/plural handling for "bottle"/"bottles"
- Verse separation with empty strings for multi-verse output
- All 7 test cases passing

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Supporting numbers beyond 10
- Error handling for invalid inputs (tests don't test for this)
- Any changes outside `go/exercises/practice/bottle-song/bottle_song.go`

## Dependencies

- No external packages; only Go standard library (`fmt`)
- The `Title` function from `bottle_song_test.go` is available at test time for title-casing (the solution can use its own approach or `strings.Title`)
