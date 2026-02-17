# Goal: polyglot-go-octal

## Problem Statement

Implement the `ParseOctal` function in `go/exercises/practice/octal/octal.go` that converts an octal (base-8) number string to its decimal (base-10) `int64` equivalent using first principles. No built-in or external conversion libraries may be used.

## Acceptance Criteria

1. `ParseOctal("1")` returns `(1, nil)`
2. `ParseOctal("10")` returns `(8, nil)`
3. `ParseOctal("1234567")` returns `(342391, nil)`
4. `ParseOctal("carrot")` returns `(0, error)` — invalid input yields 0 and an error
5. `ParseOctal("35682")` returns `(0, error)` — digits 8 and 9 are invalid in octal
6. All tests in `octal_test.go` pass (`go test ./...`)
7. Code passes `go vet ./...`
8. No use of built-in conversion functions (e.g., `strconv.ParseInt`)

## Key Constraints

- The function signature must be: `func ParseOctal(input string) (int64, error)`
- Conversion must be implemented from first principles
- Invalid input (non-octal characters, including digits 8-9) must return `0` and an error
- The solution goes in `go/exercises/practice/octal/octal.go` in package `octal`
