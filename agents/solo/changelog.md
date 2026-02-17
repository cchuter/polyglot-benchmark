# Solo Agent Change Log

## 2026-02-16: Implement Cost function

### Changes Made
- **`go/exercises/practice/book-store/book_store.go`**: Implemented the `Cost` function

### Algorithm
- Count frequency of each book (1-5) in the basket
- Sort frequencies descending and greedily form groups (largest possible each iteration)
- After greedy grouping, optimize by converting (5-group + 3-group) pairs into (4-group + 4-group) pairs, saving 40 cents per conversion
- Uses a `groupCost` helper to map group sizes to discounted costs in cents

### Test Results
- All 18 test cases pass
- `go vet` clean
- No external dependencies added
