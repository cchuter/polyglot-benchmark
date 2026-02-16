# Change Log — Implementer

## hexadecimal.go (new implementation)

- Added sentinel errors `ErrRange` and `ErrSyntax` for structured error classification
- Added `ParseError` struct with `Num` and `Err` fields, plus `Error()` method producing formatted messages
- Implemented `ParseHex(s string) (int64, error)`:
  - Empty string returns syntax error
  - Byte-by-byte loop with switch for 0-9, a-f, A-F character ranges
  - Two-stage overflow detection: pre-multiply check (`n >= math.MaxInt64/16+1`) and post-add check (`n1 < n`)
  - Invalid characters return syntax error
- Implemented `HandleErrors(tests []string) []string`:
  - Calls `ParseHex` on each input
  - Type-asserts error to `*ParseError` and compares `.Err` against sentinels
  - Returns "none", "syntax", or "range" per input

## Tests

All tests pass (`TestParseHex`, `TestHandleErrors`).

## Commit

`9c9e056` — feat: implement hexadecimal to decimal converter
