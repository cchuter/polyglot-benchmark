# Context: Forth Evaluator Implementation

## Architecture

Compile-then-execute pattern:
1. `parse(phrase, userDefs)` → tokenize, resolve words to operator list
2. `Forth(input)` → for each phrase: parse then execute operators against stack

## Key Types

- `operatorFn func(stack *[]int) error` — operator function signature
- `operatorTyp struct { fn operatorFn; id operatorID }` — operator with type tag
- `operatorID byte` — enum (opAdd through opEndDef)

## User-Defined Words

- Stored as `map[string][]operatorTyp`
- Expanded at parse-time (definition-time capture)
- Checked before builtinOps for override support
- Cannot redefine numbers (strconv.Atoi check on name)

## Error Handling

- `errNotEnoughOperands` — empty stack on pop
- `errDivideByZero` — division by zero
- `errEmptyUserDef` — malformed definition
- `errInvalidUserDef` — attempting to redefine a number

## Files

- Solution: `go/exercises/practice/forth/forth.go` (231 lines)
- Tests: `forth_test.go` + `cases_test.go` (46 subtests, all pass)
- Branch: `issue-297`
- Commit: `da98bbe` — "Implement Forth evaluator with compile-then-execute architecture"
