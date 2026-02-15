# Implementation Plan: bottle-song (Revised)

## Overview

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go`. This is a single-file change following the reference implementation pattern in `.meta/example.go`.

## File to Modify

**`go/exercises/practice/bottle-song/bottle_song.go`** — Replace the stub with a working implementation.

## Architectural Decisions

1. **Follow the reference implementation exactly**: The `.meta/example.go` provides the canonical solution. Our implementation mirrors its structure.

2. **Number-to-word map**: Use a `map[int]string` for integers 1-10 only (not 0). The n==1 case is hard-coded so 0 is never looked up.

3. **Verse helper function**: Private `verse(n int)` function generates a single verse as `[]string` with three cases via switch.

4. **Title case convention**: `Title()` is defined in the test file (`bottle_song_test.go`) and is accessible from `bottle_song.go` because they share the `bottlesong` package. This is an accepted Exercism convention — the implementation compiles at test time. Lines 1-2 of each verse use `Title(numberToWord[n])` (capitalized: "Ten"), while line 4 uses `numberToWord[n-1]` (lowercase: "nine").

5. **Only `fmt` is imported** — no other stdlib packages needed.

## Implementation Code

### Step 1: Number-to-word mapping

```go
var numberToWord = map[int]string{
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
```

### Step 2: Verse helper (three cases)

```go
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

### Step 3: Recite function with separator logic

```go
func Recite(startBottles, takeDown int) []string {
    verses := []string{}
    for i := startBottles; i > startBottles-takeDown; i -= 1 {
        verses = append(verses, verse(i)...)
        if i > startBottles-takeDown+1 {
            verses = append(verses, "")
        }
    }
    return verses
}
```

The separator condition `i > startBottles-takeDown+1` inserts an empty string after every verse **except the last one**.

## Ordering

1. Write the implementation to `bottle_song.go`
2. Run `go test` to verify all 7 test cases pass
3. Commit
