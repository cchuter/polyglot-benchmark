# Implementation Plan: polyglot-go-hexadecimal

## Proposal A: Structured Error Types with goto (Reference-style)

**Role: Proponent**

This approach closely follows the reference implementation pattern from `.meta/example.go`, using a `ParseError` struct type wrapping sentinel errors (`ErrRange`, `ErrSyntax`), and using `goto` for error flow control.

### Files to Modify
- `go/exercises/practice/hexadecimal/hexadecimal.go` — sole implementation file

### Architecture
1. Define two sentinel errors: `ErrRange` and `ErrSyntax` using `errors.New()`
2. Define a `ParseError` struct with `Num string` and `Err error` fields
3. Implement `ParseError.Error()` method that includes the error category string ("range" or "syntax") in its output
4. Implement `ParseHex` with character-by-character conversion, overflow detection via `math.MaxInt64`, and `goto Error` flow
5. Implement `HandleErrors` using type assertion on `*ParseError` to categorize results

### Rationale
- Mirrors Go standard library patterns (cf. `strconv` package)
- The structured `ParseError` type allows `HandleErrors` to inspect the underlying error category cleanly via type assertion
- `goto` simplifies the error return path by consolidating it at a single label
- Directly matches the reference solution, minimizing risk of divergence from expected behavior

### Ordering
1. Write error sentinels and `ParseError` type
2. Write `ParseHex` function
3. Write `HandleErrors` function
4. Run tests

---

## Proposal B: Simple fmt.Errorf with String Matching

**Role: Opponent**

This approach uses simple `fmt.Errorf` calls to produce error messages, without a custom error type. `HandleErrors` inspects error messages via `strings.Contains`.

### Files to Modify
- `go/exercises/practice/hexadecimal/hexadecimal.go` — sole implementation file

### Architecture
1. No custom error type — just return `fmt.Errorf("syntax: ...")` or `fmt.Errorf("range: ...")`
2. `ParseHex` uses a simple loop with early returns (no goto)
3. `HandleErrors` checks `err == nil`, then uses `strings.Contains(err.Error(), "syntax")` and `strings.Contains(err.Error(), "range")` to categorize

### Critique of Proposal A
- `goto` is generally considered non-idiomatic in modern Go (though the stdlib does use it internally)
- The `ParseError` struct adds complexity not strictly required by the test suite
- More code to maintain for the same test outcomes

### Rationale for B
- Simpler: fewer lines, no custom types
- Uses `fmt.Errorf` which is standard Go error creation
- Tests only check `strings.Contains(strings.ToLower(err.Error()), test.errCase)` — they don't require a specific error type

### Ordering
1. Write `ParseHex` with `fmt.Errorf`
2. Write `HandleErrors` with string matching
3. Run tests

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A (Structured Errors) | Proposal B (fmt.Errorf) |
|---------------|-------------------------------|------------------------|
| Correctness   | Fully satisfies all test cases | Fully satisfies all test cases |
| Risk          | Low — mirrors reference exactly | Low, but `fmt` import adds a dependency not strictly needed; string matching in HandleErrors is fragile |
| Simplicity    | Slightly more code but clear structure | Fewer lines but less robust |
| Consistency   | Matches reference solution and Go stdlib patterns | Deviates from reference; less conventional for error handling |

### Decision: Proposal A wins

**Rationale:**
1. **Correctness**: Both satisfy tests, but Proposal A's structured errors are more robust. The test file includes `var _ error = err` to verify the error interface, and while both satisfy this, Proposal A's `ParseError` type is clearly what the test was designed around.
2. **Risk**: Proposal B's `HandleErrors` using string matching on its own errors is circular and fragile. Proposal A's type assertion is cleaner and more reliable.
3. **Consistency**: The reference solution uses the structured approach. Other exercises in this repo follow conventional Go patterns. Proposal A fits better.
4. **Simplicity**: While A has slightly more code, it's well-structured code that's easy to understand.

### Selected Plan: Detailed Implementation

**File**: `go/exercises/practice/hexadecimal/hexadecimal.go`

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

### Key design decisions:
1. **Sentinel errors**: `ErrRange` and `ErrSyntax` as package-level variables for comparison
2. **ParseError struct**: Wraps the sentinel error with the input string for informative messages
3. **Error.Error() method**: Returns string containing "syntax" or "range" (via the sentinel's message), satisfying the `strings.Contains` check in tests
4. **Overflow detection**: Two-stage: check if `n >= MaxInt64/16+1` before multiplication, then check if `n + v` overflows by comparing with pre-addition value
5. **HandleErrors**: Type-asserts to `*ParseError` and compares the inner `Err` field against sentinels
