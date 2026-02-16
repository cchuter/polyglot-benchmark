# Plan Review: Book Store Discount Calculator

## Reviewer: Plan Agent (codex unavailable)

## 1. Algorithm Correctness - PASS

The recursive approach is correct. By trying all group sizes (1 to k) at each step and taking the minimum, the algorithm performs an exhaustive search that guarantees finding the optimal grouping. The frequency vector representation is the canonical approach for this problem — which specific title has which count doesn't matter, only the sorted frequency vector matters.

The critical insight is correctly identified: two groups of 4 (5120) beats group of 5 + group of 3 (5160).

## 2. Edge Cases - PASS

- **Empty basket**: `counts` is empty, returns 0 immediately
- **Single book**: `counts = [1]`, only groupSize=1 possible, returns 800
- **All same books**: `counts = [2]`, only groupSize=1 possible each step, returns 1600 for two copies
- All 18 test cases will be handled correctly

## 3. Performance - ACCEPTABLE

Without memoization, worst case is the 22-book test case with frequencies `[6,6,6,2,2]`. The reference solution also uses no memoization and follows the same recursive structure. Both will complete within test/benchmark timeouts for the given test suite.

**Suggestion**: Memoization keyed on the sorted frequency vector could be added if benchmark performance becomes an issue, but is not required.

## 4. Code Correctness - PASS (no bugs found)

- Discount table indexing: Correct (6-element array indexed by group size 0-5)
- MaxInt initialization: `int(^uint(0) >> 1)` is standard Go idiom, correct
- Frequency counting: Standard map-based counting, correct
- Trailing zero removal: Correct (zeros at end after descending sort)
- Group formation: Correctly copies, decrements top k, re-sorts, recurses
- Integer arithmetic: All divisions by 100 are exact for given discount values
- No out-of-bounds access possible (groupSize ≤ len(counts) ≤ 5, discounts array has 6 elements)

## 5. Comparison with Reference

The plan's frequency-vector approach is cleaner than the reference's raw-array approach. Both share the same performance characteristics. The plan's code is more straightforward.

## Verdict: APPROVED - Ready for implementation

No blocking issues. Plan is sound, correct, and will pass all 18 test cases.
