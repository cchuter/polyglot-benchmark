# Goal: polyglot-go-book-store

## Problem Statement

Implement the `Cost` function for the book-store exercise in Go. The function calculates the optimal (cheapest) total price for a basket of books from a 5-book series, applying volume discounts for groups of distinct books.

### Discount Structure

| Group Size | Discount |
|-----------|----------|
| 1 book    | 0%       |
| 2 different books | 5% |
| 3 different books | 10% |
| 4 different books | 20% |
| 5 different books | 25% |

Base price per book: $8.00 (800 cents).

The key challenge is finding the optimal grouping of books to minimize total cost. A greedy approach (always making the largest group) is not always optimal — e.g., two groups of 4 ($51.20) is cheaper than one group of 5 + one group of 3 ($51.60).

## Function Signature

```go
func Cost(books []int) int
```

- Input: a slice of integers where each integer (1-5) represents a book title
- Output: the minimum total cost in cents (integer)

## Acceptance Criteria

1. All 18 test cases in `cases_test.go` pass
2. The `TestCost` function in `book_store_test.go` passes
3. The `BenchmarkCost` function runs without error
4. The solution is in `book_store.go` within package `bookstore`
5. The function correctly handles edge cases:
   - Empty basket → 0
   - Single book → 800
   - All same books → 800 × count
   - All different books → appropriate discount applied
6. The function finds the globally optimal grouping, not just a greedy grouping

## Key Constraints

- Return value is in cents (integer arithmetic only)
- Must be in package `bookstore`
- Must match the exact function signature: `func Cost(books []int) int`
- Books are identified by integers (1-5)
- Must handle baskets of varying sizes efficiently
