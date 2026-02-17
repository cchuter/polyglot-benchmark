# Plan Review (Codex)

## Overall Assessment: GOOD — proceed with minor clarifications

### Critical Issue Resolved: Board Width After `prepare()`

The reviewer flagged variable-width boards as a HIGH concern. **This is a non-issue.** The `prepare()` function in `connect_test.go` strips ALL spaces (including indentation). After stripping:
- `. O . X .` → `.O.X.` (5 chars)
- ` . X X O .` → `.XXO.` (5 chars)
- `  O O O X .` → `OOOX.` (5 chars)

All rows have equal width after space removal. The reference implementation correctly uses `len(lines[0])` as the fixed width.

### Issues to Address in Implementation

1. **Method receivers**: All board methods use value receivers (including `markConnected` and `evaluate`). This works because Go slices are reference types — the slice header is copied but the underlying array is shared, so mutations to `fields[y][x]` through a value receiver still affect the original data.

2. **Return values**: Return `"X"` and `"O"` (not "black"/"white"). The comment in the reference is misleading but the code is correct.

3. **Coordinate system**: x = column (left-to-right), y = row (top-to-bottom). Standard matrix indexing.

4. **1x1 boards**: A single cell is simultaneously start and target for both players. The DFS correctly handles this — the cell is both the start coord and satisfies `isTargetCoord`.

5. **Input validation**: Only validate `len(lines) < 1` and `len(lines[0]) < 1`. No need for character validation.

### Recommendations Incorporated

- Document that value receivers work for mutation because slices share underlying arrays
- Explicitly use "X"/"O" return values
- No changes needed to the plan's architecture — it's sound

### Verdict: Plan is ready for implementation.
