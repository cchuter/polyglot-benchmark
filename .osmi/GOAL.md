# Goal: DnD Character Generator (Go)

## Problem Statement

Implement a Dungeons & Dragons character generator in Go. The stub file `go/exercises/practice/dnd-character/dnd_character.go` currently contains only the package declaration. It needs to be filled with:

1. A `Character` struct with fields for six abilities (Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma) and Hitpoints.
2. A `Modifier(score int) int` function that computes the ability modifier: `floor((score - 10) / 2)`.
3. An `Ability() int` function that rolls 4d6, drops the lowest, and sums the remaining 3.
4. A `GenerateCharacter() Character` function that creates a character with random ability scores and calculates hitpoints as `10 + Modifier(constitution)`.

## Acceptance Criteria

1. `Modifier(score)` returns correct values for all scores 3-18 (verified by `modifierTests` in `cases_test.go`).
2. `Ability()` returns values in the range [3, 18] across 1000 iterations.
3. `GenerateCharacter()` produces a `Character` with all six ability scores in [3, 18] and `Hitpoints == 10 + Modifier(Constitution)`.
4. All tests pass: `go test ./...` in the exercise directory.
5. `go vet ./...` reports no issues.

## Key Constraints

- The solution file is `dnd_character.go` in `package dndcharacter`.
- Test files (`dnd_character_test.go`, `cases_test.go`) are read-only; do not modify them.
- The module uses Go 1.18 (`go.mod`).
- Random dice rolls should use `math/rand`.
