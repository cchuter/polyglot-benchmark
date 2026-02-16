# Scope: polyglot-go-beer-song

## In Scope

- Implement `Verse(n int) (string, error)` in `beer_song.go`
- Implement `Verses(start, stop int) (string, error)` in `beer_song.go`
- Implement `Song() string` in `beer_song.go`
- Handle all edge cases: singular/plural, verse 0, verse 1, verse 2, invalid inputs
- Pass all existing tests in `beer_song_test.go`

## Out of Scope

- Modifying test files
- Modifying `go.mod`
- Adding new dependencies
- Other exercises in the repository
- Bonus point optimizations (polymorphism, etc.) - only if time permits

## Dependencies

- Go standard library only (`fmt`, `bytes`)
- No external packages required
