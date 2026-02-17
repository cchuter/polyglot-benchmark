# Goal: Implement Crypto Square Exercise in Go

## Problem Statement

Implement the classic "square code" cipher method for composing secret messages in Go. The solution must provide an `Encode` function in the `cryptosquare` package that:

1. **Normalizes** input text by removing spaces and punctuation, and converting to lowercase
2. **Arranges** the normalized characters into a rectangle as square as possible
3. **Encodes** by reading down columns left-to-right
4. **Outputs** space-separated chunks with trailing space padding where needed

## Acceptance Criteria

1. **Function signature**: `func Encode(pt string) string` in package `cryptosquare`
2. **All 18 test cases pass** in `crypto_square_test.go`, including:
   - Empty string returns empty string
   - Single character returns that character
   - Simple numeric inputs (e.g., "12" → "1 2", "123456789" → "147 258 369")
   - Text with spaces removed (e.g., "Cedric the Entertainer" → correct encoding)
   - Text with punctuation removed
   - Complex sentences with mixed case, punctuation, and spaces
   - Proper trailing-space padding for non-perfect rectangles
3. **Benchmark passes**: `BenchmarkEncode` runs without error
4. **Build succeeds**: `go build` completes with no errors
5. **Rectangle sizing**: For normalized text of length `n`, find smallest `c` such that:
   - `r * c >= n`
   - `c >= r`
   - `c - r <= 1`

## Key Constraints

- Package name: `cryptosquare`
- Module name: `cryptosquare` (go 1.18)
- Solution file: `crypto_square.go`
- Must not modify test files or metadata files
- Only alphanumeric characters survive normalization (a-z, 0-9)
