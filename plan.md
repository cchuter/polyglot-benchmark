# Implementation Plan: Book Store Discount Calculator

## Algorithm Design

### Approach: Greedy with 5→4 Redistribution

The key insight is that a group of 5 + a group of 3 (cost: 3000 + 2160 = 5160) is always more expensive than two groups of 4 (cost: 2560 + 2560 = 5120). This is the only non-obvious optimization needed.

**Algorithm:**

1. **Count frequencies**: Count how many copies of each book (1-5) are in the basket.
2. **Greedy grouping**: Build groups greedily from largest to smallest. Sort the frequency counts in descending order. Repeatedly take one book from each non-zero count to form groups.
3. **Redistribute 5+3 → 4+4**: After greedy grouping, count groups by size. For every pair of (group of 5, group of 3), convert them into two groups of 4 if it's cheaper.
4. **Calculate cost**: Sum the cost of all groups using the discount table.

### Discount Table (cost per group in cents)

| Group Size | Discount | Cost per book | Group cost |
|-----------|----------|--------------|------------|
| 1 | 0% | 800 | 800 |
| 2 | 5% | 760 | 1520 |
| 3 | 10% | 720 | 2160 |
| 4 | 20% | 640 | 2560 |
| 5 | 25% | 600 | 3000 |

### Why 5+3 → 4+4 is the only redistribution needed

- 5+3 costs 5160, 4+4 costs 5120 → save 40 cents (convert)
- 5+2 costs 4520, 4+3 costs 4720 → 4+3 is MORE expensive (don't convert)
- 5+1 costs 3800, 4+2 costs 4080 → don't convert
- 4+2 costs 4080, 3+3 costs 4320 → don't convert

So the ONLY beneficial redistribution is 5+3 → 4+4.

### Implementation Steps

1. Count frequency of each book number in the basket
2. Sort frequencies in descending order
3. Build groups greedily: in each pass, take one from each non-zero frequency
4. Count how many groups of size 5 and size 3 were formed
5. Convert pairs of (5, 3) into pairs of (4, 4)
6. Sum total cost using the group cost lookup

## File Changes

### `go/exercises/practice/book-store/book_store.go`

```go
package bookstore

import "sort"

// Cost calculates the total cost in cents for a basket of books,
// applying the best possible group discounts.
func Cost(books []int) int {
    // Count frequency of each book
    freq := make(map[int]int)
    for _, b := range books {
        freq[b]++
    }

    // Extract counts and sort descending
    counts := make([]int, 0, len(freq))
    for _, c := range freq {
        counts = append(counts, c)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(counts)))

    // Build groups greedily
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

    // Count groups of 5 and 3; redistribute to 4+4
    fives, threes := 0, 0
    for _, g := range groups {
        if g == 5 { fives++ }
        if g == 3 { threes++ }
    }
    redistribute := min(fives, threes)

    // Calculate cost using group cost table
    groupCost := [6]int{0, 800, 1520, 2160, 2560, 3000}
    total := 0
    for _, g := range groups {
        total += groupCost[g]
    }
    // Apply redistribution: each 5+3 → 4+4 saves 40 cents
    total -= redistribute * 40

    return total
}

func min(a, b int) int {
    if a < b { return a }
    return b
}
```

## Verification

- `go test ./...` in the book-store directory should pass all 18 test cases
- `go build` should succeed with no errors
- `go vet` should report no issues
