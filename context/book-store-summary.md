# Context: Book Store Exercise

## Status
Complete. All acceptance criteria met.

## Branch
`issue-7` — pushed to origin

## Commit
`130e9f1` — feat: implement book-store Cost function with greedy grouping and 5+3→4+4 optimization

## Algorithm
Frequency-based greedy grouping with 5+3→4+4 optimization:
1. Count book frequencies
2. Greedily form groups (each round takes one of each available title)
3. Replace pairs of (group-of-5, group-of-3) with (group-of-4, group-of-4) — saves 40 cents per swap
4. Sum group costs using lookup table

## Key Facts
- Package: `bookstore`
- File: `go/exercises/practice/book-store/book_store.go`
- Go version: 1.18 (no min/max builtins)
- No imports needed
- 18/18 tests pass
- Benchmark: ~11,212 ns/op

## Codex Reviews
- Design review: Algorithm confirmed correct, caught Go 1.18 min() issue
- Implementation review: Exhaustive DP verification showed 0 mismatches across 161,051 states
