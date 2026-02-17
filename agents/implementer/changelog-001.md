# Changelog 001: Pig Latin Translator

## What was implemented

Implemented a Pig Latin translator in Go that converts English sentences and individual words to Pig Latin, following the standard Pig Latin rules:

1. **Rule 1**: Words beginning with a vowel sound (`a`, `e`, `i`, `o`, `u`), `xr`, or `yt` get "ay" appended.
2. **Rules 2 & 3**: Words beginning with consonant clusters (including `qu` combinations) have the cluster moved to the end, followed by "ay".
3. **Rule 4**: Words with consonant clusters followed by `y` treat `y` as a vowel, moving only the preceding consonants.

Two exported functions are provided:
- `Sentence(s string) string` — translates a full sentence (space-separated words).
- `Word(s string) string` — translates a single word.

## Approach

Regex-based pattern matching using three compiled regular expressions:
- `vowel`: detects words starting with vowel sounds, `xr`, or `y` followed by a non-vowel.
- `cons`: captures leading consonant clusters, including optional `qu`.
- `containsy`: captures consonant clusters followed by `y` (Rule 4 special case).

The order of rule evaluation matters: Rule 4 (`y` as vowel) is checked first, then Rule 1 (vowel start), then Rules 2 & 3 (consonant clusters).

## Files modified

- `go/exercises/practice/pig-latin/pig_latin.go` — replaced stub with full implementation.
