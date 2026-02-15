# Context Summary: bowling (Issue #47)

## Key Decisions
- Used roll-array architecture with frame tracking (matches reference solution)
- Record-then-validate pattern for Roll() method
- 10th frame handled as special case with up to 3 rolls and contextual pin validation
- Score calculation walks through frames using frameStart index with strike/spare look-ahead

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — Full implementation of Game struct with NewGame, Roll, Score methods

## Test Results
- 31/31 tests pass (`go test -v -count=1`)
  - 10 Roll validation tests
  - 21 Score calculation tests
- No go vet warnings
- Clean build

## Commit
- `aa2fb8d` - "feat: implement bowling game scorer"
- Branch: `issue-47`
- Pushed to origin
