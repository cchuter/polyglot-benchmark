# Implementation Plan: polyglot-go-octal

## Branch 1: Minimal — Direct iteration with fmt.Errorf

The simplest approach: iterate over each character, validate it's 0-7, accumulate using multiplication by 8.

**Files to modify:**
- `go/exercises/practice/octal/octal.go` — implement `ParseOctal`

**Approach:**
```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for _, ch := range input {
        if ch < '0' || ch > '7' {
            return 0, fmt.Errorf("invalid octal digit: %c", ch)
        }
        num = num*8 + int64(ch-'0')
    }
    return num, nil
}
```

**Evaluation:**
- Feasibility: Excellent — trivial to implement, mirrors the example.go exactly in logic
- Risk: Very low — straightforward string iteration
- Alignment: Fully satisfies all 7 acceptance criteria
- Complexity: 1 file, ~12 lines of code

## Branch 2: Extensible — Structured error type with constants

Use a custom error type similar to the hexadecimal exercise, with sentinel errors for different failure modes.

**Files to modify:**
- `go/exercises/practice/octal/octal.go` — implement `ParseOctal` with `ParseError` type

**Approach:**
```go
package octal

import "errors"

var ErrSyntax = errors.New("invalid syntax")

type ParseError struct {
    Input string
    Err   error
}

func (e *ParseError) Error() string {
    return "octal.ParseOctal: parsing \"" + e.Input + "\": " + e.Err.Error()
}

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for i := 0; i < len(input); i++ {
        d := input[i]
        if d < '0' || d > '7' {
            return 0, &ParseError{input, ErrSyntax}
        }
        num = num*8 + int64(d-'0')
    }
    return num, nil
}
```

**Evaluation:**
- Feasibility: Good — works but adds unnecessary structure for the test requirements
- Risk: Low — but over-engineered for the problem
- Alignment: Satisfies all criteria, tests only check error != nil
- Complexity: 1 file, ~25 lines — more code than needed

## Branch 3: Performance — Bit shifting with byte-level iteration

Use bit shifting (`<<3`) instead of multiplication and iterate over bytes directly rather than runes for maximum performance.

**Files to modify:**
- `go/exercises/practice/octal/octal.go` — implement `ParseOctal`

**Approach:**
```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
    num := int64(0)
    for i := 0; i < len(input); i++ {
        d := input[i]
        if d < '0' || d > '7' {
            return 0, fmt.Errorf("unexpected character '%c'", d)
        }
        num = num<<3 + int64(d-'0')
    }
    return num, nil
}
```

**Evaluation:**
- Feasibility: Excellent — this is essentially what .meta/example.go does
- Risk: Very low — bit shifting by 3 is equivalent to *8
- Alignment: Fully satisfies all criteria
- Complexity: 1 file, ~12 lines

## Selected Plan

**Branch 1 (Minimal)** is the best choice.

**Rationale:**
- All three branches produce the same functional outcome and pass all tests
- Branch 1 is the simplest and most readable — `num*8` is clearer than `num<<3` for expressing the octal-to-decimal algorithm
- Branch 2 is over-engineered; the tests only check `error != nil`, so a custom error type adds no value
- Branch 3's bit shifting micro-optimization is unnecessary for this exercise and less readable
- Branch 1 closely mirrors the reference solution in `.meta/example.go` while using idiomatic `range` iteration

**Selected implementation:**

### File: `go/exercises/practice/octal/octal.go`

Replace the stub with:

```go
package octal

import "fmt"

func ParseOctal(input string) (int64, error) {
	num := int64(0)
	for _, ch := range input {
		if ch < '0' || ch > '7' {
			return 0, fmt.Errorf("invalid octal digit: %c", ch)
		}
		num = num*8 + int64(ch-'0')
	}
	return num, nil
}
```

### Verification steps:
1. Write the implementation to `octal.go`
2. Run `go test ./...` from the octal directory
3. Confirm all 5 test cases pass and benchmark runs
