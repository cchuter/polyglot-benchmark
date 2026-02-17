# Goal: Implement palindrome-products Go Exercise

## Problem Statement

Implement the `palindrome-products` exercise in Go. The solution must detect palindrome products in a given range of factor values.

A palindromic number is a number that remains the same when its digits are reversed (e.g., `121` is palindromic, `112` is not).

Given a range `[fmin, fmax]`, find the largest and smallest palindromes which are products of two numbers within that range, along with all factor pairs that produce each palindrome.

## Required API

The solution must implement in `palindrome_products.go`:

1. **`Product` struct** with fields:
   - `Product int` — the palindrome value
   - `Factorizations [][2]int` — all factor pairs `[a, b]` where `a <= b`

2. **`Products(fmin, fmax int) (pmin, pmax Product, err error)`** function:
   - Returns the smallest palindrome product (`pmin`) and largest palindrome product (`pmax`)
   - Returns an error prefixed with `"fmin > fmax"` when `fmin > fmax`
   - Returns an error prefixed with `"no palindromes"` when no palindrome products exist in the range

## Acceptance Criteria

1. **All 5 test cases pass** in `palindrome_products_test.go`:
   - Range `[1, 9]`: largest palindrome is `9` with factors `{(1,9), (3,3)}`
   - Range `[10, 99]`: smallest is `121` with factors `{(11,11)}`, largest is `9009` with factors `{(91,99)}`
   - Range `[100, 999]`: smallest is `10201` with factors `{(101,101)}`, largest is `906609` with factors `{(913,993)}`
   - Range `[4, 10]`: no palindromes — error with prefix `"no palindromes"`
   - Range `[10, 4]`: invalid range — error with prefix `"fmin > fmax"`

2. **Code compiles** with `go build ./...`

3. **All tests pass** with `go test ./...` (exit code 0)

4. **Benchmark runs** without error: `go test -bench=. -short`

## Key Constraints

- Package name must be `palindrome` (matches existing stub and test file)
- Module is `palindrome` with Go 1.18
- The `Product` struct and `Products` function must be exported
- Factor pairs must have the smaller factor first: `[a, b]` where `a <= b`
- Multiple factorizations for the same product must all be included
- Error messages must start with the expected prefix strings
