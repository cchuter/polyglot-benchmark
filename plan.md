# Implementation Plan: polyglot-go-pig-latin

## Branch 1: Regex-Based Approach (Simplicity)

Use compiled regular expressions to match patterns and rearrange word parts. This mirrors the reference solution in `.meta/example.go`.

### Files to modify
- `go/exercises/practice/pig-latin/pig_latin.go`

### Approach
1. Define three compiled regexes at package level:
   - `vowel`: matches words starting with a vowel, `xr`, or `yt` (but not `y` followed by a vowel)
   - `containsy`: matches consonant cluster followed by `y`
   - `cons`: matches consonant cluster (including `qu`) at word start
2. `Sentence(s string) string`: Split on whitespace, translate each word, rejoin with spaces.
3. `Word(s string) string`: Check `containsy` first (Rule 4), then `vowel` (Rule 1), then `cons` (Rules 2 & 3). Use submatch indices to rearrange and append `"ay"`.

### Evaluation
- **Feasibility**: High — regex is well-supported in Go stdlib
- **Risk**: Low — pattern directly follows proven reference solution
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: 1 file, ~35 lines of code

---

## Branch 2: Iterative Character Scanning (Extensibility)

Use a loop to scan characters from the start of each word, classifying vowels/consonants and building the consonant prefix explicitly.

### Files to modify
- `go/exercises/practice/pig-latin/pig_latin.go`

### Approach
1. `Sentence`: Split on spaces, translate each word, rejoin.
2. `Word`: Iterate through characters:
   - Check for `xr`/`yt` prefix → Rule 1
   - Check if first char is vowel → Rule 1
   - Otherwise, accumulate consonants until hitting a vowel or `y` (after position 0) or `qu` sequence
   - Handle `qu` specially: include both in the moved prefix
   - Rearrange word and append `"ay"`

### Evaluation
- **Feasibility**: High — pure string operations, no regex
- **Risk**: Medium — more manual character handling means more edge cases to get right
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: 1 file, ~50 lines; more explicit but more code

---

## Branch 3: Single-Regex with Named Groups (Performance)

Use a single compiled regex with named capture groups to handle all rules in one match.

### Files to modify
- `go/exercises/practice/pig-latin/pig_latin.go`

### Approach
1. Craft a single regex with alternations and named groups:
   ```
   ^(?P<vowel>(?:[aeiou]|xr|yt).*)$|^(?P<yprefix>[^aeiou]+)(?P<ysuffix>y.*)$|^(?P<quprefix>[^aeiou]*qu)(?P<qusuffix>.*)$|^(?P<cprefix>[^aeiou]+)(?P<csuffix>.*)$
   ```
2. Match once per word, check which named group matched, rearrange accordingly.
3. `Sentence`: Split, translate, rejoin.

### Evaluation
- **Feasibility**: Medium — Go's regex engine is RE2, which supports named groups but complex alternations can be hard to debug
- **Risk**: High — single complex regex is fragile, hard to read, hard to maintain
- **Alignment**: Satisfies acceptance criteria if regex is correct
- **Complexity**: 1 file, ~30 lines; conceptually elegant but practically fragile

---

## Selected Plan

**Branch 1 (Regex-Based)** is selected.

### Rationale
- **Lowest risk**: Directly follows the proven reference solution pattern
- **Simplest**: Minimal code, clean separation of concerns with three focused regexes
- **Most maintainable**: Each regex maps clearly to a rule, making the logic transparent
- Branch 2 adds unnecessary complexity for no benefit. Branch 3 is clever but fragile.

### Detailed Implementation

**File**: `go/exercises/practice/pig-latin/pig_latin.go`

```go
package piglatin

import (
	"regexp"
	"strings"
)

var vowel = regexp.MustCompile(`^([aeiou]|y[^aeiou]|xr)[a-z]*`)
var cons = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`)
var containsy = regexp.MustCompile(`^([^aeiou]+)y([a-z]*)`)

func Sentence(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = Word(strings.ToLower(w))
	}
	return strings.Join(words, " ")
}

func Word(s string) string {
	if containsy.MatchString(s) {
		pos := containsy.FindStringSubmatchIndex(s)
		return s[pos[3]:] + s[:pos[3]] + "ay"
	}
	if vowel.MatchString(s) {
		return s + "ay"
	}
	if x := cons.FindStringSubmatchIndex(s); x != nil {
		return s[x[3]:] + s[:x[3]] + "ay"
	}
	return s
}
```

### Steps
1. Create feature branch `issue-231`
2. Write the implementation to `pig_latin.go`
3. Run `go test ./...` in the exercise directory
4. Verify all 22 tests pass
5. Commit with descriptive message
