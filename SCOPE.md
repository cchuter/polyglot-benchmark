# Scope: bottle-song Exercise

## In Scope

- Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Number-to-word mapping for integers 0–10
- Singular/plural handling for "bottle"/"bottles"
- Verse generation with correct punctuation and capitalization
- Multi-verse output with empty-string separators
- All existing tests must pass (`go test` in the exercise directory)

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Supporting numbers outside the 0–10 range
- Adding external dependencies
- Any other exercises or languages in the repository

## Dependencies

- Go toolchain (1.18+)
- No external packages — standard library only
- The `Title` function is defined in `bottle_song_test.go` and available at test time
