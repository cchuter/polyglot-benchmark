# Context Summary: polyglot-go-hexadecimal

## Issue

#215 — Implement hexadecimal-to-decimal converter in Go using first principles.

## Solution

Implemented `ParseHex(string) (int64, error)` and `HandleErrors([]string) []string` in `go/exercises/practice/hexadecimal/hexadecimal.go`.

### Architecture

- **Sentinel errors**: `ErrRange` and `ErrSyntax` (package-level `errors.New`)
- **ParseError struct**: Wraps input string and sentinel error, implements `error` interface
- **ParseHex**: Byte-by-byte loop with switch for hex digit ranges. Two-stage overflow detection:
  1. Pre-multiply: `n >= math.MaxInt64/16+1`
  2. Post-add: `n1 < n` (belt-and-suspenders, technically unreachable)
- **HandleErrors**: Type-asserts error to `*ParseError`, compares `.Err` against sentinels

### Key Facts

- Imports: `errors`, `math` only (no strconv, no fmt)
- Handles: 0-9, a-f, A-F (case insensitive)
- Error messages contain "syntax" or "range" substrings as required by tests
- All 11 test cases pass + HandleErrors test passes

## Branch

- Feature branch: `issue-215`
- Pushed to origin
- Single commit: `feat: implement hexadecimal to decimal converter`

## Status

Complete. All acceptance criteria verified. Ready for PR/merge.
