# Implementation Plan: Book Store Discount Calculator

## Overview

Implement the `Cost` function in `book_store.go` to calculate the minimum cost for a basket of books with volume discounts. The key challenge is finding the optimal grouping since a greedy approach (always forming the largest group) doesn't always produce the cheapest result.

## Algorithm

The approach uses a frequency-based dynamic programming / recursive optimization:

1. **Count frequencies**: Count how many copies of each distinct book title are in the basket
2. **Sort frequencies in descending order**: e.g., [3, 2, 2, 1, 1] meaning one title appears 3 times, two titles appear 2 times, etc.
3. **Recursively try all group sizes**: At each step, form a group of size `k` (from 1 to the number of titles with remaining copies), deduct one copy from each of the `k` most-frequent titles, calculate the group cost, and recurse on the remaining frequencies
4. **Return the minimum cost** across all possible group sizes at each step

### Why This Works

- The frequency representation captures all the information needed (which specific title is in which group doesn't matter, only the counts)
- By trying all group sizes at each level and taking the minimum, we find the globally optimal grouping
- The key optimization insight: two groups of 4 (2 × 2560 = 5120) is cheaper than one group of 5 + one group of 3 (3000 + 2160 = 5160)

### Cost Table (integer cents)

| Group Size | Discount | Cost per book | Group Cost |
|-----------|----------|--------------|------------|
| 1 | 0% | 800 | 800 |
| 2 | 5% | 760 | 1520 |
| 3 | 10% | 720 | 2160 |
| 4 | 20% | 640 | 2560 |
| 5 | 25% | 600 | 3000 |

## File to Modify

- `go/exercises/practice/book-store/book_store.go`

## Implementation Steps

1. Replace the stub `package bookstore` with the full implementation
2. Define constants: `bookPrice = 800`, discount table
3. Implement helper: `groupCost(size int) int` — returns cost of a group of `size` different books
4. Implement `Cost(books []int) int`:
   - Count frequency of each book
   - Sort frequencies descending
   - Recursively find minimum cost by trying all possible group sizes
   - Return the minimum

## Code Structure

```go
package bookstore

import "sort"

const bookPrice = 800

var discounts = [6]int{0, 0, 5, 10, 20, 25} // indexed by group size

func Cost(books []int) int {
    // Count frequencies
    freq := make(map[int]int)
    for _, b := range books {
        freq[b]++
    }
    // Extract and sort frequencies descending
    counts := make([]int, 0, len(freq))
    for _, v := range freq {
        counts = append(counts, v)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(counts)))
    return minCost(counts)
}

func minCost(counts []int) int {
    // Remove trailing zeros
    for len(counts) > 0 && counts[len(counts)-1] == 0 {
        counts = counts[:len(counts)-1]
    }
    if len(counts) == 0 {
        return 0
    }

    best := int(^uint(0) >> 1) // MaxInt
    for groupSize := 1; groupSize <= len(counts); groupSize++ {
        // Form a group of groupSize by taking one from each of the top groupSize frequencies
        next := make([]int, len(counts))
        copy(next, counts)
        for i := 0; i < groupSize; i++ {
            next[i]--
        }
        sort.Sort(sort.Reverse(sort.IntSlice(next)))
        cost := groupCost(groupSize) + minCost(next)
        if cost < best {
            best = cost
        }
    }
    return best
}

func groupCost(size int) int {
    return bookPrice * size * (100 - discounts[size]) / 100
}
```

## Testing

- Run `go test` in `go/exercises/practice/book-store/`
- All 18 test cases must pass
- Benchmark test should also run successfully
