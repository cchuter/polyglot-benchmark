# Changes

## Implement ParseOctal function

- **File**: `go/exercises/practice/octal/octal.go`
- **Change**: Implemented `ParseOctal(octal string) (int64, error)` function
- **Details**: Converts an octal string to its decimal int64 representation. Iterates over each rune, validates it is a valid octal digit (0-7), and accumulates the result using bit-shifting (`num<<3`). Returns an error for any invalid character.
