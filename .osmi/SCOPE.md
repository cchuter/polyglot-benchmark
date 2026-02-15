# Scope: polyglot-go-beer-song (Issue #15)

## In Scope

- Verify the existing beer-song implementation at `go/exercises/practice/beer-song/`
- Ensure all tests pass: `TestBottlesVerse`, `TestSeveralVerses`, `TestEntireSong`
- Ensure benchmarks run: `BenchmarkSeveralVerses`, `BenchmarkEntireSong`
- Create feature branch `issue-15` and commit any necessary changes
- Close issue #15 via PR

### Files in scope
- `go/exercises/practice/beer-song/beer_song.go` — Implementation
- `go/exercises/practice/beer-song/beer_song_test.go` — Tests (read-only, do not modify)
- `go/exercises/practice/beer-song/go.mod` — Module definition

## Out of Scope

- Modifying test files (they are part of the exercise specification)
- Adding new exercises or modifying other exercises
- Changing the repository structure or build system
- Implementing "bonus points" refactoring suggestions from the exercise description
- Modifying any files outside `go/exercises/practice/beer-song/`

## Dependencies

- Go 1.18+ toolchain
- No external Go packages (standard library only: `bytes`, `fmt`)
- No dependencies on other exercises in the repository
