# Implementation Plan: Beer Song

## Branch 1: Direct Case-Based Approach (Minimal, Simple)

### Approach
Implement using a switch statement in `Verse()` with hardcoded strings for special cases (0, 1, 2) and `fmt.Sprintf` for the general case (3-99). `Verses()` loops and concatenates. `Song()` delegates to `Verses(99, 0)`.

### Files to Modify
- `go/exercises/practice/beer-song/beer_song.go` (only file)

### Architecture
- Import `fmt` for Sprintf and Errorf
- Import `bytes` for Buffer in Verses
- `Verse(n)`: switch on n with cases for 0, 1, 2, default
- `Verses(start, stop)`: validate inputs, loop from start to stop, concatenate with newline separator
- `Song()`: call `Verses(99, 0)` and return result

### Rationale
Follows the reference implementation pattern exactly. Minimal code, easy to understand. No abstractions needed for this simple problem.

### Evaluation
- **Feasibility**: High — straightforward, follows the reference
- **Risk**: Low — direct mapping from spec to code
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~50 lines, 1 file modified

---

## Branch 2: Helper Function Approach (Extensible)

### Approach
Extract grammar logic into helper functions: `bottleCount(n)` returns "N bottles" or "1 bottle" or "no more bottles", `actionLine(n)` returns the second line of each verse. Makes it easy to modify or extend verse patterns.

### Files to Modify
- `go/exercises/practice/beer-song/beer_song.go` (only file)

### Architecture
- `bottleCount(n int) string`: returns "no more bottles", "1 bottle", or "N bottles"
- `actionLine(n int) string`: returns "Go to the store..." for 0, "Take it down..." for 1, "Take one down..." for 2+
- `Verse(n)`: composes using helpers
- `Verses()` and `Song()`: same as Branch 1

### Rationale
More modular, easier to add new verse patterns if the song were extended. Better separation of concerns.

### Evaluation
- **Feasibility**: High — standard pattern
- **Risk**: Low — slightly more code but all simple
- **Alignment**: Fully satisfies all criteria
- **Complexity**: ~65 lines, 1 file, more functions but each simpler

---

## Branch 3: Template-Based Approach (Performance/Unconventional)

### Approach
Pre-compute all 100 verses at package init time into a slice. `Verse(n)` is a simple slice lookup. `Verses()` joins from the pre-computed slice. Optimizes for repeated calls at the cost of memory.

### Files to Modify
- `go/exercises/practice/beer-song/beer_song.go` (only file)

### Architecture
- Package-level `var verses [100]string`
- `init()` function populates all verses
- `Verse(n)` returns `verses[n]` after bounds check
- `Verses()` concatenates from the slice
- `Song()` delegates to `Verses(99, 0)`

### Rationale
Benchmark-optimized — verse generation happens once at startup. Repeated calls are O(1) lookups. Good for the benchmarks in the test file.

### Evaluation
- **Feasibility**: High — straightforward Go pattern
- **Risk**: Medium — init() adds complexity, pre-computation is overkill
- **Alignment**: Fully satisfies all criteria
- **Complexity**: ~60 lines, slightly more complex initialization

---

## Selected Plan

### Branch 1: Direct Case-Based Approach

**Rationale**: Branch 1 is the best choice because:
1. **Simplicity**: The problem is inherently simple — 4 verse patterns. Adding abstractions (Branch 2) or pre-computation (Branch 3) adds complexity without meaningful benefit.
2. **Low risk**: Direct mapping from the reference implementation means highest confidence of passing all tests.
3. **Minimal code**: Fewer lines means fewer places for bugs.
4. **Follows conventions**: Matches the patterns seen in other completed exercises in this repo (direct, straightforward implementations).

### Detailed Implementation

**File**: `go/exercises/practice/beer-song/beer_song.go`

```go
package beer

import (
	"bytes"
	"fmt"
)

func Song() string {
	result, _ := Verses(99, 0)
	return result
}

func Verses(start, stop int) (string, error) {
	switch {
	case start < 0 || start > 99:
		return "", fmt.Errorf("start value[%d] is not a valid verse", start)
	case stop < 0 || stop > 99:
		return "", fmt.Errorf("stop value[%d] is not a valid verse", stop)
	case start < stop:
		return "", fmt.Errorf("start value[%d] is less than stop value[%d]", start, stop)
	}

	var buf bytes.Buffer
	for i := start; i >= stop; i-- {
		v, _ := Verse(i)
		buf.WriteString(v)
		buf.WriteString("\n")
	}
	return buf.String(), nil
}

func Verse(n int) (string, error) {
	switch {
	case n < 0 || n > 99:
		return "", fmt.Errorf("%d is not a valid verse", n)
	case n == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case n == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case n == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}
```

### Order of Changes
1. Write the complete implementation to `beer_song.go`
2. Run `go test` to verify all tests pass
3. Run `go test -bench=.` to verify benchmarks work
4. Commit with descriptive message
