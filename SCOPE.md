# Scope: Book Store Exercise

## In Scope

- Implementing the `Cost(books []int) int` function in `go/exercises/practice/book-store/book_store.go`
- The implementation must handle:
  - Empty baskets
  - Single books
  - Multiple copies of the same book
  - Optimal grouping to minimize total cost (not just greedy grouping)
  - Baskets of any size with any combination of books 1-5
- All 18 existing test cases must pass
- Benchmark must run successfully

## Out of Scope

- Modifying test files (`book_store_test.go`, `cases_test.go`)
- Modifying `go.mod`
- Modifying `.meta/` files (config.json, example.go, tests.toml)
- Modifying `.docs/` files
- Adding new test cases
- Adding external dependencies
- Exercises in other languages (Python, Rust, Java, JavaScript, C++)
- Performance optimization beyond passing benchmarks

## Dependencies

- Go 1.18+ toolchain
- No external packages — standard library only
- Existing test infrastructure (`cases_test.go` provides test data, `book_store_test.go` provides test harness)
