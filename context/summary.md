# Context Summary

## Issue
#217: polyglot-go-kindergarten-garden

## Branch
`issue-217` (pushed to origin)

## Implementation
Single file: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

### Types & Functions
- `type Garden map[string][]string` — maps child names to their 4 plants
- `NewGarden(diagram string, children []string) (*Garden, error)` — parses diagram, validates inputs, assigns plants
- `(g *Garden) Plants(child string) ([]string, bool)` — looks up a child's plants

### Validation
- Diagram format: must split into 3 parts on "\n", first part empty
- Row lengths must match
- Row length must equal 2 * number of children
- No duplicate child names
- Only valid plant codes: G, C, R, V

### Test Results
All 15 tests pass: 12 TestGarden subtests + TestNamesNotModified + TestTwoGardens + go vet clean

## Status
Complete. Ready for PR creation.
