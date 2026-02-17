# Implementation Plan: Forth Evaluator

## Proposal A: Operator-Type Architecture (Compile-then-Execute)

**Role: Proponent**

### Approach

Follow the reference implementation's two-phase architecture: parse input into an operator list (compile), then execute operators against the stack. This approach separates concerns cleanly and handles user-defined words by expanding them at parse time.

### Architecture

1. **Operator types**: Define an `operatorFn` function type and `operatorTyp` struct pairing a function with an ID enum
2. **Parsing phase**: Tokenize each phrase, resolve words to operator lists, handle user definitions
3. **Execution phase**: Iterate operator list, apply each function to the stack
4. **User definitions**: Store as `map[string][]operatorTyp` — when a user word is referenced, its ops are spliced into the current op list

### File: `go/exercises/practice/forth/forth.go`

```
Types:
  - operatorFn func(stack *[]int) error
  - operatorID byte (enum for operator kinds)
  - operatorTyp struct { fn, id }

Main function: Forth(input []string) ([]int, error)
  - Initialize stack and userDefs map
  - For each phrase: parse → opList, then execute each op

parse(phrase, userDefs) → ([]operatorTyp, error)
  - Tokenize with strings.FieldsFunc (split on whitespace)
  - For each token (case-insensitive via ToUpper):
    - Check userDefs map first
    - Check builtinOps map
    - If ":" → parse user definition inline
    - Otherwise try strconv.Atoi for number literal
    - Unknown word → error

Stack helpers: pop, pop2, push, binaryOp
Operator functions: add, subtract, multiply, divide, dup, drop, swap, over
Error variables: errNotEnoughOperands, errDivideByZero, errEmptyUserDef, errInvalidUserDef
```

### Strengths

- **Proven correct**: Matches the reference implementation architecture exactly
- **Clean separation**: Parse and execute are distinct phases
- **Efficient**: User word definitions are expanded at parse time; no runtime lookup overhead
- **Word definition semantics**: Naturally captures meaning at definition time because ops are resolved during parse

### Ordering

1. Define types and constants
2. Implement stack helpers (pop, pop2, push)
3. Implement operator functions
4. Implement parse function
5. Implement Forth function
6. Test

---

## Proposal B: Direct Interpreter (Token-by-Token Execution)

**Role: Opponent**

### Approach

Skip the compile step entirely. Process tokens directly from the input, maintaining a stack and a definitions map that stores word definitions as string slices (not compiled operators). When a user-defined word is encountered, recursively expand and evaluate its definition tokens.

### Architecture

1. **No operator types needed** — evaluate tokens directly
2. **Definitions stored as `map[string][]string`** — raw token lists
3. **Single recursive `eval` function** handles everything
4. **Simpler code with fewer abstractions**

### File: `go/exercises/practice/forth/forth.go`

```
Main function: Forth(input []string) ([]int, error)
  - Initialize stack and defs map
  - For each phrase: tokenize, call eval(tokens, stack, defs)

eval(tokens []string, stack *[]int, defs map[string][]string) error
  - Iterate tokens
  - If ":" → collect definition, store in defs
  - If token in defs → recursively eval(defs[token], stack, defs)
  - If builtin → execute directly
  - If number → push
  - Else → error

Stack helpers: same as Proposal A
```

### Critique of Proposal A

- **Over-engineered**: The operator ID enum is unnecessary complexity — it's only used to detect end-of-definition markers
- **More types = more code**: The operatorTyp struct and operatorFn type add indirection without clear benefit for this simple problem
- **Closure capture subtlety**: Number literals create closures to capture values, adding cognitive complexity

### Critique of Proposal B (Self-Critique)

