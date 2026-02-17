# Context Summary

## Issue
#333 — polyglot-go-dnd-character: Implement D&D character generator in Go.

## Key Decisions
- Used `sort.Ints` instead of `slices.Min` for Go 1.18 compatibility
- Used `math.Floor(float64(score-10) / 2.0)` for modifier to handle negative odd values correctly
- Solution in single file: `go/exercises/practice/dnd-character/dnd_character.go`

## Files Modified
- `go/exercises/practice/dnd-character/dnd_character.go` — full implementation

## Test Results
- All tests pass (`go test ./...`)
- `go vet` clean

## Branch
- `issue-333` pushed to origin
