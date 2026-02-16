# Implementation Plan: polyglot-go-hexadecimal

## Branch 1: Minimal Approach — Direct Character Conversion

**Approach:** Simple loop with a helper function to convert each hex character to its value. Use `fmt.Errorf` for errors with embedded "syntax"/"range" keywords.

**Files to modify:**
- `go/exercises/practice/hexadecimal/hexadecimal.go`

**Design:**
```go
package hexadecimal

import "fmt"

func hexVal(b byte) (int64, bool) {
    switch {
    case '0' <= b && b <= '9': return int64(b - '0'), true
    case 'a' <= b && b <= 'f': return int64(b - 'a' + 10), true
    case 'A' <= b && b <= 'F': return int64(b - 'A' + 10), true
    default: return 0, false
    }
}

func ParseHex(s string) (int64, error) {
    if len(s) == 0 {
        return 0, fmt.Errorf("syntax error: empty string")
    }
    var n int64
    for i := 0; i < len(s); i++ {
        v, ok := hexVal(s[i])
        if !ok {
            return 0, fmt.Errorf("syntax error: invalid char %c", s[i])
        }
        prev := n
        n = n*16 + v
        if n < prev {
            return 0, fmt.Errorf("range error: overflow")
        }
    }
    return n, nil
}

func HandleErrors(tests []string) []string {
    result := make([]string, len(tests))
    for i, s := range tests {
        _, err := ParseHex(s)
        if err == nil {
            result[i] = "none"
        } else if strings.Contains(err.Error(), "syntax") {
            result[i] = "syntax"
        } else {
            result[i] = "range"
        }
    }
    return result
}
```

**Evaluation:**
- Feasibility: High — simple, no dependencies beyond fmt/strings
- Risk: Medium — overflow detection via `n < prev` is slightly fragile; must ensure correctness for edge cases like exactly MaxInt64+1. Also requires `strings` import for HandleErrors.
- Alignment: Good — all acceptance criteria met
- Complexity: Low — single file, ~40 lines

---

## Branch 2: Structured Error Types — Matches Reference Solution Pattern

**Approach:** Follow the pattern from `.meta/example.go` with sentinel errors (`ErrRange`, `ErrSyntax`), a `ParseError` struct, and robust overflow detection using `math.MaxInt64/16+1`.

**Files to modify:**
- `go/exercises/practice/hexadecimal/hexadecimal.go`

**Design:**
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
        return 0, &ParseError{s, ErrSyntax}
    }
    for i := 0; i < len(s); i++ {
        d := s[i]
        var v byte
        switch {
        case '0' <= d && d <= '9': v = d - '0'
        case 'a' <= d && d <= 'f': v = d - 'a' + 10
        case 'A' <= d && d <= 'F': v = d - 'A' + 10
        default:
            return 0, &ParseError{s, ErrSyntax}
        }
        if n >= math.MaxInt64/16+1 {
            return 0, &ParseError{s, ErrRange}
        }
        n *= 16
        n1 := n + int64(v)
        if n1 < n {
            return 0, &ParseError{s, ErrRange}
        }
        n = n1
    }
    return n, nil
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

**Evaluation:**
- Feasibility: High — directly follows the reference solution pattern
- Risk: Low — overflow detection is proven correct (uses two-stage check from Go stdlib pattern)
- Alignment: Excellent — all acceptance criteria met, error messages naturally contain "syntax" and "range"
- Complexity: Low-Medium — single file, ~60 lines, but well-structured

---

## Branch 3: Lookup Table — Performance-Optimized

**Approach:** Use a pre-computed 256-byte lookup table for character-to-value mapping, avoiding branching in the hot loop. Same error structure as Branch 2.

**Files to modify:**
- `go/exercises/practice/hexadecimal/hexadecimal.go`

**Design:**
```go
package hexadecimal

import (
    "errors"
    "math"
)

var ErrRange = errors.New("value out of range")
var ErrSyntax = errors.New("invalid syntax")

// lookup table: 0xFF means invalid
var hexTable [256]byte

func init() {
    for i := range hexTable { hexTable[i] = 0xFF }
    for i := byte('0'); i <= '9'; i++ { hexTable[i] = i - '0' }
    for i := byte('a'); i <= 'f'; i++ { hexTable[i] = i - 'a' + 10 }
    for i := byte('A'); i <= 'F'; i++ { hexTable[i] = i - 'A' + 10 }
}

type ParseError struct { Num string; Err error }

func (e *ParseError) Error() string {
    return "hexadecimal.ParseHex: parsing \"" + e.Num + "\": " + e.Err.Error()
}

func ParseHex(s string) (int64, error) {
    if len(s) == 0 {
        return 0, &ParseError{s, ErrSyntax}
    }
    var n int64
    for i := 0; i < len(s); i++ {
        v := hexTable[s[i]]
        if v == 0xFF {
            return 0, &ParseError{s, ErrSyntax}
        }
        if n >= math.MaxInt64/16+1 {
            return 0, &ParseError{s, ErrRange}
        }
        n *= 16
        n1 := n + int64(v)
        if n1 < n {
            return 0, &ParseError{s, ErrRange}
        }
        n = n1
    }
    return n, nil
}

// HandleErrors same as Branch 2
```

**Evaluation:**
- Feasibility: High — standard optimization technique
- Risk: Low — same proven overflow logic
- Alignment: Good — all acceptance criteria met
- Complexity: Medium — adds init(), lookup table, slightly harder to read
- Performance: Better for benchmarks, but over-engineered for this exercise

---

## Selected Plan

**Selected: Branch 2 — Structured Error Types**

**Rationale:**
- Branch 2 is superior because it directly follows the reference solution pattern from `.meta/example.go`, meaning it's proven correct for all test cases
- It uses robust overflow detection (two-stage: pre-multiply check + post-add check) from the Go standard library pattern
- The structured error types (`ParseError` with sentinel `ErrRange`/`ErrSyntax`) allow `HandleErrors` to use clean type assertion rather than string matching, which is more idiomatic Go
- Branch 1's string-matching approach for HandleErrors is fragile and requires an extra `strings` import
- Branch 3 is over-engineered; the lookup table provides no meaningful benefit for this exercise's scope
- Branch 2 balances correctness, readability, and Go idioms perfectly

### Detailed Implementation Plan

**File to modify:** `go/exercises/practice/hexadecimal/hexadecimal.go`

**Step 1:** Define sentinel errors and ParseError type
- `var ErrRange = errors.New("value out of range")`  — contains "range"
- `var ErrSyntax = errors.New("invalid syntax")` — contains "syntax"
- `type ParseError struct { Num string; Err error }` with `Error()` method

**Step 2:** Implement `ParseHex`
- Empty string check → return `&ParseError{s, ErrSyntax}`
- Byte-by-byte loop with switch for 0-9, a-f, A-F
- Pre-multiply overflow check: `if n >= math.MaxInt64/16+1`
- Post-add overflow check: `if n1 < n`
- On overflow → return `&ParseError{s, ErrRange}`

**Step 3:** Implement `HandleErrors`
- Allocate result slice
- Call `ParseHex` on each input
- Type-assert error to `*ParseError` and compare `.Err` against sentinels
- Return "none", "syntax", or "range" accordingly

**Step 4:** Verify
- Run `go test` in the hexadecimal directory
- All 11 test cases and HandleErrors test must pass
- Benchmark should run without error
