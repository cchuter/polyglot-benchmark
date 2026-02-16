# Plan Review: Branch 2 (BFS with Visited Set)

**Reviewer**: codex
**Verdict**: APPROVE with minor suggestions

---

## 1. Correctness

### Hex Neighbor Directions: CORRECT

The plan uses `point{row, col}` and defines six neighbor deltas as:
```go
{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, 1}
```

Mapping these to the reference solution's `coord{x, y}` (where x=col, y=row):
| Plan (row, col) | Reference (x, y) | Direction |
|---|---|---|
| `{0, 1}` | `{1, 0}` | right |
| `{0, -1}` | `{-1, 0}` | left |
| `{1, 0}` | `{0, 1}` | down |
| `{-1, 0}` | `{0, -1}` | up |
| `{1, -1}` | `{-1, 1}` | down-left |
| `{-1, 1}` | `{1, -1}` | up-right |

All six directions match the reference solution exactly. The hex adjacency model is correct.

### Player Win Conditions: CORRECT

- X seeds from left column (`col == 0`) and wins upon reaching right column (`col == cols-1`). This is left-to-right. Correct.
- O seeds from top row (`row == 0`) and wins upon reaching bottom row (`row == rows-1`). This is top-to-bottom. Correct.

### BFS Algorithm: CORRECT

The BFS implementation is standard and correct:
- Seed cells are enqueued and marked visited before the loop.
- The target-edge check happens at dequeue time, which means seeds that are already on the target edge (e.g., 1x1 board) will be detected immediately.
- Neighbors are only enqueued if they match the player's stone and have not been visited.

### Test Case Walk-Through

All 10 test cases should pass:

1. **Empty 5x5 board** (all dots): No X or O stones on starting edges. Both BFS queues start empty. Returns `""`. PASS.
2. **1x1 "X"**: X seeds `{0,0}`, immediately `col == 0 == cols-1`. Returns `"X"`. PASS.
3. **1x1 "O"**: X finds no seeds. O seeds `{0,0}`, immediately `row == 0 == rows-1`. Returns `"O"`. PASS.
4. **Only edges, no winner**: X stones form a border but do not connect left to right through hex adjacency. O stones likewise do not connect top to bottom. Returns `""`. PASS.
5. **Illegal diagonal**: The X stones form a path that uses diagonal adjacency that is not valid in hex. BFS with correct hex neighbors will not traverse it. Returns `""`. PASS.
6. **Adjacent angles, nobody wins**: Neither player has a connected path. Returns `""`. PASS.
7. **X wins crossing left to right**: X has a valid hex-connected path from col 0 to col 3. Returns `"X"`. PASS.
8. **O wins crossing top to bottom**: O has a valid hex-connected path from row 0 to row 4. Returns `"O"`. PASS.
9. **X wins convoluted path**: X has a winding path from left to right. BFS will find it. Returns `"X"`. PASS.
10. **X wins spiral path (9x9)**: X has a spiral path from left to right. BFS will find it. Returns `"X"`. PASS.

---

## 2. Edge Cases

| Edge Case | Handled? | Notes |
|---|---|---|
| Empty board (`len(lines) == 0`) | Yes | Returns error. No test case triggers this, but it is harmless. |
| Empty first line (`len(lines[0]) == 0`) | Yes | Returns error. No test case triggers this either, but harmless. |
| 1x1 board | Yes | Both X and O cases work because the seed cell is simultaneously on the start and target edge. |
| Non-square boards | Yes | `rows` and `cols` are derived independently. The 4x4 "only edges" test exercises this relative to the 5x5 default. |
| Rows with differing lengths | Not explicitly handled | The plan assumes all rows have the same length via `cols := len(lines[0])`. If a later row is shorter, `lines[nb.row][nb.col]` could panic with an index-out-of-bounds. However, this matches the reference solution's assumption, and the test data always has uniform row lengths after space stripping. This is acceptable. |
| Both X and O win simultaneously | Handled by design | X is checked first, so X would be returned. This matches the reference solution's behavior and no test case has this scenario. |

