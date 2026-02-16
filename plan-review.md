# Plan Review: Alphametics Solver

## Summary

The plan proposes a coefficient-based backtracking solver as an alternative to the reference implementation's brute-force permutation approach. The core idea is sound and the plan identifies the right optimizations. However, there are several gaps, a correctness concern, and missed optimization opportunities that should be addressed before implementation.

## 1. Correctness of the Proposed Approach

**The coefficient-based linear equation approach is correct.** Any alphametics puzzle of the form `W1 + W2 + ... + Wn == R` can be reduced to a single linear equation by computing the positional weight of each letter across all words. For example, in `SEND + MORE == MONEY`:

- S appears at position 3 (thousands) in an addend: coefficient +1000
- M appears at position 3 in an addend (+1000) and position 4 in the result (-10000): net coefficient -9000
- Y appears at position 0 in the result: coefficient -1

The solution is a digit assignment where the weighted sum equals zero. This is mathematically equivalent to column-by-column evaluation with carries, but is simpler to check (one dot product instead of iterative column sums).

**Potential correctness issue: leading zero detection.** The plan mentions checking that "leading letters are not assigned 0," but does not explicitly state that ALL words (including single-letter addends and the result) must be checked for leading zeros. Looking at the test cases:

- `"A == B"` -- A and B are both single-letter words, both are leading letters. The only solution would require A == B, but that violates uniqueness. The plan handles this via the uniqueness constraint, not leading zeros.
- `"ACA + DD == BD"` -- This expects an error. The only valid solution would require B=0 (leading the result "BD"), which is invalid.

The plan correctly identifies that leading letters of ALL words (addends and result) must be flagged. This is important because the reference implementation has a subtle bug: it only checks the leading digit of the answer (result word), not of addends. The plan's approach of flagging all leading letters is actually more correct than the reference.

## 2. Edge Case Coverage

Reviewing each test case against the plan:

| Test Case | Letters | Key Challenge | Plan Handles? |
|-----------|---------|---------------|---------------|
| `I + BB == ILL` | 3 | Simple, small | Yes |
| `A == B` | 2 | No valid solution (uniqueness) | Yes |
| `ACA + DD == BD` | 4 | Leading zero on result | Yes |
| `A + A + ... + A + B == BCC` | 3 | Repeated addends, carry propagation | Yes -- coefficients handle repetition naturally |
| `AS + A == MOM` | 4 | Single-letter addend "A" (leading letter) | Yes |
| `NO + NO + TOO == LATE` | 6 | Medium complexity | Yes |
| `HE + SEES + THE == LIGHT` | 7 | Medium complexity | Yes |
| `SEND + MORE == MONEY` | 8 | Classic puzzle | Yes |
| 10-letter puzzle | 10 | All digits used | Yes |
| 199-addend puzzle | 10 | Performance stress test | See section 3 |

**Edge case: repeated words as addends.** The 199-addend puzzle and the `A + A + ... + B == BCC` test repeat words. The coefficient approach handles this naturally -- each occurrence of a word adds to the letter's coefficient. For example, if "A" appears 11 times as an addend, the coefficient of A gets +11 (since A is a ones-digit letter in a single-letter word). This is correct.

**Edge case: single-letter words.** Words like "A" and "I" are both the leading letter and the only letter. They must not be assigned 0. The plan handles this.

**Edge case: no solution exists.** The plan returns an error when backtracking exhausts all possibilities. This is correct.

## 3. Performance Analysis for the 199-Addend / 10-Letter Test Case

This is the critical performance concern. With 10 unique letters, the search space is 10! = 3,628,800 permutations in the worst case.

**Reference implementation performance:** The reference generates ALL permutations upfront via `permutations(decDigits, 10)`, storing 3,628,800 slices in memory, then iterates through them checking each one column-by-column against 199 addends. This is extremely slow:
- Memory: ~3.6M * 10 ints = ~290 MB just for the permutation slices
- Time: For each of 3.6M permutations, it evaluates up to 200 rows across up to 10 columns

**Plan's proposed approach (backtracking with pruning):** The plan's backtracking approach avoids generating all permutations upfront, which is a major improvement. However, the plan is vague about *what specific pruning strategies* will be used during backtracking. The plan mentions:
- Uniqueness pruning (each digit used at most once) -- this is basic and does not reduce complexity significantly
- Leading zero pruning -- minor optimization

**Missing critical optimization: partial sum pruning.** The coefficient approach enables a powerful optimization that the plan alludes to but does not explicitly describe:

After assigning digits to k of n letters, you can compute the partial weighted sum. If the remaining unassigned letters cannot possibly bring the sum to zero (given the range of available digits), you can prune. Specifically:

```
partial_sum + min_possible_remaining <= 0 <= partial_sum + max_possible_remaining
```

Where min/max are computed by greedily assigning the smallest/largest available digits to the remaining letters' coefficients (sorted by sign). This is a powerful bound that can prune large subtrees.

**Recommendation:** The plan should explicitly describe this partial-sum bounding strategy. Without it, the backtracking approach for 10 letters still explores a significant portion of the 3.6M search space.

