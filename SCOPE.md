# Scope: polyglot-go-book-store

## In Scope

- Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go`
- Package: `bookstore`
- Function signature: `func Cost(books []int) int`
- Any necessary helper functions (private) within the same file
- All 18 test cases must pass

## Out of Scope

- Modifying test files (`book_store_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying files in `.docs/` or `.meta/`
- Adding new test cases
- Changes to any other exercises
- Performance optimization beyond passing the benchmark test
- Input validation (the test cases assume valid input of integers 1-5)

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- Standard library only (no external dependencies)
- The `math` and `sort` packages may be useful (used by reference solution)

## Files to Modify

- `go/exercises/practice/book-store/book_store.go` — the only file that needs to change
