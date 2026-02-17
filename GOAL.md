# Goal: Implement Forth Evaluator (Issue #256)

## Problem Statement

Implement an evaluator for a very simple subset of the Forth stack-based programming language in Go. The solution must be placed in `go/exercises/practice/forth/forth.go`.

## Acceptance Criteria

1. **Integer Arithmetic**: Support `+`, `-`, `*`, `/` operators performing integer arithmetic on stack values
2. **Stack Manipulation**: Support `DUP`, `DROP`, `SWAP`, `OVER` words
3. **User-Defined Words**: Support defining new words with `: word-name definition ;` syntax
4. **Negative Numbers**: Push negative numbers (e.g., `-1`) onto the stack correctly
5. **Case Insensitivity**: All words (built-in and user-defined) are case-insensitive
6. **Error Handling**:
   - Return error for arithmetic/stack operations on empty or insufficient stack
   - Return error for division by zero
   - Return error for redefining numbers (e.g., `: 1 2 ;`)
   - Return error for executing non-existent words
7. **User-Defined Word Semantics**:
   - Definitions can use built-in words
   - Definitions can override built-in words and operators
   - Definitions can override other user-defined words
   - Definitions capture the meaning of words at definition time (not at execution time)
   - Words defined with the same name as a prior word use the prior word's definition at definition time
8. **All 42 test cases in `cases_test.go` must pass**

## Key Constraints

- Only data type needed: signed integers (at least 16 bits)
- A number is a sequence of one or more ASCII digits (possibly preceded by `-`)
- A word is a sequence of non-whitespace characters that is not a number
- Function signature: `func Forth(input []string) ([]int, error)`
- Package name: `forth`
- Input is a slice of strings, each string is a "phrase" to evaluate in order
