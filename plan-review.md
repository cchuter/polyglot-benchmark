# Plan Review: Bottle Song

## Verdict: APPROVE

The selected plan (Proposal B - Generalized approach) is correct and will pass all 7 test cases. No revisions are needed.

---

## Detailed Analysis

### 1. Test Case Tracing

**Test Case 1: "first generic verse" (startBottles=10, takeDown=1)**
- i=0, n=10
- `word` = Title("ten") = "Ten"
- `nextWord` = numberWord(9) = "nine"
- `bottleStr(10)` = "bottles", `bottleStr(9)` = "bottles"
- Output: `["Ten green bottles hanging on the wall,", "Ten green bottles hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be nine green bottles hanging on the wall."]`
- PASS: Matches expected output exactly.

**Test Case 4: "verse with 1 bottle" (startBottles=1, takeDown=1)**
- i=0, n=1
- `word` = Title("one") = "One"
- `nextWord` = numberWord(0) = "no"
- `bottleStr(1)` = "bottle" (singular), `bottleStr(0)` = "bottles" (plural)
- Output: `["One green bottle hanging on the wall,", "One green bottle hanging on the wall,", "And if one green bottle should accidentally fall,", "There'll be no green bottles hanging on the wall."]`
- PASS: Matches expected output exactly. Singular/plural and "no" at zero are both correct.

**Test Case 6: "last three verses" (startBottles=3, takeDown=3)**
- i=0, n=3: word="Three", nextWord="two", bottleStr(3)="bottles", bottleStr(2)="bottles"
  - Lines 1-4: Three green bottles..., There'll be two green bottles...
- i=1 (separator ""), n=2: word="Two", nextWord="one", bottleStr(2)="bottles", bottleStr(1)="bottle"
  - Lines 6-9: Two green bottles..., There'll be one green bottle...
- i=2 (separator ""), n=1: word="One", nextWord="no", bottleStr(1)="bottle", bottleStr(0)="bottles"
  - Lines 11-14: One green bottle..., There'll be no green bottles...
- PASS: Matches the 14-element expected array (including two empty-string separators) exactly.

**Remaining test cases verified by inspection:**
- Test Case 2 ("last generic verse", startBottles=3, takeDown=1): Same pattern as TC1 with n=3, next=2. Correct.
- Test Case 3 ("verse with 2 bottles", startBottles=2, takeDown=1): n=2 uses "bottles", n-1=1 uses "bottle". Correct.
- Test Case 5 ("first two verses", startBottles=10, takeDown=2): Two verses with separator. Correct.
- Test Case 7 ("all verses", startBottles=10, takeDown=10): All 10 verses from 10 down to 1, with 9 separators. Correct.

### 2. Edge Cases

All edge cases present in the test suite are handled:

- **n=1 (singular)**: `bottleStr(1)` returns "bottle". The first two lines correctly say "One green bottle". Covered by Test Case 4.
- **n=0 (zero bottles)**: Only appears in the last line of a verse (n-1 when n=1). `numberWord(0)` returns "no" and `bottleStr(0)` returns "bottles" (plural). This matches the expected "no green bottles". Covered by Test Cases 4, 6, and 7.
- **n=2 transitioning to n=1**: The last line correctly uses singular "bottle" for n-1=1. Covered by Test Cases 3 and 6.
- **Multiple verses with separators**: Empty string separators are correctly inserted between verses (only when `i > 0`). Covered by Test Cases 5, 6, and 7.
- **Single verse (no separators)**: When takeDown=1, no separator is added. Covered by Test Cases 1-4.

No edge cases are missing. The domain is constrained to startBottles in [1,10] and takeDown in [1,startBottles], which the `numbers` slice fully covers (indices 0-10).

### 3. Title Function Usage

The `Title` function is defined in `bottle_song_test.go` within the `bottlesong` package. Since `bottle_song.go` is also in the `bottlesong` package, `Title` is accessible at package scope during test execution. The plan correctly uses `Title(numberWord(n))` only for the first two lines of each verse (capitalized), while the last line uses `numberWord(n-1)` without `Title` (lowercase). This matches the expected output where verse-starting lines use "Ten", "Nine", etc. and the closing line uses "nine", "eight", etc.

Note: `Title` is defined in the test file, not in the implementation file. This is intentional by the exercise design -- Go test files in the same package share scope. The implementation correctly relies on this.

### 4. Singular/Plural Logic

The `bottleStr` function:
- `n == 1` returns "bottle" (singular) -- correct
- `n != 1` (including n=0) returns "bottles" (plural) -- correct

This covers all cases:
- n=0: "bottles" (as in "no green bottles") -- correct per test cases
- n=1: "bottle" (as in "one green bottle") -- correct per test cases
- n=2 through n=10: "bottles" -- correct per test cases

### 5. Code Quality Notes

- The `numbers` slice approach (index-based) is cleaner than a map for sequential integers 0-10.
- The loop structure with `i > 0` for separator insertion is clean and correct.
- The `fmt` import is required and included.
- No external dependencies are needed.

### Summary

The plan is sound, correct, and complete. All 7 test cases will pass. The singular/plural logic, the `Title` function usage, and the verse separator logic are all correct. No revisions needed.
