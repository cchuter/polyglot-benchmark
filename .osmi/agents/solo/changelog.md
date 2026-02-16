# Solo Agent Change Log

## Change 1: Implement kindergarten garden solution

**File modified:** `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

**What was done:**
- Defined `Garden` type as `map[string][]string`
- Implemented `NewGarden` constructor with full validation:
  - Diagram format (must have 2 rows preceded by empty line)
  - Row length equality
  - Cup count matches children count (2 cups per child per row)
  - Duplicate name detection via map length comparison
  - Invalid plant code detection (only G, C, R, V allowed)
- Implemented `Plants` method as pointer-receiver map lookup
- Children are sorted alphabetically via a copy (input not mutated)

**Test results:** All 15 tests pass, `go vet` clean
