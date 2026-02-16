# Plan Review (Codex)

## Verdict: APPROVED — Algorithm is correct for all 18 test cases

## Summary

The proposed algorithm (frequency histogram with 5+3→4+4 optimization) was manually traced through all 18 test cases and produces the correct expected output in every case.

## Key Findings

1. **Histogram layer-peeling decomposition** is mathematically optimal for the greedy step. It naturally produces group sizes favoring wider groups (better discounts).

2. **The 5+3→4+4 optimization is the only correction needed.** The cost per book for each group size is 800, 760, 720, 640, 600 cents. The only pair where combining yields a cheaper result is (5,3)→(4,4), saving 40 cents.

3. **All edge cases handled**: empty basket, single books, all identical books, complex optimization cases.

## Implementation Notes

- Ensure only non-zero frequency books are included in the sorted frequency array
- Be careful with off-by-one in histogram indexing: `freq[w-1] - freq[w]` with `freq[n] = 0`
- Integer arithmetic only — pre-computed discount table in cents
- The `sort` package is fine for sorting at most 5 frequency values

## Tricky Test Cases Verified

- `[1,1,2,2,3,3,4,5]` → 5120 (2×4 beats 5+3)
- `[1,1,2,2,3,3,4,5,1,1,2,2,3,3,4,5]` → 10240 (4×4 beats 2×5+2×3)
- `[1,1,1,1,1,1,2,2,2,2,2,2,3,3,3,3,3,3,4,4,5,5]` → 14560 (4×4+2×3)
- `[1,2,2,3,3,3,4,4,4,4,5,5,5,5,5]` → 10000 (3×4+1×2+1×1)
