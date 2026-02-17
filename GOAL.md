# Goal: Implement Pig Latin Translator in Go

## Problem Statement

Implement the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go` that translates English text to Pig Latin. The stub file currently only contains the package declaration and needs a complete implementation.

## Translation Rules

1. **Rule 1 (Vowel start)**: If a word begins with a vowel (`a`, `e`, `i`, `o`, `u`), or starts with `"xr"` or `"yt"`, add `"ay"` to the end.
2. **Rule 2 (Consonant start)**: If a word begins with one or more consonants, move those consonants to the end and add `"ay"`.
3. **Rule 3 (Qu rule)**: If a word starts with zero or more consonants followed by `"qu"`, move the consonants and `"qu"` to the end and add `"ay"`.
4. **Rule 4 (Y rule)**: If a word starts with one or more consonants followed by `"y"`, move the consonants preceding `"y"` to the end and add `"ay"` (treating `"y"` as a vowel in this position).

## Acceptance Criteria

1. All 22 test cases in `cases_test.go` pass (`go test ./...` in the exercise directory)
2. The `Sentence` function is exported and has signature `func Sentence(string) string`
3. The function handles single words and multi-word phrases (space-separated)
4. Only `pig_latin.go` is modified; test files remain untouched

## Key Constraints

- Package name must be `piglatin`
- Go 1.18 module
- No external dependencies allowed (standard library only)
