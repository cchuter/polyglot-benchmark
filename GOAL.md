# Goal: Implement Book Store Pricing Exercise (Go)

## Problem Statement

Implement the `Cost` function for the book-store exercise in Go. The function calculates the optimal (lowest) price for a shopping basket of books from a 5-book series, applying bulk discounts for sets of distinct books.

### Pricing Rules

- One copy of any book costs **$8.00** (800 cents)
- Discounts apply to groups of **different** books:
  - 2 different books: **5% discount**
  - 3 different books: **10% discount**
  - 4 different books: **20% discount**
  - 5 different books: **25% discount**
- The algorithm must find the **optimal grouping** that yields the lowest total price
- Key insight: greedy grouping (always forming the largest possible set) is NOT always optimal. For example, two groups of 4 ($51.20) is cheaper than one group of 5 + one group of 3 ($51.60).

### Function Signature

```go
func Cost(books []int) int
```

- Input: `[]int` — list of book IDs (1-5), with duplicates representing multiple copies
- Output: `int` — total cost in **cents** after optimal discounts

## Acceptance Criteria

All 18 test cases in `cases_test.go` must pass:

1. Single book → 800
2. Two same books → 1600
3. Empty basket → 0
4. Two different books → 1520
5. Three different books → 2160
6. Four different books → 2560
7. Five different books → 3000
8. Two groups of four cheaper than 5+3 → 5120
9. Two groups of four cheaper than 5+3 (variant) → 5120
10. Group of 4+2 cheaper than two groups of 3 → 4080
11. Two each of first four + one of fifth → 5560
12. Two copies of each → 6000
13. Three of first + two each of remaining → 6800
14. Three each of first two + two each of rest → 7520
15. Four groups of four cheaper than two 5+3 pairs → 10240
16. Proper group-of-4 creation with complex distribution → 14560
17. One group of 1+4 cheaper than 2+3 → 3360
18. Complex 15-book basket → 10000

Additional requirements:
- `go test` passes with zero failures
- `go test -bench=.` runs the benchmark without errors
- Code compiles with `go build`
- All calculations use integer arithmetic (cents)

## Key Constraints

- Package name: `bookstore`
- File: `go/exercises/practice/book-store/book_store.go`
- Go version: 1.18 (as specified in go.mod)
- Only the `book_store.go` file should be modified — test files are read-only
