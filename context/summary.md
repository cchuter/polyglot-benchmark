# Context Summary: Issue #111 - polyglot-go-bowling

## Status: COMPLETE

## Key Decisions
- Used flat roll array `[21]int` approach for tracking rolls
- Followed reference solution pattern from `.meta/example.go`
- Single file modification: `go/exercises/practice/bowling/bowling.go`
- No external dependencies; only uses `errors` package

## Files Modified
- `go/exercises/practice/bowling/bowling.go` — full bowling scorer implementation

## Test Results
- All 31 tests pass (10 Roll + 21 Score)
- Build successful with Go 1.18

## Branch
- Feature branch: `issue-111`
- Pushed to origin
- Base: `bench/polyglot-go-bowling`

## Architecture
- `Game` struct: rolls array, roll count, frame count, frame start index
- `Roll()`: validates pin count, tracks frames 1-9 (strike/spare/open) and frame 10 (bonus rolls)
- `Score()`: walks frame by frame, adding bonuses for strikes (+next 2 rolls) and spares (+next 1 roll)
