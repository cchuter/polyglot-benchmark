# Context Summary: book-store (Issue #36)

## Key Decisions

- Used frequency-based dynamic programming with memoization to find optimal book grouping
- Represent state as sorted [5]int frequency array (canonical form for memoization)
- Recursively try all group sizes 1..distinct at each step, pick minimum cost
- Discount table indexed by group size: [6]int{0, 0, 5, 10, 20, 25}
- Integer arithmetic only (costs in cents) — no floating point

## Files Modified

- `go/exercises/practice/book-store/book_store.go` — Implemented `Cost`, `minCost`, and `groupCost` functions (56 lines)

## Test Results

- 18/18 tests pass (`go test -v -count=1`)
- No go vet warnings (`go vet ./...`)
- Build succeeds (`go build ./...`)
- Benchmark: ~282-289 us/op

## Commit

- `9cf8f0c` — "feat: implement Cost function for book-store exercise"
- Branch: `issue-36`
- Pushed to origin

## Architecture Notes

- The algorithm avoids the greedy trap where 5+3 grouping costs more than 4+4
- Memoization on sorted frequency tuples keeps the state space small (~59 unique states for 22 books)
- Go [5]int arrays are comparable and work as map keys (unlike slices)
