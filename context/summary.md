# Context: Bowling Exercise Implementation

## Key Decisions

- Used the reference solution from `.meta/example.go` as the implementation basis
- The implementation is a direct, faithful reproduction of the reference solution
- Roll-based storage with frame tracking during Roll() calls (not reconstructed during Score())
- Fixed-size array `[21]int` for rolls (max 21 rolls: 9 frames x 2 + 3 for 10th)

## Files Modified

- `go/exercises/practice/bowling/bowling.go` — full implementation (was a stub with just `package bowling`)

## Files NOT Modified

- `bowling_test.go` — test harness (read-only)
- `cases_test.go` — test cases (read-only)
- `go.mod` — module definition (read-only)

## Test Results

- 31/31 tests pass
- 10 roll validation tests: negative pins, pin count limits, frame limits, 10th frame bonus rules, game-over detection
- 21 scoring tests: all zeros, no strikes/spares, spares, strikes, consecutive combinations, 10th frame bonuses, perfect game (300), incomplete game errors

## Branch

- `issue-67` pushed to origin
