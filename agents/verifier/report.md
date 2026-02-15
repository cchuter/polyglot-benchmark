# Verification Report: Go Alphametics Solver

**Date:** 2026-02-14
**Verifier:** verifier agent
**Verdict:** **PASS**

---

## Independent Test Run

```
$ cd go/exercises/practice/alphametics && go test -v ./...

=== RUN   TestSolve
--- PASS: TestSolve (1.13s)
    --- PASS: TestSolve/puzzle_with_three_letters (0.00s)
    --- PASS: TestSolve/solution_must_have_unique_value_for_each_letter (0.00s)
    --- PASS: TestSolve/leading_zero_solution_is_invalid (0.00s)
    --- PASS: TestSolve/puzzle_with_two_digits_final_carry (0.00s)
    --- PASS: TestSolve/puzzle_with_four_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_six_letters (0.02s)
    --- PASS: TestSolve/puzzle_with_seven_letters (0.08s)
    --- PASS: TestSolve/puzzle_with_eight_letters (0.23s)
    --- PASS: TestSolve/puzzle_with_ten_letters (0.43s)
    --- PASS: TestSolve/puzzle_with_ten_letters_and_199_addends (0.37s)
PASS
ok  alphametics  1.13s
```

Results match executor's reported output exactly.

## Acceptance Criteria Verification

### AC1: All 10 test cases pass — PASS

All 10 test cases in `cases_test.go` pass. Independently confirmed by running `go test -v ./...`. The executor's test results are accurate.

### AC2: Puzzles with 3-10 unique letters solved correctly — PASS

The test suite covers the full range:
| Test | Unique Letters | Result |
|------|---------------|--------|
| `I + BB == ILL` | 3 (B, I, L) | PASS |
| `A + A + ... + B == BCC` | 3 (A, B, C) | PASS |
| `AS + A == MOM` | 4 (A, M, O, S) | PASS |
| `NO + NO + TOO == LATE` | 6 (A, E, L, N, O, T) | PASS |
| `HE + SEES + THE == LIGHT` | 7 (E, G, H, I, L, S, T) | PASS |
| `SEND + MORE == MONEY` | 8 (D, E, M, N, O, R, S, Y) | PASS |
| `AND + A + STRONG + ...` | 10 (A, D, E, F, G, N, O, R, S, T) | PASS |
| 199-addend puzzle | 10 (A, E, F, H, I, L, O, R, S, T) | PASS |

### AC3: Error returned for no-solution puzzles — PASS

Test case `"A == B"` expects `errorExpected: true`. The solver correctly returns an error because no permutation can satisfy A == B with unique digits. Test passes.

### AC4: Leading zeros rejected — PASS

Test case `"ACA + DD == BD"` expects `errorExpected: true`. The solver correctly rejects this because B would need to be 0 (a leading zero). The column-length mismatch in `isPuzzleSolution` rejects all permutations, and the result-word leading zero check in `solvePuzzle` (lines 80-83) provides an additional guard. Test passes.

### AC5: Each letter maps to a unique digit — PASS

Test case `"solution must have unique value for each letter"` (`A == B`) directly tests this. Additionally, the permutation-based approach inherently guarantees uniqueness — each permutation assigns distinct digits to distinct letters. All solution maps in passing tests have unique values per letter.

### AC6: Handles up to 199 addends — PASS

Test case `"puzzle with ten letters and 199 addends"` passes in 0.37s. The solver correctly parses all 199 addends and computes the correct solution: `{A:1, E:0, F:5, H:8, I:7, L:2, O:6, R:3, S:4, T:9}`.

### AC7: Completes within reasonable time — PASS

Total test suite completes in ~1.13s. The most expensive test (10 unique letters, 10! = 3,628,800 permutations) completes in 0.43s. The 199-addend puzzle completes in 0.37s. No timeouts observed.

## Additional Checks

- **Test files unmodified:** Confirmed via `git diff main` — `alphametics_test.go` and `cases_test.go` have zero changes from main.
- **Function signature matches:** `func Solve(puzzle string) (map[string]int, error)` — confirmed.
- **Package name:** `alphametics` — confirmed.
- **Challenger review alignment:** The challenger's review also concludes PASS. The implementation matches the Exercism reference solution.

## Conclusion

**PASS** — All 7 acceptance criteria are met. The implementation is correct, complete, and performant. No changes required.
