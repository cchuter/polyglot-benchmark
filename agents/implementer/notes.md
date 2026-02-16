# Implementer Notes: Forth Evaluator

## Approach
Direct Interpreter with Token Expansion (Branch 1 from plan).

## Implementation Summary
- **File**: `go/exercises/practice/forth/forth.go` (140 lines)
- **Architecture**: Single `Forth()` entry point → `eval()` token processor → `execBuiltin()` for ops
- **Data structures**: `[]int` stack, `map[string][]string` defs map

## Key Design Decisions
1. **Token expansion at definition time**: When `: foo bar baz ;` is defined, `bar` and `baz` are expanded using current defs immediately, achieving snapshot semantics without runtime lookup.
2. **Eval order**: defs → number (Atoi) → built-in → error. This ensures user overrides of built-ins work correctly.
3. **Case insensitivity**: All tokens uppercased via `strings.ToUpper` before processing.
4. **Recursive eval for user words**: When a user-defined word is encountered, its expanded body is evaluated recursively. Since expansion only produces numbers and built-in names, recursion depth is bounded.

## Test Results
- All 42 test cases pass
- Covers: arithmetic, stack ops, user definitions, snapshot semantics, case insensitivity, error handling

## Commit
`9292b84` - feat: implement Forth evaluator
