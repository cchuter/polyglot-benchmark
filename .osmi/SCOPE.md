# Scope: polyglot-go-alphametics (Issue #21)

## In Scope

- Verify the existing alphametics implementation at `go/exercises/practice/alphametics/`
- Ensure all 10 test cases pass: `go test -v ./...`
- Ensure benchmarks run successfully
- Create feature branch `issue-21` and commit any necessary changes
- Close issue #21 via PR

### Files in scope
- `go/exercises/practice/alphametics/alphametics.go` — Implementation (verify, fix if needed)
- `go/exercises/practice/alphametics/alphametics_test.go` — Tests (read-only, do not modify)
- `go/exercises/practice/alphametics/cases_test.go` — Test cases (read-only, do not modify)
- `go/exercises/practice/alphametics/go.mod` — Module definition (verify)

## Out of Scope

- Modifying test files (they are part of the exercise specification)
- Adding new exercises or modifying other exercises
- Changing the repository structure or build system
- Performance optimization beyond what's needed to pass tests
- Modifying any files outside `go/exercises/practice/alphametics/` and `.osmi/`

## Dependencies

- Go 1.18+ toolchain
- No external Go packages (standard library only: `errors`, `strings`, `unicode`)
- No dependencies on other exercises in the repository