- **Definition-time semantics are WRONG**: Storing definitions as raw strings means word lookups happen at call time, not definition time. This breaks test case: `": foo 5 ;", ": bar foo ;", ": foo 6 ;", "bar foo"` which expects `[5, 6]`. With raw string storage, `bar` would evaluate `foo` at call time and get 6 instead of 5.
- **Fix needed**: Must snapshot definitions at definition time by expanding/copying the current meaning. This effectively reinvents the compile step.
- **Recursive evaluation risk**: Deep nesting of user-defined words could cause stack overflow

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | Fully correct — handles definition-time semantics naturally | **Broken** without significant rework — raw string storage gets definition-time semantics wrong |
| Risk | Low — follows proven reference implementation | High — must reinvent compilation to fix semantics |
| Simplicity | Moderate — some boilerplate with types/enums | Appears simpler but hidden complexity in fixing semantics |
| Consistency | Matches reference implementation style | Diverges from reference |

### Decision: Proposal A wins

Proposal B's fundamental flaw is that storing definitions as raw token strings produces wrong results for the "definitions capture meaning at definition time" requirement. Fixing this requires expanding definitions into resolved operations at definition time — which is essentially what Proposal A does from the start. Proposal A's type system (operatorTyp, operatorFn) is lightweight and serves a clear purpose.

### Final Implementation Plan

**File to modify**: `go/exercises/practice/forth/forth.go`

**Step 1: Types and Constants**
```go
type operatorFn func(stack *[]int) error
type operatorID byte

const (
    opAdd operatorID = iota
    opSub
    opMul
    opDiv
    opDrop
    opDup
    opSwap
    opOver
    opConst
    opUserDef
    opEndDef
)

type operatorTyp struct {
    fn operatorFn
    id operatorID
}
```

**Step 2: Error Variables**
```go
var errNotEnoughOperands = errors.New("not enough operands")
var errDivideByZero = errors.New("attempt to divide by zero")
var errEmptyUserDef = errors.New("empty user definition")
var errInvalidUserDef = errors.New("invalid user def word")
```

**Step 3: Stack Helpers**
- `pop(stack *[]int) (int, error)` — pop top value or error if empty
- `pop2(stack *[]int) (int, int, error)` — pop two values
- `push(stack *[]int, v int)` — push a value
- `binaryOp(stack *[]int, op func(a, b int) int) error` — pop two, apply op, push result

**Step 4: Operator Functions**
- `add`, `subtract`, `multiply` — use `binaryOp`
- `divide` — custom (check for zero divisor)
- `dup` — pop one, push twice
- `drop` — pop one, discard
- `swap` — pop two, push in reverse order
- `over` — pop two, push second, push first, push second again

**Step 5: Built-in Operators Map**
```go
var builtinOps = map[string]operatorTyp{
    "+": {add, opAdd}, "-": {subtract, opSub},
    "*": {multiply, opMul}, "/": {divide, opDiv},
    "DUP": {dup, opDup}, "DROP": {drop, opDrop},
    "SWAP": {swap, opSwap}, "OVER": {over, opOver},
    ":": {nil, opUserDef}, ";": {nil, opEndDef},
}
```

**Step 6: Parse Function**
```go
func parse(phrase string, userDefs map[string][]operatorTyp) ([]operatorTyp, error)
```
- Tokenize phrase using `strings.FieldsFunc` (split on whitespace/control chars)
- For each token (converted to uppercase):
  1. Check `userDefs` map — if found, append all ops from the definition
  2. Check `builtinOps` map — if found:
     - If `opUserDef` (":"): parse user definition inline (read name, validate not a number, collect body ops until ";", store in userDefs)
     - Otherwise: append the operator
  3. Try `strconv.Atoi` — if succeeds, create a closure-based constant operator
  4. Otherwise: return error (undefined word)

**Step 7: Main Forth Function**
```go
func Forth(input []string) ([]int, error)
```
- Handle empty input: return `[]int{}`
- Initialize stack and userDefs
- For each phrase: parse, then execute each operator against the stack
- Return final stack

**Key design decisions**:
- User definitions checked before builtins in parse, allowing overrides
- Definition bodies are compiled (resolved to ops) at definition time, capturing current word meanings
- Numbers create closures that capture the parsed integer value
- All words converted to uppercase for case-insensitivity

**Imports needed**: `errors`, `strconv`, `strings`, `unicode`
