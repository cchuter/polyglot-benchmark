# Goal: Palindrome Products (Go)

## Problem Statement

Implement the `palindrome-products` Exercism exercise in Go. The stub file `palindrome_products.go` currently only declares the package — it needs a `Product` type and a `Products` function that detect palindrome products in a given range.

A palindromic number reads the same forwards and backwards (e.g., 121). Given a range `[fmin, fmax]`, find the smallest and largest palindromic products of two numbers within that range, along with all factor pairs that produce them.

## Acceptance Criteria

1. **`Product` struct** is defined with fields:
   - `Product int` — the palindromic product value
   - `Factorizations [][2]int` — all factor pairs `[a, b]` where `a <= b`

2. **`Products(fmin, fmax int) (pmin, pmax Product, err error)`** function:
   - Returns the smallest palindrome product as `pmin` and largest as `pmax`
   - Returns all factor pairs for each palindrome product
   - Returns an error prefixed with `"fmin > fmax"` when `fmin > fmax`
   - Returns an error prefixed with `"no palindromes"` when no palindrome products exist in the range

3. **All 5 test cases pass** (`go test ./...`):
   - Range [1, 9]: pmax = 9 with factors {1,9} and {3,3}
   - Range [10, 99]: pmin = 121 with factors {11,11}; pmax = 9009 with factors {91,99}
   - Range [100, 999]: pmin = 10201 with factors {101,101}; pmax = 906609 with factors {913,993}
   - Range [4, 10]: error "no palindromes..."
   - Range [10, 4]: error "fmin > fmax..."

4. **`go vet ./...` passes** with no issues.

## Key Constraints

- Solution must be in package `palindrome`
- Must match the exact type signatures expected by the test file
- Factor pairs must be sorted with the smaller factor first: `[a, b]` where `a <= b`
- Module is `go 1.18`
