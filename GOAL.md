# Goal: Implement D&D Character Generator (Go)

## Problem Statement

Implement a D&D character generator in Go that creates characters with six randomly-generated ability scores (Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma) and calculated hitpoints.

Each ability score is determined by rolling four 6-sided dice and summing the three highest rolls. Hitpoints are calculated as `10 + constitution modifier`, where the constitution modifier is `floor((constitution - 10) / 2)`.

## Acceptance Criteria

1. **`Modifier(score int) int`** — Given an ability score, returns the ability modifier: `floor((score - 10) / 2)`. Must pass all 15 test cases in `cases_test.go` (scores 3-18).

2. **`Ability() int`** — Returns a randomly-generated ability score by rolling 4d6 and summing the top 3. Must always return a value in [3, 18] inclusive (validated over 1000 iterations).

3. **`Character` struct** — Must have fields: `Strength`, `Dexterity`, `Constitution`, `Intelligence`, `Wisdom`, `Charisma`, `Hitpoints` (all `int`).

4. **`GenerateCharacter() Character`** — Returns a Character with all six ability scores randomly generated via `Ability()`, and `Hitpoints` set to `10 + Modifier(Constitution)`. Validated over 1000 iterations.

5. All tests pass: `go test ./...` in the exercise directory.
6. Code passes `go vet ./...`.

## Key Constraints

- Must use the `dndcharacter` package name.
- Solution goes in `dnd_character.go`.
- Must not modify test files.
- Go 1.18 module.
