# Verification Report: bottle-song

**Verdict: PASS**

## Independent Test Run

All 7 test cases pass:

```
=== RUN   TestRecite
=== RUN   TestRecite/first_generic_verse       --- PASS
=== RUN   TestRecite/last_generic_verse         --- PASS
=== RUN   TestRecite/verse_with_2_bottles       --- PASS
=== RUN   TestRecite/verse_with_1_bottle        --- PASS
=== RUN   TestRecite/first_two_verses           --- PASS
=== RUN   TestRecite/last_three_verses          --- PASS
=== RUN   TestRecite/all_verses                 --- PASS
PASS
ok  bottlesong
```

## Acceptance Criteria Checklist

### 1. All 7 test cases pass
**PASS** — Verified independently via `go test -v ./...`. All 7 subtests pass.

### 2. Number words are capitalized in first two lines of each verse
**PASS** — `strings.Title(numberToWord[n])` is used on lines 22-23 of `bottle_song.go`. The `numberToWord` map stores lowercase words ("ten", "one", etc.) and `strings.Title` capitalizes the first letter. Test expectations confirm: "Ten green bottles", "Three green bottles", "One green bottle", etc.

### 3. Singular/plural "bottle"/"bottles" is correct
**PASS** — The `bottleStr(n)` function (lines 13-18) returns "bottle" when n==1, "bottles" otherwise. Verified against test expectations:
- n=10: "Ten green **bottles**" ✓
- n=2: "Two green **bottles**" (first lines), "one green **bottle**" (There'll be line, n-1=1) ✓
- n=1: "One green **bottle**" (first lines), "no green **bottles**" (There'll be line, n-1=0) ✓

### 4. "no green bottles" for zero case
**PASS** — `numberToWord[0]` = "no" and `bottleStr(0)` = "bottles", producing "no green bottles" in the "There'll be..." line. Confirmed by test case "verse with 1 bottle" expected output: `"There'll be no green bottles hanging on the wall."`.

### 5. Verses separated by empty string
**PASS** — Lines 33-35 of `Recite()` append an empty string `""` between verses. The condition `i > startBottles-takeDown+1` correctly avoids a trailing separator. Confirmed by multi-verse test cases ("first two verses", "last three verses", "all verses").

### 6. Code compiles with go build
**PASS** — `go build ./...` completes without errors. Note: `strings.Title` is deprecated but still functional; this is acceptable per exercism constraints (no external dependencies allowed).

## Additional Checks

- Package name is `bottlesong` ✓
- Only `bottle_song.go` was modified ✓
- No external dependencies ✓
- Implementation is clean and concise (38 lines) ✓
