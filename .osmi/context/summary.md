# Context Summary: polyglot-go-octal

## Key Decisions
- Used bit-shifting (`<<3`) instead of explicit multiplication (`*8`) to match reference solution pattern
- Used `fmt.Errorf` for error creation to include the invalid character in the message
- Used `range` iteration over string (rune-based) for idiomatic Go

## Files Modified
- `go/exercises/practice/octal/octal.go` — implemented `ParseOctal`

## Test Results
- All 5 test cases pass (1, 10, 1234567, carrot, 35682)
- `go vet` clean

## Branch
- Feature branch: `issue-353`
- Pushed to origin
