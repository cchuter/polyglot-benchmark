# Implementation Context: Kindergarten Garden

## Key Decisions
- Used `map[string][]string` as the Garden type (simplest approach)
- Pre-compute all plant assignments during construction (no lazy evaluation)
- Copy children slice before sorting to avoid side effects

## Files Modified
- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` — Complete implementation (57 lines)

## Test Results
- 15/15 tests pass
- `go vet` clean
- Branch: `issue-260`, pushed to origin

## Architecture Notes
- `for cx := range rows[1:]` is a hack to iterate 0,1 (two cups per child per row)
- Duplicate detection uses map length vs slice length comparison
- Validation order: format → row match → cup count → duplicates → plant codes
