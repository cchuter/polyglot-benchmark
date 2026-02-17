# Code Review: dnd-character Implementation

**Reviewer**: challenger
**Status**: APPROVED
**File**: `go/exercises/practice/dnd-character/dnd_character.go`

## Summary

The implementation is correct, clean, and passes all tests and benchmarks.

## Detailed Review

### 1. Correctness: Modifier

**Verdict**: PASS

`Modifier` uses `math.Floor(float64(score-10) / 2.0)` which correctly handles floor division for negative scores. Go's native integer division truncates toward zero (e.g., `-7/2 = -3`), but D&D requires floor division (e.g., `floor(-7/2) = -4`). The float conversion + `math.Floor` approach handles this correctly.

Verified against all 16 test cases (scores 3-18): all produce expected results.

### 2. Edge Cases: Ability Range [3, 18]

**Verdict**: PASS

- Rolls 4d6 via `rand.Intn(6) + 1` producing values in [1, 6].
- Drops lowest via `sum - slices.Min(scores)`.
- Minimum: all 1s -> 1+1+1+1 - 1 = 3.
- Maximum: all 6s -> 6+6+6+6 - 6 = 18.
- Range is guaranteed to be [3, 18].

The test runs 1000 iterations and confirms this empirically.

### 3. API Conformance

**Verdict**: PASS

- `Character` struct: 7 fields (`Strength`, `Dexterity`, `Constitution`, `Intelligence`, `Wisdom`, `Charisma`, `Hitpoints`), all `int`. Matches test expectations exactly.
- `Modifier(score int) int`: correct signature.
- `Ability() int`: correct signature.
- `GenerateCharacter() Character`: correct signature. Sets `Hitpoints = 10 + Modifier(Constitution)`.

### 4. Code Quality

**Verdict**: PASS

- Clean, idiomatic Go.
- Appropriate use of standard library (`math`, `math/rand`, `slices`).
- No unnecessary complexity or over-engineering.
- No security concerns.

### Minor Observation (Non-blocking)

The `go.mod` specifies `go 1.18`, but `slices.Min` requires Go 1.21+. This works fine as long as the build toolchain is Go 1.21+, which it is (tests compile and pass). Not a functional issue, but `go.mod` could be updated to `go 1.21` for accuracy. Since the tests are read-only and the constraint says "must only modify dnd_character.go", this is acceptable as-is.

## Test Results

- All 3 test functions pass (16 modifier subtests + ability range test + character generation test).
- All 3 benchmarks run without error.

## Final Verdict

**APPROVED** - No changes required. Implementation is correct, complete, and idiomatic.
