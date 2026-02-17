# Goal: Kindergarten Garden (Go)

## Problem Statement

Implement the Kindergarten Garden exercise in Go. Given a diagram of a kindergarten class garden and a list of children's names, determine which plants each child is responsible for.

The garden has two rows of cups. Each child gets 4 cups (2 per row), assigned alphabetically by name. Plants are encoded as single uppercase letters: G (grass), C (clover), R (radishes), V (violets).

## Acceptance Criteria

1. **`NewGarden(diagram string, children []string) (*Garden, error)`** must:
   - Parse a two-row diagram (prefixed by a newline) into a Garden
   - Sort children alphabetically to assign cups (without modifying the input slice)
   - Return an error for invalid diagrams:
     - Diagram not starting with a newline (wrong format)
     - Mismatched row lengths
     - Odd number of cups (rows not divisible by 2 per child)
     - Duplicate child names
     - Invalid plant codes (anything other than G, C, R, V)

2. **`(*Garden).Plants(child string) ([]string, bool)`** must:
   - Return the 4 plants belonging to the named child as full plant names
   - Return `(plants, true)` for valid children
   - Return `(nil, false)` for children not in the garden

3. **Garden instances must be self-contained** (no package-level state)

4. **All tests must pass**: `go test` in the exercise directory must succeed, including:
   - TestGarden (11 test cases: valid gardens, error cases)
   - TestNamesNotModified (input slice must not be mutated)
   - TestTwoGardens (two independent Garden instances)
   - BenchmarkNewGarden and BenchmarkGarden_Plants

## Key Constraints

- Package name: `kindergarten`
- Module: `kindergarten` (Go 1.18)
- The `Garden` type must be a pointer receiver type (tests use `*Garden`)
- Must not modify the `children` slice passed to `NewGarden`
- Plant code mapping: G→"grass", C→"clover", R→"radishes", V→"violets"
