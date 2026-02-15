# Scope: bottle-song Go Exercise

## In Scope

- Implementing the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`
- Number-to-word mapping for integers 0 through 10
- Handling singular/plural "bottle"/"bottles" correctly
- Capitalizing the number word at the start of each verse's first two lines
- Separating multiple verses with an empty string element
- Passing all 7 test cases defined in `cases_test.go`

## Out of Scope

- Modifying test files (`bottle_song_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Handling numbers outside 0-10
- Adding any external dependencies
- Modifying `.meta/` or `.docs/` files
- Any other exercises in the repository

## Dependencies

- Go toolchain (1.18+)
- No external packages — standard library only
- The test file provides a `Title` helper function (reimplements `strings.Title`) that is available within the package for use
