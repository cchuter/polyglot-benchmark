# Implementation Plan: palindrome-products

## File to Modify

- `go/exercises/practice/palindrome-products/palindrome_products.go`

No other files need to be created or modified.

## Approach

Implement the solution directly in `palindrome_products.go`, following the reference solution's approach from `.meta/example.go`. The implementation consists of:

### 1. `isPal` helper function
- Converts an integer to a string
- Checks if the string reads the same forwards and backwards
- Uses two-pointer approach (front and back converging)

### 2. `Product` struct
- `Product int` — the palindrome value
- `Factorizations [][2]int` — sorted factor pairs

### 3. `Products(fmin, fmax int) (pmin, pmax Product, err error)` function

**Algorithm:**
1. Validate: if `fmin > fmax`, return error with prefix `"fmin > fmax"`
2. Iterate `x` from `fmin` to `fmax`, `y` from `x` to `fmax` (ensures `x <= y`)
3. Compute `p = x * y`; skip if not palindrome
4. For each palindrome found, update `pmin` and `pmax`:
   - If no palindrome found yet (`Factorizations == nil`) or `p` is smaller/larger, replace with new Product
   - If `p` equals current min/max, append factor pair
5. After loop, if no palindromes found (both still nil factorizations), return error with prefix `"no palindromes"`

### Ordering of Changes

1. Write the complete implementation in `palindrome_products.go`
2. Run `go test` to verify all 5 tests pass
3. Run `go vet` to verify no issues
4. Commit

## Rationale

- The reference solution in `.meta/example.go` is clean and correct; our implementation follows the same algorithmic approach
- Using `strconv.Itoa` for palindrome check is simple and readable
- The nested loop with `y >= x` avoids duplicate factor pairs and ensures `[a, b]` where `a <= b`
- Using a closure (`compare`) keeps the min/max update logic DRY
