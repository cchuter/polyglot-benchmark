# Challenger Review: alphametics implementation

**Reviewer:** challenger
**File:** `go/exercises/practice/alphametics/alphametics.go`
**Verdict:** PASS — Implementation is correct. No bugs or security issues found.

---

## 1. Correctness Analysis

### Parsing (`parsePuzzle`, lines 29-86)

- **Token handling**: Splits on whitespace, correctly skips `+` and `==` operators, collects all other tokens as value strings. The last value string is treated as the answer (right-hand side of `==`). This is correct for all test puzzle formats.
- **Character validation**: Line 42 checks every character is an uppercase letter via `unicode.IsUpper(r)`. Returns `nil` (triggering "invalid puzzle" error) for any invalid character. Sound.
- **Letter indexing**: Uses `r - 'A'` as index into a `[26]int` array. Since only uppercase ASCII letters pass validation, this is always in bounds [0, 25]. Correct.
- **Column-indexed digits** (`vDigits`): Each word is reversed into column order (rightmost digit = index 0). Values are `r - 'A' + 1` (1-26) for letters, 0 for empty positions. This encoding is used consistently in `isPuzzleSolution`. Correct.

### Solver (`solvePuzzle`, lines 88-106)

- Generates all P(10, nLetters) permutations and tests each. Since permutations produce distinct digit assignments, the "each letter maps to a unique digit" constraint is inherently satisfied. Correct.
- Leading zero check (lines 92-101): After finding a mathematically valid solution, checks all leading letters of multi-digit words. If any has value 0, the solution is skipped (`continue`), and search proceeds. If no non-leading-zero solution exists, returns error. Correct.

### Column-by-column verification (`isPuzzleSolution`, lines 109-143)

- **Carry propagation**: `carry` starts at 0, and for each column `sum = carry + addend_digits`. After checking, `carry = sum / 10`. Correct.
- **Addend summation**: Row 0 is the first addend (lines 120-125), rows 1 through n-2 are remaining addends (lines 127-133), row n-1 is the answer (line 137). The loop correctly separates addends from the answer.
- **Result check**: `sum % 10` must equal the answer's digit in this column (line 138). Returns false if mismatch or if answer has no character in this position (`r == 0`). Correct.
- **Final carry check** (line 142): `return carry == 0` ensures no overflow beyond the result word's length. This catches cases like `A + A = B` where A=9 would produce carry=1 with no column to absorb it. Critical and correct.

### Permutations (`permutations`, lines 158-230)

- Implements Python's `itertools.permutations()` algorithm faithfully. Generates r-length permutations from a pool of n elements.
- Capacity pre-calculation (lines 161-171): Computes P(n,r) = n!/(n-r)! for pre-allocation. For n=10, max P(10,10) = 3,628,800 which fits in an int. No overflow risk.

---

## 2. Edge Case Analysis

| Edge Case | Handling | Correct? |
|---|---|---|
| Leading zeros on multi-digit words | `leadingLetters` tracks first letter of words with `len > 1`; checked after each solution candidate | Yes |
| Single-letter words (e.g., "I", "A") | Not added to `leadingLetters` (line 53: `len(field) > 1`), so they CAN be zero | Yes |
| No solution exists | Exhausts all permutations, returns `errors.New("no solution")` | Yes |
| `A == B` (unique value constraint) | Permutations always produce distinct values; A != B always, so no column match possible | Yes |
| Carry overflow | Final `carry == 0` check on line 142 | Yes |
| Result word longer than addends | Higher columns have `r == 0` for addends (sum = carry only); result digit checked against carry | Yes |
| 199 addends (stress test) | Same algorithm, just more rows to sum per column. Bounded by P(10,10) permutations | Yes |
| Non-uppercase input | `unicode.IsUpper(r)` check rejects; returns nil -> "invalid puzzle" error | Yes |

---

## 3. Security Review

- **Input validation**: All non-operator characters validated as uppercase letters. No injection risk.
- **Bounded memory**: At most P(10,10) = 3.6M permutations (10-element int slices). Memory usage is bounded by the digit space, not by input size.
- **No external I/O**: No file, network, or OS operations. Pure computation.
- **No panics**: Array indices are bounded by the `[26]int` letterValues array and validated input. No out-of-bounds risk.

**No security vulnerabilities found.**

---

## 4. Adherence to Plan and Spec

| Requirement | Status |
|---|---|
| Function signature: `Solve(puzzle string) (map[string]int, error)` | Met (line 20) |
| Package name: `alphametics` | Met (line 1) |
| Module: `alphametics` with `go 1.18` | Met (go.mod) |
| Standard library only | Met (imports: errors, strings, unicode) |
| Column-by-column addition with carry | Met (isPuzzleSolution) |
| Permutation-based brute-force | Met (permutations function) |
| Leading zero validation | Met (solvePuzzle lines 92-101) |
| No-solution error | Met (solvePuzzle line 105) |

---

## 5. Minor Observations (Not Bugs)

1. **Memory**: All permutations are materialized upfront. For 10 letters, this allocates ~3.6M slices. A lazy/generator approach would reduce peak memory, but this works within test constraints (~1.2s).
2. **Parser permissiveness**: Any non-`+`/`==` token is treated as a value word. This means malformed puzzles like `FOO BAR` (no operator) would still be processed. Not a bug for the exercism context since test inputs are well-formed.
3. **`rune(field[0])`** on line 54: `field[0]` is a byte, cast to rune. Safe for ASCII uppercase letters (all single-byte in UTF-8).

---

## Conclusion

The implementation is correct, handles all edge cases properly, has no security vulnerabilities, and adheres to the plan and specification. All 10 test cases should pass. **Approved without changes.**
