# Scope: Beer Song Implementation

## In Scope

- Implement `Verse(n int) (string, error)` function
- Implement `Verses(start, stop int) (string, error)` function
- Implement `Song() string` function
- Input validation with descriptive error messages
- Handle all edge cases (0, 1, 2 bottles)
- All existing tests in `beer_song_test.go` must pass

## Out of Scope

- Modifying test files (`beer_song_test.go`)
- Modifying metadata files (`.meta/`, `.docs/`)
- Modifying `go.mod`
- Adding new test cases
- Performance optimization beyond passing benchmarks
- Bonus point challenges (polymorphism, duplication removal experiments)

## Dependencies

- Go 1.18+ (as specified in go.mod)
- Standard library only (`fmt`, `bytes` or `strings`)
- No external packages required
