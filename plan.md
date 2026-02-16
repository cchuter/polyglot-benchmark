# Implementation Plan: polyglot-go-octal

## Overview

Implement the `ParseOctal` function in `go/exercises/practice/octal/octal.go` that converts an octal string to its decimal int64 equivalent.

## File to Modify

- `go/exercises/practice/octal/octal.go` — the only file that needs changes

## Implementation Approach

### Algorithm

1. Initialize a result accumulator `num` as `int64(0)`
2. Iterate over each character (rune) in the input string
3. For each character:
   - Validate it is a digit between '0' and '7'
   - If invalid, return `(0, error)` immediately
   - If valid, shift the accumulator left by 3 bits (equivalent to multiplying by 8) and add the digit value
4. Return `(num, nil)` on success

### Implementation Details

```go
package octal

import "fmt"

func ParseOctal(octal string) (int64, error) {
    num := int64(0)
    for _, digit := range octal {
        if digit < '0' || digit > '7' {
            return 0, fmt.Errorf("unexpected rune '%c'", digit)
        }
        num = num<<3 + int64(digit-'0')
    }
    return num, nil
}
```

### Design Rationale

- **Bit shifting (`<<3`)** instead of multiplication by 8 — idiomatic for base-power-of-2 conversions and matches the reference solution in `.meta/example.go`
- **`fmt.Errorf`** for error creation — lightweight, no need for custom error types given the test only checks `err != nil`
- **Range over string** — iterates runes, which is correct for ASCII digit characters
- **Early return on invalid input** — returns 0 immediately as specified

## Verification

Run `go test ./...` from the `go/exercises/practice/octal/` directory. All 5 test cases plus the benchmark should pass.
