# Implementation Plan: Crypto Square

## File to Modify

- `go/exercises/practice/crypto-square/crypto_square.go`

## Approach

Implement a single `Encode` function in the `cryptosquare` package. The function uses only Go standard library packages.

### Algorithm Steps

1. **Normalize input**: Iterate over the input string, keeping only alphanumeric characters (`unicode.IsLetter` or `unicode.IsDigit`), converting letters to lowercase (`unicode.ToLower`). Build the result with `strings.Builder`.

2. **Handle edge case**: If the normalized string is empty, return `""`.

3. **Compute dimensions**:
   - `c = int(math.Ceil(math.Sqrt(float64(len(normalized)))))`
   - `r = int(math.Ceil(float64(len(normalized)) / float64(c)))`
   - This gives the smallest rectangle where `c >= r` and `c - r <= 1`.

4. **Read down columns**: For each column index `col` (0 to c-1), read characters at positions `col`, `col+c`, `col+2c`, ... up to `r` characters. If a position exceeds the normalized string length, use a space.

5. **Join chunks**: Join the `c` chunks with a single space separator.

## Imports Required

- `math`
- `strings`
- `unicode`

## Code Structure

```go
package cryptosquare

import (
    "math"
    "strings"
    "unicode"
)

func Encode(pt string) string {
    // 1. Normalize
    // 2. Handle empty
    // 3. Compute c, r
    // 4. Build column chunks
    // 5. Join and return
}
```

## Rationale

- Single function, no unnecessary abstractions.
- Uses `strings.Builder` for efficient string building.
- Standard `math.Ceil(math.Sqrt(...))` for dimension calculation — clean and correct.
- Padding with spaces happens naturally when column index exceeds normalized string length.
