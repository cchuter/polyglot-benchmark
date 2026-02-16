# Implementation Review: bottle-song

**Reviewer**: Challenger
**Date**: 2026-02-16
**File Reviewed**: `go/exercises/practice/bottle-song/bottle_song.go`

## Executive Summary

✅ **APPROVED** - The implementation is correct and fully satisfies all requirements.

## Review Criteria

### 1. Correctness: Edge Case Handling ✓

The switch statement correctly handles all three cases:
- **n==1**: Hardcoded singular verse with "no green bottles" ending
- **n==2**: Hardcoded verse with singular "one green bottle" ending
- **default (n≥3)**: Uses `fmt.Sprintf` with `strings.Title(numberToWord[n])` and `numberToWord[n-1]`

All edge cases are properly covered.

### 2. Singular/Plural Usage ✓

**n==1 case (lines 36-41):**
- Lines 1-2: "One green bottle" (singular) ✓
- Line 3: "if one green bottle" (singular) ✓
- Line 4: "no green bottles" (plural) ✓

**n==2 case (lines 43-48):**
- Lines 1-2: "Two green bottles" (plural) ✓
- Line 3: "if one green bottle" (singular) ✓
- Line 4: "one green bottle" (singular) ✓

**default case (lines 50-55):**
- Lines 1-2: "%s green bottles" (plural for n≥3) ✓
- Line 3: "if one green bottle" (singular) ✓
- Line 4: "%s green bottles" (plural, correctly handles n-1≥2) ✓

All singular/plural forms are correctly applied.

### 3. Capitalization ✓

- **n==1 and n==2 cases**: Manually capitalized ("One", "Two") ✓
- **default case**: Uses `strings.Title(numberToWord[n])` for lines 1-2 (capitalized) ✓
- **Line 4**: Uses lowercase `numberToWord[n-1]` ✓

Capitalization follows the specification: capitalized in the first two lines, lowercase in line 4.

### 4. Zero Count Handling ✓

The n==1 case correctly outputs "There'll be no green bottles hanging on the wall." (line 40). The `numberToWord` map includes `0: "no"` for completeness, though it's only used in the hardcoded string.

### 5. Verse Separators ✓

The `Recite` function (lines 8-17) correctly adds empty string separators:

```go
if i > startBottles-takeDown+1 {
    verses = append(verses, "")
}
```

This condition ensures:
- Empty strings are added **between** verses
- No empty string is added **after** the last verse

Tested logic:
- start=10, take=2: i=10,9 → separator after verse(10), none after verse(9) ✓
- start=3, take=3: i=3,2,1 → separators after verse(3) and verse(2), none after verse(1) ✓

### 6. Adherence to Plan ✓

The implementation faithfully follows **Branch 1: Switch-based** from `.osmi/plan.md`:

1. ✓ Defines `numberToWord` map (0→"no" through 10→"ten")
2. ✓ Creates `verse(n int) []string` helper with switch statement
3. ✓ `Recite` loops from `startBottles` down for `takeDown` iterations
4. ✓ Appends verse results with `""` separators between (not after last)
5. ✓ Uses `strings.Title` from standard library as suggested in plan revision

## Test Case Validation

All 7 test cases from `cases_test.go` are correctly handled:

1. ✓ **first generic verse** (n=10): Capitalizes "Ten", uses "nine"
2. ✓ **last generic verse** (n=3): Capitalizes "Three", uses "two"
3. ✓ **verse with 2 bottles** (n=2): Singular "one green bottle" in line 4
4. ✓ **verse with 1 bottle** (n=1): "no green bottles" in line 4
5. ✓ **first two verses**: Separator added between verses
6. ✓ **last three verses**: Covers n=3,2,1 with correct separators
7. ✓ **all verses**: Full song from 10→1 with proper separators

## Code Quality

- **Simplicity**: Clean, minimal code (~58 lines)
- **Readability**: Clear switch cases, descriptive variable names
- **Maintainability**: Follows established patterns from reference solution
- **Style**: Proper Go idioms, appropriate use of standard library

## Issues Found

**None.** The implementation is correct and complete.

## Recommendation

✅ **APPROVE FOR MERGE** - This implementation fully satisfies all acceptance criteria and test cases.
