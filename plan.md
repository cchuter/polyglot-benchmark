# Implementation Plan: Beer Song

## Branch 1: Switch-Case with String Literals (Simple)

Implement `Verse` using a switch statement with hardcoded strings for special cases (0, 1, 2) and `fmt.Sprintf` for the general case. `Verses` loops from start to stop, appending each verse with a newline separator. `Song` delegates to `Verses(99, 0)`.

**Files to modify:** `beer_song.go` only

**Approach:**
1. Add imports: `fmt`, `bytes`
2. `Verse(n)`: switch on n with cases for 0, 1, 2, default (3-99), and validation
3. `Verses(start, stop)`: validate bounds, loop and concatenate with `bytes.Buffer`
4. `Song()`: call `Verses(99, 0)`

**Evaluation:**
- Feasibility: High - straightforward, matches the reference solution pattern
- Risk: Very low - simple logic, easy to verify
- Alignment: Fully satisfies all acceptance criteria
- Complexity: Single file, ~55 lines of code

## Branch 2: Template-Based with Helper Functions (Extensible)

Define verse templates as constants and use helper functions for bottle pluralization and action text. More modular but adds abstraction layers.

**Files to modify:** `beer_song.go` only

**Approach:**
1. Create `bottleStr(n)` helper returning "bottle" or "bottles"
2. Create `countStr(n)` returning "no more" or the number
3. Create `actionStr(n)` returning the action line
4. `Verse` composes from helpers
5. `Verses` and `Song` same as Branch 1

**Evaluation:**
- Feasibility: High - standard Go
- Risk: Low but more code surface area for bugs
- Alignment: Satisfies all criteria
- Complexity: Single file, ~70 lines - more abstraction than needed

## Branch 3: Precomputed Verse Array (Performance)

Precompute all 100 verses at init time into a slice. `Verse` is a simple lookup. `Verses` joins slices. Maximum runtime performance.

**Files to modify:** `beer_song.go` only

**Approach:**
1. `init()` builds `[100]string` of all verses
2. `Verse(n)` returns `verses[n]` with bounds check
3. `Verses(start, stop)` joins slice entries
4. `Song()` returns precomputed full song

**Evaluation:**
- Feasibility: High
- Risk: Low but uses `init()` which is heavier than needed
- Alignment: Satisfies all criteria
- Complexity: Single file, ~60 lines - over-engineered for this problem

## Selected Plan

**Branch 1** is the best choice. It is the simplest, most readable, and directly matches the established patterns in the codebase (e.g., the reference solution in `.meta/example.go`). The problem doesn't warrant abstraction layers (Branch 2) or precomputation (Branch 3).

### Detailed Implementation

**File: `go/exercises/practice/beer-song/beer_song.go`**

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

### Steps:
1. Write the implementation to `beer_song.go`
2. Run `go test ./...` to verify all tests pass
3. Run `go vet ./...` to verify no issues
4. Commit
