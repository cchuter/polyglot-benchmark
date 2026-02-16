# Plan Review: bottle-song

## Review Questions & Findings

### (1) Does the selected plan correctly handle all edge cases from cases_test.go?

**Yes.** The 7 test cases are:
- Single verse for 10 (default path) ✓
- Single verse for 3 (default path) ✓
- Single verse for 2 (special case: singular "one green bottle" in result) ✓
- Single verse for 1 (special case: "no green bottles" in result, singular "bottle" in lines 1-2) ✓
- Two verses starting at 10 (separator between verses) ✓
- Three verses starting at 3 (covers 3→2→1, all paths exercised) ✓
- All 10 verses (comprehensive) ✓

### (2) Is using strings.Title safe given go.mod specifies go 1.18?

**Yes, with a caveat.** `strings.Title` is deprecated since Go 1.18 but is still available and functional. It works correctly for simple ASCII words like "ten", "nine", etc. The deprecation notice recommends `golang.org/x/text/cases` but that requires an external dependency. Using `strings.Title` is the pragmatic choice here and the test file comments explicitly endorse it.

### (3) Are there any issues with the Title function availability?

**Important finding:** The `Title` function is defined in `bottle_song_test.go`. During `go test`, test files are compiled into the same package, so `Title` IS accessible from production code **only during testing**. However, if someone tries to `go build` the package standalone (not via `go test`), referencing `Title` from `bottle_song.go` would fail because test files aren't included in normal builds.

**Recommendation:** Do NOT reference the test file's `Title` function from `bottle_song.go`. Use `strings.Title` instead, which is self-contained.

### (4) Does the verse separator logic correctly produce empty strings between verses but not after the last verse?

**Yes.** The condition `if i > startBottles-takeDown+1` ensures the empty string separator is added after every verse except the last one. For example, with `startBottles=10, takeDown=2`, the loop runs for i=10 and i=9. The separator is added after i=10 (since 10 > 10-2+1 = 9) but not after i=9 (since 9 is not > 9).

## Overall Assessment

The selected plan (Branch 1: Switch-based) is sound. The key correction is to use `strings.Title` rather than the test file's `Title` function. The implementation closely mirrors the reference solution, which gives high confidence in correctness.

**Verdict: Approved with the strings.Title correction noted above.**
