# Goal: Implement Book Store Discount Calculator

## Problem Statement

Implement the `Cost` function for the book-store exercise. A bookshop sells 5 different books at $8 each, with discounts for buying different titles together:

- 1 book: no discount (800 cents)
- 2 different books: 5% discount (1520 cents)
- 3 different books: 10% discount (2160 cents)
- 4 different books: 20% discount (2560 cents)
- 5 different books: 25% discount (3000 cents)

The key challenge is finding the optimal grouping of books to minimize the total cost. A greedy approach (always making the largest group) does not always yield the minimum cost. For example, two groups of 4 (5120 cents) is cheaper than a group of 5 + a group of 3 (5160 cents).

## Function Signature

```go
func Cost(books []int) int
```

- Input: a slice of integers representing book IDs (1-5)
- Output: the total cost in cents after applying the best possible discounts

## Acceptance Criteria

1. `Cost` returns 0 for an empty basket
2. `Cost` returns 800 for a single book
3. `Cost` returns 1600 for two copies of the same book (no discount)
4. `Cost` correctly applies discount tiers for 2, 3, 4, and 5 different books
5. `Cost` finds the optimal grouping to minimize total cost (e.g., two groups of 4 are cheaper than groups of 5 and 3)
6. All 18 test cases in `cases_test.go` pass
7. `go vet ./...` passes with no issues
8. Code is in `book_store.go` in package `bookstore`

## Key Constraints

- Only integer arithmetic is needed (costs are in cents)
- Books are identified by integers 1-5
- The solution must handle baskets with duplicate books
- The optimization problem requires considering all valid groupings
