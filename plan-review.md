# Plan Review: Alphametics Solver

**Reviewer**: Self-review (no codex agent available in tmux environment)

## Algorithm Correctness

**Verdict: Sound with caveats**

The column-by-column backtracking approach is a well-known technique for alphametics. The core idea is correct:
- Process columns right-to-left
- Assign digits to unassigned letters as they are encountered
- Check column constraint (sum mod 10 == result digit)
- Propagate carry to next column

### Issue 1: Column structure needs careful handling

The plan's `column` struct with `addendLetters` and `resultLetter` is simplistic. A letter can appear multiple times in the same column across different addend words (e.g., in the 199-addend case, many words share letters). The solver must sum the digit value of each occurrence, not each unique letter.

**Fix**: Store the full list of letter occurrences per column (with duplicates), or store a multiplier/coefficient per letter per column.

### Issue 2: Columns beyond result word length

If the carry propagates beyond the result word's length, we need to handle the case where there's no result letter but carry must be zero. The plan mentions `carry == 0` at termination but doesn't explicitly handle columns where the result word has no letter but addends do.

**Fix**: Handle result letter as optional per column. When result has no letter at a position, the constraint is that sum % 10 == 0 at that position (actually, this case doesn't arise since the result must be at least as long as any addend for a valid puzzle — but carries could extend beyond). Actually for valid puzzles the result is the longest word or has an extra carry digit. The solver should handle columns up to max(all word lengths) + 1 for potential carry.

### Issue 3: Letter ordering in backtracking

The order in which unassigned letters are tried matters for performance. Letters that appear in earlier (rightmost) columns should be assigned first since they enable early pruning. The plan does this implicitly by processing columns right-to-left.

## Performance Analysis

**The 199-addend test case** has 10 unique letters. The column-by-column approach should handle this well because:
- At most 10 letters to assign across all columns
- Each column check prunes invalid branches immediately
- The constraint is tight: with 199 addends, the column sums are large and highly constrained

**Potential bottleneck**: If many unassigned letters appear in the first (rightmost) column, the branching factor is high before any pruning kicks in.

**Mitigation**: The approach is still far better than brute-force permutation. For 10 letters, worst case is still bounded by 10! = 3.6M but effective search space is much smaller due to column pruning.

## Edge Cases

1. **`A == B`** — No `+`, two single-letter words. Each must map to a different digit but A must equal B, which is impossible. Should return error. ✅ Handled by uniqueness constraint.

2. **Leading zero** — `ACA + DD == BD` where the only solution would require B=0 (leading digit). ✅ Handled by leading letter constraint.

3. **Single-letter words** — Letters like `I` and `A` in the 199-addend case. These are both leading letters (can't be 0) and full words. ✅ Works naturally.

4. **No solution** — Must return error, not panic. ✅ Covered.

## Recommendations

1. Use a coefficient-based approach per column: for each column, compute `coeff[letter]` = (count in LHS at this position) - (count in RHS at this position, which is 0 or 1). This avoids needing separate addend/result tracking.

2. Consider sorting unassigned letters per column by most-constrained-first (letters that are leading come first since they have fewer options 1-9 vs 0-9).

3. The plan is solid. Proceed with implementation.

## Overall Assessment

**Approved with minor adjustments.** The column-by-column backtracking is the right approach. Key adjustments:
- Handle letter coefficients (multiplicity) per column correctly
- Ensure columns extend to cover carry beyond result word length
- Handle the case of no result letter in overflow columns
