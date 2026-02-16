# Challenger Review: dnd-character Implementation

## Verdict: PASS — No issues found

## Detailed Analysis

### 1. Modifier Function — CORRECT

```go
func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}
```

Verified against all 16 test cases:

| Score | Computation              | Result | Expected | Pass |
|-------|--------------------------|--------|----------|------|
| 3     | floor((3-10)/2) = floor(-3.5) | -4 | -4 | Y |
| 4     | floor((4-10)/2) = floor(-3.0) | -3 | -3 | Y |
| 5     | floor((5-10)/2) = floor(-2.5) | -3 | -3 | Y |
| 6     | floor((6-10)/2) = floor(-2.0) | -2 | -2 | Y |
| 7     | floor((7-10)/2) = floor(-1.5) | -2 | -2 | Y |
| 8     | floor((8-10)/2) = floor(-1.0) | -1 | -1 | Y |
| 9     | floor((9-10)/2) = floor(-0.5) | -1 | -1 | Y |
| 10    | floor((10-10)/2) = floor(0.0) | 0  | 0  | Y |
| 11    | floor((11-10)/2) = floor(0.5) | 0  | 0  | Y |
| 12    | floor((12-10)/2) = floor(1.0) | 1  | 1  | Y |
| 13    | floor((13-10)/2) = floor(1.5) | 1  | 1  | Y |
| 14    | floor((14-10)/2) = floor(2.0) | 2  | 2  | Y |
| 15    | floor((15-10)/2) = floor(2.5) | 2  | 2  | Y |
| 16    | floor((16-10)/2) = floor(3.0) | 3  | 3  | Y |
| 17    | floor((17-10)/2) = floor(3.5) | 3  | 3  | Y |
| 18    | floor((18-10)/2) = floor(4.0) | 4  | 4  | Y |

Key correctness point: `math.Floor` properly handles floor division for negative odd scores (e.g., score 3 yields -4 not -3, score 5 yields -3 not -2). Go's integer division truncates toward zero, so using `math.Floor` with floating-point division is the correct approach.

### 2. Ability Function — CORRECT

```go
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
```

- Rolls 4d6 (four 6-sided dice): `rand.Intn(6) + 1` produces [1, 6]. Correct.
- Tracks the minimum across all 4 rolls: `min` initialized to 7 (> max die value of 6), guaranteeing it will be updated on the first roll. Correct.
- Returns `sum - min`, which equals the sum of the three highest dice. Correct.
- **Range**: Min = 4×1 - 1 = 3. Max = 4×6 - 6 = 18. Always in [3, 18]. Correct.

### 3. Character Struct — CORRECT

All 7 required fields present with correct names and types:
- `Strength int`
- `Dexterity int`
- `Constitution int`
- `Intelligence int`
- `Wisdom int`
- `Charisma int`
- `Hitpoints int`

### 4. GenerateCharacter — CORRECT

- All 6 ability scores generated via `Ability()`. Correct.
- `Hitpoints = 10 + Modifier(c.Constitution)`. Matches the formula exactly. Correct.
- Constitution is set before Hitpoints calculation (struct literal evaluates first, then the separate assignment). Correct.

### 5. Go 1.18 Compatibility — CORRECT

- Imports: `math` and `math/rand` — both available since Go 1.0.
- No use of `slices` package (Go 1.21+). Correct.
- No use of generics beyond what Go 1.18 supports. No generics used at all. Correct.
- `rand.Intn` available since Go 1.0. Correct.

### 6. Minor Note (non-blocking)

`math/rand` is not explicitly seeded. In Go < 1.20, this means deterministic output (default seed = fixed value). However, the tests only validate range constraints and the hitpoints formula, not randomness distribution, so all tests will pass regardless. In Go >= 1.20, `math/rand` auto-seeds.

## Conclusion

The implementation is clean, correct, and fully satisfies all acceptance criteria. All 16 modifier test cases, the 1000-iteration ability range check, and the 1000-iteration character generation validation should pass. No changes needed.
