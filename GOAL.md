# Goal: Implement Book Store Discount Calculator

## Problem Statement

Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go` for the Exercism "book-store" exercise. The function calculates the optimal (cheapest) price for a basket of books from a 5-book series, applying group discounts to maximize savings.

### Pricing Rules

- 1 book: $8.00 (800 cents), no discount
- 2 different books: 5% discount on those 2
- 3 different books: 10% discount on those 3
- 4 different books: 20% discount on those 4
- 5 different books: 25% discount on all 5

The key challenge is finding the optimal grouping. A greedy approach (always making the largest group) is not optimal. For example, two groups of 4 ($51.20) is cheaper than one group of 5 plus one group of 3 ($51.60).

## Acceptance Criteria

1. `Cost(books []int) int` is defined in package `bookstore` in file `book_store.go`
2. The function returns the total cost in cents (integer)
3. All 18 test cases in `cases_test.go` pass, including:
   - Empty basket returns 0
   - Single book returns 800
   - Duplicate books are not discounted together
   - Optimal grouping is found (e.g., two groups of 4 preferred over group of 5 + group of 3)
   - Large baskets with complex groupings are handled correctly
4. `go test` passes with no failures
5. `go vet` reports no issues

## Key Constraints

- Only integer arithmetic is needed (costs are in cents)
- Books are identified by integers (1-5)
- The function must find the globally optimal grouping, not just a greedy solution
- Must handle baskets of varying sizes including empty baskets
