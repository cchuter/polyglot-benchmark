# Implementation Plan: polyglot-go-matrix

## Proposal A: Matrix as a Named Slice Type

**Role: Proponent**

### Approach

Define `Matrix` as a named slice type `type Matrix [][]int`. This is the simplest possible representation — the Matrix *is* the data.

### Files to Modify

- `go/exercises/practice/matrix/matrix.go` — implement the full solution

### Design

```go
type Matrix [][]int

func New(s string) (Matrix, error) { ... }
func (m Matrix) Rows() [][]int { ... }
func (m Matrix) Cols() [][]int { ... }
func (m Matrix) Set(row, col, val int) bool { ... }
```

### Rationale

- **Simplicity**: No struct wrapper, no indirection. The type *is* the data.
- **Nilability**: Slices are inherently nilable, so `var matrix Matrix` is nil and `matrix == nil` works.
- **Set works with value receiver**: Since slice headers share the underlying array, `m[r][c] = val` modifies the actual data even through a value receiver. No pointer receiver needed.
- **Deep copy for Rows/Cols**: Allocate new slices and copy element-by-element.
- **Minimal code**: Fewest lines, easiest to read.

### Parsing Logic

1. Split input string by `"\n"` into row strings
2. For each row string, trim leading/trailing spaces, split by whitespace
3. Validate: no empty rows, all rows same length, all values parseable as int
4. Convert each value via `strconv.Atoi`

### Weaknesses Addressed

- Value receiver on `Set` works because slices share underlying storage. This is idiomatic Go for slice-based types.

---

## Proposal B: Matrix as a Struct with Pointer Semantics

**Role: Opponent**

### Approach

Define `Matrix` as a pointer to a struct that holds the data. Use an unexported struct with an exported pointer type alias, or use an interface.

### Option B1: Pointer to struct

The test benchmark uses `var matrix Matrix` and then `matrix == nil`. If `Matrix` is a struct, this can't be nil. So we'd need `Matrix` to be `*matrixData`:

```go
type matrixData struct {
    data [][]int
}
type Matrix = *matrixData
```

But Go doesn't allow methods on type aliases of pointer types. So this doesn't work directly.

### Option B2: Interface

```go
type Matrix interface {
    Rows() [][]int
    Cols() [][]int
    Set(row, col, val int) bool
}

type matrix struct {
    data [][]int
}
```

### Rationale for B2

- **Encapsulation**: Internal data is hidden behind an interface. Users can't accidentally access `m[0][0]` directly.
- **Pointer semantics guaranteed**: `Set` modifies through a `*matrix` receiver, no ambiguity.
- **Flexibility**: Could swap implementations.

### Critique of Proposal A

- Relying on value-receiver-modifying-underlying-data for `Set` is subtle. Someone reading the code might expect a value receiver to not mutate.
- The `Matrix` type being `[][]int` means callers could bypass methods and index directly.

### Weaknesses of Proposal B

- More code: need both an interface and a struct type.
- Over-engineered for a simple exercise.
- The interface approach adds abstraction where none is needed.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion    | Proposal A (Slice)    | Proposal B (Interface) |
|-------------|----------------------|----------------------|
| Correctness | Fully satisfies all test requirements. Slice nilability, value receiver Set, deep copy Rows/Cols all work correctly. | Also correct, but more complex setup. |
| Risk        | Low. Slice-as-type is well-established Go pattern. | Low, but more surface area for bugs in the interface/struct split. |
| Simplicity  | Minimal code, single type, no indirection. | More types, more indirection, more code. |
| Consistency | Matches the exercise's minimalist stub (`package matrix` with no existing types). | Over-engineered for the codebase convention of simple exercise solutions (see `ledger.go`). |

### Decision

**Proposal A wins.** The named slice type is the simplest, most idiomatic approach. It satisfies all test requirements with minimal code. The value-receiver-modifies-underlying-data behavior is well-understood in Go for slice types and is not a "gotcha" — it's how slices work by design.

### Final Implementation Plan

**File**: `go/exercises/practice/matrix/matrix.go`

```go
package matrix

import (
    "fmt"
    "strconv"
    "strings"
)

// Matrix represents a matrix of integers as a 2D slice.
type Matrix [][]int

// New parses a string representation of a matrix into a Matrix.
func New(s string) (Matrix, error) {
    lines := strings.Split(s, "\n")
    m := make(Matrix, len(lines))
    var numCols int
    for i, line := range lines {
        fields := strings.Fields(line)
        if len(fields) == 0 {
            return nil, fmt.Errorf("row %d is empty", i)
        }
        if i == 0 {
            numCols = len(fields)
        } else if len(fields) != numCols {
            return nil, fmt.Errorf("row %d has %d columns, expected %d", i, len(fields), numCols)
        }
        row := make([]int, len(fields))
        for j, field := range fields {
            val, err := strconv.Atoi(field)
            if err != nil {
                return nil, fmt.Errorf("invalid value %q at row %d, col %d: %w", field, i, j, err)
            }
            row[j] = val
        }
        m[i] = row
    }
    return m, nil
}

// Rows returns a deep copy of all rows.
func (m Matrix) Rows() [][]int {
    result := make([][]int, len(m))
    for i, row := range m {
        result[i] = make([]int, len(row))
        copy(result[i], row)
    }
    return result
}

// Cols returns a deep copy of all columns.
func (m Matrix) Cols() [][]int {
    if len(m) == 0 {
        return nil
    }
    numCols := len(m[0])
    result := make([][]int, numCols)
    for j := 0; j < numCols; j++ {
        col := make([]int, len(m))
        for i := range m {
            col[i] = m[i][j]
        }
        result[j] = col
    }
    return result
}

// Set sets the value at (row, col). Returns false if out of bounds.
func (m Matrix) Set(row, col, val int) bool {
    if row < 0 || row >= len(m) {
        return false
    }
    if col < 0 || col >= len(m[row]) {
        return false
    }
    m[row][col] = val
    return true
}
```

### Steps

1. Write the implementation to `matrix.go`
2. Run `go test ./...` in the matrix directory
3. Run `go vet ./...` in the matrix directory
4. Fix any issues
5. Commit
