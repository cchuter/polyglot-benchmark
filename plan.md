# Implementation Plan: Forth Evaluator

## Branch 1: Minimal Direct Implementation

A straightforward approach: parse tokens, evaluate directly against the stack. No intermediate AST or operator type system.

### Approach
- Single file `forth.go` with `Forth` function
- Use a `map[string][]string` to store user-defined word definitions as token lists
- For each phrase, split into tokens, evaluate left to right
- On encountering `:`, collect tokens until `;` as a word definition
- On encountering a known word, recursively expand and evaluate
- Case-insensitive by converting all tokens to uppercase

### Files to Modify
- `go/exercises/practice/forth/forth.go` (only file)

### Evaluation
- **Feasibility**: High — simple, well-understood approach
- **Risk**: Low — minimal moving parts. Recursive expansion could hit issues with the "definitions capture meaning at definition time" requirement. Storing as string tokens means re-lookup at execution time, which would NOT satisfy the test "can use different words with the same name" (bar uses foo=5, then foo is redefined to 6, but bar should still use 5).
- **Alignment**: FAILS — string-token storage doesn't capture definition-time semantics correctly
- **Complexity**: Low code volume but semantically incorrect

## Branch 2: Operator-List Approach (Reference Pattern)

Follow the pattern from `.meta/example.go`: compile words into lists of operator functions at definition time. This captures definition-time semantics by resolving words to their operations during definition.

### Approach
- Define an `operatorFn` type (`func(stack *[]int) error`) and an `operatorTyp` struct
- Built-in operations (+, -, *, /, DUP, DROP, SWAP, OVER) each get a function
- Parse each phrase: split on whitespace, iterate tokens
- Numbers become constant-push operators (closures capturing the value)
- User-defined words resolve to operator lists at definition time (snapshot semantics)
- `:` triggers definition mode: collect subsequent tokens, resolve each to operators, store the resulting operator list under the word name
- Each phrase's operator list is executed sequentially against the stack
- Helper functions: `pop`, `pop2`, `push`, `binaryOp`

### Files to Modify
- `go/exercises/practice/forth/forth.go` (only file)

### Evaluation
- **Feasibility**: High — proven pattern from reference implementation
- **Risk**: Low — well-tested approach, handles all edge cases
- **Alignment**: Full — snapshot semantics naturally emerge from resolving at definition time
- **Complexity**: ~230 lines, single file, no dependencies beyond stdlib

## Branch 3: Interface-Based Stack Machine

Use Go interfaces to model stack operations as a more extensible virtual machine.

### Approach
- Define `Instruction` interface with `Execute(stack *Stack) error` method
- `Stack` type wraps `[]int` with methods (Push, Pop, Peek, Size)
- Each built-in operation implements `Instruction`
- User-defined words are `CompositeInstruction` (slice of `Instruction`)
- Parser compiles input to `[]Instruction`
- Executor runs instructions sequentially

### Files to Modify
- `go/exercises/practice/forth/forth.go` (only file)

### Evaluation
- **Feasibility**: High — works but over-engineered for this problem
- **Risk**: Medium — more abstractions means more surface area for bugs
- **Alignment**: Full — can handle definition-time semantics
- **Complexity**: Higher (~300+ lines), more types and methods than needed

## Selected Plan

**Branch 2: Operator-List Approach** is selected.

### Rationale
- Branch 1 fails the definition-time semantics tests — disqualified
- Branch 3 is over-engineered for a single-function exercise with no extensibility requirements
- Branch 2 directly follows the proven reference pattern, is the right level of abstraction, and handles all test cases correctly

### Detailed Implementation Plan

#### File: `go/exercises/practice/forth/forth.go`

**Package and imports:**
```go
package forth

import (
    "errors"
    "strconv"
    "strings"
    "unicode"
)
```

**Types:**
- `operatorFn func(stack *[]int) error` — function signature for stack operations
- `operatorID byte` — enum for operator types (opAdd, opSub, opMul, opDiv, opDrop, opDup, opSwap, opOver, opConst, opUserDef, opEndDef)
- `operatorTyp struct { fn operatorFn; id operatorID }` — pairs a function with its type ID

**Main function `Forth(input []string) ([]int, error)`:**
1. Initialize empty stack (`make([]int, 0, 8)`)
2. Initialize user-defined words map (`make(map[string][]operatorTyp)`)
3. For each phrase in input, parse it into an operator list, then execute each operator
4. Return the final stack state

**Parser `parse(phrase string, userDefs map[string][]operatorTyp) ([]operatorTyp, error)`:**
1. Split phrase on whitespace using `strings.FieldsFunc`
2. Iterate tokens:
   - If token matches a user-defined word (case-insensitive), append its operator list
   - If token is `:`, enter definition mode: read word name, validate it's not a number, collect operators for the body until `;`, store in userDefs
   - If token is a built-in operator, append it
   - Otherwise, try parsing as integer; if valid, create a closure that pushes the value
   - If none match, return "undefined operation" error

**Built-in operators map:**
- `+`, `-`, `*`, `/` → arithmetic functions
- `DUP`, `DROP`, `SWAP`, `OVER` → stack manipulation functions
- `:` → user definition start marker
- `;` → user definition end marker

**Stack helper functions:**
- `pop(stack *[]int) (int, error)` — pop top value
- `pop2(stack *[]int) (int, int, error)` — pop top two values
- `push(stack *[]int, int)` — push value
- `binaryOp(stack *[]int, op func(a, b int) int) error` — generic binary operation

**Arithmetic functions:** `add`, `subtract`, `multiply`, `divide` (with zero check)

**Stack manipulation functions:** `dup`, `drop`, `swap`, `over`

**Error variables:** `errNotEnoughOperands`, `errDivideByZero`, `errEmptyUserDef`, `errInvalidUserDef`

**Key semantic detail:** When parsing a user-defined word's body, each token is resolved to its operator(s) at definition time. This means if `: foo 5 ;` then `: bar foo ;` then `: foo 6 ;`, `bar` still resolves to `[push 5]` because foo's definition was captured when bar was defined.

**Important fix from reference:** The reference has a subtle bug where DUP maps to opDrop and DROP maps to opDup (swapped IDs). Since the IDs are only used for the `;` end-of-definition check and don't affect actual execution (the function pointers are correct), this doesn't cause test failures. Our implementation should use correct IDs.

#### Implementation Order
1. Write complete `forth.go` with all types, functions, and the main `Forth` function
2. Run tests to verify all 42 pass
3. Run `go vet` for static analysis
4. Commit

#### Error Handling for Undefined Words
The reference uses `strconv.Atoi` which returns an error for non-numeric, non-defined words. We should return a more descriptive error for undefined words by checking user defs and builtins before trying number parsing.
