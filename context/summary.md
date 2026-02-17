# Context Summary: Book Store Exercise (Issue #278)

## Key Decisions

- **Algorithm**: Greedy grouping with 5+3→4+4 post-optimization (Proposal A from plan.md)
- **Why not DP**: Over-engineered for a problem with a known O(n) solution
- **Go 1.18 compatibility**: Used custom `minInt` helper instead of Go 1.21+ built-in `min`
- **0-based indexing**: `freq[b-1]` maps books 1-5 to indices 0-4
- **Pre-computed price table**: `[6]int{0, 800, 760, 720, 640, 600}` indexed by group size

## Files Modified

- `go/exercises/practice/book-store/book_store.go` — Full implementation of `Cost(books []int) int`

## Test Results

All 18 test cases pass. No build errors. Only `book_store.go` was modified.

## Branch

`issue-278` — pushed to origin, ready for PR
