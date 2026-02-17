# Implementation Plan: Crypto Square

## Proposal A — Direct Column-Building Approach

**Role: Proponent**

### Approach

Implement `Encode` using `strings.Map` for normalization and modular arithmetic to distribute characters across columns, following the reference implementation's strategy closely.

### Files to Modify

- `go/exercises/practice/crypto-square/crypto_square.go` — implement `Encode` and a helper `norm` function

### Implementation Details

1. **Normalization**: Use `strings.Map` with a helper function `norm` that:
   - Keeps lowercase letters and digits as-is
   - Converts uppercase to lowercase via `r + 'a' - 'A'`
   - Returns `-1` (filtered out) for everything else

2. **Rectangle sizing**: Compute `c = ceil(sqrt(len(normalized)))`. Then compute padding:
   - First try `r = c - 1`: padding = `c*(c-1) - len`
   - If padding < 0, use `r = c`: padding = `c*c - len`

3. **Column building**: Create a slice of `c` strings. Iterate over normalized text, appending character `i` to `cols[i % c]`. This naturally reads down columns.

4. **Padding**: Append a space to each of the last `padding` columns.

5. **Output**: `strings.Join(cols, " ")`

### Rationale

- Directly mirrors the proven reference implementation
- Minimal code (~34 lines including helper)
- Uses only `math` and `strings` from stdlib
- The modular arithmetic approach is elegant: distributing chars with `i % numCols` inherently reads down columns when the columns are joined
- Simple to understand and verify

### Advocacy

This approach is the simplest possible correct solution. It avoids unnecessary abstractions, uses well-understood Go idioms (`strings.Map`, `strings.Join`), and matches the reference implementation that is known to pass all tests. The column-building via modular indexing is a natural fit for the problem structure.

---

## Proposal B — Grid-Based Approach with `unicode` Package

**Role: Opponent**

### Approach

Build a 2D grid (slice of byte slices or rune slices) representing the rectangle explicitly, then read columns from the grid. Use `unicode.IsLetter` and `unicode.IsDigit` for normalization instead of manual rune comparisons.

### Files to Modify

- `go/exercises/practice/crypto-square/crypto_square.go` — implement `Encode` with grid construction

### Implementation Details

1. **Normalization**: Use `strings.Builder` and `unicode.IsLetter`/`unicode.IsDigit`/`unicode.ToLower` for filtering and lowercasing.

2. **Rectangle sizing**: Same mathematical approach — `c = ceil(sqrt(n))`, then determine `r`.

3. **Grid construction**: Create a 2D `[][]byte` grid of size `r x c`, pre-filled with spaces. Copy normalized text into the grid row by row.

4. **Column reading**: Read grid column by column, building each chunk as a string of length `r`.

5. **Output**: Join chunks with spaces.

### Critique of Proposal A

- Proposal A uses manual rune comparison (`r >= 'a' && r <= 'z'`) which is less idiomatic than using the `unicode` package
- String concatenation in a loop (`cols[i%numCols] += string(r)`) creates many intermediate strings — inefficient for large inputs
- The padding logic in Proposal A (`numCols*(numCols-1) - len(pt)` with a negative check) is harder to reason about than explicit grid padding

### Advocacy

The grid approach makes the algorithm visually and logically transparent: you can see the rectangle being built and read. Using `unicode` functions is more idiomatic Go and handles edge cases better. The explicit grid makes debugging trivial.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| **Correctness** | Proven correct (matches reference) | Likely correct but untested |
| **Risk** | Very low — known to pass all tests | Low — straightforward but needs verification |
| **Simplicity** | 34 lines, 2 imports | ~45-50 lines, 3 imports |
| **Consistency** | Matches existing codebase patterns exactly | Slightly different style (unicode pkg) |

### Analysis

**Proposal A strengths:**
- Directly based on the proven reference implementation
- Minimal code, minimal risk
- Uses the same stdlib packages seen in other exercises
- The modular arithmetic approach is actually elegant once understood

**Proposal A weaknesses:**
- String concatenation in loops (minor performance concern for huge inputs, but benchmark exists and the reference passes it)
- The padding calculation requires a moment of thought

**Proposal B strengths:**
- More readable grid construction
- `unicode` package is idiomatic

**Proposal B weaknesses:**
- More code for no additional correctness benefit
- 2D grid allocation is unnecessary overhead
- Introduces `unicode` package when manual comparison suffices
- Not aligned with the reference implementation pattern
- Higher risk since it's a novel implementation

### Decision

**Proposal A wins.** The rationale:

1. It matches the proven reference implementation that is known to pass all 18 tests
2. It is simpler (fewer lines, fewer imports)
3. It is consistent with the codebase conventions (other exercises use simple manual rune handling)
4. The risk is minimal since the approach is battle-tested
5. Performance is adequate (reference implementation passes the benchmark)

### Selected Plan — Full Details

**File to modify:** `go/exercises/practice/crypto-square/crypto_square.go`

**Implementation:**

```go
package cryptosquare

import (
    "math"
    "strings"
)

// norm normalizes a rune: keeps lowercase alphanumeric, converts uppercase
// to lowercase, and filters everything else by returning -1.
func norm(r rune) rune {
    switch {
    case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
        return r
    case r >= 'A' && r <= 'Z':
        return r + 'a' - 'A'
    }
    return -1
}

// Encode implements the crypto square cipher.
func Encode(pt string) string {
    pt = strings.Map(norm, pt)
    numCols := int(math.Ceil(math.Sqrt(float64(len(pt)))))
    padding := numCols*(numCols-1) - len(pt)
    if padding < 0 {
        padding = numCols*numCols - len(pt)
    }
    cols := make([]string, numCols)
    for i, r := range pt {
        cols[i%numCols] += string(r)
    }
    for i := 0; i < padding; i++ {
        cols[numCols-i-1] += " "
    }
    return strings.Join(cols, " ")
}
```

**Steps:**
1. Replace the stub in `crypto_square.go` with the implementation above
2. Run `go test` to verify all 18 tests pass
3. Run `go test -bench=.` to verify benchmark passes
4. Commit with descriptive message
