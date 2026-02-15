# Challenger Review: book-store Cost Function

## Review Date: 2026-02-15

## File Reviewed: `go/exercises/practice/book-store/book_store.go`

---

## 1. Algorithm Correctness

**Verdict: CORRECT**

The implementation uses a greedy grouping approach with a 5+3 → 4+4 optimization. This produces optimal results for the standard 5-title pricing model.

- **Greedy grouping**: Repeatedly takes one copy of each distinct title to form groups of decreasing size. This is correct because it maximizes group sizes greedily.
- **5+3 → 4+4 optimization**: The only suboptimal case in greedy grouping is when a group of 5 and a group of 3 could be replaced by two groups of 4 (cost 5120 vs 5160, saving 40 per swap). The code correctly identifies and performs this transformation.
- **Empirical verification by codex**: Exhaustive DP comparison over all count vectors `0..10` for each of 5 titles (161,051 states) produced **0 mismatches** vs the optimal solution.

### Manual Trace Verification

- **Test case 15** (`[1,1,2,2,3,3,4,5,1,1,2,2,3,3,4,5]`): Groups=[5,5,3,3] → swap 2 pairs → [4,4,4,4] = 10240 ✓
- **Test case 16** (`[1×6,2×6,3×6,4×2,5×2]`): Groups=[5,5,3,3,3,3] → swap 2 pairs → [4,4,4,4,3,3] = 14560 ✓
- **Test case 18** (`[1,2×2,3×3,4×4,5×5]`): Groups=[5,4,3,2,1] → swap 1 pair → [4,4,4,2,1] = 10000 ✓

## 2. Edge Cases

**Verdict: ALL HANDLED**

| Edge Case | Behavior | Result |
|-----------|----------|--------|
| Empty basket `[]int{}` | `freq` empty → `counts` empty → loop breaks immediately → total=0 | ✓ |
| Nil input `nil` | `range nil` is no-op → same as empty basket | ✓ |
| Single book `[]int{1}` | One group of size 1 → cost=800 | ✓ |
| All same books `[]int{2,2}` | Each round takes 1 → groups of size 1 → n×800 | ✓ |

## 3. Go 1.18 Compatibility

**Verdict: COMPATIBLE**

- No use of `min`/`max` builtins (introduced in Go 1.21) — manual comparison used instead
- No generics usage
- No external imports — pure standard Go
- Integer arithmetic avoids floating point entirely

## 4. Bugs or Issues

**Verdict: NO BUGS FOUND**

### Minor Note (Low Severity)
- If more than 5 distinct book IDs were passed (outside the problem specification), `groupCosts[g]` would panic with an index-out-of-range error since the array only has 6 entries (indices 0-5). This is acceptable for Exercism's domain where books are always in range 1-5.

### groupCosts Array Verification
All values compute exactly with no integer rounding issues:
- `groupCosts[0]` = 0
- `groupCosts[1]` = 800
- `groupCosts[2]` = 1520 (2 × 800 × 95 / 100)
- `groupCosts[3]` = 2160 (3 × 800 × 90 / 100)
- `groupCosts[4]` = 2560 (4 × 800 × 80 / 100)
- `groupCosts[5]` = 3000 (5 × 800 × 75 / 100)

## 5. Test Results

**All 18 test cases PASS**

```
=== RUN   TestCost
--- PASS: TestCost (0.00s)
    --- PASS: TestCost/Only_a_single_book
    --- PASS: TestCost/Two_of_the_same_book
    --- PASS: TestCost/Empty_basket
    --- PASS: TestCost/Two_different_books
    --- PASS: TestCost/Three_different_books
    --- PASS: TestCost/Four_different_books
    --- PASS: TestCost/Five_different_books
    --- PASS: TestCost/Two_groups_of_four_is_cheaper_than_group_of_five_plus_group_of_three
    --- PASS: TestCost/Two_groups_of_four_is_cheaper_than_groups_of_five_and_three
    --- PASS: TestCost/Group_of_four_plus_group_of_two_is_cheaper_than_two_groups_of_three
    --- PASS: TestCost/Two_each_of_first_four_books_and_one_copy_each_of_rest
    --- PASS: TestCost/Two_copies_of_each_book
    --- PASS: TestCost/Three_copies_of_first_book_and_two_each_of_remaining
    --- PASS: TestCost/Three_each_of_first_two_books_and_two_each_of_remaining_books
    --- PASS: TestCost/Four_groups_of_four_are_cheaper_than_two_groups_each_of_five_and_three
    --- PASS: TestCost/Check_that_groups_of_four_are_created_properly_even_when_there_are_more_groups_of_three_than_groups_of_five
    --- PASS: TestCost/One_group_of_one_and_four_is_cheaper_than_one_group_of_two_and_three
    --- PASS: TestCost/One_group_of_one_and_two_plus_three_groups_of_four_is_cheaper_than_one_group_of_each_size
PASS
ok  	bookstore	0.004s
```

## 6. Plan Adherence

The implementation matches `.osmi/plan.md` exactly:
- Same algorithm (greedy + 5+3→4+4 optimization)
- Same `groupCosts` array
- Same code structure (frequency counting → greedy grouping → optimization → summation)
- Removed `min()` builtin per plan
- Removed `sort` import per plan
- No external imports per plan

## Summary

**APPROVED** — The implementation is correct, handles all edge cases, is Go 1.18 compatible, and passes all 18 test cases. Codex confirmed optimality via exhaustive DP comparison across 161,051 input states with 0 mismatches.
