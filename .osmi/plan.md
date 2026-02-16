# Implementation Plan: bottle-song

## File to Modify

- `go/exercises/practice/bottle-song/bottle_song.go`

## Approach

### Data Structure

A simple map or slice mapping integers 0-10 to their English word equivalents:
- 10 → "ten", 9 → "nine", ..., 1 → "one", 0 → "no"

### Helper: `numberWord(n int) string`

Returns the lowercase English word for numbers 0-10. Returns "no" for 0.

### Helper: `bottlePlural(n int) string`

Returns "bottle" if n == 1, "bottles" otherwise (including 0).

### Main Function: `Recite(startBottles, takeDown int) []string`

1. Initialize an empty `[]string` result
2. Loop from `startBottles` down for `takeDown` iterations:
   - If not the first verse, append an empty string `""` separator
   - Get the title-cased number word for current count (e.g., "Ten")
   - Build line 1: `"{Word} green {bottle/bottles} hanging on the wall,"`
   - Line 2: same as line 1
   - Line 3: `"And if one green bottle should accidentally fall,"`
   - Get the number word for (current - 1) — lowercase for all including "no"
   - Build line 4: `"There'll be {word} green {bottle/bottles} hanging on the wall."`
   - Append all 4 lines to result
3. Return result

### Title Casing

For lines 1 and 2, the number word needs initial capital (e.g., "Ten"). The test file includes a `Title` helper but it's only used in tests. We can simply use `strings.Title` or manually capitalize the first letter. Since the words are simple ASCII, we'll capitalize the first character ourselves to avoid deprecated function warnings.

### Key Details from Test Cases

- Line 3 always says "one green bottle" (singular) — this is constant
- Line 4 uses the lowercase number word
- "no" is lowercase in "There'll be no green bottles"
- 1 bottle is singular, everything else (including 0) is plural

## Order of Changes

1. Add number-to-word mapping
2. Add `bottlePlural` helper
3. Implement `Recite` function
4. Run tests to verify
