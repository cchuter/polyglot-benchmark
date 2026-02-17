# Goal: Implement palindrome-products exercise in Go

## Problem Statement

Implement the `palindrome-products` Exercism exercise for Go as part of the polyglot benchmark suite (issue #313). The solution file `go/exercises/practice/palindrome-products/palindrome_products.go` is currently a stub containing only the package declaration. It must be completed so all tests in `palindrome_products_test.go` pass.

## What needs to be built

A Go function `Products(fmin, fmax int) (pmin, pmax Product, err error)` and supporting types/helpers in the `palindrome` package that:

1. Defines a `Product` struct with fields:
   - `Product int` — the palindromic product value
   - `Factorizations [][2]int` — all factor pairs within the range

2. Given a range `[fmin, fmax]`, finds all products of two numbers within that range that are palindromes.

3. Returns the smallest palindrome product (`pmin`) and largest palindrome product (`pmax`), each with all their factor pairs.

4. Returns an error when:
   - `fmin > fmax` — error message must start with `"fmin > fmax"`
   - No palindromes exist in the range — error message must start with `"no palindromes"`

## Acceptance Criteria

- [ ] `go test` passes all 5 test cases in `palindrome_products_test.go`
- [ ] `go vet` reports no issues
- [ ] Solution is in `go/exercises/practice/palindrome-products/palindrome_products.go`
- [ ] Package name is `palindrome`
- [ ] No modifications to test file or go.mod
- [ ] Code is clean, idiomatic Go with no unnecessary comments or documentation

## Key Constraints

- Must use Go 1.18 (as specified in go.mod)
- Must not modify `palindrome_products_test.go` or `go.mod`
- Factor pairs must be sorted with smaller factor first: `[2]int{a, b}` where `a <= b`
- Multiple factorizations for the same product must all be included
- Reference implementation exists at `.meta/example.go` for validation
