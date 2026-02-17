# Implementation Plan: Book Store

## Proposal A — Greedy with 5→3 to 4→4 Optimization

**Role: Proponent**

### Approach

Use a greedy algorithm that builds groups from largest to smallest, then applies a post-processing optimization step that converts pairs of (5-group + 3-group) into (4-group + 4-group) since 2×4-group cost (5120) < 5-group + 3-group cost (5160).

### Algorithm

1. Count frequency of each book (books 1-5)
2. Sort frequencies in descending order
3. Greedily form groups: repeatedly take one copy from each non-zero frequency (forming a group of size = number of non-zero frequencies), until all books are allocated
4. Count the resulting group sizes (how many groups of size 5, 4, 3, 2, 1)
5. **Optimization**: While there are groups of size 5 AND groups of size 3, convert one 5-group + one 3-group into two 4-groups
6. Compute total cost using the discount table

### Files to Modify

- `go/exercises/practice/book-store/book_store.go` — implement `Cost` function

### Rationale

This is the classic known-optimal solution for this problem. The greedy approach naturally produces the best groupings, and the only non-obvious optimization is the 5+3 → 4+4 conversion. This approach is O(n) where n is the number of books, highly efficient, and easy to understand.

### Implementation

```go
func Cost(books []int) int {
    // Count frequencies
    freq := make([]int, 6) // index 1-5
    for _, b := range books {
        freq[b]++
    }

    // Sort frequencies descending
    sort.Sort(sort.Reverse(sort.IntSlice(freq[1:])))

    // Greedy grouping: form groups by taking one from each non-zero freq
    var groups []int
    for {
        size := 0
        for i := 1; i <= 5; i++ {
            if freq[i] > 0 {
                freq[i]--
                size++
            }
        }
        if size == 0 {
            break
        }
        groups = append(groups, size)
        // Re-sort after each group
        sort.Sort(sort.Reverse(sort.IntSlice(freq[1:])))
    }

    // Count group sizes
    groupCount := make([]int, 6) // index = group size
    for _, g := range groups {
        groupCount[g]++
    }

    // Optimize: convert 5+3 pairs into 4+4
    pairs := min(groupCount[5], groupCount[3])
    groupCount[5] -= pairs
    groupCount[3] -= pairs
    groupCount[4] += 2 * pairs

    // Calculate cost
    discount := [6]int{0, 0, 5, 10, 20, 25}
    total := 0
    for size := 1; size <= 5; size++ {
        pricePerBook := 800 * (100 - discount[size]) / 100
        total += groupCount[size] * size * pricePerBook
    }
    return total
}
```

---

## Proposal B — Dynamic Programming / Exhaustive Search

**Role: Opponent**

### Approach

Use dynamic programming on the frequency state to find the true optimal cost. Represent the state as the sorted tuple of frequencies of each book title, and memoize the minimum cost for each state.

### Algorithm

1. Count frequency of each book
2. Use recursive DP: at each state (sorted frequency vector), try forming a group of each possible size (1 through 5) by choosing one book from the appropriate number of distinct titles
3. For each group choice, compute cost of that group + recursive cost of remaining books
4. Memoize and return the minimum

### Files to Modify

- `go/exercises/practice/book-store/book_store.go` — implement `Cost` function

### Critique of Proposal A

Proposal A relies on the specific insight that the only optimization needed is 5+3 → 4+4. While this is correct for the standard discount schedule, it's a hardcoded heuristic. If the discount schedule were different, the greedy+fix approach might miss other optimization opportunities.

### Rationale

DP guarantees finding the true optimum regardless of the discount schedule. It's more general and mathematically rigorous.

### Downsides

- Significantly more complex to implement
- The state space can grow large (though bounded for this problem)
- Over-engineered for a problem where the greedy+fix is known to be correct
- Harder to read and understand

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Greedy+Fix) | Proposal B (DP) |
|-----------|------------------------|-----------------|
| **Correctness** | Proven correct for this discount schedule | Correct by construction |
| **Risk** | Low — well-known solution | Medium — complex implementation, potential bugs |
| **Simplicity** | Simple, ~30 lines | Complex, 50+ lines with memoization |
| **Consistency** | Matches codebase style (concise, pragmatic) | Over-engineered for the exercise |

### Decision

**Proposal A wins.** The greedy approach with 5+3 → 4+4 optimization is the standard, well-tested solution for this exact problem. It's simpler, faster, and less error-prone than the DP approach. The DP approach is over-engineered for a problem with a known, elegant solution.

### Final Implementation Plan

**File to modify:** `go/exercises/practice/book-store/book_store.go`

**Algorithm:**

1. Handle empty basket → return 0
2. Count frequency of each book ID (1-5) into an array
3. Greedily form groups: repeatedly sort frequencies descending, then form a group by taking one copy from each non-zero frequency position
4. Count resulting groups by size
5. Optimize: convert min(count5, count3) pairs of (5-group, 3-group) into (4-group, 4-group) pairs
6. Compute total: for each group size, multiply count × size × discounted price per book

**Imports needed:** `sort`

**No other files modified.**
