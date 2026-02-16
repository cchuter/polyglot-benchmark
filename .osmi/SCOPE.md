# Scope: polyglot-go-beer-song (Issue #93)

## In Scope

- Implement `Verse(n int) (string, error)` in `go/exercises/practice/beer-song/beer_song.go`
- Implement `Verses(start, stop int) (string, error)` in the same file
- Implement `Song() string` in the same file
- Pass all tests in `beer_song_test.go`
- Pass `go vet`

## Out of Scope

- Modifying test files
- Modifying `go.mod`
- Adding new dependencies
- Implementing other exercises
- Bonus refactoring (polymorphism, duplication removal beyond what's needed)

## Dependencies

- Go 1.18+ (specified in go.mod)
- Standard library only (`fmt`, `bytes` or `strings`)
- No external packages
