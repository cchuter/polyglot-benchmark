# Challenger Review: Alphametics Implementation

## Verdict: PASS

The implementation is correct and faithfully follows the selected plan (Branch 1 — Direct Permutation). It also includes all four key improvements over the reference solution identified during planning.

---

## Detailed Review by Focus Area

### 1. Leading Zero Check — Does it cover ALL multi-digit words?

**PASS.** The implementation correctly checks leading zeros for ALL multi-digit words (addends AND result), not just the answer.

In `parsePuzzle` (lines 41-44):
```go
if len(field) > 1 {
    p.isLeading[field[0]-'A'] = true
}
```
This iterates over every word in the puzzle (the `fields` loop includes all tokens that aren't `+` or `==`). Every word with length > 1 has its first letter marked as a leading letter. This covers addends and the result word.

In `solvePuzzle` (lines 88-97):
```go
for v := 0; v < 26; v++ {
    if p.isLeading[v] && p.letterValues[v] == 0 {
        leadingZero = true
        break
    }
}
```
All 26 possible letter slots are checked. Any letter marked as leading AND assigned digit 0 causes the permutation to be skipped.

**Comparison with reference:** The reference solution (`example.go` lines 79-83) only checks the leading digit of the answer word. This is a known bug that the plan explicitly called out. The implementation fixes it correctly.

**Test case verification:** "ACA + DD == BD" expects an error. Leading letters are A (from ACA), D (from DD), and B (from BD). Since the equation 101A + 11D = 10B + D requires A=0 (which is forbidden as a leading letter), no solution exists.

### 2. Single-letter Words — Are they exempt from leading-zero constraint?

**PASS.** The `len(field) > 1` guard ensures single-letter words like "I" and "A" are never marked as leading.

**Test case verification:**
- "I + BB == ILL": "I" as a standalone word does NOT set isLeading. However, "ILL" (len=3) DOES set isLeading for 'I'. So I must be non-zero — which is correct (I=1 in the expected solution).
- "A + A + ... + B == BCC": "A" and "B" as single-letter words don't set isLeading. "BCC" (len=3) sets isLeading for 'B'. Expected B=1 (non-zero) and C=0 (not leading). Correct.
- "AND + A + STRONG + ...": "A" appears as a single-letter word (exempt) AND as the first letter of "AND". The `isLeading` flag for 'A' gets set by "AND". Expected A=5 (non-zero). Correct.

### 3. Column Arithmetic — Is carry propagation correct?

**PASS.** The column arithmetic in `isPuzzleSolution` (lines 106-140) is correct.

- Column iteration: `d` ranges from 0 (rightmost/ones) to `maxDigits-1` (leftmost). Correct.
- Carry initialization: `carry` starts at 0 and accumulates across columns. Correct.
- Addend summation: Rows 0 through `len(vDigits)-2` are summed. Row `len(vDigits)-1` is the result. Correct.
- vDigits encoding: 0 means no character; 1-26 means 'A'-'Z'. When accessing `letterValues`, `r-1` converts back to 0-25 index. Consistent and correct.
- Carry computation: `carry = sum / 10; sum %= 10`. Standard. Correct.
- Result comparison: `sum != p.letterValues[r-1]`. Compares column remainder against the result digit's assigned value. Correct.
- **Final carry check** (line 139): `return carry == 0`. This is a correctness improvement over the reference solution, which lacks this check.

**Manual column trace for "I + BB == ILL" (I=1, B=9, L=0):**
- maxDigits = 3
- vDigits: I->[9,0,0], BB->[2,2,0], ILL->[12,12,9]
- Col 0: sum=0+1+9=10, carry=1, sum=0. Result L=0. Match.
- Col 1: sum=1+0+9=10, carry=1, sum=0. Result L=0. Match.
- Col 2: sum=1+0+0=1, carry=0, sum=1. Result I=1. Match.
- carry=0. Return true.

### 4. Permutation Function — Does it generate correct r-permutations?

**PASS.** The `permutations` function (lines 155-227) is identical to the reference solution and faithfully implements Python's `itertools.permutations()` algorithm.

- Count calculation: `n! / (n-r)!` = P(n,r). Correct.
- Pre-allocation: `make([][]int, 0, nperm)`. Efficient.
- Each permutation is deep-copied (`make` + `copy`) before appending. No aliasing bugs.
- The cycles-based algorithm correctly generates all r-permutations without repetition.
- Edge case `r > n`: returns empty slice. Correct.

### 5. Off-by-one Errors in Column Indexing

**PASS.** No off-by-one errors found.

- `parsePuzzle` vDigits construction (lines 62-68): `j = len(word) - 1 - d` correctly maps character position `d` (left-to-right) to column index `j` (right-to-left, 0-based).
- `isPuzzleSolution` column loop: `d` ranges `[0, maxDigits)`. Column 0 is ones place. Correct.
- Addend loop: `n` ranges `[1, len(vDigits)-1)`, combined with explicit row 0 access, sums all addend rows. Row len-1 is the result. Correct.

### 6. Alignment with Plan

**PASS.** All plan requirements are implemented:

| Plan Requirement | Status |
|---|---|
| Parse puzzle: split on whitespace, skip + and == | Implemented (lines 32-53) |
| Build column-based representation (right to left) | Implemented (lines 60-68) |
| Track unique letters and leading letters | Implemented (lines 41-44, 54-77) |
| Generate r-permutations of {0..9} | Implemented (lines 155-227) |
| Column-by-column check with carry | Implemented (lines 106-140) |
| Skip permutations with leading zero for ALL words | Implemented (lines 88-97) |
| Return first valid solution or error | Implemented (lines 81-102) |
| Leading zeros checked BEFORE isPuzzleSolution | Implemented (lines 95-98) |
| isLeading [26]bool for tracking | Implemented (line 17) |
| Final carry == 0 check | Implemented (line 139) |

---

## Minor Observations (Non-blocking)

1. **Redundant letterValues assignment**: `solvePuzzle` sets `letterValues` (lines 84-86) before the leading zero check, then `isPuzzleSolution` sets them again (lines 108-109). The second assignment is redundant but makes `isPuzzleSolution` self-contained. Not a bug.

2. **Memory usage for 10-letter case**: Pre-generating all 3,628,800 permutations allocates significant memory. This matches the reference solution's approach and is acceptable for the test suite.

---

## Conclusion

The implementation is correct, complete, and improves upon the reference solution in two important ways:
1. Leading zero check covers ALL multi-digit words (reference only checked the answer)
2. Final `carry == 0` check prevents accepting solutions with overflow

No correctness issues found. Ready for testing.
