# Changes - D&D Character Generator

## Implemented: `dnd_character.go`

- **Character struct**: Defined with 7 fields — Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma, Hitpoints (all `int`).
- **Modifier(score int) int**: Calculates the ability modifier using `math.Floor(float64(score-10) / 2.0)` cast to `int`.
- **Ability() int**: Rolls 4d6 (using `rand.Intn(6)+1`), tracks the sum and minimum roll in a loop, returns `sum - min` (drop lowest die).
- **GenerateCharacter() Character**: Generates all 6 ability scores via `Ability()`, then sets `Hitpoints = 10 + Modifier(constitution)`.

Imports: `math`, `math/rand`. Go 1.18 compatible.
