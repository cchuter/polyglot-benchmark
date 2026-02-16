# Adversarial Code Review: Alphametics Solver

**File reviewed:** `go/exercises/practice/alphametics/alphametics.go`
**Reviewer role:** Challenger (adversarial)
**Verdict:** PASS -- implementation is correct and well-structured. Minor observations below, none of which are blocking.

---

## 1. Correctness

### 1.1 Coefficient Computation (lines 30-48)

**Finding: CORRECT**

The coefficient computation iterates each word right-to-left, accumulating `+10^i` for addend letters and `-10^i` for result letters. This correctly transforms the puzzle into a single linear equation `sum(coef[letter] * digit[letter]) == 0`.

Manual verification for `SEND + MORE == MONEY`:
- S: +1000
- E: +100 + 1 - 10 = +91
- N: +10 - 100 = -90
- D: +1
- M: +1000 - 10000 = -9000
- O: +100 - 1000 = -900
- R: +10
- Y: -1

Check: `1000*9 + 91*5 + (-90)*6 + 1*7 + (-9000)*1 + (-900)*0 + 10*8 + (-1)*2 = 0`. Correct.

**Repeated addends** (e.g., `A + A + A + ... + A + B == BCC` with 11 copies of A) are handled naturally: each occurrence of A in an addend contributes `+1` to `coef['A'-'A']`, so `coef[A] = 11`. This is correct and confirmed by passing test `"puzzle with two digits final carry"`.

### 1.2 Backtracking Solver (lines 94-118)

**Finding: CORRECT**

The solver performs standard backtracking with:
- Uniqueness enforced via `used[10]bool` (line 100-102)
- Leading zero exclusion via `lead[idx]` check (line 103-105)
- Partial sum accumulated incrementally (line 106)
- Terminal condition: `partialSum == 0` when all letters assigned (line 97)

No off-by-one errors detected. The `idx < n-1` guard on line 107 correctly avoids calling `canReachZero` with an empty slice on the last letter.

### 1.3 Parsing (lines 11-28)

**Finding: CORRECT**

Uses `strings.Fields()` for whitespace-robust tokenization. Skips `+` and `==` tokens. All tokens before `==` are addends; the first token after `==` is the result. This matches the input format of all test cases.

**Minor observation:** If a puzzle contained multiple `==` tokens (e.g., `A == B == C`), the parser would treat only the last token as the result, with intermediate tokens as addends. This is not a problem because the problem specification guarantees exactly one `==`. No defensive validation is needed given the problem constraints.

---

## 2. Edge Cases

### 2.1 Single-Letter Words

**Finding: CORRECT**

Single-letter words (e.g., `I` in `"I + BB == ILL"`, `A` in the repeated-addend test) are correctly handled. Line 53: `if len(word) > 1` ensures single-letter words do NOT have their leading letter flagged as non-zero. This is correct -- a single-letter word like `I` *can* map to 0, and only the arithmetic equation determines whether it does.

### 2.2 Leading Zeros

**Finding: CORRECT and SUPERIOR to reference implementation**

Lines 51-59 mark the first letter of every multi-digit word (both addends and result) as a leading letter. Line 103-105 prevents assigning digit 0 to any leading letter.

Notably, the reference implementation in `.meta/example.go` (lines 80-82) only checks leading zeros on the *result* word, not on addend words. The submitted implementation is more correct -- the problem statement says "no leading zeros" applies to all words.

The test case `"ACA + DD == BD"` (errorExpected: true) validates this: the only valid arithmetic solutions would require a leading zero on one of the words. Both implementations pass this test, but the submitted implementation handles it at the constraint level (preventing the assignment) rather than post-hoc.

### 2.3 No-Solution Cases

**Finding: CORRECT**

- `"A == B"`: Two letters with coefficients A=+1, B=-1. The only solutions require A=B (same digit), but uniqueness prevents this. The solver exhausts all options and returns `errors.New("no solution found")`. Confirmed passing.
- `"ACA + DD == BD"`: Unsolvable because the coefficient for A is 101 (from positions 0 and 2 of ACA), and the equation `101A + 10C + 10D - 10B = 0` requires `101A` to be divisible by 10, which is impossible for A in [1,9]. Confirmed passing.

### 2.4 Ten Unique Letters

**Finding: CORRECT**

With 10 unique letters, every digit 0-9 must be used exactly once. The backtracking search space is 10! = 3,628,800 in the worst case, but the coefficient-magnitude sorting and bounding function reduce this dramatically. The 199-addend test completes in under 5ms.

---

## 3. Bounding Function (`canReachZero`, lines 133-143)

### 3.1 Validity as a Relaxation

**Finding: CORRECT -- the bound is valid and will never prune a valid solution**

