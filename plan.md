# Implementation Plan: Palindrome Products

## Branch 1: Direct Iteration (Simplicity-First)

**Approach**: Iterate over all pairs `(x, y)` where `fmin <= x <= y <= fmax`, compute the product, check if it's a palindrome using string reversal, and track min/max palindromes with their factor pairs.

**Files to modify**: `palindrome_products.go` only.

**Implementation**:
1. Define `Product` struct with `Product int` and `Factorizations [][2]int`
2. Define helper `isPal(x int) bool` using `strconv.Itoa` and string comparison
3. Define `Products(fmin, fmax int) (pmin, pmax Product, err error)`:
   - Validate `fmin <= fmax`, return error if not
   - Loop `x` from `fmin` to `fmax`, `y` from `x` to `fmax`
   - For each palindromic product `p = x * y`:
     - If first palindrome found, or `p < pmin.Product`: reset pmin
     - If `p == pmin.Product`: append factor pair
     - Same logic for pmax with `>`
   - If no palindromes found, return error
4. Use a `found` boolean or check `len(pmin.Factorizations) == 0` to detect "no palindromes"

**Evaluation**:
- Feasibility: High — mirrors the reference solution exactly
- Risk: Very low — straightforward nested loop
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~50 lines in a single file

## Branch 2: Map-Based Collection (Extensibility-First)

**Approach**: First collect all palindromic products into a map keyed by product value, storing factor pairs as values. Then find min/max keys from the map.

**Files to modify**: `palindrome_products.go` only.

**Implementation**:
1. Define `Product` struct as above
2. Define `isPal` helper as above
3. Define `Products` function:
   - Validate fmin <= fmax
   - Create `palins := map[int][][2]int{}`
   - Loop pairs, for each palindromic product, append `[2]int{x, y}` to `palins[p]`
   - Find min and max keys from the map
   - Construct Product values from those entries
   - Return error if map is empty

**Evaluation**:
- Feasibility: High — simple map usage
- Risk: Low, but slightly more memory usage from the map
- Alignment: Fully satisfies all criteria
- Complexity: ~60 lines, slightly more code than Branch 1

## Branch 3: Numeric Palindrome Check (Performance-First)

**Approach**: Same iteration strategy as Branch 1 but use a purely numeric palindrome check (reverse digits mathematically, no string conversion) for better performance on large ranges.

**Files to modify**: `palindrome_products.go` only.

**Implementation**:
1. Define `Product` struct as above
2. Define `isPal(x int) bool` using numeric digit reversal:
   ```go
   func isPal(n int) bool {
       if n < 0 { return false }
       reversed, original := 0, n
       for n > 0 {
           reversed = reversed*10 + n%10
           n /= 10
       }
       return original == reversed
   }
   ```
3. `Products` function same as Branch 1

**Evaluation**:
- Feasibility: High
- Risk: Low — numeric reversal is well-known
- Alignment: Fully satisfies all criteria
- Complexity: ~50 lines, same as Branch 1 but avoids string allocation

## Selected Plan

**Selected: Branch 1 (Direct Iteration)** with string-based palindrome check.

**Rationale**: Branch 1 is the simplest and most closely matches the reference solution from `.meta/example.go`. After codex review, the string-based palindrome check is preferred over the numeric approach from Branch 3 — it matches the reference, is equally readable, and avoids an unjustified deviation. Branch 2's map approach adds unnecessary complexity.

### Detailed Implementation

**File**: `go/exercises/practice/palindrome-products/palindrome_products.go`

```go
package palindrome

import (
	"fmt"
	"strconv"
)

func isPal(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// Product holds a palindromic product value and all its factor pairs.
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products finds the smallest and largest palindromic products of two factors
// in the range [fmin, fmax].
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

### Order of Changes
1. Replace the stub `palindrome_products.go` with the full implementation
2. Run `go test ./...` to verify all 5 tests pass
3. Run `go vet ./...` to check for issues
