# Goal: Implement Book Store Discount Calculator

## Problem Statement

Implement a `Cost` function in Go that calculates the optimal price for a basket of books from a 5-book series, applying the maximum possible discount by finding the best grouping of books.

### Pricing Rules
- Single book: $8.00 (800 cents)
- 2 different books: 5% discount
- 3 different books: 10% discount
- 4 different books: 20% discount
- 5 different books: 25% discount

Books must be grouped into sets of distinct titles. The algorithm must find the grouping that minimizes total cost.

## Function Signature

```go
func Cost(books []int) int
```

- Input: `[]int` where each element is a book number (1-5)
- Output: `int` total cost in cents

## Acceptance Criteria

1. **All 18 test cases pass** in `cases_test.go` — these cover:
   - Empty basket → 0
   - Single book → 800
   - Multiple same books → 800 * count (no discount)
   - Groups of 2-5 different books with correct discounts
   - Optimal grouping: e.g., two groups of 4 is cheaper than groups of 5+3
   - Complex baskets requiring non-greedy optimization
2. **Only integer arithmetic** — no floating point
3. **Function in `book_store.go`** in package `bookstore`
4. **`go test` passes** with no errors
5. **Benchmark test runs** without issues

## Key Constraints

- The greedy approach (always making the largest group) does NOT work. For example, 8 books [1,1,2,2,3,3,4,5] optimally splits into two groups of 4 (cost 5120) rather than groups of 5+3 (cost 5160).
- The solution must handle up to ~22 books efficiently (see test case 16).
- Only integer calculations needed — all costs are in cents.
