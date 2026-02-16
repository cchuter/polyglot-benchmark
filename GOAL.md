# Goal: Implement Forth Evaluator (Issue #170)

## Problem Statement

Implement an evaluator for a simple subset of the Forth stack-based programming language in Go. The solution must be written in `go/exercises/practice/forth/forth.go` and pass all existing tests in `forth_test.go` and `cases_test.go`.

## Acceptance Criteria

1. **Arithmetic operations**: `+`, `-`, `*`, `/` work correctly on the stack (integer arithmetic, integer division)
2. **Stack manipulation**: `DUP`, `DROP`, `SWAP`, `OVER` work correctly
3. **User-defined words**: Support `: word-name definition ;` syntax for defining new words
4. **Error handling**:
   - Arithmetic on empty/insufficient stack returns error
   - Division by zero returns error
   - Cannot redefine numbers (positive or negative)
   - Executing undefined words returns error
5. **Case insensitivity**: All words (built-in and user-defined) are case-insensitive
6. **Word redefinition**: User-defined words can override built-in words and other user-defined words
7. **Snapshot semantics**: Redefined words use the definition at the time of definition, not the current definition
8. **Negative numbers**: Support pushing negative numbers onto the stack
9. **All 42 test cases pass**: `go test ./...` in the forth directory passes with zero failures

## Key Constraints

- Only signed integers need to be supported
- Package name must be `forth`
- Must export function: `Forth(input []string) ([]int, error)`
- Input is a slice of strings; each string is a "phrase" to evaluate in order
- The stack persists across phrases within a single `Forth()` call
- User definitions persist across phrases within a single `Forth()` call
