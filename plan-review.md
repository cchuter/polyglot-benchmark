# Plan Review: D&D Character Generator

## Overall Verdict: PASS — The plan is correct and should produce a working solution.

The selected plan (Proposal B) is sound. Below is a detailed evaluation against each criterion.

---

## 1. Will it compile correctly with Go 1.18?

**Yes.** The final implementation uses only these imports:

- `math` (standard library, available since Go 1.0)
- `math/rand` (standard library, available since Go 1.0)
- `sort` (standard library, available since Go 1.0)

The plan correctly identified and rejected Proposal A's use of `slices.Min`, which requires Go 1.21 and would fail to compile under the `go 1.18` directive in `go.mod`. The package declaration `package dndcharacter` matches both the test files and the existing stub.

**No issues.**

---

## 2. Does the `Modifier` function handle all test cases correctly (scores 3-18)?

**Yes.** The implementation is:

```go
func Modifier(score int) int {
    return int(math.Floor(float64(score-10) / 2.0))
}
```

Verification against every test case in `cases_test.go`:

| Score | `(score-10)/2.0` | `math.Floor(...)` | `int(...)` | Expected | Match? |
|-------|-------------------|--------------------|------------|----------|--------|
| 3     | -3.5              | -4.0               | -4         | -4       | Yes    |
| 4     | -3.0              | -3.0               | -3         | -3       | Yes    |
| 5     | -2.5              | -3.0               | -3         | -3       | Yes    |
| 6     | -2.0              | -2.0               | -2         | -2       | Yes    |
| 7     | -1.5              | -2.0               | -2         | -2       | Yes    |
| 8     | -1.0              | -1.0               | -1         | -1       | Yes    |
| 9     | -0.5              | -1.0               | -1         | -1       | Yes    |
| 10    | 0.0               | 0.0                | 0          | 0        | Yes    |
| 11    | 0.5               | 0.0                | 0          | 0        | Yes    |
| 12    | 1.0               | 1.0                | 1          | 1        | Yes    |
| 13    | 1.5               | 1.0                | 1          | 1        | Yes    |
| 14    | 2.0               | 2.0                | 2          | 2        | Yes    |
| 15    | 2.5               | 2.0                | 2          | 2        | Yes    |
| 16    | 3.0               | 3.0                | 3          | 3        | Yes    |
| 17    | 3.5               | 3.0                | 3          | 3        | Yes    |
| 18    | 4.0               | 4.0                | 4          | 4        | Yes    |

All 16 test cases pass. The critical behavior is for odd scores below 10 (e.g., score 3, 5, 7, 9) where Go's integer division would truncate toward zero (rounding up), but `math.Floor` correctly rounds down. This is the right approach.

**No issues.**

---

## 3. Does the `Ability` function produce values in [3, 18]?

**Yes.** The implementation:

```go
func Ability() int {
    rolls := make([]int, 4)
    for i := range rolls {
        rolls[i] = rand.Intn(6) + 1
    }
    sort.Ints(rolls)
    return rolls[1] + rolls[2] + rolls[3]
}
```

- Each die roll produces a value in [1, 6].
- `sort.Ints(rolls)` sorts in ascending order.
- `rolls[1] + rolls[2] + rolls[3]` sums the top 3 values (dropping the minimum at index 0).
- **Minimum possible:** All four dice roll 1. Sorted: [1, 1, 1, 1]. Sum of top 3: 1+1+1 = 3.
- **Maximum possible:** All four dice roll 6. Sorted: [6, 6, 6, 6]. Sum of top 3: 6+6+6 = 18.

The test (`TestAbility`) checks that 1000 calls all produce values in [3, 18] inclusive. This implementation satisfies that constraint.

**No issues.**

---

## 4. Does `GenerateCharacter` set hitpoints correctly?

**Yes.** The implementation:

```go
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

The test (`TestGenerateCharacter`) computes `expectedHitpoints := 10 + Modifier(character.Constitution)` and checks it against `character.Hitpoints`. The plan does exactly this: it sets all ability scores first via struct literal, then computes hitpoints from the Constitution score using the same `Modifier` function.

The two-step approach (struct literal first, then hitpoints assignment) is important because it ensures `character.Constitution` is set before it is used in the hitpoints calculation. This is correct.

**No issues.**

---

## 5. Are all exported names correct?

**Yes.** Cross-referencing the plan's exports against what the tests expect:

| Test References                     | Plan Provides         | Match? |
|-------------------------------------|-----------------------|--------|
| `Modifier(tc.input.Score)`          | `func Modifier(score int) int` | Yes |
| `Ability()`                         | `func Ability() int`  | Yes |
| `GenerateCharacter()`               | `func GenerateCharacter() Character` | Yes |
| `character.Strength`                | `Character.Strength int` | Yes |
| `character.Dexterity`               | `Character.Dexterity int` | Yes |
| `character.Constitution`            | `Character.Constitution int` | Yes |
| `character.Intelligence`            | `Character.Intelligence int` | Yes |
| `character.Wisdom`                  | `Character.Wisdom int` | Yes |
| `character.Charisma`                | `Character.Charisma int` | Yes |
| `character.Hitpoints`               | `Character.Hitpoints int` | Yes |

The package name `dndcharacter` also matches the existing stub and test files.

**No issues.**

---

## 6. Any other issues?

**None found.** Additional checks:

- **Randomness seeding:** The plan uses `math/rand` without explicit seeding. In Go 1.20+, the default source is automatically seeded. In Go 1.18, the default source starts with a fixed seed (seed 0), but this does not affect test correctness -- the tests only check that values are in range, not that they are truly random. The tests run `Ability()` 1000 times and check range bounds, which will pass regardless of seed. If deterministic output were a concern, explicit seeding could be added, but it is not required for the tests to pass.
- **No unused imports:** All three imports (`math`, `math/rand`, `sort`) are used.
- **No external dependencies:** The solution uses only standard library packages, so no `go.sum` or dependency management is needed.
- **Benchmark compatibility:** The benchmarks (`BenchmarkModifier`, `BenchmarkAbility`, `BenchmarkCharacter`) call the same exported functions and will work correctly with this implementation.

---

## Summary

The selected plan (Proposal B) is correct and complete. It will compile under Go 1.18, produces correct results for all test cases, uses proper exported names, and has no identifiable issues. The plan is ready for implementation.
