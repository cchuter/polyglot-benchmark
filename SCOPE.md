# Scope: Book Store Exercise (Issue #107)

## In Scope

- Implement the `Cost` function in `go/exercises/practice/book-store/book_store.go`
- The function must calculate optimal book basket pricing with group discounts
- Must pass all 18 existing test cases in `cases_test.go`
- Implementation confined to the single file `book_store.go`

## Out of Scope

- Modifying test files (`book_store_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Adding new test cases
- Modifying any files in `.docs/` or `.meta/`
- Implementing solutions for other exercises
- Performance optimization beyond passing the benchmark test

## Dependencies

- Go 1.18+ (as specified in go.mod)
- No external packages required — standard library only
- Test infrastructure already exists in `book_store_test.go` and `cases_test.go`

## Files to Modify

| File | Action |
|------|--------|
| `go/exercises/practice/book-store/book_store.go` | Implement `Cost` function |
