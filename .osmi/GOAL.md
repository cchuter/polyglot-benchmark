# Goal: Implement Kindergarten Garden Exercise (Go)

## Problem Statement

Implement the `kindergarten-garden` exercise in Go. Given a diagram of a kindergarten garden (two rows of plant codes), determine which plants each child is responsible for. Children are assigned cups alphabetically, with each child getting two cups per row (four total).

## Required API

The test file expects:

1. `NewGarden(diagram string, children []string) (*Garden, error)` — Constructor that parses a diagram and associates plants with children (sorted alphabetically). Must validate input and return errors for invalid diagrams.

2. `(*Garden).Plants(child string) ([]string, bool)` — Lookup method that returns a child's four plants as full lowercase names, and a boolean indicating whether the child was found.

## Acceptance Criteria

1. **Single student gardens**: Correctly parse diagrams with one student.
2. **Multi-student gardens**: Correctly assign plants to multiple students based on alphabetical ordering.
3. **Full 12-student garden**: Handle the full class of 12 children with correct plant assignments.
4. **Custom/out-of-order names**: Sort children alphabetically without modifying the input slice.
5. **Two independent gardens**: Garden instances must be self-contained (no package-level state).
6. **Invalid name lookup**: Return `ok=false` for children not in the garden.
7. **Error: wrong diagram format**: Reject diagrams not starting with a newline (3 parts when split).
8. **Error: mismatched rows**: Reject diagrams where row lengths differ.
9. **Error: odd cups**: Reject diagrams where row length is not `2 * len(children)`.
10. **Error: duplicate names**: Reject duplicate child names.
11. **Error: invalid plant codes**: Reject codes other than G, C, R, V.
12. **All tests pass**: `go test ./...` and `go vet ./...` pass cleanly.

## Plant Code Mapping

- G → "grass"
- C → "clover"
- R → "radishes"
- V → "violets"

## Key Constraints

- The `Garden` type must be a pointer receiver (tests use `*Garden`).
- Must not modify the input `children` slice (tested explicitly).
- Package name must be `kindergarten`.
- Must work with Go 1.18 (per go.mod).
