# Goal: Kindergarten Garden (Go)

## Problem Statement

Implement a Go solution for the Kindergarten Garden exercise. Given a diagram of a kindergarten garden (two rows of plant codes) and a list of children's names, determine which plants each child is responsible for.

Each child gets 4 cups (2 per row), assigned alphabetically by name from left to right. Plant codes are: G (grass), C (clover), R (radishes), V (violets).

## Acceptance Criteria

1. `NewGarden(diagram string, children []string) (*Garden, error)` constructs a Garden from a diagram and list of children names
2. `(*Garden).Plants(child string) ([]string, bool)` returns the 4 plants for a given child, or `ok=false` if the child is not in the garden
3. Children are assigned cups in alphabetical order, regardless of input order
4. The input `children` slice must NOT be modified (no in-place sorting)
5. Multiple Garden instances must be independent (no shared package-level state)
6. Error cases return non-nil error:
   - Diagram does not start with newline (wrong format)
   - Rows have different lengths (mismatched rows)
   - Odd number of cups in a row (not divisible by 2 per child)
   - Duplicate child names
   - Invalid plant codes (anything other than G, C, R, V)
7. All tests in `kindergarten_garden_test.go` pass
8. Benchmarks run without error

## Key Constraints

- Package name: `kindergarten`
- File: `kindergarten_garden.go`
- Go 1.18 module
- Must match the API expected by the test file: `NewGarden` returns `*Garden`, `Plants` is a pointer receiver method
- `Garden` type must support map-like semantics for child lookups
