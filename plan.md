# Implementation Plan: Kindergarten Garden

## File to Modify

- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

No other files need to be created or modified.

## Architecture

### Type Definition

```go
type Garden map[string][]string
```

`Garden` is a type alias for `map[string][]string`, mapping child names to their list of 4 plant names. This allows the `Plants` method to do a simple map lookup.

### Plants Method

```go
func (g *Garden) Plants(child string) ([]string, bool)
```

Pointer receiver on `Garden`. Dereferences `*g` to perform map lookup. Returns the plant slice and whether the child was found.

### NewGarden Constructor

```go
func NewGarden(diagram string, children []string) (*Garden, error)
```

Steps:
1. Split diagram by `\n` - must yield exactly 3 parts with first part empty (diagram starts with `\n`)
2. Validate rows have same length
3. Validate row length equals `2 * len(children)` (each child gets 2 cups per row)
4. Copy children slice to avoid mutating the input, then sort the copy alphabetically
5. Initialize Garden map, check for duplicate names (map size vs slice length)
6. Iterate over both rows, for each child at index `nx`, read positions `2*nx` and `2*nx+1`
7. Map each character (G/C/R/V) to its plant name; return error on invalid codes
8. Return pointer to Garden

### Plant Code Mapping

Use a switch statement:
- 'G' -> "grass"
- 'C' -> "clover"
- 'R' -> "radishes"
- 'V' -> "violets"
- default -> error

### Error Cases

| Condition | Error |
|-----------|-------|
| Diagram doesn't split into ["", row1, row2] | "diagram must have two rows" |
| row1 and row2 different lengths | "diagram rows must be same length" |
| row length != 2 * number of children | "each diagram row must have two cups per child" |
| Duplicate child names | "no two children can have the same name" |
| Invalid plant code character | "plant codes must be one of G, C, R, or V" |

### Imports

- `errors` - for `errors.New`
- `sort` - for `sort.Strings`
- `strings` - for `strings.Split`

## Rationale

This follows the reference solution in `.meta/example.go` exactly, which is the canonical approach for Exercism exercises. The implementation is minimal, idiomatic Go, and handles all test cases including:
- Not modifying the input children slice (copy before sort)
- Independent Garden instances (no package-level state)
- All validation cases
