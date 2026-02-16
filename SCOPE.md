# Scope: bottle-song Exercise (Issue #65)

## In Scope

- Implementing the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Number-to-word mapping for 0-10
- Verse generation logic with proper singular/plural handling
- Multi-verse output with blank line separators

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` files
- Any exercises other than bottle-song
- Error handling for invalid inputs (not tested)

## Dependencies

- No external Go dependencies (standard library only, `fmt` package)
- The `Title` function in the test file can be used for capitalizing number words, but since the test file already provides it, the solution can also implement its own approach
