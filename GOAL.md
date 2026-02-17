# Goal: Implement Pig Latin Translator in Go

## Problem Statement

Implement the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go` that translates English text to Pig Latin. The function must handle single words and multi-word phrases.

## Translation Rules

### Rule 1 - Vowel Start
If a word begins with a vowel (`a`, `e`, `i`, `o`, `u`), or starts with `"xr"` or `"yt"`, append `"ay"` to the end.

### Rule 2 - Consonant Start
If a word begins with one or more consonants, move those consonants to the end and append `"ay"`.

### Rule 3 - Consonants + "qu"
If a word starts with zero or more consonants followed by `"qu"`, move those consonants and `"qu"` to the end and append `"ay"`.

### Rule 4 - Consonants + "y"
If a word starts with one or more consonants followed by `"y"`, move the consonants preceding `"y"` to the end and append `"ay"`. Here `"y"` is treated as a vowel.

## Acceptance Criteria

1. All 22 test cases in `cases_test.go` pass
2. The `Sentence` function is exported from package `piglatin`
3. The function handles multi-word phrases (space-separated)
4. `go test ./...` passes with zero failures in the `pig-latin` exercise directory
5. The solution file is `pig_latin.go` in the exercise directory

## Key Constraints

- Package name must be `piglatin`
- Go module version is 1.18
- Only `pig_latin.go` should be modified (test files are read-only)
- The function signature must be: `func Sentence(s string) string`
