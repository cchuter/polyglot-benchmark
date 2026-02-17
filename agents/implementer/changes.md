# Implementer Changes

## kindergarten-garden exercise solution

- **File modified**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`
- **Commit**: `de2c7e5` — "Implement kindergarten-garden exercise solution"

### What was implemented

- `Garden` type as a `map[string][]string` mapping child names to their plant lists
- `Plants(child string) ([]string, bool)` method to look up a child's plants
- `NewGarden(diagram string, children []string) (*Garden, error)` constructor that:
  - Parses a two-row newline-delimited diagram
  - Validates row count, row length parity, matching row lengths, and cup-to-child ratio
  - Sorts children alphabetically and assigns plants left-to-right
  - Detects duplicate child names
  - Maps plant codes (G, C, R, V) to full names (grass, clover, radishes, violets)
  - Returns errors for invalid diagrams or unknown plant codes
