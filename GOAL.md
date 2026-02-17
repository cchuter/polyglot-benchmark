# Goal: Implement Octal-to-Decimal Conversion (Issue #309)

## Problem Statement

Implement the `ParseOctal` function in `go/exercises/practice/octal/octal.go` that converts an octal (base-8) number represented as a string to its decimal (base-10) equivalent using first principles. No built-in or external conversion libraries may be used.

## Acceptance Criteria

1. `ParseOctal("1")` returns `(1, nil)`
2. `ParseOctal("10")` returns `(8, nil)`
3. `ParseOctal("1234567")` returns `(342391, nil)`
4. `ParseOctal("carrot")` returns `(0, error)` — invalid input produces an error
5. `ParseOctal("35682")` returns `(0, error)` — digits 8 and 9 are invalid in octal
6. All tests in `octal_test.go` pass (`go test`)
7. Benchmark test runs without error (`go test -bench .`)
8. The conversion must be implemented from first principles — no use of `strconv.ParseInt`, `fmt.Sscanf`, or similar built-in conversion functions

## Key Constraints

- Function signature must be: `func ParseOctal(input string) (int64, error)`
- Invalid input (non-digit characters or digits 8-9) must return `(0, error)`
- Valid octal digits are 0-7
- Package name: `octal`
- Module: `octal` with `go 1.18`
