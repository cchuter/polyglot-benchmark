# Plan Review (Codex)

## Reviewer: Codex (gpt-5.3-codex)

## Summary

Algorithm is **correct** — greedy layering plus 5+3→4+4 correction is the known optimal strategy for these exact discounts and 5 titles.

## Findings

### 1. Algorithm Correctness — PASS
- Walked against tricky test cases (49, 84, 89, 99); resulting groupings match expected totals
- Should pass all 18 test cases

### 2. Critical Issue: Go 1.18 Compilation
- **`min()` is NOT available in Go 1.18** — the builtin `min` was added in Go 1.21
- Plan code at line 96 uses `min(fives, threes)` which will fail to compile
- **Fix required**: replace with manual min function or inline comparison

### 3. Edge Cases
- Invalid book IDs (not 1..5): could build group sizes >5 and panic on `groupCosts[g]` index out of bounds
- Not a concern for the test suite (all tests use valid IDs 1-5), but worth noting
- nil/empty basket handled correctly (returns 0)

### 4. Optimization Sufficiency — CONFIRMED
- 5+3→4+4 is the **only** beneficial swap for this discount table
- Other combinations verified as not beneficial:
  - c5 + c2 = 4520 vs c4 + c3 = 4720 (worse)
  - c4 + c2 = 4080 vs c3 + c3 = 4320 (worse)

### 5. Minor Suggestions
- `sort` import unnecessary — each greedy round just decrements all positive counts
- Fixed `[5]int` array simpler than `map[int]int` for constrained domain
- Group-size counting (not full rebuild) simpler for swap logic

## Verdict

**Approved with required fix**: Replace `min()` builtin with manual implementation for Go 1.18 compatibility. Algorithm is correct and complete.
