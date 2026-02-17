# Solo Agent Change Log

## Change 1: Implement D&D Character Generator

**File modified:** `go/exercises/practice/dnd-character/dnd_character.go`

**What was done:**
- Defined `Character` struct with 7 int fields (Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma, Hitpoints)
- Implemented `Modifier(score int) int` using `math.Floor` for correct floor division
- Implemented `Ability() int` that rolls 4d6, sorts, and sums the top 3
- Implemented `GenerateCharacter() Character` that generates all ability scores and computes hitpoints

**Key decisions:**
- Used `sort.Ints` instead of `slices.Min` for Go 1.18 compatibility
- Used `math.Floor` for modifier to correctly handle negative odd values

**Test results:** All tests pass (16 modifier tests, ability range test, character generation test). `go vet` clean.
