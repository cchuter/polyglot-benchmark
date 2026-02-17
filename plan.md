# Implementation Plan: Pig Latin Translator

## Proposal A: Map-Based Lookup with Iterative Scanning

**Role: Proponent**

### Approach

Use map-based vowel lookups and iterate through each word character-by-character to determine the split point. This avoids regex and keeps the logic explicit.

### Files to Modify

- `go/exercises/practice/pig-latin/pig_latin.go` — sole file to implement

### Implementation

```go
package piglatin

import "strings"

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
var specials = map[string]bool{"xr": true, "yt": true}
var vowelsY = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

func Sentence(phrase string) string {
    words := strings.Fields(phrase)
    for i, word := range words {
        words[i] = pigWord(word)
    }
    return strings.Join(words, " ")
}

func pigWord(word string) string {
    // Rule 1: starts with vowel, "xr", or "yt"
    if vowels[word[0]] || (len(word) >= 2 && specials[word[:2]]) {
        return word + "ay"
    }
    // Rule 4: consonants followed by y (check before Rule 2)
    for i := 1; i < len(word); i++ {
        if word[i] == 'y' {
            return word[i:] + word[:i] + "ay"
        }
        if vowelsY[word[i]] {
            break
        }
    }
    // Rule 3: consonants + "qu"
    // Rule 2: consonants moved to end
    for i := 0; i < len(word); i++ {
        if vowels[word[i]] {
            // Check for "qu" — if current char is 'u' preceded by 'q', include 'u' in the prefix
            if word[i] == 'u' && i > 0 && word[i-1] == 'q' {
                i++
            }
            return word[i:] + word[:i] + "ay"
        }
    }
    return word + "ay"
}
```

### Rationale

- No external regex dependency; uses only `strings` from stdlib
- Map lookups for vowels are idiomatic Go and fast
- Character-by-character iteration makes the logic transparent and debuggable
- Matches the "map lookups" approach shown in `.approaches/introduction.md`, which is a known working pattern
- Separating `pigWord` from `Sentence` improves readability

### Strengths

- Simple, readable, no regex compilation overhead
- Easy to test and debug each rule independently
- Matches established exercise approach patterns

---

## Proposal B: Regex-Based Pattern Matching

**Role: Opponent**

### Approach

Use compiled regular expressions to match each Pig Latin rule, similar to the `.meta/example.go` reference implementation.

### Files to Modify

- `go/exercises/practice/pig-latin/pig_latin.go` — sole file to implement

### Implementation

```go
package piglatin

import (
    "regexp"
    "strings"
)

var vowel = regexp.MustCompile(`^([aeiou]|y[^aeiou]|xr)[a-z]*`)
var cons = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`)
var containsY = regexp.MustCompile(`^([^aeiou]+)y([a-z]*)`)

func Sentence(s string) string {
    words := strings.Fields(s)
    for i, w := range words {
        words[i] = word(strings.ToLower(w))
    }
    return strings.Join(words, " ")
}

func word(s string) string {
    if containsY.MatchString(s) {
        pos := containsY.FindStringSubmatchIndex(s)
        return s[pos[3]:] + s[:pos[3]] + "ay"
    }
    if vowel.MatchString(s) {
        return s + "ay"
    }
    if x := cons.FindStringSubmatchIndex(s); x != nil {
        return s[x[3]:] + s[:x[3]] + "ay"
    }
    return s + "ay"
}
```

### Critique of Proposal A

- The iterative approach in Proposal A has a subtle complexity with overlapping rules (Rule 4 vs Rule 2 vs Rule 3). The separate loops for 'y' detection and vowel detection could interact in unexpected ways if not carefully ordered.
- The Rule 4 check in Proposal A scans for 'y', but must also break if a regular vowel is found first — this dual-condition loop is error-prone.

### Rationale for Proposal B

- Regex patterns encode the rules declaratively and concisely
- The `.meta/example.go` reference implementation uses this exact approach, proving it works
- Pattern matching handles edge cases like "qu" naturally within the regex
- Less procedural logic to get wrong

### Strengths

- Proven correct (matches reference implementation)
- Declarative rule encoding via regex
- Handles complex consonant clusters cleanly

### Weaknesses

- Regex is harder to read for developers unfamiliar with regex syntax
- `regexp.MustCompile` adds initialization overhead (negligible but present)
- Using submatch indices (`FindStringSubmatchIndex`) is non-trivial to understand

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion    | Proposal A (Map Lookup) | Proposal B (Regex) |
|-------------|------------------------|-------------------|
| Correctness | Likely correct but needs careful rule ordering | Proven correct via reference |
| Risk        | Medium — subtle loop interactions | Low — mirrors reference impl |
| Simplicity  | More readable for Go developers | More concise but regex-heavy |
| Consistency | Matches approach docs | Matches example.go |

### Decision

**Selected: Proposal A (Map-Based Lookup)** with refinements.

**Rationale:**
- The map-based approach is more readable and idiomatic Go
- It avoids regex complexity and the `regexp` import
- The approach from `.approaches/introduction.md` is proven to work
- The risk of rule ordering bugs is manageable by following the proven map-lookup approach exactly
- The solution is simpler to understand, maintain, and debug

### Refined Implementation Plan

**File:** `go/exercises/practice/pig-latin/pig_latin.go`

```go
package piglatin

import "strings"

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
var specials = map[string]bool{"xr": true, "yt": true}
var vowelsY = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

func Sentence(phrase string) string {
    words := strings.Fields(phrase)
    for i, word := range words {
        words[i] = translateWord(word)
    }
    return strings.Join(words, " ")
}

func translateWord(word string) string {
    // Rule 1: starts with vowel, "xr", or "yt"
    if vowels[word[0]] || specials[word[:2]] {
        return word + "ay"
    }

    // Scan consonant cluster to find split point
    for pos := 1; pos < len(word); pos++ {
        letter := word[pos]
        if vowelsY[letter] {
            // Rule 3: "qu" — include the 'u' in the consonant prefix
            if letter == 'u' && word[pos-1] == 'q' {
                pos++
            }
            return word[pos:] + word[:pos] + "ay"
        }
    }

    return word + "ay"
}
```

### Key Design Decisions

1. **`vowelsY` map**: `y` is treated as a vowel when it appears after consonants. By including `y` in the vowel check during the consonant scan, Rule 4 is handled implicitly — the loop stops at `y` and splits there, treating it as a vowel boundary.

2. **`qu` handling**: When we find `u` as a vowel and the preceding character is `q`, we advance past the `u` to include it in the consonant prefix. This handles both `"queen"` → `"eenquay"` and `"square"` → `"aresquay"`.

3. **`specials` map**: The two-character prefixes `"xr"` and `"yt"` are handled by a simple map lookup on the first two characters.

4. **Single loop**: Rules 2, 3, and 4 are all handled in a single scan loop, which is clean and efficient.

### Steps

1. Write the complete `pig_latin.go` implementation
2. Run `go test ./...` in the exercise directory
3. Verify all 22 test cases pass
4. Commit with message `Closes #315: polyglot-go-pig-latin`
