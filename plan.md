# Implementation Plan: book-store (Issue #36)

## Overview

Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go` that calculates the minimum price for a basket of books with volume discounts.

## Algorithm Approach: Frequency-based Dynamic Programming

The optimal approach works with book frequencies rather than individual books:

1. Count frequency of each book title (how many of each book 1-5)
2. Sort frequencies in descending order to get a canonical state representation
3. Use recursive search with memoization to find the minimum cost across all possible groupings

### Why This Approach

- The greedy approach (always pick largest group) fails: 5+3 grouping costs $51.60 but 4+4 costs $51.20
- We need to explore all possible group sizes at each step and pick the minimum
- Memoization on the sorted frequency tuple avoids recomputing the same states
- With only 5 book titles and typical basket sizes, the state space is small enough for this approach

### Algorithm Detail

```
Cost(books):
  1. Count frequency of each book: freq[1..5]
  2. Sort frequencies descending, trim trailing zeros
  3. Call minCost(freq) with memoization

minCost(freq):
  if all freq are 0: return 0
  for groupSize = 1 to (number of non-zero frequencies):
    // Take one book from each of the top `groupSize` frequencies
    newFreq = copy of freq
    decrement top `groupSize` entries by 1
    re-sort descending, trim zeros
    candidate = groupCost(groupSize) + minCost(newFreq)
    track minimum candidate
  return minimum

groupCost(n):
  discounts = [0, 0, 5, 10, 20, 25]  // indexed by group size
  return n * 800 * (100 - discounts[n]) / 100
```

### Discount Table

| Group Size | Discount % | Cost per book (cents) | Group cost (cents) |
|-----------|-----------|----------------------|-------------------|
| 1         | 0%        | 800                  | 800               |
| 2         | 5%        | 760                  | 1520              |
| 3         | 10%       | 720                  | 2160              |
| 4         | 20%       | 640                  | 2560              |
| 5         | 25%       | 600                  | 3000              |

## Files to Modify

1. **`go/exercises/practice/book-store/book_store.go`** — Implement `Cost` function and helpers

## Implementation Structure

```go
package bookstore

import (
    "math"
    "sort"
)

// Discount percentages indexed by group size (0-indexed, so [0]=unused, [1]=0%, ...)
var discounts = [6]int{0, 0, 5, 10, 20, 25}

// Cost calculates the minimum price for a basket of books.
func Cost(books []int) int

// minCost recursively finds minimum cost using memoization.
func minCost(freq [5]int, memo map[[5]int]int) int

// groupCost returns the cost for a group of n different books.
func groupCost(n int) int
```

**Note from review:** The `"math"` import is required for `math.MaxInt32` used in `minCost`.

## Ordering

1. Define the discount constants and `groupCost` helper
2. Implement `minCost` with memoization
3. Implement `Cost` as the entry point that counts frequencies and calls `minCost`
4. Run tests to verify all 18 cases pass
5. Run benchmark to verify performance is acceptable
