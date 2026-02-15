# Plan Review: bottle-song Go Exercise

## Overall Assessment

The plan is **mostly sound** and will likely produce a working implementation, but it contains one significant correctness issue and several misleading or inaccurate statements that could cause confusion or bugs during implementation. The plan claims to follow the reference implementation but actually proposes a design that diverges from it in an important way related to the `Title()` function. The reference implementation at `.meta/example.go` is the canonical solution and would pass all tests; the plan's description of how to use `Title()` introduces a risk that the implementer will call a function that is not accessible from the production code file.

## Correctness Issues

### 1. `Title()` accessibility

The `Title()` function defined in `bottle_song_test.go` is accessible from `bottle_song.go` at test time because they share the same package `bottlesong`. Calling `Title()` from the production code means the code will only compile when the test file is present. The reference implementation at `.meta/example.go` also calls `Title()`, so the plan is aligned with the reference on this point. The plan should explicitly acknowledge that this is a test-time-only compilation dependency and that this is intentional/acceptable per the Exercism exercise design.

### 2. Plan's verse logic description is incomplete

The plan identifies three cases (n==1, n==2, default) but does not explicitly describe the formatting difference between lines 1-2 (title case) and line 4 (lowercase). Lines 1-2 use `Title(numberToWord[n])` producing "Ten", "Nine", etc. Line 4 uses `numberToWord[n-1]` producing "nine", "eight", etc. (lowercase). This nuance needs to be explicit.

### 3. The plan's map includes key 0 but reference does not

The plan shows `0: "no"` in the map, but the reference implementation does not include 0. Since n==1 is already handled as a special case returning hard-coded strings, the 0 key would never be used. Not harmful but misleading.

## Completeness Gaps

1. No code shown for the `Recite` function looping logic
2. The separator insertion condition (`i > startBottles-takeDown+1`) is the most error-prone part and not detailed
3. Import requirements not fully specified (only `fmt` needed)

## Recommendations

1. **Remove `0: "no"` from the map** to match the reference implementation
2. **Show the separator logic explicitly**: "insert an empty string after each verse except the last one" with the boundary condition
3. **Explicitly describe title-case vs. lowercase placement** in verse lines
4. **Add the verse helper code** to reduce implementation risk
5. **Acknowledge `Title()` is from the test file** — this is an accepted Exercism convention
