# Changes: book-store exercise

## File Modified
- `go/exercises/practice/book-store/book_store.go`

## What Changed
Implemented the `Cost(books []int) int` function that calculates the cheapest price for a basket of books from a 5-book series, applying optimal group discounts.

## Algorithm
1. **Frequency counting**: Count how many copies of each book (1-5) are in the basket.
2. **Sort descending**: Sort frequencies so the most-purchased book is first.
3. **Layer-peel grouping**: Compute the number of groups of each size (1-5) using differences between adjacent sorted frequencies.
4. **5+3 → 4+4 adjustment**: Convert pairs of (group-of-5 + group-of-3) into (group-of-4 + group-of-4), since 2×2560 = 5120 < 3000+2160 = 5160.
5. **Sum costs**: Apply the discount table and sum all group costs.

## Discount Table (cents)
| Group Size | Cost |
|------------|------|
| 1          | 800  |
| 2          | 1520 |
| 3          | 2160 |
| 4          | 2560 |
| 5          | 3000 |

## Test Results
All 18 test cases pass.