**Letter ordering heuristic:** The plan does not discuss the order in which letters are assigned during backtracking. Assigning letters with the largest absolute coefficient first tends to produce tighter bounds and prune more aggressively. This should be mentioned in the plan.

**Expected performance with optimizations:** With coefficient-magnitude ordering and partial-sum bounding, the 10-letter case typically resolves in thousands to tens of thousands of nodes rather than millions. This should comfortably pass within Go test timeouts.

## 4. Gaps and Issues in the Plan

### Gap 1: No discussion of parsing edge cases

The plan says "Split puzzle on `==`" and "Split LHS on `+`", but does not mention handling whitespace. The test inputs have spaces around operators (e.g., `"I + BB == ILL"`). The parsing needs to handle:
- Spaces around `+` and `==`
- Consistent trimming of word strings

**Recommendation:** Use `strings.Fields()` to tokenize and filter out `+` and `==` tokens (similar to the reference implementation). This is more robust than splitting on `==` and `+`.

### Gap 2: No explicit pruning strategy described

As noted in section 3, the plan says "check if the partial assignment is still consistent" but does not define what consistency checks are performed beyond uniqueness and leading zeros. The coefficient-based partial sum bounds should be explicitly described.

### Gap 3: No discussion of letter assignment ordering

The order in which letters are tried during backtracking has a massive impact on performance. The plan should specify that letters are sorted by descending absolute coefficient value.

### Gap 4: Error handling specifics

The plan mentions returning an error but does not specify what error message to use. The test file checks `errorExpected bool`, so any non-nil error suffices, but this should be noted.

### Gap 5: No discussion of the test harness expectations

The test file expects `map[string]int` where keys are single uppercase letters (e.g., `"A"`, `"B"`). The plan uses `map[rune]int` internally and `map[rune]bool` for leading letters, but needs to ensure the final conversion to `map[string]int` is correct. This conversion step is not explicitly mentioned.

### Gap 6: Coefficient computation sign convention

The plan describes the coefficient convention (positive for addends, negative for result), which is correct. However, it should note that when a letter appears in multiple positions across multiple words (e.g., "E" in SEND+MORE==MONEY appears at position 2 in SEND, positions 3 and 0 in MORE, and positions 3 and 1 in MONEY), all contributions must be accumulated. The plan implies this but should be more explicit.

## 5. Specific Recommendations

### Recommendation 1: Add partial-sum bounding (critical for performance)

During backtracking, after assigning k letters, compute the partial weighted sum. For each remaining unassigned letter, compute the min and max contribution using the available digit pool. If the interval `[partial + min_remaining, partial + max_remaining]` does not contain zero, prune immediately.

```go
// After assigning some letters, check if zero is reachable
func canReachZero(partialSum int, remainingCoeffs []int, availableDigits []int) bool {
    // Sort remaining coefficients and available digits
    // Assign smallest digits to most-negative coefficients and largest to most-positive (for min)
    // Vice versa for max
    // Check if min <= 0 <= max
}
```

### Recommendation 2: Sort letters by descending absolute coefficient

```go
sort.Slice(letters, func(i, j int) bool {
    return abs(coefficients[letters[i]]) > abs(coefficients[letters[j]])
})
```

This causes the largest-impact letters to be assigned first, making bounds tighter earlier and pruning more of the tree.

### Recommendation 3: Use `strings.Fields()` for parsing

This handles arbitrary whitespace and makes it trivial to skip operator tokens:

```go
tokens := strings.Fields(puzzle)
for _, token := range tokens {
    if token == "+" || token == "==" {
        continue
    }
    // token is a word
}
```

The last non-operator token is the result word; all others are addends.

### Recommendation 4: Internal representation

Use byte-indexed arrays (e.g., `[26]int` for coefficients indexed by `letter - 'A'`) rather than maps for performance. The reference implementation does this and it avoids map overhead in the hot path.

### Recommendation 5: Validate that the approach handles the "A == B" case

With the coefficient approach: A gets coefficient +1 (addend, ones place), B gets coefficient -1 (result, ones place). The equation becomes `A - B = 0`, meaning A == B. But since each digit must be unique, there is no solution. The backtracking solver correctly rejects all assignments where A == B because of the uniqueness constraint, and returns an error. This is correct.

### Recommendation 6: Consider whether a simpler implementation suffices

The reference implementation, despite being brute-force, does work -- the test suite presumably passes with it. However, generating 3.6M permutations upfront in Go consumes significant memory and is slow. The plan's backtracking approach is a clear improvement and worth pursuing. But the implementation should be kept as simple as possible while incorporating the key optimizations (coefficient reduction, backtracking, partial-sum bounding, coefficient-magnitude ordering).

## Verdict

**The plan is fundamentally sound but needs more detail on pruning strategies to guarantee adequate performance on the 10-letter test case.** The coefficient-based reduction is the right approach. The main risks are:

1. Without partial-sum bounding and letter ordering, the 10-letter case may time out
2. Parsing needs to be robust to whitespace (use `strings.Fields`)
3. The final `map[string]int` conversion must match test expectations

With the recommended additions (partial-sum bounding, coefficient-magnitude ordering, `strings.Fields` parsing), the plan should produce a correct and performant solution for all test cases.
