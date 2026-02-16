# Goal: Implement palindrome-products exercise in Go

## Problem Statement

Implement the `palindrome-products` Exercism exercise in Go. The solution must detect palindrome products in a given range: given a range of numbers [fmin, fmax], find the largest and smallest palindromes which are products of two numbers within that range, along with all factor pairs for each.

## Acceptance Criteria

1. **`Product` struct** is defined with fields:
   - `Product int` — the palindrome value
   - `Factorizations [][2]int` — all factor pairs producing this palindrome

2. **`Products(fmin, fmax int) (Product, Product, error)`** function:
   - Returns (pmin, pmax, err)
   - For valid ranges, finds the smallest and largest palindrome products
   - Returns all factor pairs `[a, b]` where `a <= b` for each palindrome
   - Returns an error prefixed with `"fmin > fmax"` when fmin > fmax
   - Returns an error prefixed with `"no palindromes"` when no palindrome products exist in the range

3. **All 5 test cases pass** (`go test ./...` in the exercise directory):
   - valid limits 1-9
   - valid limits 10-99
   - valid limits 100-999
   - no palindromes (range 4-10)
   - fmin > fmax (10, 4)

4. **`go vet ./...`** passes with no issues.

## Key Constraints

- Package name: `palindrome`
- File: `palindrome_products.go`
- Must match the test's expected API exactly (function signature, struct fields, error message prefixes)
- The `.meta/example.go` reference solution exists but the stub `palindrome_products.go` is empty (just the package declaration)
