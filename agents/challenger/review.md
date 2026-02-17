# Challenger Review: kindergarten-garden

## Test Results
All tests pass (13/13 subtests in TestGarden, plus TestNamesNotModified and TestTwoGardens).

## Comparison with Reference Solution
The implementation closely matches `.meta/example.go`. The only difference is an additional validation check at line 24-26:

```go
if len(rows[1])%2 != 0 {
    return nil, errors.New("each diagram row must have even number of cups")
}
```

This check is technically redundant since `len(rows[1]) != 2*len(children)` (line 27) would also catch odd-length rows (since `2*len(children)` is always even). However, it provides a more specific error message for the "odd number of cups" test case and is harmless.

## Correctness Analysis

### Core Logic Verified
- **Garden type**: `map[string][]string` - correct.
- **Plants method**: Dereferences pointer and does map lookup - correct.
- **NewGarden**: Properly validates diagram format, row lengths, cup counts, and duplicate names.
- **Children not modified**: Uses `append([]string{}, children...)` to copy before sorting - passes `TestNamesNotModified`.
- **Self-contained gardens**: No package-level state - passes `TestTwoGardens`.
- **Cup iteration hack**: `for cx := range rows[1:]` iterates 0, 1 because `rows[1:]` has exactly 2 elements. Unconventional but correct and matches reference.

### Edge Cases Handled
- Wrong diagram format (not starting with newline)
- Mismatched row lengths
- Odd number of cups (extra check not in reference)
- Duplicate child names (detected via map size vs slice size)
- Invalid plant codes (lowercase or other chars)
- Looking up non-existent child (returns ok=false)

## Issues Found
None. The implementation is correct, complete, and functionally equivalent to the reference solution.

## Verdict: PASS
