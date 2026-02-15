# Changes: book-store Cost function

## What was done
Replaced the `panic` stub in `go/exercises/practice/book-store/book_store.go` with a full implementation of the `Cost` function.

## Implementation details
- **Greedy grouping**: Count frequency of each book, then greedily form the largest possible groups of distinct books.
- **5+3 → 4+4 optimization**: Two groups of 4 (2×2560 = 5120) are cheaper than a group of 5 + group of 3 (3000+2160 = 5160). The algorithm swaps paired 5-groups and 3-groups into 4-groups.
- **Lookup table**: A precomputed `groupCosts` array maps group size → cost in cents.

## Verification
- `go vet ./...` passes with no issues.
- All 18 test cases pass.
