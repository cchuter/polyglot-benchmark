# Context Summary

## Issue
#345: polyglot-go-kindergarten-garden — Implement the kindergarten garden exercise in Go.

## Solution
Single file implementation in `go/exercises/practice/kindergarten-garden/kindergarten_garden.go`:
- `Garden` type as `map[string][]string`
- `NewGarden(diagram, children)` — parses diagram, validates, returns `*Garden`
- `Plants(child)` — map lookup returning `([]string, bool)`

## Branch
`issue-345` — pushed to origin

## Tests
All pass: `go test ./...` in exercise directory

## Files Modified
- `go/exercises/practice/kindergarten-garden/kindergarten_garden.go` (56 lines, was 1 line stub)
