# Context Summary: Book Store Exercise (Issue #63)

## Key Decisions

- **Algorithm**: Frequency histogram with 5+3→4+4 optimization. This is an O(n) solution that avoids the exponential brute-force approach while still finding the optimal grouping.
- **Why not greedy alone**: Greedy (largest groups first) fails because two groups of 4 ($51.20) are cheaper than groups of 5+3 ($51.60). The 5+3→4+4 correction is the only needed post-processing.
- **Why not recursive/brute-force**: The reference solution in `.meta/example.go` uses recursion which works but is slower. Our approach is deterministic and O(n).

## Files Modified

- `go/exercises/practice/book-store/book_store.go` — Implemented `Cost` function (was just `package bookstore`)

## Test Results

- 18/18 test cases pass
- Benchmark: ~11,255 ns/op
- All acceptance criteria from GOAL.md verified

## Branch

- Feature branch: `issue-63`
- Pushed to origin: yes
- Base branch: `bench/polyglot-go-book-store`
