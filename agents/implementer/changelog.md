# Implementer Changelog

## 2026-02-16

- Implemented `ParseOctal` function in `go/exercises/practice/octal/octal.go`
  - Converts an octal string to its decimal `int64` equivalent using Horner's method
  - Validates each character is in range '0'-'7'; returns error for invalid digits
  - Committed as `ccf8e5d`: "feat: implement ParseOctal for octal-to-decimal conversion"
