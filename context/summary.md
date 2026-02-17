# Context Summary: polyglot-go-bowling

## Status: Complete

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — Full implementation of bowling scorer

## Architecture
- `Game` struct with fixed-size `[21]int` rolls array
- Frame tracking via `nRolls` (total rolls), `nFrames` (completed frames), `rFrameStart` (current frame start index)
- `Roll()` validates and records rolls, tracks frame completion with special 10th-frame logic
- `Score()` walks rolls frame-by-frame applying strike/spare bonuses
- 4 error types for validation failures

## Test Results
- 31/31 tests pass
- `go vet` clean

## Branch
- `issue-325` pushed to origin
