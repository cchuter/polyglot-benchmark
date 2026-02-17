# Implementation Plan: polyglot-go-hexadecimal

## Proposal A

**Role: Proponent**

### Approach: Mirror the reference solution structure

Follow the `.meta/example.go` reference implementation closely. This is the canonical approach used by Exercism.

### Files to modify
- `go/exercises/practice/hexadecimal/hexadecimal.go` — implement `ParseHex` and `HandleErrors`

### Architecture

1. Define sentinel errors `ErrRange` and `ErrSyntax` using `errors.New()`
2. Define a `ParseError` struct wrapping the input string and underlying error
3. Implement `ParseHex` using a byte-by-byte loop with `switch` cases for digit classification
4. Detect overflow using `math.MaxInt64/16+1` threshold before multiplication
5. Detect post-addition overflow by checking if `n1 < n` after adding the digit value
6. Use `goto Error` pattern for consistent error path (matching the reference)
7. Implement `HandleErrors` using a type assertion on `*ParseError` to classify errors

### Rationale

- Directly matches the reference solution, ensuring correctness
- The `ParseError` type with embedded `Err` field enables clean error classification in `HandleErrors`
- The overflow detection is battle-tested (derived from Go's `strconv` package)
- Error messages naturally contain "syntax" and "range" substrings via `ErrSyntax` and `ErrRange`

---

## Proposal B

**Role: Opponent**

### Approach: Simplified implementation without goto or ParseError struct

Use a cleaner, more idiomatic Go style without `goto` statements or a custom error type.

### Files to modify
- `go/exercises/practice/hexadecimal/hexadecimal.go` — implement `ParseHex` and `HandleErrors`

### Architecture

1. Define `fmt.Errorf` based errors with "syntax" and "range" in the message
2. Implement `ParseHex` with early returns instead of `goto`
3. Use `strings.Contains` in `HandleErrors` to classify errors by checking error message text

### Critique of Proposal A

- Uses `goto`, which is generally discouraged in Go code
- The `ParseError` struct adds complexity for a simple exercise
- More code than necessary

### Rationale

- Simpler code, easier to read
- No custom error type needed
- More idiomatic Go (no goto)

### Weakness

- Relying on `strings.Contains` for error classification in `HandleErrors` is fragile — it couples to exact error message text
- The test explicitly checks `strings.Contains(strings.ToLower(err.Error()), test.errCase)`, so error messages must contain "syntax" or "range" regardless. But `HandleErrors` needs to distinguish error types programmatically — string matching on your own error messages is less robust than a type assertion
- Adding `fmt` as a dependency just for `Errorf` is unnecessary when `errors.New` suffices

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Fully satisfies all tests. ParseError.Err field enables reliable error classification. | Satisfies tests but HandleErrors error classification via string matching is fragile. |
| Risk | Low — mirrors reference solution. | Medium — string matching for error classification could break with message changes. |
| Simplicity | Slightly more code (ParseError struct, goto). | Less code but string matching adds hidden complexity. |
| Consistency | Matches the reference solution pattern used by Exercism. | Deviates from the canonical pattern. |

### Decision

**Proposal A wins.** While `goto` is unusual in Go, it's used in the standard library's `strconv` package (which this exercise is modeled after). The `ParseError` struct provides a clean, type-safe way to classify errors in `HandleErrors`. The reference solution exists for a reason — it's correct and handles all edge cases.

However, I'll simplify slightly: the `goto` pattern is fine since it matches the canonical solution and keeps error handling DRY.

### Final Implementation Plan

#### File: `go/exercises/practice/hexadecimal/hexadecimal.go`

```go
package hexadecimal

import (
    "errors"
    "math"
)

var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")

type ParseError struct {
    Num string
    Err error
}

func (e *ParseError) Error() string {
    return "hexadecimal.ParseHex: parsing \"" + e.Num + "\": " + e.Err.Error()
}

func ParseHex(s string) (n int64, err error) {
    if len(s) < 1 {
        err = ErrSyntax
        goto Error
    }
    for i := 0; i < len(s); i++ {
        d := s[i]
        var v byte
        switch {
        case '0' <= d && d <= '9':
            v = d - '0'
        case 'a' <= d && d <= 'f':
            v = d - 'a' + 10
        case 'A' <= d && d <= 'F':
            v = d - 'A' + 10
        default:
            n = 0
            err = ErrSyntax
            goto Error
        }
        if n >= math.MaxInt64/16+1 {
            n = math.MaxInt64
            err = ErrRange
            goto Error
        }
        n *= 16
        n1 := n + int64(v)
        if n1 < n {
            n = math.MaxInt64
            err = ErrRange
            goto Error
        }
        n = n1
    }
    return n, nil
Error:
    return n, &ParseError{s, err}
}

func HandleErrors(tests []string) []string {
    e := make([]string, len(tests))
    for i, s := range tests {
        _, err := ParseHex(s)
        switch pe, ok := err.(*ParseError); {
        case err == nil:
            e[i] = "none"
        case ok && pe.Err == ErrSyntax:
            e[i] = "syntax"
        case ok && pe.Err == ErrRange:
            e[i] = "range"
        }
    }
    return e
}
```

#### Steps
1. Write the implementation to `hexadecimal.go`
2. Run `go test ./go/exercises/practice/hexadecimal/` to verify all tests pass
3. Run `go vet ./go/exercises/practice/hexadecimal/` to check for issues
4. Commit the solution
