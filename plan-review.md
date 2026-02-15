# Plan Review: Replace `strings.Title()` with `capitalize` helper

## Verdict: PASS

### 1. Correctness — PASS

`strings.Title` capitalizes the first letter of each word. The proposed `capitalize` function uppercases only the first byte and preserves the rest. For these inputs — single, lowercase, ASCII words like "three", "four", ..., "ten" — both functions produce identical output. There are no multi-word values in the `numberToWord` map.

### 2. Edge Cases — PASS

The `capitalize` function uses `s[:1]` which is a byte slice. For the values in this map ("one" through "ten"), all first characters are single-byte ASCII, so this is safe. The empty-string guard is good defensive practice.

### 3. Import Cleanliness — PASS

The `strings` import is retained for `strings.ToUpper`. The `fmt` import is still needed for `fmt.Sprintf`. No unused imports.

### 4. Other Observations

- The test file already contains its own `Title` function copy with a comment explaining the deprecation
- The `default` branch only handles `n` values 3-10 (values 1 and 2 have dedicated cases)
- All map lookups will succeed and all resulting capitalized strings match test expectations

### Summary

The plan is minimal, correct, and appropriately scoped. It eliminates the `strings.Title` deprecation warning without changing behavior. Approved.
