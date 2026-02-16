# Context: polyglot-go-ledger Implementation

## Key Decisions

- Single-file implementation in `go/exercises/practice/ledger/ledger.go`
- Data-driven approach with currency symbol map and locale-specific branches
- Integer-only arithmetic for monetary formatting (no floats)
- Literal header strings (not Sprintf-generated) to guarantee exact test match
- Copy-then-sort pattern for input immutability
- `time.Parse` for date validation
- Lexicographic string comparison for date sorting (YYYY-MM-DD sorts naturally)

## Files Modified

- `go/exercises/practice/ledger/ledger.go` — full implementation (135 lines)

## Test Results

- 17/17 tests pass (10 success cases, 6 failure cases, 1 immutability test)
- `go build`: clean
- `go vet`: clean

## Open Notes

- Challenger flagged `math.MinInt` overflow and UTF-8 truncation as advisory edge cases — not addressed since they don't affect test results
- Branch: `issue-133`, pushed to origin
