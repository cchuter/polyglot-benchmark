# Changes

## Implement Encode function for crypto-square exercise

- **File**: `go/exercises/practice/crypto-square/crypto_square.go`
- **What**: Implemented the `Encode` function that performs crypto square encoding.
- **Algorithm**:
  1. Normalize input by removing non-alphanumeric characters and lowercasing.
  2. Return empty string for empty normalized input.
  3. Compute rectangle dimensions: c = ceil(sqrt(len)), r = ceil(len/c).
  4. Read characters column-by-column, padding with spaces for incomplete rows.
  5. Join column chunks with single spaces.
- **Imports added**: `math`, `strings`, `unicode`
