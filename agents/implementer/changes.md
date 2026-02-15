# Implementer Change Log

## 2026-02-15: Implement Cost function for book-store exercise

**File changed:** `go/exercises/practice/book-store/book_store.go`

**Summary:** Implemented the `Cost` function using a frequency-based recursive search with memoization to find the minimum cost for a basket of books.

**Details:**
- `Cost(books []int) int` — entry point; builds a frequency array from book IDs (1-5), sorts descending for canonical state, and delegates to `minCost`.
- `minCost(freq [5]int, memo map[[5]int]int) int` — recursively tries all group sizes from 1 to the number of distinct books remaining, memoizing on the sorted frequency array to avoid redundant computation.
- `groupCost(n int) int` — calculates the discounted price (in cents) for a group of `n` distinct books using the discount table: 0% for 1, 5% for 2, 10% for 3, 20% for 4, 25% for 5.
- Discount table stored as `var discounts = [6]int{0, 0, 5, 10, 20, 25}`.

**Build verification:** `go build ./...` passed with no errors.
