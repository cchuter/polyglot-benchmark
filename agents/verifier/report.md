# Verification Report: Go Bowling Exercise

## Verdict: PASS

All acceptance criteria have been met. The implementation is correct and complete.

---

## Acceptance Criteria Checklist

### Scoring Tests (21 tests) - ALL PASS

| # | Test Case | Status |
|---|-----------|--------|
| 1 | Score a game with all zeros (0) | PASS |
| 2 | Score a game with no strikes or spares (90) | PASS |
| 3 | A spare followed by zeros (10) | PASS |
| 4 | Points after a spare counted twice (16) | PASS |
| 5 | Consecutive spares each get one roll bonus (31) | PASS |
| 6 | Spare in last frame gets one bonus roll counted once (17) | PASS |
| 7 | A strike earns 10 points in a single-roll frame (10) | PASS |
| 8 | Points after a strike counted twice as bonus (26) | PASS |
| 9 | Consecutive strikes each get two roll bonus (81) | PASS |
| 10 | Strike in last frame gets two bonus rolls counted once (18) | PASS |
| 11 | Rolling a spare with two roll bonus does not get bonus roll (20) | PASS |
| 12 | Strikes with two roll bonus do not get bonus rolls (30) | PASS |
| 13 | Last two strikes followed by non-strike bonus (31) | PASS |
| 14 | Strike with one roll bonus after spare in last frame (20) | PASS |
| 15 | Perfect game (300) | PASS |
| 16 | Two bonus rolls after last-frame strike can score >10 if one is strike (26) | PASS |
| 17 | Unstarted game cannot be scored (error) | PASS |
| 18 | Incomplete game cannot be scored (error) | PASS |
| 19 | Bonus rolls for last-frame strike must be rolled before scoring (error) | PASS |
| 20 | Both bonus rolls for last-frame strike must be rolled (error) | PASS |
| 21 | Bonus roll for last-frame spare must be rolled (error) | PASS |

### Roll Validation Tests (10 tests) - ALL PASS

| # | Test Case | Status |
|---|-----------|--------|
| 1 | Rolls cannot score negative points (error) | PASS |
| 2 | A roll cannot score more than 10 points (error) | PASS |
| 3 | Two rolls in a frame cannot score more than 10 (error) | PASS |
| 4 | Bonus roll after last-frame strike cannot exceed 10 (error) | PASS |
| 5 | Two bonus rolls after last-frame strike cannot exceed 10 (error) | PASS |
| 6 | Second bonus roll cannot be strike if first isn't strike (error) | PASS |
| 7 | Second bonus roll after last-frame strike cannot exceed 10 (error) | PASS |
| 8 | Cannot roll if game already has ten frames (error) | PASS |
| 9 | Cannot roll after bonus roll for spare (error) | PASS |
| 10 | Cannot roll after bonus rolls for strike (error) | PASS |

### Technical Criteria

| Criterion | Status |
|-----------|--------|
| All tests pass with `go test` | PASS (31/31, independently verified) |
| Package name is `bowling` | PASS |
| Module is `bowling` with `go 1.18` | PASS |
| Implementation file is `bowling.go` | PASS |
| `Game` type is defined as a struct | PASS |
| Code follows Go conventions (exported types, error handling) | PASS |
| `NewGame() *Game` constructor implemented | PASS |
| `(g *Game) Roll(pins int) error` method implemented | PASS |
| `(g *Game) Score() (int, error)` method implemented | PASS |

### File Change Constraints

| Constraint | Status |
|-----------|--------|
| Test files not modified | PASS (bowling_test.go, cases_test.go unchanged) |
| go.mod not modified | PASS (module bowling, go 1.18) |
| Only bowling.go changed (exercise files) | PASS (git diff --name-only main shows only bowling.go + .osmi metadata) |

### Challenger Review

- Verdict: NO ISSUES FOUND
- All scoring paths, validation logic, and edge cases traced and verified correct

## Independent Test Run

Tests were independently executed by the verifier with `go test -v -count=1`:
- 31/31 tests passed
- 0 failures
- No build errors
