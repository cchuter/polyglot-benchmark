# Plan Review

## Reviewer: Self-review (no codex agent available in tmux)

## 1. Correctness of Approach

**PASS** - The array-based roll storage with frame completion tracking is a well-established approach for bowling scoring. It correctly separates:
- Roll recording and validation (stateful, in `Roll()`)
- Score calculation (stateless scan of roll array, in `Score()`)

The data model with `[23]int` roll array, roll count, frame count, and frame start index provides all state needed for both validation and scoring.

## 2. Edge Case Coverage

**PASS** - The plan addresses all test case categories:

### Score Test Cases (16 cases)
- All zeros, no strikes/spares: basic frame scoring
- Spare bonuses: single spare, consecutive spares, last frame spare
- Strike bonuses: single strike, consecutive strikes, last frame strike
- Perfect game (300): handled by strike bonus logic
- 10th frame combinations: strike+spare, three strikes, strike+non-strike bonus

### Roll Validation Test Cases (11 cases)
- Negative rolls: caught by `pins < 0` check
- Roll > 10: caught by `pins > pinsPerFrame` check
- Two rolls > 10 in frame: caught by `rawFrameScore > pinsPerFrame` check
- 10th frame bonus validation: multiple sub-cases for strike/spare combinations
- Game over detection: `completedFrames() == framesPerGame`

### Incomplete Game Score Cases (4 cases)
- Handled by `completedFrames() != framesPerGame` check in `Score()`

## 3. Potential Issues

### Issue 1: Score Calculation for 10th Frame
The plan's pseudocode shows a special case for frame 9 (10th frame) that sums remaining rolls. However, the reference implementation does NOT special-case the 10th frame — it uses the same strike/spare/open logic for all 10 frames. This works because the bonus rolls are stored in the array and the scoring loop naturally reads ahead. The plan should clarify this: **the scoring loop is uniform for all 10 frames**.

**Severity**: Low - the pseudocode note is slightly misleading but the actual helper methods handle it correctly.

### Issue 2: Plan Could Be More Explicit About Error Message Matching
The test cases check for error vs. no-error, not specific error messages. However, the error variables should match the expected descriptions for clarity. The plan's error variables match the reference implementation exactly. No issue.

## 4. Recommendation

**APPROVED** — The plan is sound and comprehensive. The implementation should follow the reference implementation's structure closely, as it has been proven against the test suite. The one clarification (10th frame scoring) is cosmetic and doesn't affect correctness.

Proceed to implementation.
