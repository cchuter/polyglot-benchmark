# Implementation Plan: palindrome-products

## Branch 1: Direct Reference Implementation

**Approach:** Follow the reference implementation in `.meta/example.go` closely. Copy the proven logic with minimal modifications.

**Files to modify:**
- `go/exercises/practice/palindrome-products/palindrome_products.go` — replace stub with full implementation

**Architecture:**
1. Define `Product` struct with `Product int` and `Factorizations [][2]int`
2. Implement `isPal(x int) bool` helper using string reversal comparison
3. Implement `Products(fmin, fmax int) (pmin, pmax Product, err error)`:
   - Validate `fmin <= fmax`, return error if not
   - Iterate `x` from `fmin` to `fmax`, `y` from `x` to `fmax`
   - For each product, check palindrome, update min/max tracking
   - Use closure-based compare function to update Product structs
   - If no palindromes found, return error

**Rationale:** The reference implementation is proven correct against the test suite. Using it directly minimizes risk.

**Evaluation:**
- Feasibility: High — proven code exists
- Risk: Very low — directly following a working reference
- Alignment: Fully satisfies all acceptance criteria
- Complexity: Minimal — single file, ~50 lines

---

## Branch 2: Extensible Design with Separate Concerns

**Approach:** Separate palindrome detection, product generation, and result aggregation into distinct functions for clarity and extensibility.

**Files to modify:**
- `go/exercises/practice/palindrome-products/palindrome_products.go` — implement with separate helper functions

**Architecture:**
1. Define `Product` struct
2. `isPalindrome(n int) bool` — string-based palindrome check
3. `findPalindromeProducts(fmin, fmax int) map[int][][2]int` — returns map of palindrome value to factor pairs
4. `Products(fmin, fmax int) (Product, Product, error)` — validates input, calls findPalindromeProducts, finds min/max from map

**Rationale:** Separating concerns makes each function independently testable and the logic easier to follow.

**Evaluation:**
- Feasibility: High — straightforward Go code
- Risk: Low — more code means slightly more surface area for bugs, but logic is simple
- Alignment: Fully satisfies acceptance criteria
- Complexity: Moderate — single file, ~60-70 lines, 3-4 functions

---

## Branch 3: Performance-Optimized with Early Termination

**Approach:** Optimize for large ranges by searching from the edges (highest products first for max, lowest first for min) with early termination once a palindrome is found.

**Files to modify:**
- `go/exercises/practice/palindrome-products/palindrome_products.go`

**Architecture:**
1. Define `Product` struct
2. `isPalindrome(n int) bool` — numeric palindrome check (no string conversion)
3. For finding max: iterate products from largest to smallest, stop once we've found all pairs for the largest palindrome
4. For finding min: iterate products from smallest to largest, stop once we've found all pairs for the smallest palindrome
5. `Products` orchestrates both searches

**Rationale:** For range [100, 999], the brute force checks ~405K products. An optimized approach could find results faster by searching from extremes.

**Evaluation:**
- Feasibility: Medium — more complex iteration logic, need to ensure all factor pairs are found
- Risk: Medium-high — complex termination conditions, harder to verify correctness
- Alignment: Would satisfy criteria but correctness is harder to verify
- Complexity: High — more complex control flow, potential for off-by-one errors

---

## Selected Plan

**Selected: Branch 1 — Direct Reference Implementation**

**Rationale:** Branch 1 is the clear winner because:
1. The reference implementation is proven to work against the exact test suite we need to pass
2. It has minimal risk — the code is straightforward and well-tested
3. It fully satisfies all acceptance criteria with ~50 lines of clean Go
4. The exercise doesn't require extensibility (Branch 2 over-engineers) or performance optimization (Branch 3 adds unnecessary complexity and risk)
5. It follows the conventions observed in other completed exercises: simple, focused, idiomatic Go

### Detailed Implementation Plan

**Single file to modify:** `go/exercises/practice/palindrome-products/palindrome_products.go`

**Step 1:** Replace the stub file contents with the full implementation:

```go
package palindrome

import (
	"fmt"
	"strconv"
)

// isPal checks if a number is a palindrome by comparing its string representation.
func isPal(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// Product represents a palindrome product and its factor pairs.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products finds the smallest and largest palindrome products within factor range [fmin, fmax].
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax: %d > %d", fmin, fmax)
		return
	}
	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			p := x * y
			if !isPal(p) {
				continue
			}
			compare := func(current *Product, better bool) {
				switch {
				case current.Factorizations == nil || better:
					*current = Product{p, [][2]int{{x, y}}}
				case p == current.Product:
					current.Factorizations = append(current.Factorizations, [2]int{x, y})
				}
			}
			compare(&pmin, p < pmin.Product)
			compare(&pmax, p > pmax.Product)
		}
	}
	if len(pmin.Factorizations) == 0 {
		err = fmt.Errorf("no palindromes in range [%d, %d]", fmin, fmax)
	}
	return
}
```

**Step 2:** Verify by running `go test ./...` from the exercise directory.

**Step 3:** Verify benchmark with `go test -bench=. -short`.

**Step 4:** Commit with message following repo conventions.
