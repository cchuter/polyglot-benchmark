# Implementation Plan: bottle-song

## File to Modify

- `go/exercises/practice/bottle-song/bottle_song.go` — the only file to change

## Approach

Implement `Recite` following the reference solution pattern from `.meta/example.go`:

### 1. Number-to-word mapping

Create a `map[int]string` for numbers 0-10:
```go
var numberToWord = map[int]string{
    0: "no", 1: "one", 2: "two", 3: "three", 4: "four",
    5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
}
```

### 2. Helper: `verse(n int) []string`

Generate a single verse for `n` bottles:
- Lines 1-2: `"{N} green bottle(s) hanging on the wall,"` (N title-cased)
- Line 3: `"And if one green bottle should accidentally fall,"`
- Line 4: `"There'll be {n-1} green bottle(s) hanging on the wall."`

Special cases:
- `n == 1`: singular "bottle" in lines 1-2, result is "no green bottles" (plural)
- `n == 2`: plural "bottles" in lines 1-2, result is singular "one green bottle"
- `n >= 3`: plural throughout

### 3. Title-casing

Use `strings.Title` (deprecated but functional and used in the reference solution context) or implement a simple title-case for the first letter. Since the test file provides a `Title` function and the reference uses it, we can simply capitalize the first letter of the number word using `strings.ToUpper` on the first character + rest of string, or use `fmt.Sprintf` with manual capitalization.

Decision: Use a simple inline approach — `strings.ToUpper(word[:1]) + word[1:]` for title-casing, avoiding deprecated `strings.Title`.

### 4. `Recite(startBottles, takeDown int) []string`

Loop from `startBottles` down for `takeDown` iterations, appending each verse. Insert empty string `""` between verses (not after last).

## Ordering

1. Add imports (`fmt`, `strings`)
2. Add `numberToWord` map
3. Add `verse(n)` helper
4. Add `Recite` function
5. Run tests to verify
