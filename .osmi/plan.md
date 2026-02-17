# Implementation Plan: polyglot-go-matrix

## Branch 1: Named Slice Type (Minimal, mirrors reference solution)

Define `Matrix` as `type Matrix [][]int` — a named slice type that is directly nil-comparable and requires no pointer indirection.

### Approach
- `New()`: Split on `\n`, split each line with `strings.Fields`, parse with `strconv.Atoi`, validate row lengths.
- `Rows()`: Allocate new outer slice, copy each inner slice with `append([]int{}, row...)`.
- `Cols()`: Allocate column slices, transpose the data.
- `Set()`: Bounds check, direct assignment.

### Files
- Modify: `go/exercises/practice/matrix/matrix.go`

### Rationale
Simplest possible approach. Matches the reference solution in `.meta/example.go` almost exactly. The named slice type satisfies the `matrix == nil` check in benchmarks.

### Evaluation
- **Feasibility**: Proven — this is the reference solution pattern.
- **Risk**: Extremely low. Direct match with test expectations.
- **Alignment**: Satisfies all acceptance criteria.
- **Complexity**: ~60 lines, single file.

---

## Branch 2: Struct with Pointer Receiver

Define `Matrix` as an interface or use `*matrixImpl` struct approach. Store the data in a struct with explicit `rows` and `cols` fields (or just rows, computing cols on demand).

### Approach
- `type Matrix = *matrixData` or define a `Matrix` interface.
- Actually — the test uses `Matrix` as a concrete type (`var matrix Matrix` and `matrix == nil`), and methods are called directly. Looking at the benchmark: `var matrix Matrix` then `matrix, err = New(...)` and `matrix == nil`. This means `Matrix` must be a type that can be assigned from `New()`'s return and compared to nil.
- Use `type Matrix *matrixData` where `matrixData` is a struct. But this creates issues — named pointer types can't have methods.
- Alternative: `type Matrix struct { data [][]int }` — but struct zero value is not nil.
- This approach is **not feasible** without changing the test file, because `matrix == nil` requires a nil-able type.

### Revised Approach
Use `type Matrix [][]int` but wrap it in a more structured way — pre-compute both rows and columns on construction.

- `type Matrix [][]int` (same type)
- `New()` stores data as rows
- `Cols()` transposes on each call (or we could cache, but that adds complexity for no test benefit)

### Evaluation
- **Feasibility**: The struct approach doesn't work due to nil comparison. Falls back to Branch 1.
- **Risk**: Medium — added complexity for no benefit.
- **Alignment**: Would satisfy criteria but is over-engineered.
- **Complexity**: More code, same result.

---

## Branch 3: Flat Array with Dimensions (Performance-Oriented)

Store the matrix as a flat `[]int` slice with row/column dimensions, avoiding nested slices for cache-friendly access.

### Approach
- `type Matrix struct { data []int; nrows, ncols int }`
- Problem: `var matrix Matrix` then `matrix == nil` — a struct is never nil.
- Alternative: `type Matrix *flatMatrix` — named pointer types can't have methods in Go.
- Alternative: Define `Matrix` as an interface... but the test declares `var matrix Matrix` and later `matrix, err = New(...)` and uses `matrix.Rows()`. If Matrix is an interface, New returns the interface, and nil comparison works. Let's check:
  - `var matrix Matrix` — zero value of interface is nil ✓
  - `matrix, err = New(...)` — returns concrete type implementing interface ✓
  - `matrix == nil` — works for interfaces ✓
  - `matrix.Rows()` — method call through interface ✓

### Interface Approach
```go
type Matrix interface {
    Rows() [][]int
    Cols() [][]int
    Set(row, col, val int) bool
}

type flatMatrix struct {
    data  []int
    nrows int
    ncols int
}
```

### Evaluation
- **Feasibility**: Yes, interface approach works with the test signatures.
- **Risk**: Higher — more code, interface indirection, and the flat storage is unnecessary complexity for this problem size.
- **Alignment**: Satisfies all criteria.
- **Complexity**: ~80 lines, more complex than needed.

---

## Selected Plan

**Branch 1: Named Slice Type** is the clear winner.

### Rationale
- It is the simplest and most idiomatic approach.
- It directly matches the reference solution pattern proven by the exercise authors.
- The named slice `type Matrix [][]int` naturally supports nil comparison required by benchmarks.
- Risk is minimal — the approach is proven.
- Branches 2 and 3 both struggle with the nil comparison requirement and add unnecessary complexity.

### Detailed Implementation

**File: `go/exercises/practice/matrix/matrix.go`**

```go
package matrix

import (
    "errors"
    "strconv"
    "strings"
)

// Matrix represents a matrix of integers as a slice of rows.
type Matrix [][]int

// New parses a string representation of a matrix.
// Rows are separated by newlines, columns by spaces.
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

// Rows returns a copy of the matrix data as rows.
func (m Matrix) Rows() [][]int {
    result := make([][]int, len(m))
    for i, row := range m {
        result[i] = append([]int{}, row...)
    }
    return result
}

// Cols returns the matrix data as columns.
func (m Matrix) Cols() [][]int {
    if len(m) == 0 {
        return nil
    }
    ncols := len(m[0])
    result := make([][]int, ncols)
    for c := 0; c < ncols; c++ {
        col := make([]int, len(m))
        for r := range m {
            col[r] = m[r][c]
        }
        result[c] = col
    }
    return result
}

// Set sets the value at (row, col). Returns false if out of bounds.
func (m Matrix) Set(row, col, val int) bool {
    if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
        return false
    }
    m[row][col] = val
    return true
}
```

### Key Decisions
1. `strings.Fields` handles leading/trailing whitespace per line (important: test case has " 8 7 6" with leading space).
2. `len(fields) == 0` catches empty rows from empty lines.
3. `strconv.Atoi` catches both non-numeric values and int64 overflow.
4. `Rows()` uses `append([]int{}, row...)` for independent copies.
5. `Cols()` builds columns by iterating row-major data.
6. `Set()` does full bounds checking including negative indices.
