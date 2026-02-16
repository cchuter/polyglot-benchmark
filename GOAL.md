# Goal: Implement Forth Evaluator (Issue #213)

## Problem Statement

Implement an evaluator for a simple subset of the Forth stack-based programming language in Go. The solution must be written in `go/exercises/practice/forth/forth.go` and export a single function `Forth(input []string) ([]int, error)` that evaluates Forth programs and returns the resulting stack state.

## Acceptance Criteria

1. **Integer arithmetic**: Support `+`, `-`, `*`, `/` operators performing integer arithmetic on the top two stack values
2. **Division behavior**: Integer division (truncation toward zero); division by zero returns an error
3. **Stack manipulation**: Support `DUP`, `DROP`, `SWAP`, `OVER` operations
4. **User-defined words**: Support defining new words via `: word-name definition ;` syntax
5. **Word redefinition**: User-defined words can override built-in words and other user-defined words
6. **Snapshot semantics**: When a word is redefined, existing definitions that referenced the old word retain the old behavior (definitions are resolved at definition time, not execution time)
7. **Case insensitivity**: All words (built-in and user-defined) are case-insensitive
8. **Number restriction**: Numbers cannot be redefined (`: 1 2 ;` is an error)
9. **Negative numbers**: Negative numbers (e.g., `-1`) are pushed onto the stack
10. **Error handling**: Return errors for:
    - Insufficient stack operands for an operation
    - Division by zero
    - Executing a non-existent word
    - Attempting to redefine a number
11. **Multi-phrase input**: The `input` parameter is a slice of strings; each string is a phrase that may contain definitions or operations. Definitions persist across phrases.
12. **All 42 test cases in `cases_test.go` and `forth_test.go` must pass**

## Key Constraints

- Solution file: `go/exercises/practice/forth/forth.go`
- Package: `forth`
- Go module: `go 1.18`
- Only data type: signed integers (at least 16 bits)
- A number is a sequence of one or more ASCII digits (possibly prefixed with `-`)
- A word is a sequence of one or more letters, digits, symbols, or punctuation that is not a number
