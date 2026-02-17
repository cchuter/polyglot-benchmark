# Changes

## Implement dnd-character solution

- Wrote complete implementation in `go/exercises/practice/dnd-character/dnd_character.go`
- Added `Character` struct with fields: Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma, Hitpoints
- Added `Modifier(score int) int` function using floor division formula
- Added `Ability() int` function that rolls 4d6 and sums the top 3
- Added `GenerateCharacter() Character` function that generates a character with random abilities and calculated hitpoints
