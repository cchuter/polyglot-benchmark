# Implementation Plan: bottle-song

## File to Modify

- `go/exercises/practice/bottle-song/bottle_song.go`

## Approach

The reference solution in `.meta/example.go` provides a clear pattern. We will implement the `Recite` function following the same approach, since the test file already provides the `Title` helper.

### Implementation Details

**1. Number-to-word map**

Create a map from int to lowercase word string for 0-10:
```go
var numberToWord = map[int]string{
    0: "no", 1: "one", 2: "two", 3: "three", 4: "four",
    5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
}
```

**2. Plural helper**

A simple function to return "bottle" or "bottles":
```go
func bottleStr(n int) string {
    if n == 1 { return "bottle" }
    return "bottles"
}
```

**3. Verse function**

Generate a single verse for n bottles:
```go
func verse(n int) []string {
    return []string{
        fmt.Sprintf("%s green %s hanging on the wall,", Title(numberToWord[n]), bottleStr(n)),
        fmt.Sprintf("%s green %s hanging on the wall,", Title(numberToWord[n]), bottleStr(n)),
        "And if one green bottle should accidentally fall,",
        fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberToWord[n-1], bottleStr(n-1)),
    }
}
```

**4. Recite function**

Loop from `startBottles` down for `takeDown` verses, joining with empty strings:
```go
func Recite(startBottles, takeDown int) []string {
    var result []string
    for i := startBottles; i > startBottles-takeDown; i-- {
        result = append(result, verse(i)...)
        if i > startBottles-takeDown+1 {
            result = append(result, "")
        }
    }
    return result
}
```

## Ordering

1. Write the complete implementation in `bottle_song.go`
2. Run `go test` to verify all 7 tests pass
3. Commit

## Rationale

- Using `Title()` from the test file for capitalization (available since it's in the same package)
- Using `fmt.Sprintf` for string construction — simple and readable
- Handling singular/plural via a helper function rather than special-casing each verse count
- The "no" case for zero bottles is handled naturally through the `numberToWord` map
