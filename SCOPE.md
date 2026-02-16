# Scope: polyglot-go-alphametics

## In Scope

- Implement the `Solve` function in `go/exercises/practice/alphametics/alphametics.go`
- Parse puzzle strings to extract words and identify addends vs. result
- Implement constraint-based digit assignment with backtracking
- Handle all edge cases covered by the test suite:
  - Simple 3-letter puzzles
  - Puzzles that have no solution (return error)
  - Leading zero constraint enforcement
  - Puzzles with many addends (up to 199)
  - Puzzles with up to 10 unique letters

## Out of Scope

- Modifying test files (`alphametics_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding external dependencies
- Benchmark optimization beyond passing tests within timeout
- Changes to any other exercise or language directory

## Dependencies

- Go 1.18+ toolchain
- Standard library only (no external packages)
- Existing test infrastructure in the repository
