# Plan Review (Self-Review — no Codex agent available)

## Correctness Assessment

The weighted-sum approach is mathematically sound:
- For `SEND + MORE == MONEY`: S*1000 + E*100 + N*10 + D + M*1000 + O*100 + R*10 + E - (M*10000 + O*1000 + N*100 + E*10 + Y) = 0
- This correctly reduces the problem to: assign unique digits to letters such that the weighted sum equals zero, subject to leading-zero constraints.
- Edge case: `A == B` correctly requires two different digits with weight A: +1, B: -1, and since they must be different, the sum can never be zero → error. ✓
- Edge case: `ACA + DD == BD` → leading zeros on multi-digit words. If A leads ACA and D leads DD and B leads BD — need to check which letters are leading. ✓

## Performance Assessment

- 10 unique letters → worst case 10! = 3,628,800 permutations
- With early pruning (sorting by weight magnitude), most branches will be cut early
- The 199-addend test has 10 letters but the weights will be very large, making pruning very effective
- **Risk**: If pruning is insufficient, the 199-addend test might be slow. Mitigation: sorting letters by descending weight magnitude ensures the largest-impact assignments are tried first, cutting invalid branches quickly.

## Potential Issues

1. **Parsing edge cases**: Need to handle extra whitespace carefully. Use `strings.TrimSpace`.
2. **Letter extraction for leading zeros**: Must identify the first letter of each word that has length > 1 (single-letter words CAN be zero unless other constraints apply). Wait — actually looking at the test: `A == B` expects error because unique values, not because of leading zeros. Single-letter words should be allowed to be zero. Only multi-digit words have the leading-zero constraint. ✓
3. **Weight computation**: The `E` in `SEND + MORE == MONEY` appears in multiple positions — weights accumulate correctly via addition. ✓

## Verdict

The plan is sound. Proceed with implementation.
