# Implementation Plan: bottle-song (Issue #38)

## Overview

Implement the `Recite` function in `bottle_song.go` to generate "Ten Green Bottles" lyrics. The function accepts a starting bottle count and number of verses to produce.

## Files to Modify

1. **`go/exercises/practice/bottle-song/bottle_song.go`** — The only file to change. Currently contains only the package declaration.

## Approach

### 1. Number-to-word mapping

Create a `map[int]string` for numbers 0-10:
- 0 → "no", 1 → "one", 2 → "two", ..., 10 → "ten"

### 2. Capitalize helper

Implement a simple `capitalize` function that uppercases the first byte of a string. All number words are ASCII, so byte-level operation is safe. This avoids using the deprecated `strings.Title`.

### 3. Verse generator

Create a `verse(n int) []string` function that returns 4 lines for bottle count `n`:
- Lines 1-2: `"{Capitalized_word} green bottle(s) hanging on the wall,"`
- Line 3: `"And if one green bottle should accidentally fall,"`
- Line 4: `"There'll be {word} green bottle(s) hanging on the wall."`

Key logic:
- Use "bottle" (singular) when count is exactly 1, "bottles" otherwise
- Line 4 uses `n-1` for the remaining count
- Numbers in line 4 are lowercase

### 4. Recite function

`Recite(startBottles, takeDown int) []string`:
- Loop from `startBottles` down for `takeDown` iterations
- Append each verse's 4 lines
- Between verses, append an empty string `""`
- Return the accumulated slice

## Rationale

- Using explicit cases for n==1 and n==2 in the verse function handles the singular/plural transitions cleanly (n==1 needs singular in lines 1-2 and "no bottles" in line 4; n==2 needs plural in lines 1-2 but singular in line 4)
- The approach matches the reference solution pattern in `.meta/example.go`
- No external dependencies needed

## Ordering

1. Write the complete implementation in `bottle_song.go`
2. Run `go test -v` to verify all 7 tests pass
3. Run `go vet` to check for issues
