# Goal: Implement Book Store Discount Calculator (Issue #107)

## Problem Statement

Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go` for the Exercism "book-store" exercise. The function calculates the optimal price for a basket of books from a 5-book series, applying group discounts to minimize total cost.

### Pricing Rules
- 1 book: $8.00 (800 cents) — no discount
- 2 different books: 5% discount → 2 × 760 = 1520 cents
- 3 different books: 10% discount → 3 × 720 = 2160 cents
- 4 different books: 20% discount → 4 × 640 = 2560 cents
- 5 different books: 25% discount → 5 × 600 = 3000 cents

### Key Insight
The optimal grouping is not always the greedy approach (largest groups first). For example, two groups of 4 (5120 cents) is cheaper than one group of 5 + one group of 3 (3000 + 2160 = 5160 cents). The algorithm must find the grouping that produces the minimum total cost.

## Function Signature

```go
func Cost(books []int) int
```

- Input: `books []int` — a list of book numbers (1-5), may contain duplicates
- Output: `int` — the total cost in cents after applying the best possible discounts

## Acceptance Criteria

1. All 18 test cases in `cases_test.go` pass
2. The `TestCost` test function in `book_store_test.go` passes
3. The function handles edge cases: empty basket (returns 0), single book (returns 800)
4. The function finds the globally optimal grouping, not just a greedy local optimum
5. Only integer arithmetic is used (no floating point)
6. Code compiles with `go build` and tests pass with `go test`

## Key Constraints

- Package name must be `bookstore`
- Function must be exported as `Cost`
- Return value is in cents (integer)
- Books are identified by numbers 1-5
- Must handle baskets of any size
