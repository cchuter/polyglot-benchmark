# Changes

## kindergarten_garden.go

- Implemented `Garden` type as `map[string][]string`
- Implemented `Plants` method to look up a child's plants by name
- Implemented `NewGarden` constructor that:
  - Parses a two-row diagram (newline-separated, leading newline)
  - Validates row count, row length equality, and cups-per-child count
  - Sorts children alphabetically and assigns plants left-to-right
  - Detects duplicate child names
  - Maps plant codes (G, C, R, V) to full names (grass, clover, radishes, violets)
  - Returns error for invalid plant codes
