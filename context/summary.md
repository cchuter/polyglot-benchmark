# Context Summary: Issue #107 — polyglot-go-book-store

## Status: COMPLETE

## Problem
Implement a `Cost(books []int) int` function that calculates the optimal price for a basket of books from a 5-book series with group discounts (5% for 2 different, 10% for 3, 20% for 4, 25% for 5).

## Solution
Greedy grouping algorithm with 5+3→4+4 redistribution:
1. Count frequency of each book title
2. Sort frequencies descending
3. Greedily form groups by taking one from each non-zero count per pass
4. Count groups of size 5 and 3; convert pairs to 4+4 (saves 40 cents each)
5. Sum costs using lookup table, subtract redistribution savings

## Key Files
- `go/exercises/practice/book-store/book_store.go` — Implementation (single file, 66 lines)
- `go/exercises/practice/book-store/book_store_test.go` — Test runner
- `go/exercises/practice/book-store/cases_test.go` — 18 test cases

## Branch
- Feature branch: `issue-107`
- Pushed to origin

## Verification
- 18/18 tests pass
- go build clean
- go vet clean
- All acceptance criteria met
