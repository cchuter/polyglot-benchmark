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

## 2026-02-15: Implement Recite function for bottle-song exercise

**File changed:** `go/exercises/practice/bottle-song/bottle_song.go`

**Summary:** Implemented the `Recite` function to generate "Ten Green Bottles" song lyrics.

**Details:**
- Added imports: `fmt`, `strings`
- `numberToWord` map (`map[int]string`) mapping integers 0-10 to English words
- `capitalize(s string) string` helper to uppercase the first character using `strings.ToUpper`
- `verse(n int) []string` function with switch cases for n==1 (singular bottle, "no" remaining), n==2 (plural bottles, singular remaining), and default (plural both)
- `Recite(startBottles, takeDown int) []string` loops from `startBottles` downward for `takeDown` iterations, assembling verses separated by empty strings

**Test verification:** All 7 tests pass (`go test -v`).
