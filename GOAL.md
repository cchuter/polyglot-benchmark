# Goal: Implement Bottle Song (Ten Green Bottles)

## Problem Statement

Implement the `Recite` function in `go/exercises/practice/bottle-song/bottle_song.go` that generates the lyrics to the children's song "Ten Green Bottles." The function takes two parameters:
- `startBottles` (int): the number of bottles to start from
- `takeDown` (int): how many verses to recite

Each verse follows this pattern:
```
{N} green bottle(s) hanging on the wall,
{N} green bottle(s) hanging on the wall,
And if one green bottle should accidentally fall,
There'll be {N-1} green bottle(s) hanging on the wall.
```

Numbers must be spelled out as English words with title case (e.g., "Ten", "Nine").
The word "bottle" must use correct singular/plural form ("bottle" for 1, "bottles" for 0 and 2+).
When the count reaches 0, use "no green bottles".

## Acceptance Criteria

1. `Recite(n, 1)` returns a single verse starting from `n` bottles
2. `Recite(n, k)` returns `k` consecutive verses starting from `n`, separated by empty strings (`""`)
3. Numbers are spelled out as capitalized English words (Ten, Nine, ..., One)
4. "bottle" is singular when referring to exactly one bottle; "bottles" is plural otherwise
5. The last line of the final verse (1 bottle) uses "no green bottles"
6. All 7 test cases pass: first generic verse, last generic verse, verse with 2 bottles, verse with 1 bottle, first two verses, last three verses, all verses
7. `go vet` produces no warnings

## Key Constraints

- Function signature: `func Recite(startBottles int, takeDown int) []string`
- Returns `[]string` where each element is one line of the song
- Verses are separated by an empty string element (`""`) in the returned slice
- Package name: `bottlesong`
- No external dependencies (go.mod specifies only `go 1.18`)
