# Goal: Implement Bowling Game Scoring in Go

## Problem Statement

Issue #77 requires implementing a bowling game scorer in Go as part of the polyglot benchmark suite. The solution file `go/exercises/practice/bowling/bowling.go` currently contains only a package declaration stub. A complete implementation must be provided that passes all 36+ test cases defined in `bowling_test.go` and `cases_test.go`.

## What Needs to Be Built

A Go package `bowling` that provides:

1. **`Game` struct** - Tracks the state of a bowling game
2. **`NewGame() *Game`** - Constructor that returns a fresh game
3. **`Roll(pins int) error`** - Records each ball roll with validation
4. **`Score() (int, error)`** - Returns the total game score (only valid at game end)

## Bowling Rules Summary

- 10 frames per game, each frame starts with 10 pins
- **Open frame**: Less than 10 pins knocked down in two rolls; score = pins knocked down
- **Spare**: All 10 pins in two rolls; score = 10 + next 1 roll bonus
- **Strike**: All 10 pins in first roll; score = 10 + next 2 rolls bonus
- **10th frame special**: Strike or spare earns fill balls (1 or 2); no additional bonuses on fill balls
- **Perfect game** (all strikes) = 300 points

## Acceptance Criteria

1. **All score test cases pass** (16 cases in `cases_test.go`):
   - All zeros = 0
   - No strikes/spares = 90
   - Spare followed by zeros = 10
   - Points after spare counted twice = 16
   - Consecutive spares = 31
   - Spare in last frame with bonus = 17
   - Strike earns 10 points = 10
   - Points after strike counted twice = 26
   - Consecutive strikes = 81
   - Strike in last frame with bonus = 18
   - Spare with two-roll bonus = 20
   - Three strikes in last frame = 30
   - Last two strikes + bonus = 31
   - Strike after spare in last frame = 20
   - Perfect game = 300
   - Two bonus rolls after last frame strike > 10 = 26

2. **All incomplete game score cases return error** (4 cases):
   - Unstarted game cannot be scored
   - Incomplete game cannot be scored
   - Bonus rolls for last-frame strike must be completed
   - Bonus roll for last-frame spare must be completed

3. **All roll validation test cases pass** (11 cases):
   - Negative rolls return error
   - Roll > 10 returns error
   - Two rolls in frame > 10 returns error
   - Bonus roll validation in 10th frame
   - Cannot roll after game over (multiple scenarios)

4. **Code compiles** with `go build ./...`
5. **All tests pass** with `go test -v ./...`
6. **No new dependencies** beyond standard library

## Key Constraints

- Must use package name `bowling`
- Must export: `Game` struct, `NewGame`, `Roll`, `Score`
- Must work with Go 1.18+
- Must not modify test files (`bowling_test.go`, `cases_test.go`)
- Solution goes in `bowling.go` only
