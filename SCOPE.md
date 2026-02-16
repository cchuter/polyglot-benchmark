# Scope: Beer Song Exercise

## In Scope

- Implement three exported functions in `go/exercises/practice/beer-song/beer_song.go`:
  - `Verse(n int) (string, error)`
  - `Verses(start, stop int) (string, error)`
  - `Song() string`
- Handle all grammatical edge cases for verses 0, 1, 2, and 3–99
- Input validation (return errors for out-of-range verse numbers)
- Pass all existing tests in `beer_song_test.go`

## Out of Scope

- Modifying test files (`beer_song_test.go`)
- Modifying `go.mod`
- Adding new dependencies
- Performance optimization beyond what's needed to pass benchmarks
- Bonus/stretch goals mentioned in the issue (polymorphism, duplication experiments)

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Standard library only (`fmt`, `bytes` or `strings`)
- No external packages required
