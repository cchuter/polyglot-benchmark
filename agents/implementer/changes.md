# Change Log: Implementer

## Commit ccfbeff — Implement ParseHex and HandleErrors for hexadecimal exercise

**File modified**: `go/exercises/practice/hexadecimal/hexadecimal.go`

### Changes
- Added sentinel errors `ErrRange` and `ErrSyntax` for error classification
- Added `ParseError` struct wrapping input string and underlying error, with `Error()` method
- Implemented `ParseHex(s string) (int64, error)`: character-by-character hex-to-int64 conversion with overflow detection using `math.MaxInt64`, supports `0-9`, `a-f`, `A-F`
- Implemented `HandleErrors(tests []string) []string`: categorizes ParseHex results as `"none"`, `"syntax"`, or `"range"` via type assertion on `*ParseError`
