# Implementation Plan: Forth Evaluator

## File to Modify

- `go/exercises/practice/forth/forth.go` — the sole implementation file

## Architecture

The solution follows a two-phase approach per phrase: **parse** then **execute**.

### Data Structures

1. **`operatorFn`** — `func(stack *[]int) error` — a function that operates on the stack
2. **`operatorID`** — byte enum identifying operator type (used to detect `:` and `;` during parsing)
3. **`operatorTyp`** — struct combining `fn` and `id`
4. **Stack** — `[]int` slice, manipulated via `push`, `pop`, `pop2` helpers
5. **User definitions** — `map[string][]operatorTyp` mapping uppercased word names to their expanded operator lists

### Core Algorithm

**`Forth(input []string) ([]int, error)`**:
1. Initialize empty stack and user-definitions map
2. For each phrase in input:
   a. Parse the phrase into an operator list (expanding user defs inline, recording new defs)
   b. Execute each operator in the list against the stack
3. Return the stack

**`parse(phrase string, userDefs map[string][]operatorTyp) ([]operatorTyp, error)`**:
1. Split phrase on whitespace
2. For each token:
   - If token matches a user-defined word → expand (append its operator list)
   - If token is a built-in operator → append it
   - If token is `:` → begin user definition:
     - Next token is the word name (must not be a number)
     - Collect tokens until `;`, recursively parsing each one
     - Store the resulting operator list in userDefs (snapshot semantics: uses current defs at definition time)
   - Otherwise → parse as integer literal, create a const-push operator

### Built-in Operations

| Word | Stack Effect | Notes |
|------|-------------|-------|
| `+` | `(a b -- a+b)` | |
| `-` | `(a b -- a-b)` | |
| `*` | `(a b -- a*b)` | |
| `/` | `(a b -- a/b)` | Integer division; error on zero divisor |
| `DUP` | `(a -- a a)` | |
| `DROP` | `(a -- )` | |
| `SWAP` | `(a b -- b a)` | |
| `OVER` | `(a b -- a b a)` | |

### Error Cases

- Insufficient operands → error
- Division by zero → error
- Redefining a number → error
- Undefined word → error (token is not a number and not in builtins or user defs)

### Key Design Decisions

1. **Snapshot semantics for word definitions**: When a word is defined, its body is expanded immediately using the *current* definitions. This means redefining `foo` after using `foo` in the body of `bar` doesn't change `bar`'s behavior.
2. **Case insensitivity**: All tokens are uppercased before lookup.
3. **Negative number support**: `strconv.Atoi` naturally handles negative numbers with a leading `-`. The key is that `-1` is parsed as a number (not as minus followed by 1) because it's a single token.

## Implementation Order

1. Write helper functions: `push`, `pop`, `pop2`, `binaryOp`
2. Write arithmetic operators: `add`, `subtract`, `multiply`, `divide`
3. Write stack operators: `dup`, `drop`, `swap`, `over`
4. Define `builtinOps` map and error variables
5. Implement `parse` function
6. Implement `Forth` function
7. Run tests
