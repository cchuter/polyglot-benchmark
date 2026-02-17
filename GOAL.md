# Goal: DnD Character Generator (Go)

## Problem Statement

Implement a D&D character generator in Go that creates characters with six randomly-determined ability scores and calculated hitpoints. The stub file `go/exercises/practice/dnd-character/dnd_character.go` currently contains only a package declaration and must be filled with the complete implementation.

## What Needs to Be Built

1. **`Character` struct** with fields: `Strength`, `Dexterity`, `Constitution`, `Intelligence`, `Wisdom`, `Charisma`, `Hitpoints` (all `int`)

2. **`Modifier(score int) int`** - Calculates the ability modifier: `floor((score - 10) / 2)`

3. **`Ability() int`** - Rolls four 6-sided dice and returns the sum of the three highest rolls (value range: 3-18)

4. **`GenerateCharacter() Character`** - Creates a new Character with random ability scores (each via `Ability()`) and hitpoints calculated as `10 + Modifier(constitution)`

## Acceptance Criteria

- [ ] `Modifier` returns correct values for all scores 3-18 (per `cases_test.go` table)
- [ ] `Ability()` always returns values in [3, 18] range (tested 1000 times)
- [ ] `GenerateCharacter()` produces characters where all six ability scores are in [3, 18]
- [ ] `GenerateCharacter()` produces characters with `Hitpoints == 10 + Modifier(Constitution)`
- [ ] All tests pass: `go test ./...` in the exercise directory
- [ ] Benchmarks run without error

## Key Constraints

- Package name: `dndcharacter`
- Must use Go 1.18+ (per go.mod)
- Solution must only modify `dnd_character.go`
- Tests are read-only; solution must conform to the test API
