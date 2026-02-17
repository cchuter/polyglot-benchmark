# Goal: Implement Crypto Square (Go)

## Problem Statement

Implement the classic "square code" cipher in Go. Given an English text, produce the encoded version by:

1. **Normalize** the input: remove all non-alphanumeric characters and downcase.
2. **Arrange** the normalized text into a rectangle with `r` rows and `c` columns, where `c` is the smallest integer such that `r * c >= length`, `c >= r`, and `c - r <= 1`.
3. **Encode** by reading columns top-to-bottom, left-to-right, producing `c` chunks of `r` characters each, separated by spaces. Pad the last `n` chunks with a trailing space if the text is `n` characters short of filling the full `r * c` rectangle.

## Acceptance Criteria

1. The function `Encode(string) string` is exported from package `cryptosquare` in `crypto_square.go`.
2. All 19 test cases in `crypto_square_test.go` pass.
3. `go vet ./...` reports no issues.
4. Empty input returns empty string.
5. Single character input returns that character.
6. Padding is applied correctly: last `n` chunks get a trailing space when the text doesn't fill the rectangle.

## Key Constraints

- Solution file: `go/exercises/practice/crypto-square/crypto_square.go`
- Package name: `cryptosquare`
- Must export `Encode(string) string`
- No modifications to test file or go.mod
- Go 1.18 compatibility
