# Context: Kindergarten Garden Implementation

## Status: COMPLETE

## Branch: issue-131 (pushed to origin)

## What Was Built

Implemented the Kindergarten Garden exercise in Go. Given a garden diagram (two rows of plant codes) and a list of children's names, the solution determines which 4 plants each child is responsible for.

## Key Implementation Details

- **File**: `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`
- **Type**: `Garden` is `map[string][]string`
- **Constructor**: `NewGarden(diagram, children)` validates input, sorts children alphabetically, maps plant codes to names
- **Lookup**: `Plants(child)` does a map lookup returning `([]string, bool)`

## Validation Handled

- Diagram format (must start with newline, exactly 2 rows)
- Mismatched row lengths
- Row length vs number of children
- Duplicate child names
- Invalid plant codes (only G/C/R/V allowed)

## Test Results

All 14 tests and 2 benchmarks pass.
