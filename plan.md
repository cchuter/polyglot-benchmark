# Implementation Plan: Forth Evaluator

## Proposal A: Compiled Operator List Approach

**Role: Proponent**

This approach pre-compiles each phrase into a list of operator functions, then executes them. User-defined words are stored as lists of operators that are inlined at definition-time.

### Architecture

- **Operator type**: A struct pairing an `operatorFn func(stack *[]int) error` with an `operatorID` enum for identification (needed to detect end-of-definition markers).
- **Parse phase**: Each phrase is tokenized by splitting on whitespace. Tokens are resolved to operators: numbers become push-closures, built-in words map to predefined functions, user-defined words are expanded inline from a definition map, and `: name ... ;` blocks register new definitions.
- **Execute phase**: Walk the operator list, calling each function with the stack.
- **User definitions**: Stored as `map[string][]operatorTyp`. When a definition references another word, it captures the *current* operator list for that word (snapshot semantics).

### Files to Modify

- `go/exercises/practice/forth/forth.go` — Full implementation

### Rationale

This mirrors the reference solution architecture exactly. It's proven to pass all tests. The compile-then-execute model cleanly separates parsing from evaluation, making the code easy to follow. Snapshot semantics for user definitions fall out naturally because operators are resolved at parse time.

### Ordering

1. Define types (`operatorFn`, `operatorID`, `operatorTyp`)
2. Implement stack helpers (`push`, `pop`, `pop2`)
3. Implement built-in operators (`add`, `subtract`, `multiply`, `divide`, `dup`, `drop`, `swap`, `over`)
4. Define the `builtinOps` map
5. Implement `parse()` function
6. Implement `Forth()` function
7. Test

---

## Proposal B: Direct Interpretation Approach

**Role: Opponent**

This approach interprets tokens directly without a separate compile step. Each token is immediately executed against the stack.

### Architecture

- **Single-pass evaluation**: For each phrase, split into tokens and process each token immediately.
- **User definitions**: Stored as `map[string][]string` (lists of token strings). When a user-defined word is encountered during execution, its token list is recursively evaluated.
- **No operator types**: No need for operator ID enums or operator structs — just switch on the token string.

### Files to Modify

- `go/exercises/practice/forth/forth.go` — Full implementation

### Critique of Proposal A

Proposal A introduces unnecessary complexity with `operatorID` enums, `operatorTyp` structs, and a separate parse phase. For this simple subset of Forth, we don't need a compiled representation.

### Rationale

This approach is simpler — fewer types, less indirection. A single `eval` function with a switch statement is easier to understand. Storing definitions as token lists is more straightforward.

### Weakness

However, storing definitions as token strings means word resolution happens at *execution* time, not *definition* time. This breaks the test "can use different words with the same name" which requires snapshot semantics: when `bar` is defined as `foo` and `foo` is later redefined, `bar` should still use the old `foo`. Fixing this requires either:
1. Expanding user-defined words recursively at definition time (effectively re-implementing the compile step), or
2. Versioning the definition map

This significantly complicates what was supposed to be the "simpler" approach.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| **Correctness** | Proven correct (matches reference) | Requires complex fixes for snapshot semantics |
| **Risk** | Low — follows known-working pattern | Medium — snapshot semantics hard to retrofit |
| **Simplicity** | Moderate complexity but well-structured | Simpler surface but hidden complexity for definitions |
| **Consistency** | Matches `.meta/example.go` conventions exactly | Diverges from established patterns |

### Decision

**Proposal A wins.** The compiled operator list approach is the clear choice:

1. It correctly handles snapshot semantics for user-defined words without any special logic.
2. It matches the reference solution's architecture, ensuring consistency.
3. The "complexity" of operator types is minimal and pays for itself in correctness.
4. Proposal B's apparent simplicity is misleading — it would need equivalent or more complex code to handle definition semantics correctly.

### Detailed Implementation Plan

**File**: `go/exercises/practice/forth/forth.go`

#### Step 1: Package declaration and imports
```go
package forth

import (
    "errors"
    "strconv"
    "strings"
    "unicode"
)
```

#### Step 2: Type definitions
- `operatorFn func(stack *[]int) error` — function type for stack operations
- `operatorID byte` — enum for operator identification
- Constants: `opAdd`, `opSub`, `opMul`, `opDiv`, `opDrop`, `opDup`, `opSwap`, `opOver`, `opConst`, `opUserDef`, `opEndDef`
- `operatorTyp struct { fn operatorFn; id operatorID }` — pairs function with ID

#### Step 3: Stack helpers
- `push(stack *[]int, v int)` — append value
- `pop(stack *[]int) (int, error)` — remove and return top, error if empty
- `pop2(stack *[]int) (int, int, error)` — pop two values

#### Step 4: Built-in operator functions
- `add`, `subtract`, `multiply` — use `binaryOp` helper
- `divide` — special case for division by zero
- `dup` — pop one, push twice
- `drop` — pop one, discard
- `swap` — pop two, push in reverse order
- `over` — pop two, push back with copy of second

#### Step 5: Built-in operator map
- Map from uppercase string to `operatorTyp` for `+`, `-`, `*`, `/`, `DUP`, `DROP`, `SWAP`, `OVER`, `:`, `;`

#### Step 6: Parse function
- `parse(phrase string, userDefs map[string][]operatorTyp) ([]operatorTyp, error)`
- Split phrase on whitespace using `strings.FieldsFunc`
- For each token (uppercased):
  - Check user definitions first (inline the operator list)
  - Check built-in ops
  - If `:` (user def), parse word name, collect body operators until `;`, store in userDefs
  - Otherwise parse as integer, create push-closure

#### Step 7: Forth function
- `Forth(input []string) ([]int, error)`
- Initialize empty stack and user definitions map
- For each phrase, parse then execute all operators
- Return final stack

#### Step 8: Error variables
- `errNotEnoughOperands`
- `errDivideByZero`
- `errEmptyUserDef`
- `errInvalidUserDef`

#### Verification
- Run `go test ./...` in the exercise directory
- Run `go vet ./...`
- All 42 test cases must pass
