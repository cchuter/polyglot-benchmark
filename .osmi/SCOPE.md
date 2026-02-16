# Scope: polyglot-go-beer-song (Issue #149)

## In Scope

- Implement `Verse(n int) (string, error)` in `beer_song.go`
- Implement `Verses(start, stop int) (string, error)` in `beer_song.go`
- Implement `Song() string` in `beer_song.go`
- Handle edge cases: verse 0, verse 1, verse 2, invalid inputs
- Pass all existing tests in `beer_song_test.go`

## Out of Scope

- Modifying test files
- Modifying `go.mod`
- Adding new dependencies
- Bonus refactoring challenges mentioned in the issue
- Changes to any other exercises

## Dependencies

- Go standard library only (`fmt`, `bytes` or `strings`)
- No external packages
