# Goal: Implement Kindergarten Garden (Go)

## Problem Statement

Implement the `kindergarten-garden` exercise in Go. Given a diagram of a kindergarten garden (two rows of plant codes), determine which plants each child is responsible for. Children are assigned cups alphabetically by name, with each child getting two cups per row (four cups total).

## What Needs to Be Built

The stub file `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` currently only declares the package. It needs a complete implementation with:

1. A `Garden` type (map from child name to plant list)
2. A `NewGarden(diagram string, children []string) (*Garden, error)` constructor that:
   - Parses a two-row diagram (prefixed by newline) into plant assignments
   - Sorts children alphabetically to assign cups left-to-right
   - Does NOT modify the input `children` slice (makes a copy before sorting)
   - Validates the diagram format and returns errors for invalid input
3. A `Plants(child string) ([]string, bool)` method on `*Garden` that looks up a child's plants

## Acceptance Criteria

1. **All tests pass**: `go test ./...` in the exercise directory passes all tests in `kindergarten_garden_test.go`
2. **No vet issues**: `go vet ./...` produces no warnings
3. **Correct plant mapping**: G=grass, C=clover, R=radishes, V=violets
4. **Alphabetical assignment**: Children are sorted alphabetically; cups assigned left-to-right
5. **No side effects**: The input `children` slice must not be modified (TestNamesNotModified)
6. **Self-contained**: Two separate Garden instances with different children must work independently (TestTwoGardens)
7. **Error handling**: Returns errors for:
   - Wrong diagram format (not starting with newline, not exactly 2 rows)
   - Mismatched row lengths
   - Odd number of cups (not 2 per child)
   - Duplicate child names
   - Invalid plant codes (anything other than G, C, R, V)

## Key Constraints

- Package name must be `kindergarten`
- Must use Go 1.18+ (as specified in go.mod)
- The `Garden` type is used as a pointer receiver (`*Garden`) in tests
- Plant names are lowercase full words: "grass", "clover", "radishes", "violets"
