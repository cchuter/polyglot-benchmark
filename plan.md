# Implementation Plan: D&D Character Generator

## Proposal A (Proponent)

**Approach: Use `math.Floor` with float conversion and `slices.Min` for dice logic**

### Files to Modify
- `go/exercises/practice/dnd-character/dnd_character.go` — implement all required types and functions

### Architecture
Follow the reference solution pattern closely. Use `math.Floor` for the modifier calculation and the `slices` package (available in Go 1.21+, but also in `golang.org/x/exp/slices` for 1.18) for finding the minimum die roll.

### Implementation
1. Define `Character` struct with 7 int fields.
2. `Modifier(score int) int` — uses `math.Floor(float64(score-10) / 2.0)` and casts to int.
3. Helper `RollDice() int` — returns `rand.Intn(6) + 1`.
4. `Ability() int` — rolls 4 dice, sums all, subtracts the minimum using `slices.Min`.
5. `GenerateCharacter() Character` — creates Character with 6 `Ability()` calls and sets Hitpoints.

### Rationale
- Closely follows the reference solution in `.meta/example.go`.
- Clean, readable code.
- `slices.Min` is elegant for finding the lowest die.

### Weakness
- `slices` package requires Go 1.21, but `go.mod` specifies Go 1.18. This would cause a compilation error.

---

## Proposal B (Opponent)

**Approach: Use integer arithmetic and `sort.Ints` for dice logic**

### Files to Modify
- `go/exercises/practice/dnd-character/dnd_character.go` — implement all required types and functions

### Architecture
Use only packages available in Go 1.18. For Modifier, use integer division which naturally floors for positive dividends, and handle negative values correctly via integer math. For finding the top 3 dice, sort the 4 rolls and sum the last 3.

### Implementation
1. Define `Character` struct with 7 int fields.
2. `Modifier(score int) int` — use `(score - 10) / 2` with a correction for odd negative values: `if (score-10)%2 != 0 && score < 10 { result-- }`. Or more simply: use `(score - 10)` and if negative and odd, subtract 1, then divide by 2. Even simpler: `(score - 10) / 2` works correctly for all cases when score >= 10. For score < 10 with odd difference, Go's integer division truncates toward zero (rounds up for negatives), so we need floor. The cleanest Go 1.18 approach: `math.Floor(float64(score-10) / 2.0)` — this import is fine in Go 1.18.
3. Helper: roll a die with `rand.Intn(6) + 1`.
4. `Ability() int` — roll 4 dice into a slice, `sort.Ints(slice)`, sum the last 3 elements.
5. `GenerateCharacter() Character` — same pattern as Proposal A.

### Critique of Proposal A
Proposal A's use of `slices.Min` will fail on Go 1.18 since the `slices` package was added in Go 1.21. This is a **compilation-breaking issue**.

### Rationale
- Uses only Go 1.18-compatible standard library packages (`math`, `math/rand`, `sort`).
- `sort.Ints` + summing top 3 is clear and idiomatic.
- `math.Floor` for the modifier is safe and correct across all Go versions.

---

## Selected Plan (Judge)

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Fails to compile on Go 1.18 (`slices` unavailable) | Correct across all versions |
| Risk | High — build failure | Low |
| Simplicity | Slightly simpler min logic | Sort + slice is equally clear |
| Consistency | Matches reference but incompatible with go.mod | Compatible with constraints |

**Winner: Proposal B** — Proposal A has a critical compatibility issue. Proposal B is correct, simple, and compatible with the Go 1.18 module constraint.

### Final Implementation Plan

**File:** `go/exercises/practice/dnd-character/dnd_character.go`

```go
package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

// Character represents a D&D character with ability scores and hitpoints.
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score.
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability.
func Ability() int {
	rolls := make([]int, 4)
	for i := range rolls {
		rolls[i] = rand.Intn(6) + 1
	}
	sort.Ints(rolls)
	return rolls[1] + rolls[2] + rolls[3]
}

// GenerateCharacter creates a new Character with random scores for abilities.
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	character.Hitpoints = 10 + Modifier(character.Constitution)
	return character
}
```

### Steps
1. Write the solution to `dnd_character.go`.
2. Run `go test ./...` to verify all tests pass.
3. Run `go vet ./...` to verify no issues.
4. Commit with descriptive message.
