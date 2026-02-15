# Challenger Review: bottle-song Implementation

## Verdict: PASS — Implementation is correct

## Test Case Analysis

### Test 1: "first generic verse" — Recite(10, 1)
- `Title("ten")` → `"Ten"`, `bottleStr(10)` → `"bottles"` → first two lines correct
- `numberToWord[9]` → `"nine"`, `bottleStr(9)` → `"bottles"` → fourth line correct
- Single verse, no separator needed
- **PASS**

### Test 2: "last generic verse" — Recite(3, 1)
- `Title("three")` → `"Three"`, `bottleStr(3)` → `"bottles"`
- `numberToWord[2]` → `"two"`, `bottleStr(2)` → `"bottles"`
- **PASS**

### Test 3: "verse with 2 bottles" — Recite(2, 1)
- `Title("two")` → `"Two"`, `bottleStr(2)` → `"bottles"`
- `numberToWord[1]` → `"one"`, `bottleStr(1)` → `"bottle"` (singular)
- Expected: `"There'll be one green bottle hanging on the wall."` — matches
- **PASS**

### Test 4: "verse with 1 bottle" — Recite(1, 1)
- `Title("one")` → `"One"`, `bottleStr(1)` → `"bottle"` (singular)
- `numberToWord[0]` → `"no"`, `bottleStr(0)` → `"bottles"` (plural for zero)
- Expected: `"One green bottle hanging on the wall,"` and `"There'll be no green bottles hanging on the wall."` — matches
- **PASS**

### Test 5: "first two verses" — Recite(10, 2)
- Loop: i=10 (separator), i=9 (no separator)
- Separator condition: `i=10 > 10-2+1=9` → true, adds `""`
- Separator condition: `i=9 > 9` → false, no trailing separator
- **PASS**

### Test 6: "last three verses" — Recite(3, 3)
- Loop: i=3, i=2, i=1
- Separators after verses for i=3 and i=2 (both > 1); none after i=1
- Includes transition from plural to singular to zero — all handled correctly
- **PASS**

### Test 7: "all verses" — Recite(10, 10)
- Full loop from 10 to 1 with separators between each verse
- All number words, singular/plural transitions, and capitalization correct
- **PASS**

## Edge Cases Verified

| Edge Case | Mechanism | Correct? |
|---|---|---|
| Singular "bottle" for n=1 | `bottleStr(1)` returns `"bottle"` | Yes |
| Plural "bottles" for n=0 | `bottleStr(0)` returns `"bottles"` | Yes |
| "no" for zero bottles | `numberToWord[0]` returns `"no"` | Yes |
| Capitalization on lines 1-2 | `Title()` applied to number word | Yes |
| No capitalization on line 4 | Raw `numberToWord[n-1]` used (lowercase) | Yes |
| Verse separation | Empty string `""` between verses, not after last | Yes |

## Potential Issues Reviewed

1. **Map completeness**: All values 0-10 are covered in `numberToWord`. Since the exercise only uses bottles 1-10, this is sufficient.
2. **Title function access**: `Title` is defined in `bottle_song_test.go` in the same `bottlesong` package — accessible during `go test`.
3. **Loop bounds**: `i > startBottles-takeDown` correctly produces exactly `takeDown` iterations.
4. **Separator logic**: `i > startBottles-takeDown+1` correctly skips the separator after the final verse.
5. **No off-by-one errors** detected in any path.

## Conclusion

The implementation faithfully follows the plan and correctly handles all 7 test cases, including all edge cases around singular/plural forms, zero bottles, capitalization, and verse separation. No bugs or issues found.
