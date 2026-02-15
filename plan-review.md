# Plan Review: Bowling Scoring (Issue #47)

## Review Method

Detailed comparison of plan against all 25 test cases in `cases_test.go`, the reference solution in `.meta/example.go`, and the test harness in `bowling_test.go`.

## Verdict: PASS (with minor notes)

### 1. Data Structure — PASS

The `[21]int` rolls array with `nRolls`, `nFrames`, `rFrameStart` tracking fields correctly matches the reference solution and handles all test cases including perfect game (12 rolls), 10th-frame spares (21 rolls), and 10th-frame strikes with bonus rolls.

### 2. Error Variables — PASS

All four error variable names match the reference solution. Tests only check `err == nil` vs `err != nil`, so error message strings are not asserted upon, but matching the reference is good practice.

### 3. Validation Logic — PASS

All 10 roll test cases are handled:
- Negative roll: basic validation catches it
- Roll > 10: basic validation catches it
- Two rolls exceeding 10 in frame: frame completion logic catches it
- 10th frame bonus roll validation: all 5 test cases covered
- Game-over detection: all 3 test cases covered

### 4. Score Calculation — PASS

All 16 score test cases verified:
- All scoring formulas (strike, spare, open) correctly handle look-ahead across frames
- 10th frame scoring works because bonus rolls are stored contiguously in the rolls array
- Perfect game (300) correctly calculated
- Premature score (5 test cases) correctly returns error

### 5. Notes for Implementation

1. **Use constants** (like the reference): `pinsPerFrame=10`, `framesPerGame=10`, `maxRollsPerFrame=2`, `maxRollsLastFrame=3`, `maxRolls=21` instead of magic numbers
2. **Check order**: Reference checks `pins > 10` before `pins < 0`. Either order works for tests.
3. **Record-then-validate**: The roll should be recorded in the array BEFORE frame-completion validation, because helpers like `rawFrameScore()` and `strikeBonus()` read from the array
4. **10th frame 2-roll logic**: After 2 rolls, if `rawFrameScore > 10`, it's only valid if the first roll was a strike (pins reset). Otherwise error.
5. **Defensive guard**: Add check at 3 rolls in 10th frame for the case where neither strike nor spare occurred (though normally prevented by completing the frame at 2 rolls when score < 10)

### Summary

The plan is correct and complete. All test cases are covered. The minor notes above are implementation details that should be addressed during coding, but the architectural approach is sound. Approved for implementation.
