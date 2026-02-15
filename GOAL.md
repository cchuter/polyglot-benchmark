# Goal: Implement Go Bowling Exercise

## Problem Statement

Implement a bowling game scorer in Go. The implementation must define a `Game` type with two operations:
- `Roll(pins int) error` — called each time the player rolls a ball, with the number of pins knocked down
- `Score() (int, error)` — called at the end of the game to return the total score

The implementation must correctly handle all bowling scoring rules including open frames, spares, strikes, and the special 10th frame with bonus rolls.

## Acceptance Criteria

All of the following must pass:

### Scoring Tests (scoreTestCases)
1. Score a game with all zeros → 0
2. Score a game with no strikes or spares → 90
3. A spare followed by zeros → 10
4. Points after a spare counted twice → 16
5. Consecutive spares each get one roll bonus → 31
6. Spare in last frame gets one bonus roll counted once → 17
7. A strike earns 10 points in a single-roll frame → 10
8. Points after a strike counted twice as bonus → 26
9. Consecutive strikes each get two roll bonus → 81
10. Strike in last frame gets two bonus rolls counted once → 18
11. Rolling a spare with two roll bonus does not get bonus roll → 20
12. Strikes with two roll bonus do not get bonus rolls → 30
13. Last two strikes followed by non-strike bonus → 31
14. Strike with one roll bonus after spare in last frame → 20
15. Perfect game (all strikes) → 300
16. Two bonus rolls after last-frame strike can score >10 if one is strike → 26
17. Unstarted game cannot be scored → error
18. Incomplete game cannot be scored → error
19. Bonus rolls for last-frame strike must be rolled before scoring → error
20. Both bonus rolls for last-frame strike must be rolled → error
21. Bonus roll for last-frame spare must be rolled → error

### Roll Validation Tests (rollTestCases)
1. Rolls cannot score negative points → error
2. A roll cannot score more than 10 points → error
3. Two rolls in a frame cannot score more than 10 → error
4. Bonus roll after last-frame strike cannot exceed 10 → error
5. Two bonus rolls after last-frame strike cannot exceed 10 → error
6. Second bonus roll cannot be strike if first isn't strike → error
7. Second bonus roll after last-frame strike cannot exceed 10 → error
8. Cannot roll if game already has ten frames → error
9. Cannot roll after bonus roll for spare → error
10. Cannot roll after bonus rolls for strike → error

### Technical Criteria
- All tests in `bowling_test.go` and `cases_test.go` pass with `go test`
- Package name is `bowling`
- Module is `bowling` with `go 1.18`
- Implementation file is `bowling.go`
- `Game` type is defined as a struct
- Code follows Go conventions (exported types, error handling)

## Key Constraints
- Must not modify test files (`bowling_test.go`, `cases_test.go`)
- Must not modify `go.mod`
- Must implement the exact API: `NewGame() *Game`, `(g *Game) Roll(pins int) error`, `(g *Game) Score() (int, error)`
