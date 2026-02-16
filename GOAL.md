# Goal: Implement Bowling Game Scoring (Go)

## Problem Statement

Implement a bowling game scorer in Go. The solution must track rolls and compute the final score according to standard 10-pin bowling rules, including proper handling of strikes, spares, and the special 10th frame.

The stub file `go/exercises/practice/bowling/bowling.go` currently contains only a package declaration. It must be filled in with a working implementation that passes all tests in `bowling_test.go` and `cases_test.go`.

## Required API

- `NewGame() *Game` — Returns a new Game instance
- `(g *Game) Roll(pins int) error` — Records a roll; returns error for invalid rolls
- `(g *Game) Score() (int, error)` — Returns the total score; only valid at end of game

## Acceptance Criteria

1. **All zeros game**: 20 rolls of 0 → score 0
2. **No strikes/spares**: e.g., 10 frames of (3,6) → score 90
3. **Spare scoring**: Spare adds next roll as bonus (e.g., 6,4,3,0... → 16)
4. **Consecutive spares**: Each spare gets its own one-roll bonus
5. **Strike scoring**: Strike adds next two rolls as bonus
6. **Consecutive strikes**: Each strike gets its own two-roll bonus
7. **Perfect game**: 12 strikes → score 300
8. **10th frame spare**: Gets one bonus roll, counted once
9. **10th frame strike**: Gets two bonus rolls, counted once
10. **Error: negative pins**: Roll(-1) returns error
11. **Error: pins > 10**: Roll(11) returns error
12. **Error: frame > 10 pins**: Two rolls in a frame exceed 10 → error
13. **Error: game over**: Rolling after game is complete → error
14. **Error: premature score**: Score() before game is complete → error
15. **Error: incomplete bonus rolls**: Score() with missing 10th-frame bonus rolls → error
16. **10th frame bonus validation**: Bonus rolls in the 10th frame must respect pin limits

All 36 test cases (21 score tests + 15 roll tests) in `cases_test.go` must pass.

## Key Constraints

- Must use Go 1.18+ (module declares `go 1.18`)
- Package name must be `bowling`
- Only modify `bowling.go`; test files are read-only
- Must match the exact API expected by the test harness: `NewGame`, `Roll`, `Score`
