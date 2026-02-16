# Context: Kindergarten Garden Implementation

## Status: Complete

## Solution file
`go/exercises/practice/kindergarten-garden/kindergarten_garden.go`

## Approach
- `Garden` is `map[string][]string` (child name → 4 plant names)
- `NewGarden` validates diagram format, row equality, cup count, duplicates, plant codes
- Children sorted alphabetically via a copy (input not mutated)
- `Plants` is a pointer-receiver map lookup returning `([]string, bool)`

## Test results
All 15 tests pass: 13 in TestGarden, TestNamesNotModified, TestTwoGardens. `go vet` clean.

## Branch
`issue-174` pushed to origin
