# Context: polyglot-go-hexadecimal

## Files Modified
- `go/exercises/practice/hexadecimal/hexadecimal.go` — full implementation

## Key Decisions
- Character-by-character hex parsing with byte-level switch statement
- Overflow detection: pre-multiply check (`n >= MaxInt64/16+1`) and post-add check (`n1 < n`)
- `ParseError` wraps sentinel errors (`ErrSyntax`, `ErrRange`) with input string
- `HandleErrors` uses type assertion switch on `*ParseError`

## Test Results
- `go test ./...` — PASS (all 11 test cases)
- `go vet ./...` — clean

## Branch
- `issue-172` pushed to origin
