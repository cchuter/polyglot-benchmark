# Scope: Bottle Song Exercise

## In Scope

- Implementing the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Handling all verse variations: generic verses (3-10), verse with 2 bottles, verse with 1 bottle
- Number-to-word mapping for integers 0-10
- Proper singular/plural handling for "bottle"/"bottles"
- Verse separation with empty string elements for multi-verse output

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `.meta/` or `.docs/` directories
- Supporting numbers outside 0-10
- Adding error handling for invalid inputs (tests don't test for this)
- External dependencies or module changes
- Changes to any other exercises

## Dependencies

- Go 1.18+ (as specified in go.mod)
- `Title` function defined in `bottle_song_test.go` (available at package scope during testing)
- `fmt` package from standard library (for string formatting)
