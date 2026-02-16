# Scope: Beer Song Exercise

## In Scope

- Implement `Verse(n int) (string, error)` in `go/exercises/practice/beer-song/beer_song.go`
- Implement `Verses(start, stop int) (string, error)` in the same file
- Implement `Song() string` in the same file
- Handle all edge cases: singular "bottle", verse 0, verse 1, error conditions
- Pass all tests in `beer_song_test.go`
- Pass `go vet`

## Out of Scope

- Modifying test files (`beer_song_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` or `.docs/` files
- Adding additional files or packages
- Bonus refactoring beyond what's needed to pass tests

## Dependencies

- Go standard library only (`fmt`, `bytes` or `strings`)
- No external packages required
