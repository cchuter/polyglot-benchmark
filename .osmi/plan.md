# Implementation Plan: Beer Song (Issue #319)

## Proposal A

**Role: Proponent**

### Approach: Switch-based with `fmt.Sprintf`

Edit `beer_song.go` to implement three functions using a switch statement to handle the four verse cases (0, 1, 2, default) with `fmt.Sprintf` for the generic case and string literals for special cases. Use `bytes.Buffer` for concatenation in `Verses()`.

### Files to Modify
- `go/exercises/practice/beer-song/beer_song.go` — the only file to change

### Implementation

```go
package beer

import (
    "bytes"
    "fmt"
)

func Song() string {
    s, _ := Verses(99, 0)
    return s
}

func Verses(start, stop int) (string, error) {
    if start < 0 || start > 99 {
        return "", fmt.Errorf("invalid start: %d", start)
    }
    if stop < 0 || stop > 99 {
        return "", fmt.Errorf("invalid stop: %d", stop)
    }
    if start < stop {
        return "", fmt.Errorf("start %d < stop %d", start, stop)
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
        return "", fmt.Errorf("invalid verse: %d", n)
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

- Directly mirrors the reference implementation in `.meta/example.go`
- Minimal code, easy to verify against test expectations
- Uses `bytes.Buffer` for efficient string concatenation
- Clear switch cases for each grammar variation
- Error handling matches what tests expect

---

## Proposal B

**Role: Opponent**

### Approach: Template-based with helper functions

Use helper functions to compute the bottle count string and action string, then compose verses using a single template. Avoids hardcoded string literals for verses 1 and 2.

### Files to Modify
- `go/exercises/practice/beer-song/beer_song.go` — the only file to change

### Implementation

```go
package beer

import (
    "fmt"
    "strings"
)

func bottleStr(n int) string {
    switch n {
    case 0:
        return "no more bottles"
    case 1:
        return "1 bottle"
    default:
        return fmt.Sprintf("%d bottles", n)
    }
}

func capitalBottleStr(n int) string {
    if n == 0 {
        return "No more bottles"
    }
    return bottleStr(n)
}

func actionStr(n int) string {
    switch n {
    case 0:
        return "Go to the store and buy some more"
    case 1:
        return "Take it down and pass it around"
    default:
        return "Take one down and pass it around"
    }
}

func nextN(n int) int {
    if n == 0 {
        return 99
    }
    return n - 1
}

func Verse(n int) (string, error) {
    if n < 0 || n > 99 {
        return "", fmt.Errorf("invalid verse: %d", n)
    }
    return fmt.Sprintf("%s of beer on the wall, %s of beer.\n%s, %s of beer on the wall.\n",
        capitalBottleStr(n), bottleStr(n), actionStr(n), bottleStr(nextN(n))), nil
}

func Verses(start, stop int) (string, error) {
    // validation...
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

### Critique of Proposal A

- Uses string literals that are hard to verify without running the code
- Has some duplication (the "bottles of beer on the wall" phrase appears in every case)
- If the song format ever changed, you'd need to update multiple literals

### Rationale for B

- DRY — the verse template is defined once
- Each grammar rule is isolated in its own function
- Easier to reason about correctness of each component
- More extensible

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A                                    | Proposal B                                         |
|---------------|-----------------------------------------------|---------------------------------------------------|
| Correctness   | Directly matches reference and test expectations | Higher risk — composing from parts could introduce subtle format mismatches |
| Risk          | Low — literals are copy-paste verifiable       | Medium — template composition could produce wrong spacing/capitalization |
| Simplicity    | Very simple, ~30 lines                         | More complex, 5 functions, ~50 lines               |
| Consistency   | Matches `.meta/example.go` exactly             | Different pattern from reference implementation    |

### Decision: Proposal A

Proposal A wins on all criteria. This is a straightforward exercise where the primary risk is getting the exact string format wrong. Hardcoded literals for special cases make it trivially verifiable. The reference implementation uses this same pattern, confirming it's the idiomatic approach for this exercise. Proposal B's template approach is over-engineered for a fixed-format song.

### Final Implementation Plan

1. Edit `go/exercises/practice/beer-song/beer_song.go`
2. Add imports: `bytes`, `fmt`
3. Implement `Verse(n int) (string, error)` with switch on n:
   - `n < 0 || n > 99`: return error
   - `n == 0`: return verse 0 literal
   - `n == 1`: return verse 1 literal
   - `n == 2`: return verse 2 literal
   - default: return `fmt.Sprintf(...)` with n, n, n-1
4. Implement `Verses(start, stop int) (string, error)`:
   - Validate start and stop are in [0, 99] and start >= stop
   - Loop from start down to stop, appending each verse + "\n" to a `bytes.Buffer`
   - Return buffer contents
5. Implement `Song() string`:
   - Call `Verses(99, 0)` and return result
6. Run `go test ./...` and `go vet ./...` to verify
7. Commit
