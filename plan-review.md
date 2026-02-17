# Plan Review

## Methodology

Manually traced the selected plan against all 7 test cases in `cases_test.go`.

## Test Case Verification

### 1. "first generic verse" — Recite(10, 1)
- i=10: numberWord[10]="ten" → capitalize → "Ten", bottle(10)="bottles"
- Lines 1-2: "Ten green bottles hanging on the wall," ✓
- Line 3: "And if one green bottle should accidentally fall," ✓
- nextN=9: numberWord[9]="nine", bottle(9)="bottles"
- Line 4: "There'll be nine green bottles hanging on the wall." ✓
- **PASS**

### 2. "last generic verse" — Recite(3, 1)
- i=3: "Three green bottles" ✓
- nextN=2: "two green bottles" ✓
- **PASS**

### 3. "verse with 2 bottles" — Recite(2, 1)
- i=2: "Two green bottles" ✓
- nextN=1: numberWord[1]="one", bottle(1)="bottle" (singular) ✓
- Line 4: "There'll be one green bottle hanging on the wall." ✓
- **PASS**

### 4. "verse with 1 bottle" — Recite(1, 1)
- i=1: numberWord[1]="one" → "One", bottle(1)="bottle" (singular) ✓
- Lines 1-2: "One green bottle hanging on the wall," ✓
- nextN=0: numberWord[0]="no", bottle(0)="bottles" (plural) ✓
- Line 4: "There'll be no green bottles hanging on the wall." ✓
- **PASS**

### 5. "first two verses" — Recite(10, 2)
- Verse 1 (i=10): same as test 1 ✓
- Separator: "" ✓
- Verse 2 (i=9): "Nine green bottles...eight green bottles" ✓
- **PASS**

### 6. "last three verses" — Recite(3, 3)
- Verse 1 (i=3): Three→two ✓
- Separator: "" ✓
- Verse 2 (i=2): Two→one (singular) ✓
- Separator: "" ✓
- Verse 3 (i=1): One (singular)→no (plural) ✓
- **PASS**

### 7. "all verses" — Recite(10, 10)
- All 10 verses with 9 separators
- Singulars/plurals verified at boundaries (2→1, 1→0) ✓
- **PASS**

## Issues Found

None. The plan correctly handles:
- Number-to-word conversion for 0-10
- Title case capitalization for first two lines of each verse
- Lowercase for the "There'll be" line
- Singular "bottle" for count == 1
- Plural "bottles" for count != 1 (including 0)
- Empty string separators between verses (not after last verse)

## Recommendation

**Approved.** The plan is sound and will pass all test cases. Proceed with implementation.
