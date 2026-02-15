# Implementation Plan: Fix bottle-song (Issue #26)

## Overview

Replace deprecated `strings.Title()` calls in `bottle_song.go` with a simple manual capitalization helper that uppercases the first letter of a string.

## Files to Modify

1. **`go/exercises/practice/bottle-song/bottle_song.go`** — The only file to change.

## Changes

### 1. Add a `capitalize` helper function

Replace the `strings.Title` usage with a simple helper:

```go
func capitalize(s string) string {
    if s == "" {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

This works because all number words in the map are ASCII lowercase strings, so `s[:1]` safely captures the first byte/rune.

### 2. Replace `strings.Title(numberToWord[n])` with `capitalize(numberToWord[n])`

In the `verse` function's default case (lines 39-40), change:
- `strings.Title(numberToWord[n])` → `capitalize(numberToWord[n])`

### 3. Remove unused import

After removing `strings.Title`, the `strings` package import may still be needed if we use `strings.ToUpper`. Check and keep only necessary imports.

## Rationale

- `strings.Title` is deprecated since Go 1.18 (SA1019)
- The recommended replacement `golang.org/x/text/cases` requires an external dependency, which is not allowed
- A manual `capitalize` function is the simplest correct approach since all input strings are simple ASCII lowercase words
- No API changes needed — the `Recite` function signature stays the same

## Ordering

1. Add `capitalize` helper function
2. Replace `strings.Title` calls with `capitalize`
3. Clean up imports if needed
4. Run tests and staticcheck to verify
