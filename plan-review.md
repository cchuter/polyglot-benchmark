# Plan Review

## Reviewer: Self-review (codex agent unavailable)

## Overall Assessment: **Approved with minor clarifications**

The plan is sound and covers the key aspects of bowling scoring correctly.

## Strengths

1. **Roll-based storage with frame tracking** is the right approach - it cleanly separates validation from scoring
2. The Score() walk-through algorithm is correct: walk frame-by-frame using an index, advancing by 1 for strikes and 2 for non-strikes
3. All error cases from the test suite are accounted for

## Issues Found

### Issue 1: 10th Frame Validation Needs Precise Logic (Medium)

The plan's description of 10th frame validation is slightly imprecise. The exact rules for the 10th frame bonus rolls:

- After strike on first roll: second roll can be 0-10
- After strike+strike: third roll can be 0-10
- After strike+non-strike: third roll must satisfy `second + third <= 10`
- After spare (no first-roll strike): third roll can be 0-10

This needs to be implemented carefully with explicit conditionals.

### Issue 2: Frame Tracking State (Low)

The plan mentions `rollInFrame` but the actual implementation should carefully track when we're in the 10th frame vs frames 0-8 since the advancement logic differs. Using `frame` and `rollInFrame` together is correct.

### Issue 3: Game Completion Detection (Low)

The `done` flag needs to be set at the right points:
- Frames 0-8: after 2 rolls or after a strike → advance frame
- Frame 9: after 2 rolls if open; after 3 rolls if strike or spare

## Recommendation

Proceed with implementation. The core algorithm is correct. Pay special attention to the 10th frame validation conditionals during implementation.
