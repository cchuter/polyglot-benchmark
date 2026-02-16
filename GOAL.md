# Goal: Implement Kindergarten Garden Exercise (Go)

## Problem Statement

Implement a Go solution for the Kindergarten Garden exercise. Given a two-row diagram of plant codes and a list of children's names, determine which four plants each child is responsible for. Children are assigned cups alphabetically (sorted by name), with each child getting two cups per row (four total).

## Acceptance Criteria

1. **`NewGarden(diagram string, children []string) (*Garden, error)`** — Constructs a Garden from a diagram and list of children.
   - Diagram format: newline + row1 + newline + row2 (e.g., `"\nVRCG\nVVRC"`)
   - Children are sorted alphabetically; each child maps to 2 consecutive cups per row
   - Plant codes: G=grass, C=clover, R=radishes, V=violets
   - Must NOT modify the input `children` slice (must copy before sorting)
   - Must return errors for:
     - Wrong diagram format (not exactly 3 parts when split on newline, or first part non-empty)
     - Mismatched row lengths
     - Odd number of cups (row length != 2 * number of children)
     - Duplicate child names
     - Invalid plant codes (anything other than G, C, R, V)

2. **`(*Garden).Plants(child string) ([]string, bool)`** — Looks up a child's plants.
   - Returns the child's 4 plants as full names: `["violets", "radishes", "violets", "radishes"]`
   - Returns `(nil, false)` if the child doesn't exist in the garden

3. **Multiple Garden instances are independent** (no package-level state).

4. **All tests pass**: `go test` in the exercise directory must succeed, including `TestGarden`, `TestNamesNotModified`, `TestTwoGardens`, and both benchmarks.

## Key Constraints

- Package name: `kindergarten`
- File: `kindergarten_garden.go`
- Go 1.18 module
- No external dependencies
