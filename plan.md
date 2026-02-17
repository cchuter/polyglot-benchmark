# Implementation Plan: polyglot-go-octal (Issue #309)

## Proposal A — Bit-Shift Approach (Proponent)

### Approach
Use bit-shifting (`<<3`) instead of multiplication by 8 for the base conversion. Iterate over each rune in the input string, validate it is in the range '0'-'7', and accumulate the result.

### Files to Modify
- `go/exercises/practice/octal/octal.go` — implement `ParseOctal`

### Implementation
```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for _, digit := range input {
        if digit < '0' || digit > '7' {
            return 0, fmt.Errorf("invalid octal digit '%c'", digit)
        }
        num = num<<3 + int64(digit-'0')
    }
    return num, nil
}
```

### Rationale
- **Matches the reference solution** in `.meta/example.go` almost exactly
- Bit-shifting is idiomatic for base-8/base-2/base-16 conversions
- Simple, single-pass, O(n) algorithm
- Minimal imports (only `fmt` for error formatting)
- Concise — under 15 lines of code

### Strengths
- Proven correct (matches example.go)
- Maximum simplicity
- Performant (bit shift vs multiply)

---

## Proposal B — Explicit Multiplication with Math.Pow Approach (Opponent)

### Approach
Process the string from right to left, using explicit multiplication by powers of 8. Use a position counter to track the current power.

### Files to Modify
- `go/exercises/practice/octal/octal.go` — implement `ParseOctal`

### Implementation
```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    result := int64(0)
    power := int64(1)
    for i := len(input) - 1; i >= 0; i-- {
        c := input[i]
        if c < '0' || c > '7' {
            return 0, fmt.Errorf("invalid octal digit '%c'", c)
        }
        result += int64(c-'0') * power
        power *= 8
    }
    return result, nil
}
```

### Rationale
- More explicit and readable — the math is visible (multiply by powers of 8)
- Mirrors the mathematical definition from the problem description
- No bit-shifting "magic"

### Critique of Proposal A
- Bit-shifting is a micro-optimization that obscures intent for readers unfamiliar with the trick
- Using `range` over the string allocates runes; indexing bytes directly is more efficient

### Strengths
- Clearer mapping to the mathematical definition
- Byte indexing avoids rune allocation overhead

### Weaknesses
- Slightly more code (power variable, reverse iteration)
- Diverges from the reference solution pattern

---

## Selected Plan (Judge)

### Evaluation

| Criterion    | Proposal A (Bit-Shift) | Proposal B (Explicit Multiply) |
|-------------|------------------------|-------------------------------|
| Correctness | Matches reference solution exactly | Correct but untested variant |
| Risk        | Very low — proven approach | Low — straightforward math but diverges from reference |
| Simplicity  | Simpler — fewer variables, forward iteration | Slightly more complex — reverse iteration, power variable |
| Consistency | Matches `.meta/example.go` and hexadecimal exercise patterns | Different style from other exercises |

### Decision: Proposal A wins

**Rationale:**
1. It matches the reference solution in `.meta/example.go`, which is the gold standard for correctness
2. It is simpler with fewer moving parts (no power variable, no reverse iteration)
3. It follows the same pattern used in the hexadecimal exercise
4. Bit-shifting by 3 for octal is a well-known idiom in Go

### Final Implementation Plan

**File:** `go/exercises/practice/octal/octal.go`

**Changes:**
1. Add `import "fmt"` for error formatting
2. Implement `ParseOctal(octal string) (int64, error)`:
   - Initialize `num` as `int64(0)`
   - Iterate over each rune in `octal` using `range`
   - For each rune, validate it is between '0' and '7'; if not, return `(0, fmt.Errorf("unexpected rune '%c'", digit))`
   - Accumulate: `num = num<<3 + int64(digit-'0')`
   - Return `(num, nil)`

**Note:** Parameter name `octal` and error message `"unexpected rune '%c'"` match the reference solution exactly (per codex review).

**Verification:**
1. Run `go test` in the octal directory — all 5 test cases must pass
2. Run `go test -bench .` — benchmark must run without error
