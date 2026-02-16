# Implementation Plan: Forth Evaluator

## Branch 1: Direct Interpreter (Minimal, Simple)

**Approach**: A straightforward interpreter that processes tokens directly without any intermediate representation. Uses a map of word names to token lists for user-defined words.

**Architecture**:
- Single file: `forth.go`
- `Forth(input []string) ([]int, error)` as the entry point
- A `stack` (slice of int) and a `defs` map (string → []string) for user word definitions
- For each phrase, split into tokens, iterate through them
- When encountering a defined word, recursively expand it inline
- All tokens lowercased for case insensitivity

**Key design decisions**:
- Store user definitions as expanded token lists (snapshot at definition time)
- No type system for operators — just string matching in a switch statement
- Simple stack operations via helper functions

**Files**: Only `go/exercises/practice/forth/forth.go`

**Evaluation**:
- Feasibility: High — simple approach, well-understood pattern
- Risk: Low — no complex data structures; risk of infinite recursion on bad definitions, but test cases don't include that
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~120-150 lines of Go code, single file

---

## Branch 2: Compiled Operator List (Extensible, Reference-Based)

**Approach**: Parse each phrase into a list of operator structs (similar to the `.meta/example.go`), where each operator is a function closure. User-defined words are compiled into operator lists at definition time, achieving snapshot semantics naturally.

**Architecture**:
- Single file: `forth.go`
- Define `operatorFn func(stack *[]int) error` type
- Define operator ID enum for distinguishing operator types
- `parse()` function converts token strings into operator lists
- User-defined words stored as `map[string][]operator` — when a word is defined, its body is immediately compiled into an operator list
- At execution time, just iterate the operator list and call each function

**Key design decisions**:
- Closures capture constant values at parse time (snapshot semantics for free)
- Operator IDs distinguish definition start/end markers from executable ops
- Follows the same proven pattern as the example solution

**Files**: Only `go/exercises/practice/forth/forth.go`

**Evaluation**:
- Feasibility: High — proven pattern from example.go
- Risk: Low — well-tested approach, straightforward closure semantics
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~150-200 lines, single file, slightly more complex type definitions but cleaner execution

---

## Branch 3: Token-Stream Interpreter with Flat Expansion

**Approach**: Preprocess all user definitions by expanding them into flat token streams (no nested definitions in execution). Then execute the expanded token stream in a single pass.

**Architecture**:
- Single file: `forth.go`
- Two-pass approach: first pass collects definitions, second pass executes
- Definitions stored as `map[string][]string` (word name → expanded token list)
- When defining a word, immediately expand any referenced user-defined words into their token equivalents
- Execution is a simple loop over tokens with a switch on the lowercased token

**Key design decisions**:
- Eager expansion ensures snapshot semantics and eliminates runtime lookup overhead
- Single execution pass with no recursion
- Potentially faster for repeated word usage (all expansions resolved upfront per-phrase)

**Files**: Only `go/exercises/practice/forth/forth.go`

**Evaluation**:
- Feasibility: High — straightforward token manipulation
- Risk: Medium — two-pass approach per phrase adds complexity; must handle inter-phrase definitions carefully
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~130-160 lines, single file

---

## Selected Plan

**Selected: Branch 1 — Direct Interpreter**

**Rationale**: Branch 1 is the simplest approach with the lowest implementation complexity and risk. It achieves snapshot semantics by expanding user-defined words into token lists at definition time. Unlike Branch 2, it avoids the overhead of defining operator types and ID enums. Unlike Branch 3, it doesn't require a two-pass approach. The direct interpreter pattern is easy to understand, debug, and verify correctness.

### Detailed Implementation Plan

**File**: `go/exercises/practice/forth/forth.go`

**Structure**:

```go
package forth

import (
    "errors"
    "strconv"
    "strings"
)

// Error variables
var (
    errInsufficientOperands = errors.New("insufficient operands")
    errDivisionByZero       = errors.New("division by zero")
    errUndefinedWord        = errors.New("undefined word")
    errIllegalOperation     = errors.New("illegal operation")
)

// Forth evaluates a sequence of Forth phrases and returns the stack.
func Forth(input []string) ([]int, error) {
    stack := []int{}
    defs := map[string][]string{} // user-defined word → expanded token list

    for _, phrase := range input {
        tokens := strings.Fields(phrase)
        err := eval(tokens, &stack, defs)
        if err != nil {
            return nil, err
        }
    }
    return stack, nil
}

// eval processes a list of tokens, modifying the stack and defs.
func eval(tokens []string, stack *[]int, defs map[string][]string) error {
    for i := 0; i < len(tokens); i++ {
        token := strings.ToUpper(tokens[i])

        if token == ":" {
            // Parse user definition: : word-name body ;
            // Find matching ;
            // Extract word name (must not be a number)
            // Expand body tokens using current defs
            // Store in defs map
        } else if expanded, ok := defs[token]; ok {
            // Execute expanded tokens for user-defined word
            err := eval(expanded, stack, defs)
            if err != nil { return err }
        } else if n, err := strconv.Atoi(token); err == nil {
            // Push number
            *stack = append(*stack, n)
        } else {
            // Try built-in operation
            err := execBuiltin(token, stack)
            if err != nil { return err }
        }
    }
    return nil
}

// execBuiltin executes built-in arithmetic and stack operations.
func execBuiltin(op string, stack *[]int) error {
    // Switch on op for +, -, *, /, DUP, DROP, SWAP, OVER
    // Return errUndefinedWord for unknown operations
}
```

**Key implementation details**:

1. **User definition parsing**: When `:` is encountered, scan forward to find `;`. The first token after `:` is the word name. Validate it's not a number. The remaining tokens (before `;`) are the body. Expand each body token: if it's a known user-defined word, replace it with its expansion; otherwise keep it as-is. Store the expanded token list.

2. **Snapshot semantics**: By expanding user-defined words at definition time (not execution time), we get correct snapshot behavior. When `: foo 10 ;` then `: foo foo 1 + ;`, the second definition expands `foo` to `10` at definition time, resulting in body `[10, 1, +]`.

3. **Negative numbers**: `strconv.Atoi` handles negative numbers naturally (`"-1"` → -1).

4. **Case insensitivity**: All tokens uppercased before processing.

5. **Error handling**: Each built-in op checks stack depth before operating.

**Implementation order**:
1. Write error variables and helper functions (push/pop)
2. Write `execBuiltin` for arithmetic operators (+, -, *, /)
3. Add stack manipulation operations (DUP, DROP, SWAP, OVER)
4. Write `eval` with number pushing and built-in execution
5. Add user-defined word parsing and expansion
6. Write the top-level `Forth` function
7. Run tests and iterate
