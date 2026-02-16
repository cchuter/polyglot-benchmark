# Changes: kindergarten-garden

## Summary
Implemented the kindergarten garden exercise in Go.

## Files Changed
- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` — wrote full solution

## Implementation Details
- **Type**: `Garden map[string][]string` — maps child names to their plant lists
- **Plants method**: Simple map lookup returning `([]string, bool)`
- **NewGarden function**: Constructs a Garden from a two-row diagram string and a list of children
  - Splits diagram on `\n`, validates exactly 3 parts with empty first part
  - Validates row lengths match and equal `2 * len(children)`
  - Copies and sorts children alphabetically (does not mutate input)
  - Detects duplicate names via map length check
  - Decodes plant codes: G→grass, C→clover, R→radishes, V→violets
  - Returns descriptive errors for all invalid inputs

## Test Results
All 14 tests pass (13 subtests in TestGarden + TestNamesNotModified + TestTwoGardens).
