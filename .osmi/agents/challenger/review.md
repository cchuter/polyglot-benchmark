# Challenger Review: bottle-song Implementation

## Verdict: PASS

The implementation is correct and will pass all 7 test cases.

## Test Case Trace

### Test 1: "first generic verse" (startBottles=10, takeDown=1)
- `verse(10)` hits `default` case: capitalize("ten") = "Ten", numberToWord[9] = "nine"
- Output: `["Ten green bottles hanging on the wall,", ..., "There'll be nine green bottles hanging on the wall."]`
- **PASS**

### Test 2: "last generic verse" (startBottles=3, takeDown=1)
- `verse(3)` hits `default` case: capitalize("three") = "Three", numberToWord[2] = "two"
- Output: `["Three green bottles hanging on the wall,", ..., "There'll be two green bottles hanging on the wall."]`
- **PASS**

### Test 3: "verse with 2 bottles" (startBottles=2, takeDown=1)
- `verse(2)` hits `case n == 2`: hardcoded with plural "bottles" in lines 1-2, singular "bottle" in line 4
- Output: `["Two green bottles hanging on the wall,", ..., "There'll be one green bottle hanging on the wall."]`
- **PASS**

### Test 4: "verse with 1 bottle" (startBottles=1, takeDown=1)
- `verse(1)` hits `case n == 1`: hardcoded with singular "bottle" in lines 1-2, "no green bottles" in line 4
- Output: `["One green bottle hanging on the wall,", ..., "There'll be no green bottles hanging on the wall."]`
- **PASS**

### Test 5: "first two verses" (startBottles=10, takeDown=2)
- Loop: i=10 (verse + separator), i=9 (verse, no trailing separator)
- Separator logic: `10 > 10-2+1=9` is true, so empty string appended after verse 10
- `9 > 9` is false, so no trailing separator
- **PASS**

### Test 6: "last three verses" (startBottles=3, takeDown=3)
- Loop: i=3 (default, separator), i=2 (case 2, separator), i=1 (case 1, no trailing separator)
- Separator logic: `3 > 1` true, `2 > 1` true, `1 > 1` false
- **PASS**

### Test 7: "all verses" (startBottles=10, takeDown=10)
- Loop i=10 down to 1. Each verse correct. Separators between all consecutive verses.
- i=10..2 get separators, i=1 does not (since `1 > 1` is false)
- **PASS**

## Edge Case Analysis

| Edge Case | Handling | Status |
|-----------|----------|--------|
| Singular "bottle" (n=1 in lines 1-2) | Explicit `case n == 1` with hardcoded singular | Correct |
| Singular "bottle" (n=2 remaining, line 4) | Explicit `case n == 2` with hardcoded singular in line 4 | Correct |
| Plural "bottles" (n >= 2 in lines 1-2) | Default case uses plural | Correct |
| "no green bottles" (n=0 remaining) | Handled in `case n == 1` with hardcoded line 4 | Correct |
| Capitalization (first line of verse) | `capitalize()` uppercases first byte | Correct |
| Lowercase in line 4 | Uses `numberToWord[n-1]` directly (lowercase) | Correct |
| Verse separation | Empty string `""` between verses, none after last | Correct |

## Code Quality

- **Structure**: Clean separation into `numberToWord` map, `capitalize` helper, `verse` generator, and `Recite` orchestrator. Matches the plan exactly.
- **Imports**: `fmt` (used in `Sprintf`) and `strings` (used in `ToUpper`) -- both used, no unused imports.
- **No code smells**: No dead code, no unnecessary complexity.
- **Plan adherence**: Implementation follows the plan precisely -- explicit cases for n==1 and n==2, default for n>=3, loop with separator logic.

## Minor Observations (non-blocking)

1. The `capitalize` function would panic on an empty string, but it's only called with values from `numberToWord` which are all non-empty. This is acceptable for an internal helper.
2. The `Title` function defined in the test file is unused by the implementation -- this is fine since it's provided by the test harness for students who want it.

## Conclusion

The implementation is correct, clean, and follows the plan. No changes needed. Ready for test execution.
