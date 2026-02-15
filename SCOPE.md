# Scope: bottle-song Exercise

## In Scope

- Implementing the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Number-to-word conversion for integers 0-10
- Singular/plural handling for "bottle"/"bottles"
- Verse generation with proper formatting
- Multi-verse output with blank line separators
- Passing all 7 existing test cases

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying metadata files (`.meta/config.json`, `.meta/tests.toml`)
- Modifying `go.mod`
- Adding external dependencies
- Handling numbers beyond 10 (tests only cover 1-10)
- Adding additional exported functions beyond `Recite`
- Error handling for invalid inputs (no test cases for this)
- Implementing exercises in other languages (python, rust, java, etc.)

## Dependencies

- No external Go packages required (stdlib only: `fmt`)
- The `Title()` function defined in `bottle_song_test.go` is available at test time
- `go.mod` already configured with `module bottlesong` and `go 1.18`
