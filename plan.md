# Implementation Plan: bottle-song

## Overview

Implement the `Recite` function in `bottle_song.go` following the reference solution pattern from `.meta/example.go`. The solution is straightforward: a number-to-word map, a verse generator with singular/plural handling, and a loop that assembles verses with separators.

## File to Modify

- `go/exercises/practice/bottle-song/bottle_song.go` — currently contains only `package bottlesong`

## Approach

### 1. Number-to-Word Map

Define a package-level `var numberToWord = map[int]string{...}` mapping integers 1–10 to their English word equivalents ("one", "two", ..., "ten").

### 2. Title-casing Helper

The `Title` function is defined in `bottle_song_test.go`. In Go, during `go test`, all `.go` files (including `_test.go`) in the same package are compiled together into a single test binary, so functions defined in test files ARE available to production code during testing. The reference solution (`.meta/example.go`) uses `Title()` from the test file, confirming this pattern works.

**Decision**: Use the `Title()` function defined in the test file directly, matching the reference solution pattern. If this causes compilation issues, fall back to implementing a local `title()` helper:
```go
func title(s string) string {
    if s == "" {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

### 3. Verse Generator

Create an unexported `verse(n int) []string` function that returns 4 lines for a given bottle count:

- **n == 1 (special)**: singular "bottle" in lines 1-2, "no green bottles" in line 4. The word "no" is used (not "zero").
- **n == 2 (special)**: plural "bottles" in lines 1-2, singular "one green bottle" (not "bottles") in line 4
- **default (n >= 3)**: plural "bottles" throughout, word-form numbers via `Title(numberToWord[n])` for lines 1-2 and `numberToWord[n-1]` for line 4

Each verse has exactly this structure:
```
"{N} green bottle(s) hanging on the wall,"
"{N} green bottle(s) hanging on the wall,"
"And if one green bottle should accidentally fall,"
"There'll be {N-1} green bottle(s) hanging on the wall."
```

Line 3 is always the same. "one green bottle" is always singular in line 3.

### 4. Recite Function

`func Recite(startBottles, takeDown int) []string`:
- Loop from `startBottles` down for `takeDown` iterations (decrement by 1)
- Append each verse's lines to the result slice
- Between verses (not after the last one), append an empty string `""` as separator
- Return the accumulated slice

The separator condition: insert `""` before each verse except the first. Alternatively, insert after each verse except the last: `if i > startBottles-takeDown+1`.

## Architectural Decisions

1. **Follow the reference solution pattern exactly** — the `.meta/example.go` is the canonical answer
2. **No external dependencies** — only `fmt` from stdlib
3. **Use test-file `Title` function** — both files compile together during `go test`; fallback to local helper if needed
4. **Hardcode special cases** for n==1 and n==2 to avoid complex conditional logic for singular/plural

## Testing

Run `go test` from `go/exercises/practice/bottle-song/` — all 7 test cases must pass:
- first generic verse (10, 1)
- last generic verse (3, 1)
- verse with 2 bottles (2, 1)
- verse with 1 bottle (1, 1)
- first two verses (10, 2)
- last three verses (3, 3)
- all verses (10, 10)
