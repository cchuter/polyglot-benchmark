# Goal: Implement Pig Latin Translator in Go

## Problem Statement

Implement the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go` that translates English text into Pig Latin. The stub file exists with only a `package piglatin` declaration; the function body needs to be written.

## Pig Latin Rules

1. **Rule 1 (Vowel start):** If a word begins with a vowel (`a`, `e`, `i`, `o`, `u`), or starts with `"xr"` or `"yt"`, append `"ay"` to the end.
2. **Rule 2 (Consonant start):** If a word begins with one or more consonants, move those consonants to the end of the word and append `"ay"`.
3. **Rule 3 (`qu` handling):** If a word starts with zero or more consonants followed by `"qu"`, move the consonants and the `"qu"` to the end and append `"ay"`.
4. **Rule 4 (`y` as vowel):** If a word starts with one or more consonants followed by `"y"`, move the consonants preceding the `"y"` to the end and append `"ay"` (i.e., `"y"` acts as a vowel when preceded by consonants).

## Acceptance Criteria

- [ ] `Sentence(input)` returns the correct Pig Latin translation for all 22 test cases in `cases_test.go`
- [ ] All tests pass: `go test ./...` succeeds with exit code 0
- [ ] `go vet ./...` reports no issues
- [ ] The function handles multi-word sentences (splitting on spaces, translating each word)
- [ ] Package name remains `piglatin`
- [ ] Only `pig_latin.go` is modified (test files are read-only)

## Key Constraints

- The function signature must be: `func Sentence(s string) string`
- Words are space-separated
- All input is lowercase English text
- The solution must handle edge cases like `"y"` at the beginning of a word (treated as consonant) vs after consonants (treated as vowel)
