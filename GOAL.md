# Goal: Implement bottle-song Go Exercise

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` to produce the lyrics of the children's song "Ten Green Bottles".

The function signature is:
```go
func Recite(startBottles, takeDown int) []string
```

- `startBottles`: the number to start counting down from (1-10)
- `takeDown`: how many verses to produce
- Returns a slice of strings, where each string is one line of the song

## Acceptance Criteria

1. **All 7 test cases pass** (`go test` in the bottle-song directory returns PASS):
   - "first generic verse" — `Recite(10, 1)` returns the verse for 10 bottles
   - "last generic verse" — `Recite(3, 1)` returns the verse for 3 bottles
   - "verse with 2 bottles" — `Recite(2, 1)` returns verse with singular "bottle" in the last line
   - "verse with 1 bottle" — `Recite(1, 1)` returns verse with singular "bottle" and "no green bottles" in last line
   - "first two verses" — `Recite(10, 2)` returns two verses separated by an empty string
   - "last three verses" — `Recite(3, 3)` returns three verses including singular edge cases
   - "all verses" — `Recite(10, 10)` returns all 10 verses of the complete song

2. **Number words are capitalized** in the first two lines of each verse (e.g., "Ten green bottles")

3. **Singular/plural "bottle"/"bottles"** is correct:
   - 1 bottle uses singular "bottle"
   - 0 and 2+ bottles use plural "bottles"

4. **"no green bottles"** is used for the zero case in "There'll be..." lines

5. **Verses are separated by an empty string** `""` in the returned slice

6. **Code compiles** with `go build`

## Key Constraints

- Package name must be `bottlesong`
- Only the file `bottle_song.go` should be modified
- Go 1.18 module
- No external dependencies
