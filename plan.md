# Implementation Plan: Pig Latin (Issue #145)

## Overview

Implement the Pig Latin translator in `go/exercises/practice/pig-latin/pig_latin.go`. The solution uses regex-based pattern matching to classify words and apply the appropriate transformation rule.

## File to Modify

- `go/exercises/practice/pig-latin/pig_latin.go` — the only file that needs changes

## Architecture

### Approach: Regex-based pattern matching

Use three compiled regex patterns to classify and transform words:

1. **vowel pattern** — `^([aeiou]|y[^aeiou]|xr)[a-z]*` — matches words starting with a vowel, `xr`, or `yt` (where `yt` is matched by `y` followed by non-vowel)
2. **consonant-y pattern** — `^([^aeiou]+)y([a-z]*)` — matches consonant cluster followed by `y` (Rule 4)
3. **consonant-qu pattern** — `^([^aeiou]?qu|[^aeiou]+)([a-z]*)` — matches consonant clusters including `qu` handling (Rules 2 and 3)

### Functions

1. **`Sentence(s string) string`** (exported) — Splits input on whitespace using `strings.Fields`, translates each word via `Word`, and rejoins with spaces.

2. **`Word(s string) string`** (exported) — Translates a single word:
   - First check Rule 4: if consonant cluster + `y`, move consonants before `y` to end + `"ay"`
   - Then check Rule 1: if starts with vowel/xr/yt pattern, append `"ay"`
   - Otherwise apply Rules 2/3: find consonant cluster (including `qu`), move to end + `"ay"`

### Rationale

- Regex approach is clean and concise for this problem
- Order of checks matters: `y`-as-vowel must be checked before the general consonant rule
- The `cons` regex handles both plain consonant clusters and `qu` patterns in a single expression
- Using `strings.Fields` handles any whitespace splitting naturally
- `strings.ToLower` ensures case normalization (though tests only use lowercase)

## Implementation Steps

1. Add `import` block for `regexp` and `strings`
2. Declare three compiled regex patterns as package-level variables
3. Implement `Word` function with the three-branch logic
4. Implement `Sentence` function that splits, translates, and joins
5. Run `go test` to verify all 22 test cases pass

## Testing

```bash
cd go/exercises/practice/pig-latin && go test -v ./...
```

Expected: all 22 test cases pass, 0 failures.
