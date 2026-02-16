# Goal: polyglot-go-octal

## Problem Statement

Implement an octal (base-8) to decimal (base-10) converter in Go. The function `ParseOctal` receives an octal number as a string and returns its decimal equivalent as an `int64`, along with an error if the input is invalid.

This is part of the Exercism polyglot benchmark suite. The solution file `go/exercises/practice/octal/octal.go` currently contains only the package declaration and needs the `ParseOctal` function implemented.

## Acceptance Criteria

1. `ParseOctal("1")` returns `(1, nil)`
2. `ParseOctal("10")` returns `(8, nil)`
3. `ParseOctal("1234567")` returns `(342391, nil)`
4. `ParseOctal("carrot")` returns `(0, error)` — invalid input containing non-octal characters
5. `ParseOctal("35682")` returns `(0, error)` — invalid input containing digits 8 and 9
6. All tests in `octal_test.go` pass (`go test ./...`)
7. The conversion must be implemented from first principles — no use of `strconv.ParseInt` or similar built-in/external conversion functions

## Key Constraints

- The function signature must be: `func ParseOctal(string) (int64, error)`
- Valid octal digits are 0-7 only
- Invalid input must return 0 and a non-nil error
- Must use first-principles implementation (manual digit-by-digit conversion)
- Package name: `octal`
- Go module: `octal` with `go 1.18`
