# Change Log: Pig Latin Implementation

## Modified File
- `go/exercises/practice/pig-latin/pig_latin.go`

## Changes
- Added imports for `regexp` and `strings`
- Defined three compiled regex patterns as package-level vars:
  - `vowel`: matches words starting with a vowel, `xr`, or `yt`
  - `containsy`: matches consonant cluster followed by `y` (Rule 4)
  - `cons`: matches consonant clusters including `qu` handling (Rules 2/3)
- Implemented `Word(s string) string` with three-branch regex logic
- Implemented `Sentence(s string) string` that splits, translates, and joins words

## Test Results
All 22 test cases pass.

## Commit
`c052b92` — "Implement pig-latin Sentence function for issue #145"
