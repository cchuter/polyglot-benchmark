# Implementation Plan: DnD Character Generator

## Overview

Implement the D&D character generator in a single file: `go/exercises/practice/dnd-character/dnd_character.go`.

## Files to Modify

- `go/exercises/practice/dnd-character/dnd_character.go` — the only file to change

## Implementation Details

### 1. Character Struct

```go
type Character struct {
    Strength     int
    Dexterity    int
    Constitution int
    Intelligence int
    Wisdom       int
    Charisma     int
    Hitpoints    int
}
```

Fields match what the tests expect (exported fields accessed as `character.Strength`, etc.).

### 2. Modifier Function

```go
func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}
```

Uses `math.Floor` for correct rounding toward negative infinity (important for odd scores below 10, e.g. score=3 → -4, not -3).

### 3. Ability Function

Roll 4 six-sided dice, drop the lowest, sum the top 3:

```go
func Ability() int {
    rolls := make([]int, 4)
    sum := 0
    min := 7
    for i := 0; i < 4; i++ {
        rolls[i] = rand.Intn(6) + 1
        sum += rolls[i]
        if rolls[i] < min {
            min = rolls[i]
        }
    }
    return sum - min
}
```

Using a manual min instead of `slices.Min` to avoid compatibility issues with `go 1.18` in the module file. This is simpler and has no import dependency.

### 4. GenerateCharacter Function

```go
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

### 5. Imports

Only need `math` and `math/rand`.

## Approach and Ordering

1. Write the complete solution in `dnd_character.go`
2. Run `go test ./...` in the exercise directory to verify all tests pass
3. Run `go vet ./...` to check for issues
4. Commit the change

## Architectural Decisions

- **Manual min instead of `slices.Min`**: The `go.mod` specifies `go 1.18` but `slices` is a Go 1.21 package. While the runtime supports it, keeping compatibility with the declared module version is safer. Tracking min during iteration is O(1) space and simpler.
- **`math.Floor` for modifier**: Integer division in Go truncates toward zero, which gives wrong results for negative values. `math.Floor` correctly rounds toward negative infinity.
