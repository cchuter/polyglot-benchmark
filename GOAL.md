# Goal: Implement Forth Evaluator (Issue #341)

## Problem Statement

Implement an evaluator for a simple subset of the Forth stack-based programming language in Go. The stub file `go/exercises/practice/forth/forth.go` contains only a package declaration and needs a complete implementation of the `Forth(input []string) ([]int, error)` function.

## Acceptance Criteria

1. **Integer arithmetic**: Support `+`, `-`, `*`, `/` operators performing integer arithmetic on the stack (pop two, push result).
2. **Stack manipulation**: Support `DUP` (duplicate top), `DROP` (remove top), `SWAP` (swap top two), `OVER` (copy second-from-top to top).
3. **User-defined words**: Support defining new words with `: word-name definition ;` syntax.
4. **Case insensitivity**: All words (built-in and user-defined) are case-insensitive.
5. **Number parsing**: Numbers (sequences of ASCII digits, optionally with a leading `-`) are pushed onto the stack.
6. **Error handling**:
   - Arithmetic/stack ops on insufficient stack elements return an error.
   - Division by zero returns an error.
   - Redefining numbers (e.g., `: 1 2 ;`) returns an error.
   - Referencing undefined words returns an error.
7. **User-defined word semantics**: Definitions capture the meaning of words at definition time (not execution time). Redefining a word used in another definition does not retroactively change the other definition.
8. **All 42 test cases in `cases_test.go` pass** (`go test ./...` in the exercise directory).
9. **`go vet ./...` passes** with no issues.

## Key Constraints

- Only data type is signed integers (at least 16 bits).
- Input is a slice of strings; each string is a "phrase" to evaluate sequentially.
- The stack persists across phrases within a single `Forth()` call.
- A reference solution exists in `.meta/example.go` for guidance on expected behavior.
