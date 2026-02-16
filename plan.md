# Implementation Plan: bottle-song

## Branch 1: Switch-based (Minimal, matches example.go)

Mirror the reference implementation in `.meta/example.go` closely. Use a map for number-to-word conversion and a switch statement in `verse()` to handle the three cases (n==1, n==2, default).

### Files to modify
- `go/exercises/practice/bottle-song/bottle_song.go`

### Approach
1. Define a `numberToWord` map (0→"no", 1→"one", ..., 10→"ten")
2. Create a `verse(n int) []string` helper using a switch:
   - `n == 1`: hardcoded singular verse with "no green bottles" ending
   - `n == 2`: hardcoded verse with singular "one green bottle" ending
   - `default`: use fmt.Sprintf with Title(numberToWord[n]) and numberToWord[n-1]
3. `Recite` loops from `startBottles` down for `takeDown` iterations, appending verse results with `""` separators

### Evaluation
- **Feasibility**: High — directly mirrors the proven example.go
- **Risk**: Very low — known working pattern
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~55 lines, 1 file, very straightforward

## Branch 2: Generalized with helper functions (Extensible)

Create a more structured approach with separate helper functions for pluralization, number formatting, and line generation.

### Files to modify
- `go/exercises/practice/bottle-song/bottle_song.go`

### Approach
1. Define `numberWord(n int) string` returning the word for a number
2. Define `bottlePhrase(n int) string` returning "N green bottle(s)" with correct pluralization
3. Define `capitalizedBottlePhrase(n int) string` returning capitalized version
4. `verse(n int) []string` builds 4 lines using these helpers
5. `Recite` orchestrates verse collection with separators

### Evaluation
- **Feasibility**: High — standard Go patterns
- **Risk**: Low — more functions but each is simple
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~70 lines, 1 file, more abstracted but more functions than needed

## Branch 3: Array-based lookup (Performance/Unconventional)

Pre-compute all 10 verses as constant string slices at package init time, then Recite just slices into the pre-built array.

### Files to modify
- `go/exercises/practice/bottle-song/bottle_song.go`

### Approach
1. Define a `var allVerses [11][]string` containing pre-built verse lines for bottles 1-10
2. In `init()`, populate all verses
3. `Recite` simply concatenates slices from the pre-built array with separators

### Evaluation
- **Feasibility**: High — straightforward
- **Risk**: Low, but init() is unnecessary complexity for this problem
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~80 lines, 1 file, over-engineered for the problem size

## Selected Plan

**Branch 1: Switch-based** is the best choice.

### Rationale
- **Simplest**: Fewest lines of code, fewest abstractions
- **Proven**: Directly mirrors the reference solution in `.meta/example.go`
- **Low risk**: The pattern is well-established in this exercise suite
- **Fully aligned**: Satisfies every acceptance criterion
- Branch 2 adds unnecessary abstraction for a problem with only 10 numbers. Branch 3 over-engineers with pre-computation that provides no real benefit.

### Detailed Implementation

**File**: `go/exercises/practice/bottle-song/bottle_song.go`

```go
package bottlesong

import "fmt"

// Recite returns the lyrics for the specified verses of "Ten Green Bottles".
func Recite(startBottles, takeDown int) []string {
	verses := []string{}
	for i := startBottles; i > startBottles-takeDown; i-- {
		verses = append(verses, verse(i)...)
		if i > startBottles-takeDown+1 {
			verses = append(verses, "")
		}
	}
	return verses
}

var numberToWord = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func verse(n int) []string {
	switch {
	case n == 1:
		return []string{
			"One green bottle hanging on the wall,",
			"One green bottle hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be no green bottles hanging on the wall.",
		}
	case n == 2:
		return []string{
			"Two green bottles hanging on the wall,",
			"Two green bottles hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be one green bottle hanging on the wall.",
		}
	default:
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberToWord[n])),
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberToWord[n])),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottles hanging on the wall.", numberToWord[n-1]),
		}
	}
}
```

Note: The `Title` function is defined in `bottle_song_test.go` and is accessible within the same package. We use it for capitalizing number words. We also define our own copy to avoid depending on the test file at build time — actually, since the test file is in the same package, `Title` is available. However, to avoid relying on test-only code, we should use `strings.Title` directly (deprecated but functional for Go 1.18) or implement our own capitalization inline.

**Decision**: Use `strings.Title` from the standard library since the test file comments explicitly say "Students should still be able to use strings.Title if they want." This keeps the implementation self-contained.

### Revised implementation detail for capitalization:
- Use `strings.Title(numberToWord[n])` instead of the test file's `Title` function
- Import `strings` and `fmt`
