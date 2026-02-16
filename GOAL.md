# Goal: Implement Pig Latin Translator in Go

## Problem Statement

Implement a Go solution for the Pig Latin exercise (issue #145). The solution must translate English text to Pig Latin by implementing the `Sentence` function in `go/exercises/practice/pig-latin/pig_latin.go`. The function must handle single words and multi-word phrases according to four translation rules.

## Translation Rules

1. **Vowel start**: If a word begins with a vowel (`a`, `e`, `i`, `o`, `u`), or starts with `"xr"` or `"yt"`, append `"ay"` to the end.
2. **Consonant cluster**: If a word begins with one or more consonants, move those consonants to the end and append `"ay"`.
3. **Qu handling**: If a word starts with zero or more consonants followed by `"qu"`, move consonants + `"qu"` to the end and append `"ay"`.
4. **Y as vowel**: If a word starts with one or more consonants followed by `"y"`, move the consonants preceding `"y"` to the end and append `"ay"` (treating `"y"` as a vowel in this position).

## Acceptance Criteria

1. All 22 test cases in `cases_test.go` and `pig_latin_test.go` pass (`go test ./...` exits 0).
2. The `Sentence(string) string` function is exported from package `piglatin`.
3. Rule 1 works: vowel-initial words get `"ay"` appended (`apple` -> `appleay`).
4. Rule 1 works: `"xr"` and `"yt"` prefixes treated as vowel starts (`xray` -> `xrayay`, `yttria` -> `yttriaay`).
5. Rule 2 works: consonant clusters moved to end (`pig` -> `igpay`, `chair` -> `airchay`).
6. Rule 3 works: `"qu"` (with optional preceding consonants) moved to end (`queen` -> `eenquay`, `square` -> `aresquay`).
7. Rule 4 works: `"y"` treated as vowel after consonant cluster (`my` -> `ymay`, `rhythm` -> `ythmrhay`).
8. Multi-word phrases are handled correctly (`quick fast run` -> `ickquay astfay unray`).
9. The solution compiles without errors.
10. No modifications to test files (`cases_test.go`, `pig_latin_test.go`).

## Key Constraints

- Package name must be `piglatin`.
- Only the file `pig_latin.go` should be modified.
- Go module is `piglatin` with Go 1.18.
- Must export `Sentence(s string) string` function (this is what the tests call).
