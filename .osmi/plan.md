# Implementation Plan: bottle-song (Issue #45)

## Overview

Implement the `Recite` function in `bottle_song.go` to generate "Ten Green Bottles" lyrics. The function accepts a starting bottle count and number of verses to produce.

## Files to Modify

1. **`go/exercises/practice/bottle-song/bottle_song.go`** — The only file to change. Currently contains only the package declaration.

## Approach

### 1. Number-to-word mapping

Create a `map[int]string` for numbers 0-10:
- 0 -> "no", 1 -> "one", 2 -> "two", ..., 10 -> "ten"

### 2. Capitalize helper

Implement a simple `capitalize` function that uppercases the first letter of a string using `strings.ToUpper(s[:1]) + s[1:]`. All number words are ASCII, so byte-level operation is safe. This avoids using the deprecated `strings.Title`.

### 3. Pluralization helper

Create a `bottleWord(n int) string` helper:
- Returns "bottle" when n == 1
- Returns "bottles" otherwise (including n == 0)

### 4. Verse generator

Create a `verse(n int) []string` function that returns 4 lines for bottle count `n`:
- Lines 1-2: `"{capitalize(word)} green {bottleWord(n)} hanging on the wall,"`
- Line 3: `"And if one green bottle should accidentally fall,"`
- Line 4: `"There'll be {word(n-1)} green {bottleWord(n-1)} hanging on the wall."`

Key logic:
- Use capitalize for the first word in lines 1-2
- Line 4 uses `n-1` for the remaining count (lowercase word)
- The bottleWord helper handles singular/plural automatically

### 5. Recite function

`Recite(startBottles, takeDown int) []string`:
- Loop from `startBottles` down for `takeDown` iterations
- Append each verse's 4 lines
- Between verses (not after the last), append an empty string `""`
- Return the accumulated slice

## Rationale

- Using a general `verse` function with `bottleWord` helper avoids special-casing n==1 and n==2 as separate switch cases, making the code cleaner
- The capitalize helper is self-contained and avoids deprecated `strings.Title`
- Matches the pattern from `.meta/example.go` reference solution

## Ordering

1. Write the complete implementation in `bottle_song.go`
2. Run `go test -v` to verify all 7 tests pass
3. Run `go vet` to check for issues
