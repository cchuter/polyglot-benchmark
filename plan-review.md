# Plan Review: Bowling Scorer

## Overall Assessment

The plan is solid and should produce a correct, passing implementation. It correctly identifies the single file to modify, describes a data structure and algorithm that match the reference solution, and enumerates the major validation rules. Below is a detailed analysis across four dimensions.

---

## 1. Correctness

**Verdict: Likely correct, with one area that needs more precision.**

The plan's data structure (`Game` struct with `rolls [21]int`, `nRolls`, `nFrames`, `rFrameStart`) exactly mirrors the reference solution. The Score() logic described (walk frames, check strike/spare/open, advance `frameStart` by 1 for strikes and 2 otherwise) is correct.

### Issue: 10th Frame Validation Is Under-Specified

The plan says:

> Handle 10th frame special cases:
> - After 2 rolls: if score < 10, complete frame (no bonus rolls)
> - After 3 rolls: validate bonus roll constraints, complete frame

This is too vague for the most error-prone part of the implementation. The reference solution has several distinct validation branches at the 3-roll point in the 10th frame:

1. **First roll is a strike, second is NOT a strike**: the sum of the 2nd and 3rd rolls must not exceed 10 (test: "the second bonus rolls after a strike in the last frame cannot be a strike if the first one is not a strike" -- previousRolls `[..., 10, 6]`, roll `10` is invalid; and "two bonus rolls after a strike in the last frame cannot score more than 10 points" -- previousRolls `[..., 10, 5]`, roll `6` is invalid).

2. **First roll is a strike, second IS a strike**: the 3rd roll can be up to 10 (but not 11+). The bonus sum can be up to 20, but only if at least one of the bonus rolls is itself a strike (test: "two bonus rolls after a strike in the last frame can score more than 10 points if one is a strike" -- previousRolls `[..., 10, 10, 6]` is valid with score 26).

3. **First roll is NOT a strike but IS a spare**: the 3rd roll is a bonus roll with no special constraint beyond 0-10 (test: "a strike with the one roll bonus after a spare in the last frame does not get a bonus" -- previousRolls `[..., 7, 3, 10]` is valid).

4. **First two rolls are neither strike nor spare**: attempting a 3rd roll should be treated as "game already over" because the frame was already completed at the 2-roll point.

The plan mentions "validate bonus roll constraints" but does not enumerate these branches. An implementer working solely from the plan (without the reference solution) could easily miss case 1 or case 2's subtlety, producing incorrect validation. The plan says it follows the reference solution "closely," which mitigates this in practice, but as a standalone specification, this section is insufficient.

**Recommendation**: Expand the 10th frame validation section to explicitly enumerate the four cases above, including the specific conditions and the error to return for each.

### Validation Order

The plan lists validation in this order:
1. pins < 0
2. pins > 10
3. game over check

The reference solution checks `pins > 10` before `pins < 0`. This ordering difference does not affect correctness because a single pin value cannot be both negative and greater than 10 simultaneously. The tests only check that *an error* is returned, not which specific error. So this is fine, but worth noting the divergence from the reference.

---

## 2. Completeness

**Verdict: Complete for passing all tests, with minor gaps in specification detail.**

### Test Coverage Analysis

I verified each test case against the plan:

**Score test cases (20 cases):**
- All zeros, no strikes/spares, spares, strikes, consecutive variants, 10th frame bonuses, perfect game -- all covered by the described Score() walk-through logic.
- Incomplete/unstarted game errors -- covered by the `nFrames != 10` check returning `ErrPrematureScore`.
- Bonus rolls not yet complete (strike in last frame with 1 or 2 rolls, spare with 0 bonus rolls) -- covered because `nFrames` won't reach 10 until the 10th frame is fully completed.

**Roll test cases (12 cases):**
- Negative roll: covered by step 1.
- Roll > 10: covered by step 2.
- Two rolls in frame > 10: covered by step 6 (validate sum <= 10).
- Bonus roll after strike in last frame > 10: covered, but plan is vague on mechanism.
- Two bonus rolls after strike exceeding 10: covered by 10th frame logic (but under-specified).
- Second bonus roll is strike when first is not: this subtle case IS handled by the reference but the plan does not call it out.
- Second bonus roll after strike-strike exceeding 10: covered by pins > 10 check.
- Cannot roll after game over (3 variants): covered by the `nFrames == 10` check.

All 32 test cases should be covered by an implementation following this plan plus the reference solution.

### Missing Details

