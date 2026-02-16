# Goal: Kindergarten Garden (Go)

## Problem Statement

Implement the Kindergarten Garden exercise in Go. Given a diagram of a kindergarten garden (two rows of plant codes) and a list of children's names, determine which plants each child is responsible for.

Each child gets 4 cups (2 per row), assigned alphabetically by name from left to right. Plant codes are single uppercase letters: G (grass), C (clover), R (radishes), V (violets).

## Acceptance Criteria

1. `NewGarden(diagram string, children []string) (*Garden, error)` constructs a Garden from a diagram and list of children names
2. `(*Garden).Plants(child string) ([]string, bool)` returns the list of plant names for a given child, or `ok=false` if the child is not in the garden
3. Children are assigned cups alphabetically (sorted by name), but the original `children` slice must NOT be modified
4. Garden objects are self-contained (no package-level state shared between instances)
5. Error cases handled:
   - Diagram must start with `\n` (3 parts when split on `\n`, first part empty)
   - Rows must be equal length
   - Row length must equal `2 * number of children`
   - No duplicate children names
   - Only valid plant codes (G, C, R, V)
6. All tests pass: `TestGarden`, `TestNamesNotModified`, `TestTwoGardens`, benchmarks

## Key Constraints

- Package name: `kindergarten`
- Solution file: `kindergarten_garden.go`
- Must pass all tests in `kindergarten_garden_test.go` without modification
- Go 1.18 module
