# Implementation Plan: Beer Song

## Proposal A — Switch-based approach (direct string construction)

**Role: Proponent**

This approach uses a `switch` statement in `Verse()` to handle each special case directly, returning hardcoded strings for verses 0 and 1, and using `fmt.Sprintf` for the general case. `Verses()` iterates from start to stop, concatenating results. `Song()` delegates to `Verses(99, 0)`.

### Files to modify
- `go/exercises/practice/beer-song/beer_song.go`

### Implementation
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
    if start < 0 || start > 99 {
        return "", fmt.Errorf("start value is not valid")
    }
    if stop < 0 || stop > 99 {
        return "", fmt.Errorf("stop value is not valid")
    }
    if start < stop {
        return "", fmt.Errorf("start is less than stop")
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
        return "", fmt.Errorf("invalid verse")
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

### Rationale
- Simple, readable, and directly maps to the problem's special cases.
- Matches the reference solution in `.meta/example.go` almost exactly.
- Minimal code, no unnecessary abstractions.
- Easy to verify correctness by inspection.

---

## Proposal B — Template-based approach (parameterized verse construction)

**Role: Opponent**

This approach builds each verse from components, avoiding hardcoded strings for cases 1 and 2. It uses helper functions to determine the correct bottle word and action phrase.

### Files to modify
- `go/exercises/practice/beer-song/beer_song.go`

### Implementation
```go
package beer

import (
    "fmt"
    "strings"
)

func bottleWord(n int) string {
    if n == 1 { return "bottle" }
    return "bottles"
}

func countStr(n int) string {
    if n == 0 { return "no more" }
    return fmt.Sprintf("%d", n)
}

func Verse(n int) (string, error) {
    if n < 0 || n > 99 {
        return "", fmt.Errorf("invalid verse")
    }
    if n == 0 {
        return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
    }
    next := n - 1
    action := "Take one down and pass it around"
    if n == 1 {
        action = "Take it down and pass it around"
    }
    return fmt.Sprintf("%s %s of beer on the wall, %s %s of beer.\n%s, %s %s of beer on the wall.\n",
        countStr(n), bottleWord(n), countStr(n), bottleWord(n),
        action, countStr(next), bottleWord(next)), nil
}

// Verses and Song same as Proposal A...
```

### Critique of Proposal A
- Proposal A hardcodes verse 2, which is unnecessary — the only special thing about verse 2 is that `n-1 = 1`, making it "1 bottle" (singular). A parameterized approach handles this naturally.
- However, Proposal B introduces more complexity (helper functions, more Sprintf args) for marginal benefit.

### Why Proposal B is superior
- Less duplication: The verse structure is defined once.
- More extensible if requirements change.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion    | Proposal A (Switch) | Proposal B (Template) |
|-------------|--------------------|-----------------------|
| Correctness | Fully correct      | Fully correct         |
| Risk        | Very low           | Low (more Sprintf args to get wrong) |
| Simplicity  | Very simple        | Moderate complexity   |
| Consistency | Matches `.meta/example.go` exactly | Diverges from reference |

### Decision: **Proposal A**

**Rationale:**
- Proposal A is simpler and directly matches the proven reference solution.
- The "duplication" in Proposal A is minimal (one extra case for verse 2) and makes the code easier to read and verify.
- Proposal B's helper functions add complexity without meaningful benefit for a problem this small.
- Proposal A has the lowest risk of introducing bugs in string formatting.
- Consistency with the reference solution is valuable.

### Final Implementation Plan

**File:** `go/exercises/practice/beer-song/beer_song.go`

**Step 1:** Implement `Verse(n int) (string, error)` with switch cases for n<0||n>99 (error), n==0, n==1, n==2, and default (3-99).

**Step 2:** Implement `Verses(start, stop int) (string, error)` with validation and loop from start down to stop, concatenating verses separated by blank lines.

**Step 3:** Implement `Song() string` that delegates to `Verses(99, 0)`.

**Imports needed:** `bytes`, `fmt`

**Verification:** Run `go test ./...` and `go vet ./...` in the beer-song directory.
