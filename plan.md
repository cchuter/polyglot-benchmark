# Implementation Plan: palindrome-products

## Proposal A (Proponent)

### Approach: Direct brute-force with nested loops (mirror the reference solution)

**Files to modify:**
- `go/exercises/practice/palindrome-products/palindrome_products.go`

**Architecture:**
1. Define `Product` struct with `Product int` and `Factorizations [][2]int`
2. Implement `isPal(x int) bool` helper using string reversal comparison
3. Implement `Products(fmin, fmax int) (pmin, pmax Product, err error)`:
   - Validate `fmin <= fmax`, return error if not
   - Iterate `x` from `fmin` to `fmax`, `y` from `x` to `fmax` (avoiding duplicate pairs)
   - For each product `x*y`, check if palindrome
   - Track min/max palindrome products and their factorizations
   - If no palindromes found, return error

**Rationale:**
- This mirrors the reference implementation in `.meta/example.go` almost exactly
- It's proven correct by the existing test suite
- Simple O(n^2) nested loop is sufficient for the test ranges (up to 999)
- Uses `strconv.Itoa` for palindrome check — clear and idiomatic
- The inner closure pattern (`compare`) keeps the min/max tracking DRY

**Ordering:**
1. Write the `Product` struct
2. Write `isPal` helper
3. Write `Products` function with validation, iteration, and palindrome tracking

### Why this is best:
- Matches the established reference solution pattern exactly
- Minimal complexity — easy to verify correctness
- Handles all edge cases naturally (no palindromes, invalid ranges)
- Factor pairs are naturally ordered since `y >= x`

---

## Proposal B (Opponent)

### Approach: Generate palindromes first, then find factors

**Files to modify:**
- `go/exercises/practice/palindrome-products/palindrome_products.go`

**Architecture:**
1. Define `Product` struct same as Proposal A
2. Compute the product range: `[fmin*fmin, fmax*fmax]`
3. Generate all palindromes in that numeric range
4. For each palindrome, find all factor pairs `(a, b)` where `fmin <= a <= b <= fmax` and `a*b == palindrome`
5. Track smallest and largest palindromes that have valid factors

**Rationale:**
- For large ranges, palindromes are sparse, so this could be faster
- Separates concerns: palindrome generation vs. factorization

**Critique of Proposal A:**
- Proposal A checks every single product pair even when most aren't palindromes
- For very large ranges, Proposal A would be slow

### Why this is better:
- Could be asymptotically faster for large ranges
- Cleaner separation of concerns

### Weaknesses acknowledged:
- Significantly more complex to implement
- Generating palindromes in a range and finding their factors is non-trivial
- The test cases only go up to 999, so performance gains are irrelevant
- More code means more surface area for bugs
- Doesn't match the reference solution pattern

---

## Selected Plan (Judge)

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | High — mirrors proven reference | Medium — more complex logic to get right |
| Risk | Low — straightforward implementation | Higher — palindrome generation + factoring is tricky |
| Simplicity | Simple — ~50 lines | Complex — ~80+ lines with multiple phases |
| Consistency | Matches reference and codebase conventions | Over-engineered for the use case |

### Decision: Proposal A wins

**Rationale:** The test suite only goes up to range [100, 999]. Performance optimization is unnecessary. Proposal A is simpler, proven correct by the reference implementation, and follows the established codebase pattern. Proposal B introduces unnecessary complexity with no practical benefit.

### Final Plan

**File:** `go/exercises/practice/palindrome-products/palindrome_products.go`

**Implementation (in order):**

1. **Package declaration and imports:**
   ```go
   package palindrome

   import (
       "fmt"
       "strconv"
   )
   ```

2. **`isPal` helper function:**
   - Convert int to string via `strconv.Itoa`
   - Compare characters from both ends moving inward
   - Return true if all pairs match

3. **`Product` struct:**
   ```go
   type Product struct {
       Product        int
       Factorizations [][2]int
   }
   ```

4. **`Products` function:**
   - Check `fmin > fmax`, return error with prefix `"fmin > fmax"`
   - Nested loop: outer `x` from `fmin` to `fmax`, inner `y` from `x` to `fmax`
   - For each palindromic product, update `pmin`/`pmax` using a comparison closure:
     - If first palindrome found or product is better (smaller/larger), replace
     - If product equals current, append factorization
   - After loops, if no palindromes found (`Factorizations` is nil), return error with prefix `"no palindromes"`
   - Return `pmin`, `pmax`, `nil`

**Verification:** Run `go test` in the exercise directory. All 5 test cases must pass.
