# Implementation Plan: Book Store Cost Function

## Overview

Implement `Cost(books []int) int` in `go/exercises/practice/book-store/book_store.go`. The function must find the optimal grouping of books into sets of distinct titles to minimize total cost after discounts.

## Algorithm Design

### Approach: Frequency-based greedy with 5→3 group splitting optimization

The key insight is that a greedy approach (always forming the largest possible group) fails because two groups of 4 ($25.60 × 2 = $51.20) is cheaper than one group of 5 + one group of 3 ($30.00 + $21.60 = $51.60). The discount jump from 3→4 (10%→20%) is larger than from 4→5 (20%→25%).

**Algorithm:**

1. Count the frequency of each book title in the basket
2. Build groups greedily: repeatedly take one copy of each available distinct title to form groups from largest to smallest
3. Apply the 5+3 → 4+4 optimization: whenever there's a group of 5 and a group of 3, replace them with two groups of 4 (since 2×groupCost(4) < groupCost(5) + groupCost(3))
4. Sum the costs of all groups using the discount tiers

**Why this works:**
- The greedy step produces groups sorted by size (largest first)
- The only suboptimal case in greedy grouping is 5+3 vs 4+4
- All other group size pairs are already optimal under greedy
- The 5+3→4+4 transformation is the only correction needed (confirmed by codex review)

**Discount tiers (per book in group):**
| Group Size | Discount | Per Book | Group Cost |
|------------|----------|----------|------------|
| 1          | 0%       | 800      | 800        |
| 2          | 5%       | 760      | 1520       |
| 3          | 10%      | 720      | 2160       |
| 4          | 20%      | 640      | 2560       |
| 5          | 25%      | 600      | 3000       |

**Verification: 5+3 vs 4+4:**
- 5+3: 3000 + 2160 = 5160
- 4+4: 2560 + 2560 = 5120
- Saving: 40 cents per swap

## File Changes

### Modified: `go/exercises/practice/book-store/book_store.go`

Replace the panic stub with the full implementation:

```go
package bookstore

var groupCosts = [6]int{
    0,
    1 * 800,
    2 * 800 * 95 / 100,
    3 * 800 * 90 / 100,
    4 * 800 * 80 / 100,
    5 * 800 * 75 / 100,
}

func Cost(books []int) int {
    // Count frequency of each book
    freq := make(map[int]int)
    for _, b := range books {
        freq[b]++
    }

    // Extract frequencies into a slice
    counts := make([]int, 0, len(freq))
    for _, c := range freq {
        counts = append(counts, c)
    }

    // Build groups greedily: each round, take one from each non-zero count
    var groups []int
    for {
        size := 0
        for i := range counts {
            if counts[i] > 0 {
                counts[i]--
                size++
            }
        }
        if size == 0 {
            break
        }
        groups = append(groups, size)
    }

    // Optimize: replace 5+3 pairs with 4+4
    fives, threes := 0, 0
    for _, g := range groups {
        if g == 5 {
            fives++
        }
        if g == 3 {
            threes++
        }
    }
    swaps := fives
    if threes < swaps {
        swaps = threes
    }
    if swaps > 0 {
        newGroups := make([]int, 0, len(groups))
        fivesLeft, threesLeft := swaps, swaps
        for _, g := range groups {
            if g == 5 && fivesLeft > 0 {
                newGroups = append(newGroups, 4)
                fivesLeft--
            } else if g == 3 && threesLeft > 0 {
                newGroups = append(newGroups, 4)
                threesLeft--
            } else {
                newGroups = append(newGroups, g)
            }
        }
        groups = newGroups
    }

    // Sum costs
    total := 0
    for _, g := range groups {
        total += groupCosts[g]
    }
    return total
}
```

**Key changes from initial draft (per codex review):**
- Removed `min()` builtin — not available in Go 1.18; replaced with manual comparison
- Removed `sort` import — unnecessary since greedy loop works on all positive counts regardless of order
- No external imports needed at all

## Ordering

1. Replace `book_store.go` with the implementation
2. Run `go test` to verify all 18 tests pass
3. Run `go test -bench=.` to verify benchmark runs
4. Commit with descriptive message

## Rationale

- **Frequency-based greedy + optimization** is simpler and more efficient than the recursive approach in example.go
- No recursion means no risk of stack overflow on large inputs
- The 5+3→4+4 swap is a well-known optimization for this specific discount structure
- O(n) time complexity where n is the number of books
- No imports needed — pure standard Go
