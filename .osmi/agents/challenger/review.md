# Bottle Song Implementation Review

## Verdict: PASS

The implementation is correct and should pass all 7 test cases.

## Detailed Review

### Number-to-word mapping (lines 8-20)
- Maps 0-10 correctly: no, one, two, three, four, five, six, seven, eight, nine, ten
- All values needed by tests are present
- Status: PASS

### Singular/plural handling (lines 26-31)
- `bottleWord(1)` returns "bottle" (singular)
- `bottleWord(n)` for n != 1 returns "bottles" (plural)
- Correctly applied to both current count (lines 1-2) and next count (line 4)
- Status: PASS

### Capitalization (lines 22-24, 33-42)
- Lines 1-2 of each verse use `capitalize()` on the number word (uppercase first letter)
- Line 3 is hardcoded with "And" capitalized
- Line 4 uses lowercase `nextWord` directly (no capitalization)
- Status: PASS

### Verse separators (lines 44-52)
- Empty string "" inserted between consecutive verses
- Not added before first verse or after last verse
- Logic: separator added when `i < startBottles` (i.e., not the first verse)
- Status: PASS

### Edge cases
- verse(1): word="one", bottleWord(1)="bottle" (singular), nextWord="no", bottleWord(0)="bottles" (plural)
- Produces: "One green bottle...no green bottles" - correct
- Status: PASS

### Test case trace-through
1. first generic verse (10,1): Ten...nine - PASS
2. last generic verse (3,1): Three...two - PASS
3. verse with 2 bottles (2,1): Two bottles...one bottle - PASS
4. verse with 1 bottle (1,1): One bottle...no bottles - PASS
5. first two verses (10,2): verses 10,9 with separator - PASS
6. last three verses (3,3): verses 3,2,1 with separators - PASS
7. all verses (10,10): all 10 verses with separators - PASS

## No issues found.