---

## 3. API Alignment

### Function Signature: CORRECT

```go
func ResultOf(lines []string) (string, error)
```

This matches exactly what the tests expect:
```go
actual, err := ResultOf(prepare(tc.board))
```

### Package: CORRECT

The plan specifies `package connect`, which matches the test files.

### Return Values: CORRECT

- `"X"` for X wins, `"O"` for O wins, `""` for no winner, all with `nil` error.
- Error returns for invalid input (not tested but harmless).

---

## 4. Potential Bugs

### No bugs found in the proposed implementation.

Specific checks performed:

1. **Off-by-one in bounds checking**: `nr >= 0 && nr < rows && nc >= 0 && nc < cols` is correct. No off-by-one.
2. **Off-by-one in target edge check**: `cur.col == cols-1` and `cur.row == rows-1` are correct.
3. **Visited before enqueue**: The plan correctly marks cells as visited when they are enqueued (not when dequeued). This prevents duplicate enqueuing and is the standard correct BFS pattern.
4. **Seed cell on target edge**: Checked at dequeue time, so a 1x1 board (where the seed is also on the target edge) is handled correctly.
5. **Byte comparison**: `lines[r][0] == player` where `player` is `byte` and `lines[r][0]` indexes a string to get a `byte`. This is correct in Go for ASCII characters.

### One minor robustness note (not a bug):

The `import "errors"` is included for the error-returning guard clauses. If the implementer decides to remove the guard clauses (since no test expects errors), the import should be removed too, or the Go compiler will reject it as an unused import. The plan as written is self-consistent, so this is not a problem.

---

## 5. Code Quality

### Strengths

- **Clean and idiomatic Go**: The code uses standard Go idioms -- struct types, slice-based queue, map for visited set, byte indexing into strings.
- **Good separation of concerns**: `neighbors()` is a standalone pure function. `hasWon()` encapsulates all traversal logic. `ResultOf()` is a thin orchestrator.
- **Reasonable naming**: `point`, `row`, `col`, `visited`, `queue`, `cur`, `nb` are all clear and conventional.
- **Estimated ~100 lines**: Significantly simpler than the ~170-line reference solution. Lower complexity means fewer places for bugs.

### Suggestions for improvement (non-blocking)

1. **Pre-allocate the visited map**: For large boards, `make(map[point]bool, rows*cols)` would reduce map resizing. This is a minor performance improvement and not necessary for correctness.

2. **Consider using a slice-based visited set instead of a map**: A `[][]bool` indexed by `[row][col]` would be faster than `map[point]bool` due to avoiding hashing. Example:
   ```go
   visited := make([][]bool, rows)
   for i := range visited {
       visited[i] = make([]bool, cols)
   }
   ```
   This is optional and the map-based approach is perfectly fine for the test case sizes.

3. **The `errors` import and guard clauses are unnecessary for passing the tests**: The test file explicitly states "We don't expect errors for any of the test cases." However, keeping them is good defensive programming and does not hurt. If kept, consider whether the guard clauses could be hit by accident (they cannot with the provided test data).

4. **Comment on hex adjacency model**: The inline comment `// Hex grid adjacency: right, left, down, up, down-left, up-right` is good. Consider adding a brief note about why exactly these six directions (and not e.g. down-right) to help future readers unfamiliar with hex grids on parallelogram boards.

---

## Summary

The selected plan (Branch 2: BFS with Visited Set) is **correct, complete, and ready for implementation**. The hex neighbor directions are verified against the reference solution. The BFS algorithm handles all test cases including the 1x1 edge case. The function signature and return values match the test expectations exactly. No bugs or logic errors were found. The code is clean, idiomatic, and simpler than the reference solution.

**Recommendation**: Proceed with implementation as planned.
