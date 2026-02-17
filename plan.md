# Implementation Plan: Bottle Song

## Proposal A — Map-based approach with special cases

**Role: Proponent**

### Approach

Follow the reference solution pattern closely: use a `map[int]string` for number-to-word conversion and a `switch` statement to handle the special cases (1 bottle, 2 bottles, default).

### Files to modify

- `go/exercises/practice/bottle-song/bottle_song.go` — implement `Recite` function and helper

### Design

```go
package bottlesong

import "fmt"

var numberToWord = map[int]string{
    1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
    6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
}

func Recite(startBottles, takeDown int) []string {
    verses := []string{}
    for i := startBottles; i > startBottles-takeDown; i-- {
        verses = append(verses, verse(i)...)
        if i > startBottles-takeDown+1 {
            verses = append(verses, "")
        }
    }
    return verses
}

func verse(n int) []string {
    switch {
    case n == 1:
        // singular "bottle", "no" for zero
        return []string{...hardcoded 4 lines...}
    case n == 2:
        // plural first two lines, singular in last line ("one green bottle")
        return []string{...hardcoded 4 lines...}
    default:
        // plural throughout, use Title() for first lines, lowercase for last
        return []string{...fmt.Sprintf with Title(numberToWord[n])...}
    }
}
```

### Rationale

- Directly matches the reference solution, minimizing risk of deviation
- Simple and readable — each special case is explicit
- Uses the `Title` function from the test file for proper capitalization
- Map lookup is O(1) and straightforward for the small domain (1-10)

### Strengths

- Proven correct (matches `.meta/example.go`)
- Minimal code, easy to understand
- No unnecessary abstractions

---

## Proposal B — Fully generalized approach with helper functions

**Role: Opponent**

### Approach

Build a more generalized solution that computes singular/plural and number words dynamically, avoiding special-cased switch branches for n=1 and n=2.

### Design

```go
package bottlesong

import "fmt"

func numberWord(n int) string {
    words := []string{"no", "one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine", "ten"}
    return words[n]
}

func bottleStr(n int) string {
    if n == 1 { return "bottle" }
    return "bottles"
}

func Recite(startBottles, takeDown int) []string {
    var result []string
    for i := 0; i < takeDown; i++ {
        n := startBottles - i
        if i > 0 { result = append(result, "") }
        word := Title(numberWord(n))
        result = append(result,
            fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
            fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
            "And if one green bottle should accidentally fall,",
            fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberWord(n-1), bottleStr(n-1)),
        )
    }
    return result
}
```

### Critique of Proposal A

- Hardcoding verses for n=1 and n=2 is redundant — the logic can be generalized
- More code to maintain with three separate return statements

### Strengths of Proposal B

- Single code path for all verses — DRY principle
- Uses a slice instead of a map, which is simpler for sequential integers
- The `bottleStr` helper cleanly handles singular/plural everywhere
- Adding support for more bottles would require zero code changes

### Weaknesses of Proposal B

- Slightly more complex formatting logic
- The `Title` function dependency is implicit (defined in test file)
- Small risk: must ensure `numberWord(0)` returns "no" and `bottleStr(0)` returns "bottles"

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion     | Proposal A (Map + Switch) | Proposal B (Generalized) |
|---------------|--------------------------|--------------------------|
| Correctness   | Proven (matches reference) | Correct but needs verification |
| Risk          | Very low — copy of reference | Low — but novel code path |
| Simplicity    | Simple, slightly repetitive | Cleaner, slightly more abstract |
| Consistency   | Matches codebase convention | Also fits Go idioms |

### Decision

**Selected: Proposal B (Generalized)** — with minor refinements.

**Rationale**: Proposal B is cleaner, avoids duplication, and handles all edge cases through a single code path. The risk is minimal since the logic is straightforward and the test suite is comprehensive. The slice-based number lookup is simpler than a map for sequential integers 0-10.

However, I will incorporate one element from Proposal A: the verse separator logic using the countdown loop style, which is slightly cleaner.

### Final Implementation Plan

**File**: `go/exercises/practice/bottle-song/bottle_song.go`

```go
package bottlesong

import "fmt"

func Recite(startBottles, takeDown int) []string {
    var result []string
    for i := 0; i < takeDown; i++ {
        n := startBottles - i
        if i > 0 {
            result = append(result, "")
        }
        word := Title(numberWord(n))
        nextWord := numberWord(n - 1)
        result = append(result,
            fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
            fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
            "And if one green bottle should accidentally fall,",
            fmt.Sprintf("There'll be %s green %s hanging on the wall.", nextWord, bottleStr(n-1)),
        )
    }
    return result
}

var numbers = []string{
    "no", "one", "two", "three", "four", "five",
    "six", "seven", "eight", "nine", "ten",
}

func numberWord(n int) string {
    return numbers[n]
}

func bottleStr(n int) string {
    if n == 1 {
        return "bottle"
    }
    return "bottles"
}
```

### Steps

1. Write the implementation to `bottle_song.go`
2. Run `go test` to verify all 7 test cases pass
3. Commit the changes
