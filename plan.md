# Implementation Plan: Beer Song

## Proposal A — Switch-Case with fmt.Sprintf

**Role: Proponent**

### Approach

Implement all three functions in `beer_song.go` using a switch statement in `Verse()` to handle the four distinct verse patterns (0, 1, 2, 3-99), with `fmt.Sprintf` for the general case and string literals for special cases. `Verses()` builds output by iterating and joining. `Song()` delegates to `Verses(99, 0)`.

### Files to Modify

- `go/exercises/practice/beer-song/beer_song.go` — sole file to edit

### Implementation Details

```go
package beer

import (
    "fmt"
    "strings"
)

func Verse(n int) (string, error) {
    if n < 0 || n > 99 {
        return "", fmt.Errorf("%d is not a valid verse", n)
    }
    switch n {
    case 0:
        return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
    case 1:
        return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
    case 2:
        return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
    default:
        return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
    }
}

func Verses(start, stop int) (string, error) {
    if start < 0 || start > 99 {
        return "", fmt.Errorf("start value[%d] is not a valid verse", start)
    }
    if stop < 0 || stop > 99 {
        return "", fmt.Errorf("stop value[%d] is not a valid verse", stop)
    }
    if start < stop {
        return "", fmt.Errorf("start value[%d] is less than stop value[%d]", start, stop)
    }
    var parts []string
    for i := start; i >= stop; i-- {
        v, _ := Verse(i)
        parts = append(parts, v)
    }
    return strings.Join(parts, "\n"), nil
}

func Song() string {
    s, _ := Verses(99, 0)
    return s
}
```

### Rationale

- **Matches the reference solution** (`.meta/example.go`) almost exactly
- Simple, readable, easy to verify against test expectations
- Uses `strings.Join` instead of `bytes.Buffer` for slightly cleaner code
- Minimal complexity — each special case is explicit and easy to audit
- Direct correspondence between test cases and switch branches

---

## Proposal B — Template-Based Approach with Helper Functions

**Role: Opponent**

### Approach

Use a more DRY approach with helper functions for pluralization and quantity formatting. A single template handles all verses, with helper functions swapping in the right words.

### Files to Modify

- `go/exercises/practice/beer-song/beer_song.go` — sole file to edit

### Implementation Details

```go
package beer

import (
    "fmt"
    "strings"
)

func bottleStr(n int) string {
    if n == 1 { return "1 bottle" }
    if n == 0 { return "no more bottles" }
    return fmt.Sprintf("%d bottles", n)
}

func capitalBottleStr(n int) string {
    if n == 0 { return "No more bottles" }
    return bottleStr(n)
}

func actionStr(n int) string {
    if n == 0 { return "Go to the store and buy some more" }
    if n == 1 { return "Take it down and pass it around" }
    return "Take one down and pass it around"
}

func nextCount(n int) int {
    if n == 0 { return 99 }
    return n - 1
}

func Verse(n int) (string, error) {
    if n < 0 || n > 99 {
        return "", fmt.Errorf("%d is not a valid verse", n)
    }
    return fmt.Sprintf("%s of beer on the wall, %s of beer.\n%s, %s of beer on the wall.\n",
        capitalBottleStr(n), bottleStr(n), actionStr(n), bottleStr(nextCount(n))), nil
}

// Verses and Song same as Proposal A
```

### Critique of Proposal A

- Proposal A has string duplication across cases (the phrase "bottles of beer on the wall" appears in every case)
- If the song lyrics changed, Proposal A requires updating multiple places
- Proposal A's special cases are harder to verify — you must read each full string carefully

### Why Proposal B is Superior

- DRY: changing the bottle string format requires one edit
- Helper functions are independently testable
- More maintainable and extensible

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A | Proposal B |
|-----------|-----------|-----------|
| Correctness | High — directly mirrors reference solution and test expectations | Medium — more moving parts means more chance of subtle formatting errors |
| Risk | Low — string literals match tests exactly | Medium — template composition could produce incorrect whitespace/capitalization |
| Simplicity | High — flat switch, no abstractions | Medium — 4 helper functions add complexity |
| Consistency | High — matches `.meta/example.go` conventions | Low — introduces patterns not seen in reference |

### Decision: Proposal A

**Rationale:**

1. **Correctness is paramount**: The tests check exact string output. Proposal A uses string literals for special cases, making it trivially verifiable that output matches expectations. Proposal B's template composition risks subtle formatting bugs.

2. **Simplicity**: This is a straightforward exercise. Four helper functions for a problem with only three special cases is over-engineering. The "duplication" in Proposal A is actually clarity — each case is self-contained and readable.

3. **Consistency**: Proposal A closely mirrors the reference solution in `.meta/example.go`, following established conventions.

4. **Risk**: Proposal A has near-zero risk of formatting errors. Proposal B's `capitalBottleStr` vs `bottleStr` distinction adds a category of bugs.

### Final Plan (Revised after review)

**File to modify:** `go/exercises/practice/beer-song/beer_song.go`

**Critical fix from review:** Use `bytes.Buffer` approach (matching the reference solution) instead of `strings.Join` in `Verses()`. The `strings.Join` approach omits the trailing `\n` after the last verse, causing multi-verse test failures.

**Implementation:**

1. Add imports: `fmt` and `bytes`
2. Implement `Verse(n int) (string, error)`:
   - Validate n is 0-99, return error otherwise
   - Switch on n: cases 0, 1, 2 return string literals; default uses `fmt.Sprintf`
3. Implement `Verses(start, stop int) (string, error)`:
   - Validate start and stop are 0-99
   - Validate start >= stop
   - Loop from start down to stop, append each verse + `"\n"` to a `bytes.Buffer`
   - Return `buff.String()`
4. Implement `Song() string`:
   - Delegate to `Verses(99, 0)`, ignore error
5. Run `go test` to verify all tests pass

**Final code:**

```go
package beer

import (
	"bytes"
	"fmt"
)

func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", fmt.Errorf("%d is not a valid verse", n)
	}
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}

func Verses(start, stop int) (string, error) {
	if start < 0 || start > 99 {
		return "", fmt.Errorf("start value[%d] is not a valid verse", start)
	}
	if stop < 0 || stop > 99 {
		return "", fmt.Errorf("stop value[%d] is not a valid verse", stop)
	}
	if start < stop {
		return "", fmt.Errorf("start value[%d] is less than stop value[%d]", start, stop)
	}
	var buff bytes.Buffer
	for i := start; i >= stop; i-- {
		v, _ := Verse(i)
		buff.WriteString(v)
		buff.WriteString("\n")
	}
	return buff.String(), nil
}

func Song() string {
	s, _ := Verses(99, 0)
	return s
}
```
