# Implementation Plan: Book Store Cost Function

## Branch 1: Greedy with 5→3 Pair Optimization

**Approach:** Use a greedy algorithm that forms groups from largest to smallest, then apply a well-known optimization: convert pairs of (5-group + 3-group) into pairs of (4-group + 4-group) since 2×4-group is cheaper.

**Rationale:** This is the classic approach for this problem. The greedy algorithm on its own isn't optimal, but the only case where greedy fails is when a 5+3 split should be 4+4 instead. This fix handles all known edge cases.

**Steps:**
1. Count frequency of each book (books 1-5)
2. Sort frequencies in descending order
3. Greedily form groups: repeatedly take one copy from each non-zero frequency (forms the largest possible group), decrement counts
4. This produces a list of group sizes (e.g., [5, 3] or [4, 4, 2])
5. Count groups of size 5 and size 3. Convert min(count5, count3) pairs from (5,3) into (4,4)
6. Calculate total cost using the discount table

**Files to modify:** `book_store.go` only

**Evaluation:**
- Feasibility: High — straightforward algorithm, no external deps
- Risk: Low — well-known optimization, covers all edge cases in the test suite
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~30-40 lines of Go code in one file

## Branch 2: Dynamic Programming / Memoized Search

**Approach:** Model the problem as an optimization over all possible ways to partition books into valid discount groups. Use dynamic programming with the book frequency counts as state.

**Steps:**
1. Count frequency of each book (books 1-5)
2. Sort frequencies to normalize state (reduce memoization space)
3. Use recursive function with memoization: at each step, try forming a group of size 1, 2, 3, 4, or 5 (picking one book from each of the top-N most frequent titles)
4. For each choice, compute cost of that group + recursively solve the remaining books
5. Return the minimum cost across all choices
6. Memoize on the sorted frequency tuple

**Files to modify:** `book_store.go` only

**Evaluation:**
- Feasibility: High — works correctly by exhaustive search
- Risk: Medium — more code, potential performance concerns for very large baskets, though test cases are small
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~50-60 lines, more complex logic with memoization map

## Branch 3: Frequency-Based Direct Computation

**Approach:** Directly compute the optimal grouping analytically from the frequency distribution. After counting and sorting frequencies, compute group sizes layer by layer (like a histogram), then apply the 5→3 pair fix.

This is essentially the same as Branch 1 but computed differently — instead of iteratively peeling off groups, compute the number of groups of each size directly from the sorted frequency array using differences between adjacent frequencies.

**Steps:**
1. Count frequency of each book, sort descending to get freq[0] >= freq[1] >= ... >= freq[4]
2. Compute layers: the number of groups of size >= k is freq[k-1] (for k=1..5)
3. Groups of exactly size k = (groups of size >= k) - (groups of size >= k+1)
4. Apply 5↔3 → 4+4 optimization
5. Sum costs

**Files to modify:** `book_store.go` only

**Evaluation:**
- Feasibility: High — mathematically sound
- Risk: Low-Medium — the layer computation is slightly tricky to get right
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~25-35 lines, compact but less intuitive

---

## Selected Plan

**Selected: Branch 1 — Greedy with 5→3 Pair Optimization**

**Rationale:** Branch 1 is the simplest, most readable, and most well-established solution to this problem. It's easy to verify correctness, has no performance concerns, requires minimal code, and the 5→3→4+4 optimization is a well-known technique for this exact exercise. Branch 2 is over-engineered for the problem size. Branch 3 is mathematically equivalent to Branch 1 but harder to understand.

### Detailed Implementation

**File: `go/exercises/practice/book-store/book_store.go`**

```go
package bookstore

import "sort"

func Cost(books []int) int {
    // Count frequency of each book
    freq := make([]int, 5)
    for _, b := range books {
        freq[b-1]++
    }

    // Sort frequencies descending
    sort.Sort(sort.Reverse(sort.IntSlice(freq)))

    // Greedily form groups (largest possible each time)
    var groups []int
    for freq[0] > 0 {
        size := 0
        for i := 0; i < 5; i++ {
            if freq[i] > 0 {
                freq[i]--
                size++
            }
        }
        groups = append(groups, size)
    }

    // Optimize: convert (5+3) pairs into (4+4) pairs
    fives, threes := 0, 0
    for _, g := range groups {
        switch g {
        case 5:
            fives++
        case 3:
            threes++
        }
    }
    convert := min(fives, threes)
    // Rebuild groups with the conversion applied
    // Remove `convert` fives and `convert` threes, add 2*convert fours
    // But simpler: just compute cost directly
    cost := 0
    for _, g := range groups {
        cost += groupCost(g)
    }
    // Each (5+3)→(4+4) conversion saves: (groupCost(5)+groupCost(3)) - 2*groupCost(4)
    // = (3000 + 2160) - 2*2560 = 5160 - 5120 = 40
    cost -= convert * 40

    return cost
}

func groupCost(n int) int {
    switch n {
    case 1: return 800
    case 2: return 1520
    case 3: return 2160
    case 4: return 2560
    case 5: return 3000
    default: return 0
    }
}
```

**Key design decisions:**
1. Use a frequency array of size 5 (one per book title)
2. Sort descending to greedily form largest groups first
3. The 5+3→4+4 optimization is applied as a cost adjustment (saves 40 cents per conversion) rather than modifying the groups array
4. The `groupCost` function encodes the discount table directly

**Verification:**
- Run `go test ./...` in the book-store directory
- Run `go vet ./...` in the book-store directory
- All 18 test cases must pass
