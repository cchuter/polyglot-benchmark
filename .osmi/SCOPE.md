# Scope: DnD Character Generator

## In Scope

- Implement `Character` struct in `dnd_character.go`
- Implement `Modifier(score int) int` function
- Implement `Ability() int` function (roll 4d6, drop lowest, sum top 3)
- Implement `GenerateCharacter() Character` function
- Ensure all existing tests pass

## Out of Scope

- Modifying test files (`dnd_character_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Adding CLI or main package
- Any other exercises in the repository

## Dependencies

- `math` (for `math.Floor`)
- `math/rand` (for dice rolling)
- `slices` (for `slices.Min` — available in Go 1.21+, but the example uses it; alternatively use a manual min)
- No external dependencies
