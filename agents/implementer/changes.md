# Change Log: Book Store Exercise

## Implemented `Cost` function in `book_store.go`

**File modified:** `go/exercises/practice/book-store/book_store.go`

### What was done
- Implemented `func Cost(books []int) int` that calculates the minimum cost for a basket of books with volume discounts
- Added helper `minCost(counts []int) int` for recursive optimization over frequency counts
- Added helper `groupCost(size int) int` to compute the discounted price for a group of distinct books

### Algorithm
1. Count frequency of each book title in the basket
2. Sort frequencies in descending order
3. Recursively try all possible group sizes (1 to k distinct titles) at each step
4. Return the minimum cost across all groupings

### Discount table
- 1 book: 0% discount (800 cents)
- 2 different books: 5% discount (1520 cents)
- 3 different books: 10% discount (2160 cents)
- 4 different books: 20% discount (2560 cents)
- 5 different books: 25% discount (3000 cents)

### Test results
All 18 test cases pass.
