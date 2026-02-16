# Changes: Implement Cost function for book-store exercise

## File Modified
- `go/exercises/practice/book-store/book_store.go`

## Summary
Implemented the `Cost` function that calculates the total price (in cents) for a basket of books, applying the best possible group discounts.

## Algorithm
1. Count frequency of each book in the basket
2. Sort frequencies descending and greedily build groups (each pass takes one book from each non-zero count)
3. Count groups of size 5 and size 3, redistribute pairs into two groups of 4 (saves 40 cents per pair)
4. Sum total cost using the discount table: `[0, 800, 1520, 2160, 2560, 3000]`

## Commit
`fa018aa` - "Implement Cost function for book-store exercise"
