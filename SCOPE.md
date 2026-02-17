# Scope: palindrome-products (issue #313)

## In Scope

- Implementing the `Products` function in `palindrome_products.go`
- Defining the `Product` struct with `Product int` and `Factorizations [][2]int` fields
- Implementing a palindrome check helper function
- Error handling for invalid ranges and no-palindrome cases
- Passing all 5 test cases in the test file

## Out of Scope

- Bonus test cases (negative numbers) — these are commented out in the test file
- Modifying test files or go.mod
- Performance optimization beyond what's needed to pass tests
- Adding documentation, comments, or README files
- Changes to any files outside `go/exercises/practice/palindrome-products/palindrome_products.go`

## Dependencies

- Go 1.18+ toolchain
- Standard library only (`fmt`, `strconv`)
- No external packages
