# Implementation Plan: Book Store Cost Function

## Branch 1: Greedy + Adjustment (Simplest)

**Approach**: Use a greedy algorithm to form the largest possible groups of distinct books, then apply a known optimization: convert pairs of (group-of-5, group-of-3) into two groups-of-4, since 2x4 is cheaper than 5+3.

**Files to modify**: `go/exercises/practice/book-store/book_store.go`

**Algorithm**:
1. Count frequency of each book title
2. Greedily form groups from largest to smallest: repeatedly take one copy from each title that still has copies remaining, forming groups of distinct books
3. This produces groups sorted by size (largest first, smallest last)
4. Post-process: while there exists both a group-of-5 and a group-of-3, convert them to two groups-of-4
5. Compute cost by summing `groupCost(size)` for each group

**Rationale**: The only case where greedy is suboptimal is the 5+3 → 4+4 conversion. This is a well-known property of this specific discount structure.

**Evaluation**:
- Feasibility: High - straightforward to implement, no recursion needed
- Risk: Low - the 5+3→4+4 adjustment is the only correction needed for this discount structure
- Alignment: Fully satisfies all test cases
- Complexity: ~30-40 lines of code, single file, O(n) time

## Branch 2: Recursive with Memoization (Reference-style)

**Approach**: Mirror the `.meta/example.go` approach - recursively try all possible group sizes at each step, memoize results. Sort books by frequency first to improve cache hits.

**Files to modify**: `go/exercises/practice/book-store/book_store.go`

**Algorithm**:
1. Organize books by frequency (most frequent first)
2. At each recursive step, extract distinct books and try forming groups of size 1, 2, ..., up to the number of distinct books
3. Recursively solve the remaining books after removing each group
4. Return the minimum cost across all choices
5. Use memoization keyed on sorted frequency counts

**Rationale**: This is the general-purpose approach that works for any discount structure. It's directly based on the example solution.

**Evaluation**:
- Feasibility: High - reference solution exists
- Risk: Medium - recursive approach can be slow without proper memoization; the example.go doesn't actually memoize and relies on the organize step for efficiency
- Alignment: Fully satisfies all test cases
- Complexity: ~60-80 lines, more complex logic, potential performance concerns for large baskets

## Branch 3: Frequency-Count Dynamic Programming

**Approach**: Model the problem by frequency counts rather than individual books. Count how many titles appear 1 time, 2 times, etc. Use DP over the frequency histogram.

**Files to modify**: `go/exercises/practice/book-store/book_store.go`

**Algorithm**:
1. Count frequency of each book (5 possible titles)
2. Sort frequencies in descending order to create a canonical representation
3. Use recursive DP with memoization: state = tuple of sorted frequencies
4. At each state, try forming a group of size k (1 to number of non-zero frequencies)
5. Subtract 1 from the k largest frequencies, re-sort, and recurse
6. Memoize on the sorted frequency tuple

**Rationale**: The state space is very small (at most a few hundred distinct states) since there are only 5 book titles. This gives guaranteed optimality with excellent performance.

**Evaluation**:
- Feasibility: High - clean DP approach
- Risk: Low - state space is bounded, correct by construction
- Alignment: Fully satisfies all test cases
- Complexity: ~40-50 lines, moderate complexity but very efficient

---

## Evaluation Summary

| Criterion | Branch 1 (Greedy+Adj) | Branch 2 (Recursive) | Branch 3 (Frequency DP) |
|-----------|----------------------|---------------------|------------------------|
| Feasibility | High | High | High |
| Risk | Low | Medium | Low |
| Alignment | Full | Full | Full |
| Complexity | Low (~30 lines) | High (~70 lines) | Medium (~45 lines) |
| Performance | O(n) | Exponential worst-case | Bounded state space |

## Selected Plan

**Branch 1: Greedy + Adjustment** is the best choice.

**Rationale**: It is the simplest, most readable, and most efficient approach. The greedy algorithm with the 5+3→4+4 adjustment is a well-known correct solution for this specific discount structure. It avoids recursion, memoization, and complex state management. It's also the easiest to verify for correctness.

### Detailed Implementation Plan

**File**: `go/exercises/practice/book-store/book_store.go`

**Step 1**: Define constants and discount lookup
- `bookPrice = 800` (cents)
- Discount percentages indexed by group size: `[0, 0, 5, 10, 20, 25]`

**Step 2**: Implement `Cost(books []int) int`
1. If books is empty, return 0
2. Count frequency of each book title using a map or array
3. Sort frequencies in descending order
4. Greedily form groups: iterate frequencies from high to low, at each level the group size = number of titles with frequency >= current level. Equivalently, build a list of group sizes by peeling layers off the frequency histogram.
5. Count groups by size: `groups[size]` = number of groups of that size

**Step 3**: Apply the 5+3→4+4 adjustment
- `pairs = min(groups[5], groups[3])`
- `groups[5] -= pairs; groups[3] -= pairs; groups[4] += 2 * pairs`

**Step 4**: Compute total cost
- For each group size 1-5: `total += groups[size] * groupCost(size)`
- Return total

**Step 5**: Implement `groupCost(size int) int`
- `normalPrice = bookPrice * size`
- `discount = normalPrice * discountPercent[size] / 100`
- Return `normalPrice - discount`

### Expected group formation example
Books: [1,1,2,2,3,3,4,5] → frequencies: [2,2,2,1,1] sorted desc
- Layer 1: 5 titles have freq >= 1 → group of 5
- Layer 2: 3 titles have freq >= 2 → group of 3
- Groups: one 5, one 3
- Adjustment: convert to two groups of 4
- Cost: 2 * groupCost(4) = 2 * 2560 = 5120 ✓
