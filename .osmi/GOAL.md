# Goal: polyglot-go-alphametics (Issue #21)

## Problem Statement

Implement (and verify) the alphametics puzzle solver exercise in Go as part of the polyglot benchmark suite. Alphametics puzzles replace letters in words with digits such that an arithmetic equation holds, with constraints that each letter maps to a unique digit and leading digits of multi-digit numbers cannot be zero.

The implementation exists at `go/exercises/practice/alphametics/alphametics.go` and currently passes all tests. This issue validates the implementation meets acceptance criteria and closes issue #21 on a proper feature branch.

## Acceptance Criteria

1. **`Solve(puzzle string) (map[string]int, error)`** — Solves an alphametics puzzle
   - Accepts a puzzle string with zero or more `+` operators and one `==` operator
   - Returns a `map[string]int` mapping each letter to its unique digit (0-9)
   - Each letter must map to a different digit
   - Leading digits of multi-digit numbers must not be zero
   - Returns an error if no valid solution exists

2. **All 10 test cases pass** (`go test -v` in `go/exercises/practice/alphametics/`):
   - "puzzle with three letters": `I + BB == ILL` -> `{B:9, I:1, L:0}`
   - "solution must have unique value for each letter": `A == B` -> error
   - "leading zero solution is invalid": `ACA + DD == BD` -> error
   - "puzzle with two digits final carry": `A + A + ... + B == BCC` -> `{A:9, B:1, C:0}`
   - "puzzle with four letters": `AS + A == MOM` -> `{A:9, M:1, O:0, S:2}`
   - "puzzle with six letters": `NO + NO + TOO == LATE` -> `{A:0, E:2, L:1, N:7, O:4, T:9}`
   - "puzzle with seven letters": `HE + SEES + THE == LIGHT` -> `{E:4, G:2, H:5, I:0, L:1, S:9, T:7}`
   - "puzzle with eight letters": `SEND + MORE == MONEY` -> `{D:7, E:5, M:1, N:6, O:0, R:8, S:9, Y:2}`
   - "puzzle with ten letters": `AND + A + STRONG + ... == DEFENSE` -> `{A:5, D:3, E:4, F:7, G:8, N:0, O:2, R:1, S:6, T:9}`
   - "puzzle with ten letters and 199 addends": large puzzle -> `{A:1, E:0, F:5, H:8, I:7, L:2, O:6, R:3, S:4, T:9}`

3. **Benchmarks run** without error

4. **Package name:** `alphametics`

5. **Module:** `alphametics` with `go 1.18`

## Key Constraints

- Must follow existing exercism exercise conventions in the repository
- No external dependencies (standard library only)
- Test files (`alphametics_test.go`, `cases_test.go`) are read-only exercise specifications
- Must handle edge cases: leading zero rejection, no-solution detection, multi-addend puzzles
