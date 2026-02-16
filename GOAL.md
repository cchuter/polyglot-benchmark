# Goal: Solve Alphametics Puzzles in Go

## Problem Statement

Implement a `Solve` function in Go that solves alphametics puzzles. Alphametics puzzles replace letters in words with digits such that:
- The resulting arithmetic equation is valid
- Each letter maps to a unique digit (0-9)
- Leading digits of multi-digit numbers must not be zero

The function signature is:
```go
func Solve(puzzle string) (map[string]int, error)
```

The puzzle string contains words joined by `+` operators on the left side and `==` separating the result on the right side.

## Acceptance Criteria

1. `Solve` returns a `map[string]int` mapping each uppercase letter to its unique digit
2. `Solve` returns an error when no valid solution exists
3. All 10 test cases in `cases_test.go` pass, including:
   - Puzzle with three letters (`I + BB == ILL`)
   - Unique value constraint (`A == B` → error)
   - Leading zero constraint (`ACA + DD == BD` → error)
   - Two-digit carry (`A + A + ... + B == BCC`)
   - Four, six, seven, eight, and ten letter puzzles
   - Large puzzle with 199 addends and 10 letters
4. `go vet ./...` passes with no issues
5. The solution runs within a reasonable time for the large 199-addend test case

## Key Constraints

- Package name: `alphametics`
- File: `go/exercises/practice/alphametics/alphametics.go`
- Go 1.18 module
- Must handle up to 10 unique letters (digits 0-9)
- Must enforce no leading zeros on multi-digit words
- Must handle multiple addends (not just two)
