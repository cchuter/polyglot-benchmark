# Goal: polyglot-go-octal (Issue #225)

## Problem Statement

Implement an octal (base-8) to decimal (base-10) converter in Go. The function `ParseOctal` takes an octal number represented as a string (e.g. `"1735263"`) and returns its decimal equivalent as an `int64`. The conversion must be implemented from first principles — no built-in or external conversion libraries may be used.

## Acceptance Criteria

1. `ParseOctal("1")` returns `(1, nil)`
2. `ParseOctal("10")` returns `(8, nil)`
3. `ParseOctal("1234567")` returns `(342391, nil)`
4. `ParseOctal("carrot")` returns `(0, error)` — invalid input containing non-octal characters
5. `ParseOctal("35682")` returns `(0, error)` — invalid input containing digits 8 and 9
6. All tests in `octal_test.go` pass (`go test` exits 0)
7. The benchmark `BenchmarkParseOctal` runs without error

## Key Constraints

- Must implement conversion from first principles (no `strconv.ParseInt`, `fmt.Sscanf`, or similar)
- Invalid input (any character not in `0-7`) must return `0` and a non-nil error
- The solution goes in `go/exercises/practice/octal/octal.go`
- The test file `octal_test.go` must NOT be modified
- The `go.mod` file must NOT be modified
