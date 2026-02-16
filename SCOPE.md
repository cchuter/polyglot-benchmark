# Scope: Kindergarten Garden (Go)

## In Scope

- Implement `Garden` type in `kindergarten_garden.go`
- Implement `NewGarden(diagram string, children []string) (*Garden, error)` constructor
- Implement `(*Garden).Plants(child string) ([]string, bool)` lookup method
- Handle all validation/error cases tested in the test file
- Ensure all tests pass including `TestGarden`, `TestNamesNotModified`, `TestTwoGardens`, and benchmarks

## Out of Scope

- Modifying the test file (`kindergarten_garden_test.go`)
- Modifying `go.mod`
- Adding additional files or packages
- Changes to any other exercise
- Performance optimization beyond what is needed to pass benchmarks

## Dependencies

- Go standard library only (`errors`, `sort`, `strings`)
- No external packages
