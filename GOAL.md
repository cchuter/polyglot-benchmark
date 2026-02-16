# Goal: Implement D&D Character Generator (Go)

## Problem Statement

Implement a D&D character generator in Go that creates characters with six randomly-generated ability scores (strength, dexterity, constitution, intelligence, wisdom, charisma) and calculates hitpoints based on the constitution modifier.

The stub file `go/exercises/practice/dnd-character/dnd_character.go` currently contains only a package declaration and needs a complete implementation.

## Rules

1. **Ability Score Generation**: Roll four 6-sided dice, discard the lowest, sum the remaining three. Each score ranges from 3 to 18.
2. **Modifier Calculation**: `modifier = floor((score - 10) / 2)`
3. **Hitpoints**: `hitpoints = 10 + modifier(constitution)`

## Acceptance Criteria

1. `Modifier(score int) int` correctly computes the ability modifier using floor division: `floor((score - 10) / 2)`. Must pass all 16 test cases in `cases_test.go` (scores 3–18).
2. `Ability() int` returns a random score between 3 and 18 inclusive (simulating 4d6-drop-lowest). Must pass 1000-iteration range check.
3. `Character` struct has fields: `Strength`, `Dexterity`, `Constitution`, `Intelligence`, `Wisdom`, `Charisma`, `Hitpoints` (all `int`).
4. `GenerateCharacter() Character` returns a Character with all six abilities in [3,18] and `Hitpoints == 10 + Modifier(Constitution)`. Must pass 1000-iteration validation.
5. All tests pass: `go test ./...` in the exercise directory.
6. All benchmarks run without error.

## Key Constraints

- Package name: `dndcharacter`
- Must use Go's standard library (no external dependencies)
- go.mod specifies `go 1.18`; `slices` package requires Go 1.21+, so the reference solution's use of `slices.Min` needs adjustment or the implementation must use an alternative approach