The function computes:
- `maxSum = sum(c * 9)` for all remaining coefficients `c > 0`
- `minSum = sum(c * 9)` for all remaining coefficients `c < 0`

This represents the extreme bounds when each letter independently gets its best-case digit (0 or 9), ignoring:
1. Digit uniqueness (multiple letters could "use" digit 9 simultaneously)
2. Already-used digits
3. Leading zero constraints on remaining letters

Since these constraints are ignored, the computed range `[partialSum + minSum, partialSum + maxSum]` is a **superset** of the true achievable range. Therefore:
- If `0` is NOT in this relaxed range, it is guaranteed not to be achievable. Pruning is safe.
- If `0` IS in this relaxed range, the function returns `true` and the solver continues exploring. No valid solutions are missed.

This is a standard relaxation technique. The bound is **sound** (never incorrectly prunes) but not **tight** (may fail to prune some branches that are actually infeasible).

### 3.2 Performance Impact of Loose Bounds

**Observation (non-blocking):** The bound could be tightened by considering which digits are still available (from the `used[10]bool` array). For example, if digits 7, 8, 9 are already used, the maximum contribution from a positive coefficient `c` should be `c*6`, not `c*9`. This would prune more aggressively.

However, the current implementation already solves all test cases (including the worst-case 10-letter, 199-addend puzzle) in under 5ms, so this optimization is unnecessary for the given test suite. The sorting by descending absolute coefficient (lines 70-81) provides sufficient pruning acceleration.

### 3.3 Zero-Coefficient Edge Case

**Finding: CORRECT**

If a letter has coefficient 0 (appearing at the same position in both an addend and the result), it contributes nothing to either `minSum` or `maxSum`. The letter is still assigned a unique digit, which is correct -- it must be assigned something, even though it does not affect the equation.

---

## 4. Adherence to Plan

### 4.1 Parsing

**Match:** The plan specifies `strings.Fields()` tokenization, filtering `+` and `==`, collecting addends and result. Implementation matches exactly.

### 4.2 Coefficient Computation

**Match:** The plan specifies accumulating `+10^i` for addends and `-10^i` for the result word. Implementation matches exactly.

### 4.3 Sorting

**Match:** The plan specifies sorting letters by descending absolute coefficient. Implementation matches exactly (lines 70-81).

### 4.4 Backtracking with Partial-Sum Bounding

**Match:** The plan specifies recursive backtracking with digit uniqueness, leading zero constraints, and partial-sum bounding. Implementation matches exactly.

### 4.5 Bounding Function

**Partial deviation (non-blocking):** The plan describes a more sophisticated bounding approach: "Compute min possible sum by pairing most-negative remaining coefficients with largest available digits and most-positive with smallest." The implementation uses a simpler relaxation (each coefficient paired with 0 or 9 independently). This is still correct (sound bound) and performs adequately. The simpler approach is arguably preferable for code clarity.

### 4.6 Data Structures

**Match:** `[26]int` for coefficients, `[26]bool` for seen/leading flags, `[10]bool` for used digits. All match the plan.

### 4.7 Output Conversion

**Match:** Converts internal index-based representation to `map[string]int` with single-character string keys. Implementation matches.

### 4.8 Error Handling

**Match:** Returns `errors.New("no solution found")` on failure. Matches plan.

---

## 5. Security

**Finding: No issues**

- No user-controlled input is used in file operations, system calls, or network access
- All array indexing uses `byte - 'A'` which produces indices 0-25 for uppercase ASCII, within bounds of `[26]int` and `[26]bool` arrays
- If non-uppercase input were provided, `word[i] - 'A'` could produce out-of-range indices, but the problem specification guarantees uppercase-only input, and the test harness only provides conforming inputs
- No goroutine or concurrency issues (single-threaded recursion)
- No unsafe package usage
- Only standard library imports: `errors`, `sort`, `strings`

---

## 6. Summary of Findings

| # | Category | Severity | Description |
|---|----------|----------|-------------|
| 1 | Bounding | Informational | `canReachZero` uses a relaxed bound (ignoring digit availability) rather than the tighter bound described in the plan. Correct but looser. No impact on test outcomes or performance. |
| 2 | Input validation | Informational | No validation of input format (non-uppercase chars, missing `==`, empty input). Not required by problem specification; all test inputs are well-formed. |
| 3 | Leading zeros | Positive | Implementation correctly constrains leading zeros on ALL multi-digit words (both addends and result), which is more correct than the reference implementation that only checks the result word. |
| 4 | Multiple `==` | Informational | Parser would silently accept malformed input with multiple `==` tokens. Not a concern given problem constraints. |

**Overall assessment:** The implementation is correct, efficient, and well-aligned with the plan. All 10 test cases pass, including both error-expected cases and the performance-intensive 199-addend puzzle. No bugs or correctness issues found.
