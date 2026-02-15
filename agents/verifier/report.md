# Verification Report: book-store Exercise

**Date:** 2026-02-15
**Verifier:** verifier agent
**Exercise:** book-store (Go)
**Verdict:** **PASS**

---

## Acceptance Criteria Verification

### 1. All 18 test cases in `cases_test.go` pass

**PASS** - All 18 test cases pass. Independently confirmed by running `go test -v -count=1 ./...` (uncached). Every subtest under `TestCost` reported PASS:

| # | Test Case | Result |
|---|-----------|--------|
| 1 | Only_a_single_book | PASS |
| 2 | Two_of_the_same_book | PASS |
| 3 | Empty_basket | PASS |
| 4 | Two_different_books | PASS |
| 5 | Three_different_books | PASS |
| 6 | Four_different_books | PASS |
| 7 | Five_different_books | PASS |
| 8 | Two_groups_of_four_is_cheaper_than_group_of_five_plus_group_of_three | PASS |
| 9 | Two_groups_of_four_is_cheaper_than_groups_of_five_and_three | PASS |
| 10 | Group_of_four_plus_group_of_two_is_cheaper_than_two_groups_of_three | PASS |
| 11 | Two_each_of_first_four_books_and_one_copy_each_of_rest | PASS |
| 12 | Two_copies_of_each_book | PASS |
| 13 | Three_copies_of_first_book_and_two_each_of_remaining | PASS |
| 14 | Three_each_of_first_two_books_and_two_each_of_remaining_books | PASS |
| 15 | Four_groups_of_four_are_cheaper_than_two_groups_each_of_five_and_three | PASS |
| 16 | Check_that_groups_of_four_are_created_properly_even_when_there_are_more_groups_of_three_than_groups_of_five | PASS |
| 17 | One_group_of_one_and_four_is_cheaper_than_one_group_of_two_and_three | PASS |
| 18 | One_group_of_one_and_two_plus_three_groups_of_four_is_cheaper_than_one_group_of_each_size | PASS |

### 2. TestCost passes

**PASS** - `TestCost` passes with all 18 subtests succeeding in 0.005s.

### 3. BenchmarkCost runs without error

**PASS** - `BenchmarkCost` runs successfully. Result: 4,509 iterations at 282,251 ns/op (~282 us/op).

### 4. Solution is in `book_store.go` within package `bookstore`

**PASS** - File `go/exercises/practice/book-store/book_store.go` exists, declares `package bookstore`, and exports `func Cost(books []int) int`.

### 5. Edge cases handled

**PASS** - All edge cases verified through test results:
- **Empty basket** (`[]int{}`) → returns `0` (test case #3)
- **Single book** (`[]int{1}`) → returns `800` (test case #1)
- **All same books** (`[]int{2, 2}`) → returns `1600` (800 x 2, no discount) (test case #2)
- **All different books** (`[]int{1, 2, 3, 4, 5}`) → returns `3000` (25% discount) (test case #7)

### 6. Function finds globally optimal grouping (not just greedy)

**PASS** - The implementation uses memoized dynamic programming (`minCost` with a `memo map[[5]int]int`), exhaustively exploring all possible group sizes at each step and selecting the minimum cost. This guarantees a globally optimal result. Key test cases that specifically verify this:

- Test #8: Two groups of 4 ($51.20) is cheaper than group of 5 + group of 3 ($51.60) → correctly returns 5120
- Test #15: Four groups of 4 cheaper than two groups each of 5 and 3 → correctly returns 10240
- Test #17: One group of 1 and 4 cheaper than one group of 2 and 3 → correctly returns 3360
- Test #18: One group of 1+2 plus three groups of 4 cheaper than one group of each size → correctly returns 10000

---

## Implementation Review

The solution in `book_store.go` (56 lines) is clean and correct:

- **Algorithm:** Memoized recursion over frequency vectors (dynamic programming)
- **Frequency counting:** Counts occurrences of each book (1-5) into a `[5]int` array
- **Normalization:** Sorts frequencies in descending order for canonical memoization keys
- **Exploration:** For each state, tries all possible group sizes (1 to `distinct`) and picks the minimum
- **Discount table:** Clean lookup via `var discounts = [6]int{0, 0, 5, 10, 20, 25}`
- **Cost calculation:** `groupCost(n) = n * 800 * (100 - discounts[n]) / 100` — integer arithmetic only

No issues found. No code smells or correctness concerns.

---

## Summary

| Criterion | Verdict |
|-----------|---------|
| 18/18 test cases pass | PASS |
| TestCost passes | PASS |
| BenchmarkCost runs | PASS |
| Package & file correct | PASS |
| Edge cases handled | PASS |
| Globally optimal grouping | PASS |

**Overall Verdict: PASS**

All acceptance criteria from GOAL.md are fully met. The implementation is correct, complete, and ready.
