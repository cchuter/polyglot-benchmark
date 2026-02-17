# Implementation Plan: bottle-song

## Branch 1

**Approach: Simple lookup table with direct string construction**

Prioritizes simplicity and minimal code.

### Design
- Define a slice/map of number words: `[]string{"no", "one", "two", ..., "ten"}`.
- A helper `numberWord(n int) string` returns the lowercase word.
- A helper `plural(n int) string` returns `"bottle"` if n==1, else `"bottles"`.
- A helper `capitalize(s string) string` uppercases the first letter.
- `Recite` loops from `startBottles` down for `takeDown` iterations, building 4 lines per verse with an empty string separator between verses.

### Files
- Modify: `go/exercises/practice/bottle-song/bottle_song.go`

### Rationale
Minimal code, easy to read, no abstractions beyond what's needed.

### Evaluation
- **Feasibility**: Fully feasible, uses only standard library.
- **Risk**: Very low. Simple string operations.
- **Alignment**: Satisfies all acceptance criteria directly.
- **Complexity**: ~40 lines of code, 1 file change.

---

## Branch 2

**Approach: Template-based verse generation with fmt.Sprintf**

Prioritizes extensibility if verse structure changed.

### Design
- Define verse template constants: `line12Template = "%s green %s hanging on the wall,"`, `line3 = "And if one green bottle should accidentally fall,"`, `line4Template = "There'll be %s green %s hanging on the wall."`.
- `generateVerse(n int) []string` fills templates using `fmt.Sprintf`.
- `Recite` assembles verses with separators.
- Number words via a lookup array.

### Files
- Modify: `go/exercises/practice/bottle-song/bottle_song.go`

### Rationale
Templates make the verse structure explicit and easy to modify.

### Evaluation
- **Feasibility**: Fully feasible.
- **Risk**: Low. Slightly more code than Branch 1 due to templates.
- **Alignment**: Satisfies all acceptance criteria.
- **Complexity**: ~50 lines, 1 file change. Marginally more complex.

---

## Branch 3

**Approach: Functional pipeline with string builder**

Prioritizes performance for large outputs.

### Design
- Use `strings.Builder` to accumulate output efficiently, then split into `[]string`.
- Actually, since the return type is `[]string`, a builder doesn't help much — we need individual lines.
- Alternative: pre-allocate the result slice with exact capacity: `4*takeDown + (takeDown-1)` for separators.
- Same helpers as Branch 1 but with pre-allocated slice.

### Files
- Modify: `go/exercises/practice/bottle-song/bottle_song.go`

### Rationale
Pre-allocation avoids slice growth. Relevant if called with large inputs.

### Evaluation
- **Feasibility**: Fully feasible.
- **Risk**: Low. Slightly over-engineered for a max of 10 bottles.
- **Alignment**: Satisfies all criteria.
- **Complexity**: ~45 lines, 1 file. The pre-allocation adds minimal complexity.

---

## Selected Plan

**Branch 1 — Simple lookup table with direct string construction**

### Rationale

Branch 1 is the clear winner because:
1. The problem is inherently simple (max 10 bottles), making performance optimization (Branch 3) unnecessary.
2. Template-based approaches (Branch 2) add indirection without real benefit for such a small, fixed format.
3. Branch 1 produces the least code, is easiest to verify correct, and directly maps to the acceptance criteria.

### Detailed Implementation

**File: `go/exercises/practice/bottle-song/bottle_song.go`**

```go
package bottlesong

import "strings"

// numberWords maps integers 0-10 to their English word equivalents (lowercase).
var numberWords = []string{
    "no", "one", "two", "three", "four",
    "five", "six", "seven", "eight", "nine", "ten",
}

// Recite returns the lyrics for the specified verses of "Ten Green Bottles".
// startBottles is the number to start counting from (1-10).
// takeDown is how many verses to generate.
func Recite(startBottles, takeDown int) []string {
    var result []string
    for i := 0; i < takeDown; i++ {
        if i > 0 {
            result = append(result, "")
        }
        current := startBottles - i
        next := current - 1
        result = append(result,
            capitalize(numberWords[current])+" green "+plural(current)+" hanging on the wall,",
            capitalize(numberWords[current])+" green "+plural(current)+" hanging on the wall,",
            "And if one green bottle should accidentally fall,",
            "There'll be "+numberWords[next]+" green "+plural(next)+" hanging on the wall.",
        )
    }
    return result
}

// plural returns "bottle" if n == 1, "bottles" otherwise.
func plural(n int) string {
    if n == 1 {
        return "bottle"
    }
    return "bottles"
}

// capitalize returns s with its first letter uppercased.
func capitalize(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

### Steps
1. Write the implementation to `bottle_song.go`.
2. Run `go test ./...` in the exercise directory.
3. Run `go vet ./...` in the exercise directory.
4. Fix any issues and re-test.
5. Commit when all tests pass.
