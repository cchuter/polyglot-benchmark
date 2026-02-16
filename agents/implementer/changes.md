# Changes

## Implement Cost function for book-store exercise

- **File**: `go/exercises/practice/book-store/book_store.go`
- **Commit**: `bd0fa4d` — Implement Cost function for book-store exercise
- **Summary**: Implemented the `Cost` function which calculates the total price of a basket of books with optimal discount grouping. Uses frequency histogram layer-peeling to build groups, then applies the 5+3 → 4+4 optimization to minimize total cost.
