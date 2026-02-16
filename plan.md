# Implementation Plan: D&D Character Generator

## Branch 1: Direct Implementation (Minimal, Simple)

### Approach
Write the simplest possible implementation directly in `dnd_character.go`. Use `math` for floor division and `math/rand` for dice rolls. Find the minimum of four dice manually with a loop instead of using `slices.Min` (which requires Go 1.21+).

### Files to Modify
- `go/exercises/practice/dnd-character/dnd_character.go` — complete implementation

### Implementation Details
```go
package dndcharacter

import (
    "math"
    "math/rand"
)

type Character struct {
    Strength     int
    Dexterity    int
    Constitution int
    Intelligence int
    Wisdom       int
    Charisma     int
    Hitpoints    int
}

func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}

func Ability() int {
    dice := make([]int, 4)
    sum := 0
    min := 7
    for i := 0; i < 4; i++ {
        dice[i] = rand.Intn(6) + 1
        sum += dice[i]
        if dice[i] < min {
            min = dice[i]
        }
    }
    return sum - min
}

func GenerateCharacter() Character {
    c := Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: Ability(),
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
    }
    c.Hitpoints = 10 + Modifier(c.Constitution)
    return c
}
```

### Evaluation
- **Feasibility**: High — uses only Go 1.18-compatible stdlib
- **Risk**: Very low — straightforward logic, no edge cases
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~40 lines of code

---

## Branch 2: Sort-Based Approach (Extensible)

### Approach
Use `sort.Ints` to sort the four dice rolls, then sum the top three. This is slightly more extensible if the dice-dropping rule changes.

### Files to Modify
- `go/exercises/practice/dnd-character/dnd_character.go`

### Implementation Details
```go
package dndcharacter

import (
    "math"
    "math/rand"
    "sort"
)

type Character struct { /* same fields */ }

func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}

func Ability() int {
    dice := make([]int, 4)
    for i := range dice {
        dice[i] = rand.Intn(6) + 1
    }
    sort.Ints(dice)
    return dice[1] + dice[2] + dice[3] // drop lowest (dice[0])
}

func GenerateCharacter() Character {
    c := Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: Ability(),
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
    }
    c.Hitpoints = 10 + Modifier(c.Constitution)
    return c
}
```

### Evaluation
- **Feasibility**: High — `sort.Ints` available in all Go versions
- **Risk**: Low — sorting 4 elements is trivial
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~40 lines, one extra import

---

## Branch 3: Integer-Only Modifier (Performance)

### Approach
Avoid `math.Floor` and floating-point entirely. Use integer arithmetic for the modifier: `(score - 10) / 2` with a correction for negative odd values (Go truncates toward zero, but D&D rules require floor division).

### Files to Modify
- `go/exercises/practice/dnd-character/dnd_character.go`

### Implementation Details
```go
package dndcharacter

import "math/rand"

type Character struct { /* same fields */ }

func Modifier(score int) int {
    d := score - 10
    if d >= 0 {
        return d / 2
    }
    return (d - 1) / 2
}

func Ability() int {
    sum, min := 0, 7
    for i := 0; i < 4; i++ {
        roll := rand.Intn(6) + 1
        sum += roll
        if roll < min {
            min = roll
        }
    }
    return sum - min
}

func GenerateCharacter() Character {
    c := Character{
        Strength:     Ability(),
        Dexterity:    Ability(),
        Constitution: Ability(),
        Intelligence: Ability(),
        Wisdom:       Ability(),
        Charisma:     Ability(),
    }
    c.Hitpoints = 10 + Modifier(c.Constitution)
    return c
}
```

### Evaluation
- **Feasibility**: High — pure integer arithmetic
- **Risk**: Low-medium — the integer floor division formula needs to be correct for negative values
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: Single file, ~35 lines, one fewer import

---

## Selected Plan

**Branch 1: Direct Implementation** is selected.

### Rationale
- **Simplest and most readable**: The tracking-minimum approach is O(n) and avoids both sorting overhead and integer-division subtleties.
- **Closest to reference solution**: Mirrors the `.meta/example.go` pattern, just replacing `slices.Min` with an inline minimum tracking loop for Go 1.18 compatibility.
- **Lowest risk**: `math.Floor` for the modifier is proven correct in the reference solution and avoids the edge-case reasoning needed for Branch 3's integer arithmetic.
- **No unnecessary imports**: Only `math` and `math/rand` are needed.

### Full Implementation Plan

**File**: `go/exercises/practice/dnd-character/dnd_character.go`

1. Define the `Character` struct with 7 int fields: `Strength`, `Dexterity`, `Constitution`, `Intelligence`, `Wisdom`, `Charisma`, `Hitpoints`
2. Implement `Modifier(score int) int` using `math.Floor(float64(score-10) / 2.0)` cast to int
3. Implement `Ability() int`: roll 4d6, track sum and minimum simultaneously, return `sum - min`
4. Implement `GenerateCharacter() Character`: call `Ability()` for each of the 6 scores, compute hitpoints as `10 + Modifier(constitution)`
5. Run `go test ./...` to verify all tests pass
6. Run `go test -bench=.` to verify benchmarks work
