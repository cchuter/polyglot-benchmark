# Goal: Implement Book Store Discount Calculator

## Problem Statement

Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go` that calculates the optimal (cheapest) price for a basket of books from a 5-book series, applying volume discounts for groups of different titles.

### Pricing Rules

- Base price: 800 cents ($8.00) per book
- 2 different books: 5% discount on that group
- 3 different books: 10% discount on that group
- 4 different books: 20% discount on that group
- 5 different books: 25% discount on that group

### Key Insight

A greedy approach (always making the largest group first) does NOT produce the optimal result. For example, 2 groups of 4 (5120 cents) is cheaper than 1 group of 5 + 1 group of 3 (5160 cents). The algorithm must find the grouping that minimizes total cost.

## Function Signature

```go
func Cost(books []int) int
```

- Input: `[]int` where each element is a book ID (1-5), may contain duplicates
- Output: `int` representing the total cost in cents after optimal discounts

## Acceptance Criteria

1. `Cost([]int{})` returns `0` (empty basket)
2. `Cost([]int{1})` returns `800` (single book, no discount)
3. `Cost([]int{2, 2})` returns `1600` (two same books, no group discount)
4. `Cost([]int{1, 2})` returns `1520` (two different, 5% discount)
5. `Cost([]int{1, 2, 3})` returns `2160` (three different, 10% discount)
6. `Cost([]int{1, 2, 3, 4})` returns `2560` (four different, 20% discount)
7. `Cost([]int{1, 2, 3, 4, 5})` returns `3000` (five different, 25% discount)
8. `Cost([]int{1, 1, 2, 2, 3, 3, 4, 5})` returns `5120` (two groups of 4, not 5+3)
9. All 18 test cases in `cases_test.go` pass
10. `go test ./...` passes with no errors in the `book-store` exercise directory
11. Only integer arithmetic is used (no floating point)

## Constraints

- Only `book_store.go` should be modified
- Package name must be `bookstore`
- Must handle all edge cases in the test suite
- Must use integer calculations only
