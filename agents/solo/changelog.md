# Change Log: polyglot-go-hexadecimal

## Changes

### `go/exercises/practice/hexadecimal/hexadecimal.go`
- **Modified**: Replaced empty stub with full implementation
- Added `ErrRange` and `ErrSyntax` sentinel errors
- Added `ParseError` struct with `Error()` method
- Implemented `ParseHex(string) (int64, error)` — character-by-character hex-to-decimal conversion with overflow detection
- Implemented `HandleErrors([]string) []string` — classifies ParseHex errors as "none", "syntax", or "range"

## Test Results
- `go test ./...` — PASS
- `go vet ./...` — clean
