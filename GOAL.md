# Goal: Implement bottle-song Exercise in Go

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates lyrics to the "Ten Green Bottles" children's song.

The function signature is:
```go
func Recite(startBottles int, takeDown int) []string
```

- `startBottles`: the number of bottles to start from (1-10)
- `takeDown`: how many verses to recite (counting down from startBottles)

Each verse follows the pattern:
```
{N} green bottle(s) hanging on the wall,
{N} green bottle(s) hanging on the wall,
And if one green bottle should accidentally fall,
There'll be {N-1} green bottle(s) hanging on the wall.
```

Key variations:
- Numbers are spelled out as words with initial capital (e.g., "Ten", "Nine")
- "bottle" is singular when count is 1, "bottles" is plural otherwise
- When count reaches 0, use "no green bottles"
- Verses are separated by an empty string `""` in the returned slice

## Acceptance Criteria

1. `Recite(10, 1)` returns the first verse (Ten green bottles...)
2. `Recite(3, 1)` returns the verse starting at three
3. `Recite(2, 1)` correctly handles singular "one green bottle" in the "There'll be" line
4. `Recite(1, 1)` correctly uses singular "bottle" and "no green bottles" in the result
5. `Recite(10, 2)` returns two verses separated by an empty string
6. `Recite(3, 3)` returns last three verses with proper separators
7. `Recite(10, 10)` returns all ten verses of the complete song
8. All tests pass: `go test ./...` in the exercise directory
9. Code passes `go vet ./...`

## Key Constraints

- Package name must be `bottlesong`
- Must implement `func Recite(startBottles int, takeDown int) []string`
- Numbers must be English words with title case (Ten, Nine, ..., One, no)
- Singular/plural "bottle"/"bottles" must be correct
- Return type is `[]string` where each element is one line of lyrics
- Verses separated by empty string elements in the slice
