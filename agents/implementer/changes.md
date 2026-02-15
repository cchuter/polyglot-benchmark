# Changes: Replace deprecated strings.Title in bottle-song

## File Modified
- `go/exercises/practice/bottle-song/bottle_song.go`

## What Changed
1. **Added `capitalize` helper function** (lines 21-26): A simple function that uppercases the first letter of a string using `strings.ToUpper(s[:1]) + s[1:]`, replacing the deprecated `strings.Title()`.
2. **Replaced two `strings.Title()` calls** (lines 46-47): Both occurrences of `strings.Title(numberToWord[n])` in the `default` case of `verse()` were changed to `capitalize(numberToWord[n])`.

## Why
`strings.Title` is deprecated in Go since 1.18 because its word boundary rules do not handle Unicode punctuation properly. The `capitalize` helper is sufficient here since we only need to capitalize the first letter of single-word number strings ("one", "two", etc.).

## Imports
No import changes needed -- `strings` is still required for `strings.ToUpper` in the new helper, and `fmt` is still used for `Sprintf`.