- The plan does not mention the helper methods (`rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `spareBonus`, `strikeBonus`). While these are implementation details, they are architecturally significant since the entire Roll() and Score() logic depends on them. Listing them would improve the plan's completeness.
- The plan does not specify that `completeTheFrame()` should both increment `nFrames` AND update `rFrameStart` to the current `nRolls`. This is a critical detail -- if someone only increments `nFrames` without advancing `rFrameStart`, the frame tracking breaks entirely.
- The plan does not mention that `isSpare` checks the *sum* of two rolls equaling 10 (not that the second roll equals 10 minus the first, which is equivalent but could be misunderstood).

---

## 3. Clarity

**Verdict: Good overall structure, but the critical section (10th frame) is the least clear.**

### Strengths

- The plan clearly identifies the single file to modify.
- The data structure is presented as concrete Go code, which is unambiguous.
- The constants are defined with their derivation shown.
- The key design decisions section explains *why* certain choices were made (roll-based storage, frame tracking during Roll, etc.), which is helpful for implementation.
- The Score() logic is clearly described as a frame walk.
- The implementation order (write everything in one file, then test) is appropriate.

### Weaknesses

- **Roll() step numbering is misleading**: Steps 5 and 6 say "handle strike in frames 1-9" and "handle 2nd roll in frames 1-9," but this implies a simple if/else. In reality, the reference solution uses `rollsThisFrame()` as the primary dispatch mechanism combined with whether we're in the last frame, making the flow more nuanced.
- **Step 7 is too compressed**: "Handle 10th frame special cases" with two sub-bullets does not convey the 4+ distinct code paths in the reference solution. This is the hardest part of the exercise and deserves the most detail.
- **Error variable names are specified but not their string values**: The plan shows `ErrNegativeRollIsInvalid`, `ErrPinCountExceedsPinsOnTheLane`, `ErrPrematureScore`, `ErrCannotRollAfterGameOver` but does not give their message strings. While the tests don't check error messages, it's good practice to specify them. (The reference uses specific strings like "Negative roll is invalid", etc.)

---

## 4. Risks

**Verdict: Low risk overall, with the 10th frame validation as the primary concern.**

### Risk 1: 10th Frame Validation Complexity (Medium)

As noted above, this is the most complex part. The plan acknowledges this in its own Risk Assessment section ("The 10th frame validation is the most complex part -- multiple edge cases around when bonus rolls can exceed the normal pin limit"). However, acknowledging the risk without providing the detailed specification to mitigate it is only partially helpful.

**Specific failure scenario**: If the implementer does not handle the case where the first bonus roll after a 10th-frame strike is NOT a strike (meaning the second bonus roll + first bonus roll must sum to at most 10), the test "the second bonus rolls after a strike in the last frame cannot be a strike if the first one is not a strike" will fail. The plan does not explicitly describe this constraint.

### Risk 2: Frame Completion Tracking (Low)

The plan correctly identifies that frames are completed during Roll(), not reconstructed during Score(). However, the plan's description of `rFrameStart` as "index of the first roll in the current frame" could be misunderstood. It is actually the index *after the last completed frame's last roll*, which happens to equal the first roll index of the next frame. The nuance matters during 10th frame processing where the frame start, combined with `nRolls`, determines how many rolls have been taken in the current frame.

### Risk 3: Off-by-One in Frame Counting (Low)

The plan uses 0-indexed frame counting (frames 0-9), and correctly notes that `framesPerGame - 1` is the last frame index. This matches the reference solution. Low risk but a common source of bugs in bowling implementations.

### Risk 4: Array Bounds (Very Low)

The `[21]int` array is correctly sized. With at most 21 rolls (9 frames x 2 + 3 for 10th), there is no risk of array overflow. The `Score()` function's lookahead (`rolls[f+1]`, `rolls[f+2]`) is safe because bonus rolls for the last frame are stored in the array.

---

## Summary

| Dimension    | Rating     | Notes |
|-------------|-----------|-------|
| Correctness  | Good      | Will produce passing code if 10th frame details are handled per reference |
| Completeness | Good      | All test cases addressed; helper methods and frame-completion mechanics under-documented |
| Clarity      | Adequate  | Clean structure but the hardest part (10th frame) has the least detail |
| Risks        | Low       | Primary risk is 10th frame validation edge cases |

### Recommended Changes Before Implementation

1. **Expand 10th frame validation** in the Roll() Logic section to enumerate all four cases (strike+strike, strike+non-strike, spare, open) with their specific constraints.
2. **Document the helper methods** -- at minimum list their signatures and what they compute.
3. **Clarify `completeTheFrame()`** must both increment `nFrames` and set `rFrameStart = nRolls`.
4. **Add the 2-roll validation branch** for the 10th frame more explicitly: when in the 10th frame with `rollsThisFrame() == 2`, if the first roll was a strike (pins == 10), the raw frame score can exceed 10, so the normal "sum > 10" validation must be skipped in that specific case.

### Verdict

**Approve with minor revisions.** The plan is fundamentally sound and closely follows a known-correct reference solution. The primary gap is insufficient detail in the 10th frame validation logic. If the implementer has access to the reference solution (which the plan explicitly says to use as the basis), the plan is sufficient. If implementing without reference access, the 10th frame section needs expansion.
