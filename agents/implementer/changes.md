# Changes: Implement ParseOctal

## What was done
- Implemented the `ParseOctal` function in `go/exercises/practice/octal/octal.go`
- The function converts an octal string to its decimal `int64` representation
- Validates each character is a valid octal digit (0-7), returning an error for invalid runes
- Uses bit shifting (`<<3`) for efficient multiplication by 8 during conversion
- Committed as `c39a04f` on branch `issue-309` with message: "feat: implement ParseOctal for octal-to-decimal conversion"
