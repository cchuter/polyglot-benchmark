# Implementation Plan: polyglot-go-matrix

## Branch 1: Direct Slice-Based (Minimal / Reference Solution)

**Approach:** Define `Matrix` as `type Matrix [][]int`. Store data as a 2D slice of ints. This closely follows the reference example in `.meta/example.go`.

**Files to modify:**
- `go/exercises/practice/matrix/matrix.go` — single file, full implementation

**Architecture:**
- `Matrix` is a named type alias for `[][]int`
- `New` splits on `\n`, then `strings.Fields` each line, `strconv.Atoi` each field
- Validates: equal row lengths, no empty rows, valid integers
- `Rows()` returns a deep copy (new outer slice, each row copied via `append`)
- `Cols()` builds transposed copy
- `Set()` bounds-checks, sets value in-place

**Rationale:** Minimal code, follows existing example exactly, easy to understand.

**Evaluation:**
- Feasibility: Excellent — proven by reference solution
- Risk: Very low — straightforward implementation
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~60 lines, 1 file

---

## Branch 2: Struct-Based with Cached Dimensions

**Approach:** Wrap the data in a struct `type matrixData struct { data [][]int; nRows, nCols int }` and define `type Matrix = *matrixData`. Cache dimensions for faster bounds checking.

**Files to modify:**
- `go/exercises/practice/matrix/matrix.go`

**Architecture:**
- Store dimensions alongside data for O(1) bounds checks in `Set`
- Same parsing logic but store row/col counts
- `Rows()`/`Cols()` use cached dimensions to pre-allocate slices

**Rationale:** More extensible if we needed to add more operations later.

**Evaluation:**
- Feasibility: **Problematic** — the test file uses `Matrix` as a type that can be compared to `nil` (`got == nil`) and also uses `var matrix Matrix` in benchmarks. A struct type won't satisfy `== nil` unless it's a pointer. However the tests also call `matrix.Rows()` which requires `Matrix` to have methods. The benchmark declares `var matrix Matrix` then checks `matrix == nil`. This **only works if Matrix is a reference type** (slice, pointer, interface, map). A `type Matrix = *matrixData` type alias would work but is unusual.
- Risk: Medium — fragile type definition to satisfy nil checks
- Alignment: Can satisfy criteria but requires careful type definition
- Complexity: ~70 lines, more boilerplate

---

## Branch 3: Flat Array with Computed Indexing

**Approach:** Store matrix data as a flat `[]int` slice with row/column count, compute element positions via `row*cols + col`. Define Matrix as a pointer type.

**Files to modify:**
- `go/exercises/practice/matrix/matrix.go`

**Architecture:**
- `type matrixImpl struct { data []int; rows, cols int }`
- `type Matrix *matrixImpl` or interface-based
- `New` parses into flat array, records dimensions
- `Rows()` slices the flat array into row slices (copies)
- `Cols()` iterates with stride to build columns
- `Set()` computes flat index for O(1) access

**Rationale:** Better cache locality for large matrices, interesting from a performance angle.

**Evaluation:**
- Feasibility: Problematic — same nil-comparability issue as Branch 2. The test declares `var matrix Matrix` and checks `matrix == nil`, plus calls methods on it. Need a type that supports nil, methods, and value semantics from `New`.
- Risk: Higher — more complex indexing, potential off-by-one errors, type compatibility issues
- Alignment: Can satisfy criteria but over-engineered
- Complexity: ~80 lines, more error-prone

---

## Selected Plan

**Branch 1: Direct Slice-Based** is the clear winner.

**Rationale:**
1. The test file constrains the `Matrix` type significantly — it must be nil-comparable, support method calls, and work with `var matrix Matrix`. A `[][]int` named type satisfies all these naturally.
2. The reference solution already validates this approach works perfectly.
3. Minimal code means fewer bugs and faster implementation.
4. Branches 2 and 3 introduce type compatibility risks with no meaningful benefit for this exercise.

### Detailed Implementation Plan

**File:** `go/exercises/practice/matrix/matrix.go`

```go
package matrix

import (
    "errors"
    "strconv"
    "strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
    lines := strings.Split(s, "\n")
    m := make(Matrix, len(lines))
    for i, line := range lines {
        fields := strings.Fields(line)
        if len(fields) == 0 {
            return nil, errors.New("empty row")
        }
        if i > 0 && len(fields) != len(m[0]) {
            return nil, errors.New("uneven rows")
        }
        row := make([]int, len(fields))
        for j, f := range fields {
            val, err := strconv.Atoi(f)
            if err != nil {
                return nil, err
            }
            row[j] = val
        }
        m[i] = row
    }
    return m, nil
}

func (m Matrix) Rows() [][]int {
    result := make([][]int, len(m))
    for i, row := range m {
        result[i] = make([]int, len(row))
        copy(result[i], row)
    }
    return result
}

func (m Matrix) Cols() [][]int {
    if len(m) == 0 {
        return nil
    }
    nCols := len(m[0])
    result := make([][]int, nCols)
    for c := 0; c < nCols; c++ {
        col := make([]int, len(m))
        for r := 0; r < len(m); r++ {
            col[r] = m[r][c]
        }
        result[c] = col
    }
    return result
}

func (m Matrix) Set(row, col, val int) bool {
    if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
        return false
    }
    m[row][col] = val
    return true
}
```

### Steps:
1. Write `matrix.go` with the implementation above
2. Run tests to verify all pass
3. Commit the solution
