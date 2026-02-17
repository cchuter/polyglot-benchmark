# Goal: Implement Forth Evaluator (Issue #297)

## Problem Statement

Implement a working evaluator for a simple subset of the Forth stack-based programming language in Go. The solution file `go/exercises/practice/forth/forth.go` is currently a stub containing only `package forth`. The test file and test cases already exist with 45 test cases covering all required functionality.

## Required Function Signature

```go
func Forth(input []string) ([]int, error)
```

- **Input**: A slice of strings, where each string is a "phrase" (line) of Forth code
- **Output**: The final stack state as `[]int`, or an error

## Acceptance Criteria

1. **Number parsing**: Push positive and negative integers onto the stack
2. **Arithmetic operations**: `+`, `-`, `*`, `/` (integer arithmetic, integer division)
3. **Stack manipulation**: `DUP`, `DROP`, `SWAP`, `OVER`
4. **User-defined words**: `: word-name definition ;` syntax for defining new words
5. **Case-insensitivity**: All built-in words and user-defined words are case-insensitive
6. **Error handling**:
   - Arithmetic on insufficient stack values → error
   - Division by zero → error
   - Stack operations on insufficient stack values → error
   - Redefining numbers → error ("illegal operation")
   - Executing undefined words → error ("undefined operation")
7. **User-defined word semantics**:
   - Can consist of built-in words
   - Can override other user-defined words
   - Can override built-in words and operators
   - Definitions capture the meaning of words at definition time (not call time)
   - Words can redefine themselves using their prior meaning
8. **All 45 test cases in `cases_test.go` pass**
9. **`go test` runs successfully with no failures**

## Key Constraints

- Only modify `forth.go` — do not modify test files or auto-generated case files
- Package must be `forth`
- Go 1.18 compatibility (per go.mod)
- Data type: signed integers (Go `int` satisfies 16-bit minimum)
