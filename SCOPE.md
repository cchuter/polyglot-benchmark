# Scope: polyglot-go-beer-song (Issue #105)

## In Scope

- Implement `Verse()`, `Verses()`, and `Song()` in `go/exercises/practice/beer-song/beer_song.go`
- Handle all special-case verses (0, 1, 2) and the general case (3-99)
- Input validation with proper error returns
- Pass all existing tests in `beer_song_test.go`

## Out of Scope

- Modifying the test file (`beer_song_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Implementing bonus points (polymorphism, duplication optimization beyond clean code)
- Changes to any other exercises or non-Go files

## Dependencies

- Go toolchain (1.18+)
- Standard library only (`fmt`, `bytes` or `strings`)
- No external packages
