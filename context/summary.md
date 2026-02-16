# Context Summary — polyglot-go-octal

## Issue

#139 — Implement octal (base-8) to decimal (base-10) conversion in Go.

## Solution

Single function `ParseOctal(string) (int64, error)` in `go/exercises/practice/octal/octal.go`.

Algorithm: iterate runes, validate each is '0'-'7', accumulate via bit-shift (`num<<3 + digit`). Return error on invalid input.

## Key Decisions

- Used bit-shifting (`<<3`) instead of multiplication by 8, matching the reference solution
- Used `fmt.Errorf` for error creation (lightweight, tests only check `err != nil`)
- Empty string returns `(0, nil)` — not tested, acceptable per reference solution

## Files Modified

- `go/exercises/practice/octal/octal.go` — the only file changed

## Branch

- `issue-139` — pushed to origin

## Status

Complete. All tests pass. Ready for PR.
