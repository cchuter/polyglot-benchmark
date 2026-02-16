# Challenger Review: bottle-song Implementation

## Verdict: PASS

All 7 test cases pass. The implementation is correct.

## Detailed Findings

### 1. Number-to-word mapping (1-10)
**Status: Correct**
The `numberToWord` map covers all values 1-10. All words are lowercase, matching expected mid-line usage.

### 2. Singular/plural "bottle"/"bottles" handling
**Status: Correct**
- n >= 3 (default case): Uses "bottles" (plural) for both the current count and the result. Correct since n >= 3 and n-1 >= 2.
- n == 2 (hardcoded): Uses "bottles" for 2, "bottle" (singular) for the result (1). Correct.
- n == 1 (hardcoded): Uses "bottle" (singular) for 1, "bottles" (plural) for "no green bottles". Correct.

### 3. Verse separation with empty strings
**Status: Correct**
The condition `i > startBottles-takeDown+1` correctly adds an empty string between verses but not after the last verse.

### 4. Edge cases
**Status: All correct**
- 1 bottle: singular "bottle" + "no green bottles" (plural for zero)
- 2 bottles: plural "bottles" + singular "bottle" in result line
- 0 bottles: "no green bottles" (never a starting count, only appears as result)

### 5. Title-casing
**Status: Correct**
- `Title(numberToWord[n])` produces capitalized words at line start (e.g., "Ten", "Nine")
- `numberToWord[n-1]` stays lowercase in mid-line position (e.g., "nine", "eight")
- The `Title` function is defined in `bottle_song_test.go` and is accessible during `go test`

### 6. Function signature
**Status: Correct**
`Recite(startBottles, takeDown int) []string` matches the test expectations exactly.

## Test Results

```
=== RUN   TestRecite
=== RUN   TestRecite/first_generic_verse       PASS
=== RUN   TestRecite/last_generic_verse         PASS
=== RUN   TestRecite/verse_with_2_bottles       PASS
=== RUN   TestRecite/verse_with_1_bottle        PASS
=== RUN   TestRecite/first_two_verses           PASS
=== RUN   TestRecite/last_three_verses          PASS
=== RUN   TestRecite/all_verses                 PASS
--- PASS: TestRecite (0.00s)
PASS
```

## Notes

- The `Title` helper is defined in `bottle_song_test.go`, so `bottle_song.go` references it. This works during `go test` (Go compiles test files and package source together), but `go build` alone would fail. This is acceptable per the GOAL.md which explicitly states "The `Title` helper function is available in `bottle_song_test.go` for use" and acceptance criteria only requires `go test` to pass.
- The approach of hardcoding cases for n=1 and n=2 is clean and avoids complex singular/plural branching in the default case.
