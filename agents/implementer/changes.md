# Changes

## Implemented `Cost` function in `book_store.go`

- Added `Cost(books []int) int` function that calculates the optimal price for a basket of books
- Algorithm: greedy grouping with 5+3 to 4+4 optimization
  - Counts book frequencies, sorts descending, greedily forms groups of distinct books
  - Converts (5-group + 3-group) pairs into (4-group + 4-group) pairs for better discount
- Added `minInt` helper since Go 1.18 lacks a built-in `min` for ints
- Discount table: 0% (1 book), 5% (2), 10% (3), 20% (4), 25% (5 different)
- Returns total cost in cents using integer arithmetic only
- Commit: `770df60`
