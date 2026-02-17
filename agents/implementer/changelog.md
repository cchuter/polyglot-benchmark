# Changelog — implementer

## Implemented Encode function in crypto_square.go

- Added `norm` helper function that normalizes runes: keeps lowercase letters and digits, converts uppercase to lowercase, filters everything else
- Added `Encode` function implementing the crypto square cipher:
  - Normalizes input using `strings.Map` with the `norm` helper
  - Computes rectangle dimensions using `ceil(sqrt(len))` for columns
  - Distributes characters across columns using modular arithmetic (`i % numCols`)
  - Pads trailing columns with spaces as needed
  - Joins columns with spaces for the final output
- Uses only `math` and `strings` from stdlib
