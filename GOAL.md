# Goal: Implement Book Store Discount Calculator

## Problem Statement

Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go` for the book-store exercise. The function calculates the optimal (cheapest) total price for a basket of books from a 5-book series, applying group discounts.

### Discount Rules

| Group Size | Discount |
|-----------|----------|
| 1 book    | 0%       |
| 2 different books | 5% |
| 3 different books | 10% |
| 4 different books | 20% |
| 5 different books | 25% |

- Base price per book: $8.00 (800 cents)
- Books are identified by integers (1-5)
- The algorithm must find the grouping that yields the maximum total discount (minimum cost)
- Key insight: Two groups of 4 can be cheaper than a group of 5 + group of 3 (5120 < 5160)

### Function Signature

```go
func Cost(books []int) int
```

Returns cost in cents. Only integer arithmetic is needed.

## Acceptance Criteria

1. `Cost` compiles and passes all 18 test cases in `cases_test.go`
2. `Cost` handles the empty basket case (returns 0)
3. `Cost` handles single and duplicate books correctly
4. `Cost` optimally groups books to minimize total cost (e.g., preferring two groups of 4 over one group of 5 + one group of 3)
5. `go test ./...` passes in the `book-store` directory with no errors
6. No modifications to test files (`book_store_test.go`, `cases_test.go`, `go.mod`)

## Key Constraints

- Package name: `bookstore`
- Must use only integer calculations (cents)
- Must handle optimization: greedy grouping is NOT sufficient; the algorithm must consider breaking up 5+3 groups into 4+4 groups for lower cost
