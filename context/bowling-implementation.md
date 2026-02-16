# Context: Bowling Implementation

## Key Decisions

1. **Array-based storage**: Rolls stored in flat `[23]int` array rather than frame-based structure. Simplifies bonus calculation by indexing forward.
2. **Frame completion tracking**: `nFrames` and `rFrameStart` track game progress for both validation and score calculation.
3. **Matched reference implementation**: Solution closely follows `.meta/example.go` pattern, which is proven against the test suite.
4. **10th frame special handling**: Separate logic branches in `Roll()` for maxRollsPerFrame (2) and maxRollsLastFrame (3) roll counts.

## Files Modified

| File | Change |
|------|--------|
| `go/exercises/practice/bowling/bowling.go` | Full implementation replacing stub |

## Test Results

- 31/31 tests pass
- Build compiles cleanly
- No new dependencies

## Branch

- Feature branch: `issue-77`
- Pushed to origin
- Commit: `d3b34d7 feat: implement bowling game scorer for issue #77`

## Acceptance Criteria Status

All criteria from GOAL.md met:
- 16 score tests pass
- 5 incomplete game error tests pass
- 10 roll validation tests pass
- Code compiles
- No new dependencies
