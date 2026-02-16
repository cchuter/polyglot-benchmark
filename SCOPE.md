# Scope: Book Store Exercise

## In Scope

- Implementing `func Cost(books []int) int` in `go/exercises/practice/book-store/book_store.go`
- The function must find the optimal (minimum cost) grouping of books
- All 18 test cases defined in `cases_test.go` must pass
- Package name: `bookstore`

## Out of Scope

- Modifying test files (`book_store_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying any `.meta/` files
- Modifying any other exercises
- Adding new files to the exercise
- Performance benchmarking beyond passing tests

## Dependencies

- Go 1.18+ (as specified in `go.mod`)
- No external dependencies; standard library only
- Test infrastructure already exists in `book_store_test.go` and `cases_test.go`

## Key Files

| File | Role |
|------|------|
| `go/exercises/practice/book-store/book_store.go` | Solution file (to be implemented) |
| `go/exercises/practice/book-store/book_store_test.go` | Test runner |
| `go/exercises/practice/book-store/cases_test.go` | Test cases (18 cases) |
| `go/exercises/practice/book-store/.meta/example.go` | Reference solution (read-only) |
