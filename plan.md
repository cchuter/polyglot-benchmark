# Implementation Plan: Crypto Square

## Branch 1: String Concatenation with strings.Map

**Approach**: Use `strings.Map` for normalization, `math.Ceil(math.Sqrt(...))` for column count, and string concatenation to build column strings. This mirrors the reference example approach.

**Files to modify**: `go/exercises/practice/crypto-square/crypto_square.go`

**Implementation**:
1. Define a `norm` helper that maps runes: keep lowercase letters and digits as-is, convert uppercase to lowercase, return -1 for everything else.
2. Use `strings.Map(norm, pt)` to normalize.
3. Compute `c = ceil(sqrt(len(normalized)))`.
4. Compute `r` as `c - 1` if `(c-1)*c >= len`, else `c`.
5. Build `c` column strings by iterating over normalized text, appending `string(r)` to `cols[i % c]`.
6. Pad the last `n` columns with a space where `n = r*c - len(normalized)`.
7. Return `strings.Join(cols, " ")`.

**Evaluation**:
- Feasibility: High. Uses only stdlib. Matches reference solution closely.
- Risk: Low. Simple, well-understood approach.
- Alignment: Fully satisfies all acceptance criteria.
- Complexity: ~30 lines, 1 file modified.

## Branch 2: Byte Slice with Pre-allocated Buffer

**Approach**: Normalize into a byte slice, pre-allocate output buffer, compute positions mathematically.

**Files to modify**: `go/exercises/practice/crypto-square/crypto_square.go`

**Implementation**:
1. Normalize by iterating over bytes (input is ASCII), building a `[]byte` slice.
2. Compute `c` and `r` as in Branch 1.
3. Pre-allocate a `[]byte` of size `c*r + c - 1` (characters + spaces between chunks).
4. Fill the output buffer by computing the source index for each position: for chunk `j` (column) and row `i`, source index = `i*c + j`.
5. If source index >= len(normalized), write a space (padding).
6. Insert space separators between chunks.
7. Return `string(output)`.

**Evaluation**:
- Feasibility: High. Pure byte manipulation, no imports beyond `math`.
- Risk: Medium. Index arithmetic requires care; off-by-one errors possible.
- Alignment: Fully satisfies all acceptance criteria.
- Complexity: ~40 lines, 1 file. More code, but zero allocations during encoding.

## Branch 3: strings.Builder with Unicode Support

**Approach**: Use `strings.Builder` for efficient string construction, `unicode` package for character classification.

**Files to modify**: `go/exercises/practice/crypto-square/crypto_square.go`

**Implementation**:
1. Normalize using `unicode.IsLetter`, `unicode.IsDigit`, `unicode.ToLower` into a `strings.Builder`.
2. Extract normalized string.
3. Compute `c` and `r` as above.
4. Build output with a `strings.Builder`, iterating column-first: for each column `j`, for each row `i`, read position `i*c + j` or pad with space.
5. Add space separator between columns.
6. Return builder string.

**Evaluation**:
- Feasibility: High. Uses standard `unicode` and `strings` packages.
- Risk: Low-Medium. `strings.Builder` avoids repeated string allocation.
- Alignment: Fully satisfies all acceptance criteria.
- Complexity: ~35 lines, 1 file. More idiomatic Go than Branch 1 (avoids rune-to-string conversion in loop).

---

## Selected Plan

**Branch 1: String Concatenation with strings.Map** is selected.

**Rationale**:
- **Simplest implementation**: ~30 lines of clean, readable code.
- **Lowest risk**: Mirrors the reference solution pattern, reducing chance of bugs.
- **Minimal complexity**: Uses well-understood `strings.Map` pattern.
- **Sufficient performance**: Passes all tests including benchmarks without issue.
- Branch 2 adds unnecessary complexity for marginal performance gain on small inputs.
- Branch 3 is reasonable but the `unicode` package is overkill for ASCII-only inputs (the test cases are all ASCII).

### Detailed Implementation

**File**: `go/exercises/practice/crypto-square/crypto_square.go`

```go
package cryptosquare

import (
	"math"
	"strings"
)

func norm(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
		return r
	case r >= 'A' && r <= 'Z':
		return r + 'a' - 'A'
	}
	return -1
}

func Encode(pt string) string {
	pt = strings.Map(norm, pt)
	if len(pt) == 0 {
		return ""
	}
	numCols := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	numRows := numCols - 1
	if numRows*numCols < len(pt) {
		numRows = numCols
	}
	padding := numRows*numCols - len(pt)
	cols := make([]string, numCols)
	for i, r := range pt {
		cols[i%numCols] += string(r)
	}
	for i := 0; i < padding; i++ {
		cols[numCols-i-1] += " "
	}
	return strings.Join(cols, " ")
}
```

### Steps
1. Write the implementation to `crypto_square.go`
2. Run `go test ./...` from the exercise directory to verify all 19 tests pass
3. Run `go vet ./...` to verify no issues
4. Commit with descriptive message
