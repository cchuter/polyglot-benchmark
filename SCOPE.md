# Scope: Beer Song Exercise

## In Scope

- Implementing `Verse(n int) (string, error)` in `beer_song.go`
- Implementing `Verses(start, stop int) (string, error)` in `beer_song.go`
- Implementing `Song() string` in `beer_song.go`
- Ensuring all tests in `beer_song_test.go` pass
- Proper error handling for invalid inputs

## Out of Scope

- Modifying test files (`beer_song_test.go`)
- Modifying `go.mod`
- Modifying any files outside the `go/exercises/practice/beer-song/` directory
- Adding new dependencies
- Bonus point optimizations (polymorphism, duplication removal beyond what's needed)
- Creating additional files

## Dependencies

- Go toolchain (1.18+)
- Standard library only (`fmt`, `bytes` or `strings`)
- No external packages required
