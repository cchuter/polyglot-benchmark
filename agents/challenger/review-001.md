# Code Review: Forth Evaluator Implementation

**Reviewer:** challenger
**File:** `go/exercises/practice/forth/forth.go`
**Date:** 2026-02-16
**Verdict:** PASS — All 48 test cases pass. No correctness issues found.

---

## Test Results

```
go test -v -count=1 ./...
--- PASS: TestForth (0.00s) — 48/48 subtests passed
```

---

## Critical Test Case Walk-Through

### 1. Numbers pushed onto stack: `"1 2 3 4 5"` → [1,2,3,4,5]
**PASS.** Each token parsed via `strconv.Atoi`, closure captures value correctly.

### 2. Negative numbers: `"-1 -2 -3 -4 -5"` → [-1,-2,-3,-4,-5]
**PASS.** `-1` uppercased is still `-1`, not in builtinOps, `strconv.Atoi("-1")` returns -1 correctly. No collision with `-` operator.

### 3. Definition-time capture: `": foo 5 ;"`, `": bar foo ;"`, `": foo 6 ;"`, `"bar foo"` → [5, 6]
**PASS.** When `bar` is defined (phrase 2), `parse("foo", userDefs)` resolves `FOO` from userDefs and copies the ops (push 5) into bar's definition. When `foo` is redefined to push 6 (phrase 3), `bar`'s definition is unaffected because Go slices were copied by value via `append(oplist, udef...)`. This is the key correctness requirement for Forth definition-time semantics.

### 4. Self-referential redefinition: `": foo 10 ;"`, `": foo foo 1 + ;"`, `"foo"` → [11]
**PASS.** Second definition of `foo` references the *current* meaning of `foo` (push 10) at parse time. The new definition becomes [push 10, push 1, add]. Executing `foo` gives 10+1=11.

### 5. Cannot redefine negative numbers: `": -1 2 ;"` → error
**PASS.** `strconv.Atoi("-1")` succeeds (returns -1), triggering `errInvalidUserDef`.

### 6. Case-insensitive SWAP: `"1 2 SWAP 3 Swap 4 swap"` → [2,3,4,1]
**PASS.** All variants uppercased to "SWAP", found in builtinOps. Stack trace:
- push 1,2 → [1,2]; SWAP → [2,1]; push 3 → [2,1,3]; SWAP → [2,3,1]; push 4 → [2,3,1,4]; SWAP → [2,3,4,1].

### 7. Override built-in operators: `": + * ;"`, `"3 4 +"` → [12]
**PASS.** User definition `userDefs["+"] = [{multiply}]` is checked before builtinOps in parse. `3 4 +` evaluates as 3*4=12.

### 8. Override built-in words: `": swap dup ;"`, `"1 swap"` → [1, 1]
**PASS.** `userDefs["SWAP"] = [{dup}]`. `1 swap` executes as push 1, then dup → [1,1].

---

## Architecture Review

### Plan Adherence
The implementation faithfully follows the selected plan (Proposal A: Compile-then-Execute):
- **Types & constants:** `operatorFn`, `operatorID`, `operatorTyp` — all present as specified
- **Two-phase architecture:** `parse()` compiles to operator list, `Forth()` executes — correct
- **User definitions:** Stored as `map[string][]operatorTyp` — correct
- **Stack helpers:** `pop`, `pop2`, `push`, `binaryOp` — all present
- **Operator functions:** `add`, `subtract`, `multiply`, `divide`, `dup`, `drop`, `swap`, `over` — all present
- **Error variables:** 4 sentinel errors — all present
- **Imports:** `errors`, `strconv`, `strings`, `unicode` — matches plan exactly

### Correctness Analysis

| Category | Status | Notes |
|----------|--------|-------|
| Arithmetic (+, -, *, /) | PASS | `binaryOp` correctly pops v1 (top), v2 (second), computes `op(v2, v1)` |
| Division by zero | PASS | Explicitly checked in `divide()` before executing |
| Stack ops (dup, drop, swap, over) | PASS | All correctly manipulate stack |
| User definitions | PASS | Ops resolved at definition time (parse-time expansion) |
| Definition-time capture | PASS | `append(oplist, udef...)` copies ops by value |
| Self-referential definitions | PASS | Current meaning of word resolved before redefinition stored |
| Override builtins | PASS | `userDefs` checked before `builtinOps` in parse |
| Number redefinition rejection | PASS | `strconv.Atoi` check on definition name |
| Case insensitivity | PASS | `strings.ToUpper()` applied to all tokens |
| Empty stack errors | PASS | `pop()` returns error when stack is empty |
| Undefined word errors | PASS | Falls through to `strconv.Atoi`, which fails for non-numeric words |
| Empty input | PASS | Returns `[]int{}` for `len(input) == 0` |

### Security & Robustness

1. **No nil pointer dereferences:** All pointer operations use `*[]int` pattern safely. The `builtinOps[":"]` and `builtinOps[";"]` have `nil` fn, but these are never executed — `:` is handled by special-case branching (`op.id == opUserDef`) and `;` is detected by `oneOp[0].id == opEndDef`.

2. **Bounds checking:** `pop()` checks `len(*stack) == 0` before access. Definition parsing checks `t >= len(words)-2` before proceeding.

3. **Closure capture:** Number literals use a local variable `x` scoped within the for-loop iteration, ensuring each closure captures its own value. No shared-variable bug.

4. **Recursive parse safety:** `parse()` is called recursively for single words within definition bodies. A single non-`:` word cannot trigger further recursion, so depth is bounded to 1 level.

5. **Map mutation during iteration:** `userDefs` is modified only after the definition body is fully parsed, not during body parsing. No concurrent modification issues.

---

## Minor Observations (Non-Blocking)

1. **Error message for undefined words:** Returns the raw `strconv.Atoi` error (e.g., `strconv.Atoi: parsing "FOO": invalid syntax`) rather than a semantic error like "undefined word: FOO". This is acceptable because the test runner (`forth_test.go:12-21`) only checks error presence (nil vs non-nil), not error messages.

2. **Incomplete definitions without `;`:** Input like `": foo 5"` (missing semicolon) returns `errEmptyUserDef` due to the boundary check `t >= len(words)-2`. The error message is slightly misleading (it's a missing semicolon, not an empty definition), but no test exercises this case.

3. **operatorID enum:** `opConst` and `opUserDef`/`opEndDef` IDs are only used for sentinel detection in definition parsing. The runtime execution phase only uses `fn`. This is a minor over-abstraction but matches the plan exactly.

---

## Conclusion

**The implementation is correct, complete, and ready to merge.** All 48 test cases pass. The code faithfully implements the plan's Proposal A architecture. Definition-time capture semantics — the most critical correctness requirement — are handled correctly through parse-time operator expansion. No blocking issues identified.
