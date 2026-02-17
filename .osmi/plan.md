# Implementation Plan: Pig Latin Translator

## Branch 1: Simple Imperative Approach

**Approach**: A straightforward procedural implementation using a single helper function `translateWord` that applies rules in priority order using string indexing and a vowel-check helper.

**Files to modify**:
- `go/exercises/practice/pig-latin/pig_latin.go` — add `Sentence` and helper functions

**Architecture**:
1. `Sentence(sentence string) string` — splits input on spaces, translates each word, joins with spaces
2. `translateWord(word string) string` — applies rules in order:
   - Check Rule 1: starts with vowel, "xr", or "yt" → append "ay"
   - Check Rule 3: find consonant cluster + "qu" → move to end + "ay"
   - Check Rule 4: consonant cluster + "y" → move consonants before "y" to end + "ay"
   - Default Rule 2: move consonant cluster to end + "ay"
3. `isVowel(b byte) bool` — checks if character is a/e/i/o/u

**Rationale**: Minimal code, easy to read, no imports beyond `strings`. Rules checked in priority order (Rule 3 before Rule 2 because "qu" is a special consonant pattern).

**Evaluation**:
- Feasibility: Excellent — simple string operations only
- Risk: Low — straightforward logic, easy to debug
- Alignment: Fully satisfies all acceptance criteria
- Complexity: ~40-50 lines, single file change

## Branch 2: Regex-Based Approach

**Approach**: Use `regexp` package to match patterns for each rule, making the rule logic declarative.

**Files to modify**:
- `go/exercises/practice/pig-latin/pig_latin.go`

**Architecture**:
1. Compile regex patterns at package level for each rule:
   - Rule 1: `^([aeiou]|xr|yt)`
   - Rule 3: `^([^aeiou]*qu)(.*)`
   - Rule 4: `^([^aeiou]+)(y.*)`
   - Rule 2: `^([^aeiou]+)(.*)`
2. `Sentence` splits and joins words
3. `translateWord` tries each regex in priority order

**Rationale**: More declarative, patterns self-document the rules.

**Evaluation**:
- Feasibility: Good — `regexp` is stdlib
- Risk: Medium — regex can be tricky to get right, especially edge cases with "y"
- Alignment: Satisfies criteria but adds `regexp` dependency (still stdlib)
- Complexity: ~40-50 lines but harder to debug

## Branch 3: Index-Scanning Approach

**Approach**: Use a single-pass index scan to find the consonant cluster boundary, then apply the appropriate rule based on what follows.

**Files to modify**:
- `go/exercises/practice/pig-latin/pig_latin.go`

**Architecture**:
1. `Sentence` splits and joins as before
2. `translateWord`:
   - If starts with vowel or "xr"/"yt" → return word + "ay"
   - Scan forward to find the boundary index where consonants end:
     - If we hit "qu", include "qu" in the prefix (Rule 3)
     - If we hit "y" (not at position 0), stop before "y" (Rule 4)
     - Otherwise stop at first vowel (Rule 2)
   - Return `word[boundary:] + word[:boundary] + "ay"`

**Rationale**: Most performant — single pass, no string building, no regex compilation.

**Evaluation**:
- Feasibility: Excellent — pure byte operations
- Risk: Medium — boundary logic needs careful handling of edge cases
- Alignment: Fully satisfies criteria
- Complexity: ~35-45 lines, compact but dense

## Selected Plan

**Branch 1: Simple Imperative Approach** is selected.

**Rationale**: This exercise is a straightforward string transformation. Branch 1 is the clearest and most maintainable. Branch 2 adds regex overhead and complexity for no benefit. Branch 3 is slightly more performant but the index-scanning logic is harder to follow. Branch 1 strikes the best balance of readability, correctness, and simplicity.

### Detailed Implementation

**File**: `go/exercises/practice/pig-latin/pig_latin.go`

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

func translateWord(word string) string {
    // Rule 1: starts with vowel, "xr", or "yt"
    if isVowel(word[0]) || strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") {
        return word + "ay"
    }

    // Find consonant cluster length
    for i := 0; i < len(word); i++ {
        // Rule 3: consonant(s) + "qu"
        if i+1 < len(word) && word[i] == 'q' && word[i+1] == 'u' {
            return word[i+2:] + word[:i+2] + "ay"
        }
        // Rule 4: consonant(s) + "y" (y not at start)
        if word[i] == 'y' && i > 0 {
            return word[i:] + word[:i] + "ay"
        }
        // Found a vowel — Rule 2
        if isVowel(word[i]) {
            return word[i:] + word[:i] + "ay"
        }
    }
    // All consonants (shouldn't happen with valid input)
    return word + "ay"
}

func isVowel(b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
```

**Order of changes**:
1. Write the implementation to `pig_latin.go`
2. Run `go test ./...` in the exercise directory
3. Run `go vet ./...`
4. Fix any issues and re-test
5. Commit when all tests pass
