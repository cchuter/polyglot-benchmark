# Scope: Crypto Square (Go)

## In Scope

- Implement the `Encode` function in `go/exercises/practice/crypto-square/crypto_square.go`
- Handle normalization (remove non-alphanumeric, downcase)
- Handle rectangle sizing (c >= r, c - r <= 1, r * c >= len)
- Handle column-wise reading and space-separated output with padding
- All edge cases: empty string, single char, strings needing padding

## Out of Scope

- Modifying `crypto_square_test.go`
- Modifying `go.mod`
- Implementing a decode function
- Changes to any other exercise or language directory
- Benchmark optimization beyond passing tests

## Dependencies

- Go standard library only (`math`, `strings`, `unicode`)
- No external packages required
