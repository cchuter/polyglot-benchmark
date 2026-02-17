# Implementation Plan: Pig Latin Translator

## Proposal A: Regex-Based Approach

**Role: Proponent**

Use compiled regular expressions to match the different Pig Latin patterns, similar to the `.meta/example.go` reference solution.

### Files to Modify
- `go/exercises/practice/pig-latin/pig_latin.go`

### Approach
1. Define three compiled regex patterns at package level:
   - `vowel`: matches words starting with a vowel, `xr`, or `yt`
   - `containsy`: matches consonant cluster followed by `y`
   - `cons`: matches consonant cluster (including optional `qu`)
2. `Sentence()` splits input on spaces, translates each word, joins with spaces
3. `Word()` checks patterns in priority order: y-rule first, then vowel, then consonant

### Rationale
- Regex provides concise pattern matching for the rules
- Well-tested approach (it's the reference solution pattern)
- Compact code

### Weaknesses (self-acknowledged)
- Regex can be harder to read/debug
- Adds `regexp` import when simple string operations might suffice
- Regex compilation has overhead (though compiled at package init, it's fine)

---

## Proposal B: Iterative String-Scanning Approach

**Role: Opponent**

Use a simple loop that scans the beginning of each word character by character to find the split point, then rearranges the string. No regex needed.

### Files to Modify
- `go/exercises/practice/pig-latin/pig_latin.go`

### Approach
1. `Sentence()` splits input on spaces, translates each word via `translateWord()`, joins with spaces
2. `translateWord()` logic:
   - Check if word starts with a vowel or `xr` or `yt` → return `word + "ay"` (Rule 1)
   - Otherwise, scan forward through consonants to find the split point:
     - If we encounter `qu`, include both in the consonant prefix (Rule 3)
     - If we encounter `y` (and we've already seen at least one consonant), stop — `y` acts as vowel (Rule 4)
     - Otherwise keep scanning consonants (Rule 2)
   - Return `word[splitPoint:] + word[:splitPoint] + "ay"`

### Rationale
- **Simpler dependencies**: only needs `strings` package, no `regexp`
- **Easier to understand**: the logic maps directly to the rules described in the problem
- **Easier to debug**: each rule corresponds to a clear conditional check
- **Better performance**: no regex overhead, just simple byte comparisons
- **More maintainable**: adding a new rule is just adding another condition

### Critique of Proposal A
- The regex patterns in the example (`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`) are dense and require careful reading to verify correctness
- Three separate regex patterns checked in a specific order is fragile — easy to get wrong
- The `FindStringSubmatchIndex` API is clunky and error-prone with its index-based access
- For this problem, regex is overkill — the patterns are simple enough that character-level scanning is clearer

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion | Proposal A (Regex) | Proposal B (Iterative) |
|---|---|---|
| Correctness | Proven (matches reference) | Must verify all edge cases |
| Risk | Low (known working pattern) | Low (straightforward logic) |
| Simplicity | Medium (regex patterns are dense) | High (clear character scanning) |
| Consistency | Both consistent with Go conventions | Both consistent |

### Decision: **Proposal B (Iterative String-Scanning)**

**Rationale:**
- Both approaches are correct and feasible for this problem
- Proposal B is simpler to read and understand, with logic mapping directly to the four rules
- No regex dependency means fewer imports and cleaner code
- The iterative approach handles all edge cases naturally through the scanning loop
- Performance is better (not that it matters much here, but it's a bonus)
- Proposal A's main advantage (proven pattern) is minor — the iterative approach is also straightforward to verify

### Selected Implementation Plan

**File:** `go/exercises/practice/pig-latin/pig_latin.go`

```go
package piglatin

import "strings"

func Sentence(sentence string) string {
    words := strings.Fields(sentence)
    for i, w := range words {
        words[i] = translateWord(w)
    }
    return strings.Join(words, " ")
}

func isVowel(b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func translateWord(word string) string {
    // Rule 1: starts with vowel, "xr", or "yt"
    if isVowel(word[0]) || (len(word) >= 2 && (word[:2] == "xr" || word[:2] == "yt")) {
        return word + "ay"
    }

    // Scan consonant cluster to find split point
    for i := 0; i < len(word); i++ {
        // Rule 3: "qu" — include both in prefix
        if word[i] == 'q' && i+1 < len(word) && word[i+1] == 'u' {
            return word[i+2:] + word[:i+2] + "ay"
        }
        // Rule 4: "y" after consonants acts as vowel
        if word[i] == 'y' && i > 0 {
            return word[i:] + word[:i] + "ay"
        }
        // Rule 2: hit a vowel — split here
        if isVowel(word[i]) {
            return word[i:] + word[:i] + "ay"
        }
    }
    // Fallback (shouldn't happen with valid input)
    return word + "ay"
}
```

### Ordering
1. Write the implementation to `pig_latin.go`
2. Run `go test ./...` to verify all 22 test cases pass
3. Run `go vet ./...` to check for issues
4. Commit the change
