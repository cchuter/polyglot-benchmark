# Scope: D&D Character Generator

## In Scope

- Implement `Modifier(score int) int` function
- Implement `Ability() int` function
- Define `Character` struct with seven int fields
- Implement `GenerateCharacter() Character` function
- All code in `go/exercises/practice/dnd-character/dnd_character.go`
- Passing all existing tests in `dnd_character_test.go` and `cases_test.go`

## Out of Scope

- Modifying test files (`dnd_character_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test files
- Changing the package name or module structure
- Any other exercises or languages in the repository

## Dependencies

- Go standard library only (`math`, `math/rand`)
- No external packages
- Must compile with Go 1.18+ (go.mod constraint)
