# Goal: Implement Pig Latin Translator in Go

## Problem Statement

Implement a `Sentence` function in the `piglatin` package that translates English text to Pig Latin. The function must handle single words and multi-word phrases, applying four translation rules based on vowel/consonant patterns at the beginning of each word.

## Translation Rules

1. **Vowel start / "xr" / "yt" prefix**: Append `"ay"` to the word unchanged.
2. **Consonant cluster start**: Move leading consonants to end, append `"ay"`.
3. **Consonant(s) + "qu"**: Move consonants and `"qu"` to end, append `"ay"`.
4. **Consonant(s) + "y"**: Move consonants before `"y"` to end, append `"ay"` (treating `"y"` as a vowel when it follows consonants).

## Acceptance Criteria

- [ ] `Sentence(input string) string` function exists in `pig_latin.go` in package `piglatin`
- [ ] All 22 test cases in `cases_test.go` pass
- [ ] `go test ./...` passes with zero failures in the `pig-latin` exercise directory
- [ ] `go vet ./...` reports no issues
- [ ] Multi-word phrases are handled (words split by spaces, each translated independently)
- [ ] No external dependencies (only standard library)

## Key Constraints

- The solution file is `pig_latin.go` in `go/exercises/practice/pig-latin/`
- The package name must be `piglatin`
- The exported function signature must be `Sentence(sentence string) string`
- Vowels are: a, e, i, o, u
- "y" is treated as a consonant at the start of a word but as a vowel after a consonant cluster
- Test files (`pig_latin_test.go`, `cases_test.go`) must not be modified
