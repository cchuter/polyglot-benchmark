# Goal: Implement Crypto Square in Go

## Problem Statement

Implement the `Encode` function in `go/exercises/practice/crypto-square/crypto_square.go` that performs the classic square code cipher encoding.

The function signature is: `func Encode(pt string) string`

### Algorithm

1. **Normalize** the input: remove all non-alphanumeric characters, downcase all letters.
2. **Determine rectangle dimensions**: find the smallest `c` (columns) such that:
   - `r * c >= len(normalized)` where `r` is rows
   - `c >= r`
   - `c - r <= 1`
   (i.e., `c = ceil(sqrt(len(normalized)))`, `r = ceil(len / c)`)
3. **Fill the rectangle** row by row with the normalized text (pad with spaces if needed).
4. **Read columns** left-to-right to produce cipher chunks of length `r`.
5. **Output** the `c` chunks separated by spaces.

## Acceptance Criteria

1. All 18 test cases in `crypto_square_test.go` pass.
2. The `Encode` function is implemented in `crypto_square.go` in package `cryptosquare`.
3. `go test ./...` succeeds with no failures.
4. `go vet ./...` reports no issues.
5. The solution handles edge cases: empty string, single character, strings with only special characters.

## Key Constraints

- The solution file is `crypto_square.go` in package `cryptosquare`.
- Must not modify test files.
- Go 1.18 module.
