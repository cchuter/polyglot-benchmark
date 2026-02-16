# Solo Agent Change Log

## Change 1: Implement dnd-character solution

**File modified:** `go/exercises/practice/dnd-character/dnd_character.go`

**Changes:**
- Added `Character` struct with 7 fields: Strength, Dexterity, Constitution, Intelligence, Wisdom, Charisma, Hitpoints
- Added `Modifier(score int) int` — computes ability modifier using `math.Floor((score-10)/2)`
- Added `Ability() int` — rolls 4d6, drops lowest, sums top 3 using manual min tracking
- Added `GenerateCharacter() Character` — generates all 6 abilities randomly, computes hitpoints

**Test results:** All 3 test suites pass (TestModifier: 16/16, TestAbility: 1/1, TestGenerateCharacter: 1/1). `go vet` clean.

**Decision:** Used manual min tracking instead of `slices.Min` to stay compatible with `go 1.18` module declaration.
