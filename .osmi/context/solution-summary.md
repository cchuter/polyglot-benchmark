# Context: DnD Character Solution

## Files Modified

- `go/exercises/practice/dnd-character/dnd_character.go` — complete solution

## Test Results

- All tests pass: `go test ./...` → PASS
- `go vet ./...` → clean
- 16 modifier tests, 1 ability test (1000 iterations), 1 character test (1000 iterations)

## Key Decisions

1. **Manual min tracking** over `slices.Min` for go 1.18 compat
2. **`math.Floor`** for modifier to handle negative scores correctly

## Branch

- Feature branch: `issue-162`
- Commit: implements full solution, closes #162
