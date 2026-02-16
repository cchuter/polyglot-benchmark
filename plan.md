# Implementation Plan: Kindergarten Garden

## File to Modify

- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

## Design

### Type Definition

```go
type Garden map[string][]string
```

`Garden` is a map from child name to their list of 4 plant names. Using a map type alias keeps it simple and makes `Plants` a straightforward map lookup.

### `NewGarden(diagram string, children []string) (*Garden, error)`

1. Split diagram on `\n` — must produce exactly 3 parts with `parts[0] == ""`
2. Validate rows have equal length
3. Copy children slice (to avoid mutating the input), sort the copy alphabetically
4. Validate row length equals `2 * len(sortedChildren)`
5. Initialize a `Garden` map, inserting each child with an empty slice (capacity 4)
6. Check for duplicates: if `len(garden) != len(sortedChildren)`, there were duplicates
7. Iterate over both rows, for each child at index `nx`, extract their 2 cups per row (positions `2*nx` and `2*nx+1`)
8. Map each cup code to its plant name (G→grass, C→clover, R→radishes, V→violets); return error for invalid codes
9. Return `&garden, nil`

### `(*Garden).Plants(child string) ([]string, bool)`

Simple map lookup: `p, ok := (*g)[child]; return p, ok`

## Approach

This follows the reference solution in `.meta/example.go` closely, which is the canonical approach for this exercise. The implementation is straightforward — a single file change replacing the stub.

## Ordering

1. Write the full implementation in `kindergarten_garden.go`
2. Run tests to verify
3. Commit
