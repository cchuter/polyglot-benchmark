# Implementation Plan: Book Store Cost Function

## File to Modify

- `go/exercises/practice/book-store/book_store.go`

## Algorithm Design

### Approach: Frequency-based Greedy with 5→3 Pair Optimization

The key insight is that a greedy algorithm (always forming the largest possible group) fails because two groups of 4 (cost: 2 × 2560 = 5120) are cheaper than one group of 5 + one group of 3 (cost: 3000 + 2160 = 5160). The difference is 40 cents per such pair.

**Strategy:**

1. Count the frequency of each book title
2. Sort frequencies in descending order to form a histogram
3. Use the frequency histogram to determine group sizes greedily (largest groups first)
4. Apply the 5+3 → 4+4 optimization: for every pair of a group-of-5 and a group-of-3, convert them into two groups-of-4

This approach works because:
- The only non-optimal greedy outcome is groups of 5 paired with groups of 3
- Converting each (5,3) pair to (4,4) saves 40 cents
- All other group-size combinations are already optimal under the greedy approach

### Detailed Steps

1. **Count frequencies**: Create a frequency map of book IDs → count
2. **Build group sizes from histogram**:
   - Sort frequencies descending: e.g., [3, 3, 2, 2, 1]
   - Determine how many groups of each size by computing the differences between sorted frequency levels
   - For frequencies [3,3,2,2,1]: groups at level 1 = 5 books (5-wide), at level 2 = 4 books (4-wide), at level 3 = 2 books (2-wide)
   - This gives: 1 group of 5, 1 group of 4, 1 group of 2
3. **Optimize**: Convert (5,3) pairs to (4,4) pairs
   - Count groups of 5 and groups of 3
   - pairs = min(count_of_5, count_of_3)
   - groups_of_5 -= pairs, groups_of_3 -= pairs, groups_of_4 += 2*pairs
4. **Calculate cost**: Sum up cost for each group size using the discount table

### Discount Table (in cents)
- 1 book: 800 (0% discount)
- 2 books: 1520 (5% discount on 1600)
- 3 books: 2160 (10% discount on 2400)
- 4 books: 2560 (20% discount on 3200)
- 5 books: 3000 (25% discount on 4000)

## Implementation Structure

```go
package bookstore

func Cost(books []int) int {
    // 1. Count frequencies
    // 2. Sort frequencies descending
    // 3. Build group counts from histogram
    // 4. Apply 5+3 → 4+4 optimization
    // 5. Sum costs
}
```

No helper functions beyond what's needed inline. Use `sort` package for sorting frequencies.

## Ordering

1. Write the complete `Cost` function in `book_store.go`
2. Run `go test` to verify all 18 test cases pass
3. Run benchmark to verify performance
