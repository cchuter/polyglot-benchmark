# Implementation Plan Review: bottle-song

## Summary

**Overall Assessment: Plan has critical flaws that will prevent tests from passing.**

The plan correctly identifies the overall structure but contains a fundamental architectural error regarding the `Title()` function that will cause compilation failures. Additionally, there are subtle logic errors in the verse generation approach.

## Critical Issues

### 1. Title Function Dependency (BLOCKING)

**Problem**: The plan states "Use the `Title()` function defined in the test file directly" and claims "both files share the `bottlesong` package, this is valid Go."

**Why This Fails**: While both files are in the `bottlesong` package, `Title()` is defined in `bottle_song_test.go`, which is part of the test binary, not the production code. The production code in `bottle_song.go` cannot call functions defined in test files. This will result in a compilation error: "undefined: Title".

**Evidence from Reference Solution**: The reference solution in `.meta/example.go` does use `Title()`, but this works because `.meta/example.go` is a standalone reference implementation. When students implement `bottle_song.go`, they must provide their own title-casing logic.

**Fix Required**: The plan must either:
- Implement a local `title()` helper function in `bottle_song.go`
- Use `strings.ToUpper(string(s[0])) + s[1:]` for simple capitalization
- Import and use a standard library alternative

### 2. Verse Generation Logic Error

**Problem**: The plan's verse generator approach has hardcoded strings for n==1 and n==2 cases, which is correct. However, the description in section 3 is misleading about the "default" case.

**Actual Issue**: Looking at the reference solution's default case (lines 50-54), it uses:
- `Title(numberToWord[n])` for the first two lines
- `numberToWord[n-1]` for the fourth line

The plan description says "plural 'bottles' throughout" for n >= 3, which is correct, but doesn't clarify that when n==3, the fourth line becomes "two green bottles" (still plural), not "one green bottle" (singular). The reference solution handles this correctly in the hardcoded n==2 case.

**Actual Logic**: The reference solution is correct. When n==3, line 4 becomes "There'll be two green bottles..." When n==2, line 4 becomes "There'll be one green bottle..." (singular). The hardcoded cases handle both the singular "bottle" scenarios (n==1 and the result when n==2).

**Conclusion**: The reference solution logic is actually correct. The plan's description could be clearer, but if implemented exactly as the reference shows, it will work.

### 3. Number-to-Word Map Coverage

**Observation**: The plan specifies mapping integers 1-10, which matches the reference solution. However, it doesn't account for the case where 0 bottles should be represented as "no" (not "zero").

**Evidence from Tests**: The n==1 case expects "There'll be no green bottles..." not "There'll be zero green bottles..."

**Status**: The reference solution handles this correctly by hardcoding the n==1 case with "no green bottles" in line 4. The map doesn't need a 0 entry because it's never used (the verse generator hardcodes all scenarios where 0 would appear).

**Conclusion**: Not a bug, but the plan could explain this edge case more clearly.

## Correctness Analysis

### Test Case Coverage

Reviewing against the 7 test cases:

1. **first generic verse (10, 1)**: Will work if Title issue is fixed
2. **last generic verse (3, 1)**: Will work if Title issue is fixed
3. **verse with 2 bottles (2, 1)**: Correctly hardcoded, will work
4. **verse with 1 bottle (1, 1)**: Correctly hardcoded with "no green bottles", will work
5. **first two verses (10, 2)**: Empty string separator logic looks correct
6. **last three verses (3, 3)**: Tests transition from 3->2->1, will work if all cases correct
7. **all verses (10, 10)**: Full integration test, will work if individual verses correct

**Verdict**: All test cases should pass IF the Title function issue is resolved.

## Architecture Review

### Strengths

1. **Separation of concerns**: Verse generation is cleanly separated from the recite logic
2. **Hardcoded special cases**: Using explicit strings for n==1 and n==2 avoids complex conditional logic
3. **Simple data structure**: Map for number-to-word is straightforward and maintainable

### Weaknesses

1. **Test-time dependency**: Relying on a test-file function is architecturally unsound
2. **No validation**: The plan doesn't mention input validation (though the tests may not require it)
3. **Magic numbers**: The separator logic `i > startBottles-takeDown+1` could be clearer

## Risks

### High Risk

- **Compilation failure**: The Title function issue will prevent the code from compiling

### Medium Risk

- **Off-by-one errors**: The loop condition and separator logic could have edge cases

### Low Risk

- **Capitalization bugs**: If the Title replacement doesn't properly capitalize the first letter
- **Map lookup failures**: If the map is missing an entry (current plan covers 1-10 correctly)

## Specific Recommendations

### Required Changes

1. **Replace Title dependency**: Add this to section 2 of the plan:

```go
func title(s string) string {
    if s == "" {
        return s
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
```

Or copy the `Title` and `isSeparator` functions from the test file into the production code (they're provided as reference implementations).

2. **Update the plan's section 2**: Change from "Use the `Title()` function defined in the test file directly" to "Implement a `title()` helper function in `bottle_song.go` that capitalizes the first letter of a string."

### Optional Improvements

1. **Add input validation**: Document assumptions (e.g., "startBottles and takeDown are always valid integers 1-10")
2. **Clarify separator logic**: Add a comment explaining when empty strings are inserted
3. **Add edge case documentation**: Explicitly note that "no" is used for zero bottles, not "zero"

## Verdict

**Status: REQUIRES REVISION**

The plan's overall approach is sound and mirrors the reference solution correctly. However, the critical Title function dependency error must be fixed before implementation. With that single change, the implementation should pass all tests.

**Confidence Level**: High - The reference solution is available and correct, and the test cases are comprehensive. Once the Title function issue is resolved, success is highly probable.

## Action Items

1. Fix the Title function dependency in section 2 of the plan
2. Provide explicit code for the title helper function
3. Clarify that the production code must be self-contained (no test file dependencies)
4. Optional: Add more detail about the "no" vs "zero" edge case handling
