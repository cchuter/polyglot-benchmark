# Plan Review: polyglot-go-matrix

## Verdict: PASS (with minor observations)

The plan is well-structured, correctly analyzed, and the proposed implementation will pass all tests. The selected approach (Proposal A: named slice type) is the right choice and matches the canonical reference solution in `.meta/example.go`. No blocking issues were found.

---

## Detailed Analysis

### 1. Type Design: Correct

The choice of `type Matrix [][]int` is correct for the following reasons:

- **Nilability**: The test benchmark at line 277-288 declares `var matrix Matrix` and checks `matrix == nil`. A named slice type is inherently nilable -- its zero value is `nil`. This works correctly.
- **Matches reference**: The `.meta/example.go` canonical solution uses the exact same type definition.
- **Set with value receiver**: The plan correctly identifies that `m[row][col] = val` on a value receiver modifies the underlying data because Go slices are reference types (the slice header contains a pointer to the backing array). The test at lines 222-271 creates a matrix, calls `Set`, then verifies via `Rows()` and `Cols()` that the mutation took effect. This will work as designed.

### 2. Parsing Logic: Correct

Tracing the plan's `New` function against all invalid test cases:

| Test Case | Input | Plan Behavior | Correct? |
|-----------|-------|--------------|----------|
| int64 overflow | `"9223372036854775808"` | `strconv.Atoi` returns error (exceeds `math.MaxInt64` on 64-bit) | Yes |
| uneven rows | `"1 2\n10 20 30"` | Row 1 has 3 fields vs row 0's 2 fields; `len(fields) != numCols` error | Yes |
| first row empty | `"\n3 4\n5 6"` | `strings.Split` yields `["", "3 4", "5 6"]`; `strings.Fields("")` returns `[]`; `len(fields) == 0` error | Yes |
| middle row empty | `"1 2\n\n5 6"` | Empty string in middle; `strings.Fields("")` returns `[]`; `len(fields) == 0` error | Yes |
| last row empty | `"1 2\n3 4\n"` | Trailing `\n` creates empty last element; `strings.Fields("")` returns `[]`; `len(fields) == 0` error | Yes |
| non-integer | `"2.7"` | `strconv.Atoi("2.7")` returns error | Yes |
| non-numeric | `"cat"` | `strconv.Atoi("cat")` returns error | Yes |

All valid test cases also parse correctly. In particular, the case with leading whitespace (`" 8 7 6"` on line 73) is handled because `strings.Fields` splits on any whitespace and ignores leading/trailing whitespace. This is correct.

### 3. Deep Copy (Rows/Cols): Correct

The test at lines 189-191 modifies `rows[0][0]` after calling `Rows()` and then verifies that a subsequent `Rows()` call still returns the original data. The plan allocates new slices and uses `copy()`, producing an independent deep copy. Same logic applies to `Cols()` at lines 214-216.

The plan's implementation is equivalent to the reference solution's `append([]int{}, mr...)` idiom -- both produce independent copies. The plan's approach using `make` + `copy` is slightly more explicit and equally correct.

### 4. Set Bounds Checking: Correct

The test at lines 260-270 checks out-of-bounds Set calls with indices `{-2, -1, 0, 3, 4}` for both row and col (skipping `(0,0)` which is valid). The plan checks:
- `row < 0 || row >= len(m)` -- catches negative and too-large row indices
- `col < 0 || col >= len(m[row])` -- catches negative and too-large col indices

This correctly returns `false` for all out-of-bounds cases and `true` for all valid positions.

### 5. Minor Observations (Non-blocking)

**Observation 1: `fmt` import may be unnecessary overhead.**
The plan uses `fmt.Errorf` with `%w` wrapping for error messages. The reference solution uses `errors.New` with simple string messages. Both work -- the tests only check `err == nil` vs `err != nil`, never inspecting the error message itself. Using `fmt.Errorf` is perfectly fine but adds a dependency on the `fmt` package. This has no functional impact.

**Observation 2: The plan's error for empty rows is slightly different from what the reference does.**
The plan explicitly checks `len(fields) == 0` and returns an error with a descriptive message. The reference solution does NOT check for empty rows explicitly -- instead it relies on the `len(ws) != len(m[0])` check on subsequent rows to catch some (but not all) empty-row cases. This means the reference solution would actually fail the `"\n3 4\n5 6"` test case (first row is empty, so `m[0]` has length 0, and row 1 would compare `2 != 0` and error), but would handle it via a less descriptive error path. The plan's explicit empty-row check is actually **better** than the reference solution and handles all cases cleanly.

**Observation 3: No issues with the `strconv.Atoi` overflow behavior.**
On all Go target platforms with `go 1.18`, `int` is at least 64 bits on 64-bit systems. `strconv.Atoi` internally uses `strconv.ParseInt` with `bitSize` matching the platform's `int` size. The value `9223372036854775808` (2^63) overflows `int64`, so `Atoi` will return a `*NumError` with `Err: ErrRange`. This correctly triggers the error path. If this were a 32-bit system, the threshold would be even lower, so it would still fail. No issue here.

---

## Conclusion

The plan is sound. The implementation as written in the "Final Implementation Plan" section will pass all 9 valid test cases, all 7 invalid test cases, the `TestSet` test (including both valid and out-of-bounds positions), and all 3 benchmarks. No changes are needed. Proceed with implementation.
