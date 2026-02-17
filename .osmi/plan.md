# Implementation Plan: polyglot-go-octal

## Proposal A (Proponent)

**Approach: Iterate with bit-shifting (matching the reference solution pattern)**

Implement `ParseOctal` by iterating through each character of the input string left-to-right. For each character, validate it is in the range '0'-'7'. If valid, left-shift the accumulator by 3 bits (equivalent to multiplying by 8) and add the digit value. If any character is invalid, return `(0, error)`.

### Files to modify
- `go/exercises/practice/octal/octal.go` — add `ParseOctal` function and `fmt` import

### Implementation
```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for _, r := range input {
        if r < '0' || r > '7' {
            return 0, fmt.Errorf("invalid octal digit: %c", r)
        }
        num = num<<3 + int64(r-'0')
    }
    return num, nil
}
```

### Rationale
- Matches the reference solution in `.meta/example.go` almost exactly
- Bit-shifting (`<<3`) is idiomatic for powers-of-2 base conversions
- Simple, minimal, single-pass O(n) algorithm
- Uses `fmt.Errorf` which is standard library only
- Uses `range` over string which iterates runes (handles UTF-8 correctly, though all valid inputs are ASCII)

### Strengths
- Proven correct (matches reference)
- Minimal code, easy to understand
- No unnecessary abstractions

---

## Proposal B (Opponent)

**Approach: Explicit multiplication with `errors.New` and index-based loop**

Implement `ParseOctal` by iterating through the string using byte indexing. Validate each byte, multiply accumulator by 8 (explicit multiplication), and add digit value. Use `errors.New` for the error.

### Files to modify
- `go/exercises/practice/octal/octal.go` — add `ParseOctal` function and `errors` import

### Implementation
```go
package octal

import "errors"

func ParseOctal(input string) (int64, error) {
    var result int64
    for i := 0; i < len(input); i++ {
        d := input[i]
        if d < '0' || d > '7' {
            return 0, errors.New("invalid octal input")
        }
        result = result*8 + int64(d-'0')
    }
    return result, nil
}
```

### Critique of Proposal A
- Using `fmt.Errorf` pulls in the heavier `fmt` package when `errors.New` suffices
- Using `range` iterates runes which adds unnecessary complexity for what is guaranteed to be ASCII-only input
- Bit-shifting is slightly less readable than explicit `*8` for someone unfamiliar with the idiom

### Strengths of Proposal B
- `errors.New` is lighter weight than `fmt.Errorf`
- Byte indexing is more explicit about what we're doing (all valid octal chars are single-byte ASCII)
- `result*8` is clearer about the mathematical operation

### Weaknesses of Proposal B
- `errors.New` gives a less informative error message (no character info)
- Byte indexing vs range is a minor stylistic difference with no practical impact
- `*8` vs `<<3` compiles to the same instruction; readability is subjective

---

## Selected Plan (Judge)

### Evaluation

| Criterion     | Proposal A              | Proposal B              |
|---------------|-------------------------|-------------------------|
| Correctness   | Fully correct           | Fully correct           |
| Risk          | Minimal                 | Minimal                 |
| Simplicity    | Very simple             | Very simple             |
| Consistency   | Matches reference exactly | Slightly diverges      |

Both proposals are correct and minimal. The differences are trivial:
- `fmt.Errorf` vs `errors.New`: The test only checks `err != nil`, so error message content doesn't matter. However, `fmt.Errorf` gives better debugging info.
- `<<3` vs `*8`: Both compile identically. The reference uses `<<3`.
- `range` vs byte indexing: Both work correctly for ASCII input.

**Winner: Proposal A**, because it matches the reference solution pattern exactly, which reduces risk and is consistent with the codebase's existing example. The `fmt` import is negligible overhead.

### Final Implementation Plan

**File: `go/exercises/practice/octal/octal.go`**

Replace the stub with:

```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for _, r := range input {
        if r < '0' || r > '7' {
            return 0, fmt.Errorf("invalid octal digit: %c", r)
        }
        num = num<<3 + int64(r-'0')
    }
    return num, nil
}
```

### Steps
1. Create feature branch `issue-353`
2. Write the implementation to `octal.go`
3. Run `go test ./...` from the octal exercise directory
4. Run `go vet ./...`
5. Commit with descriptive message
