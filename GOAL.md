# Goal: polyglot-go-food-chain

## Problem Statement

Implement the Go solution for the "Food Chain" exercise: generate the lyrics of the cumulative song "I Know an Old Lady Who Swallowed a Fly" algorithmically.

The solution must implement three exported functions in `go/exercises/practice/food-chain/food_chain.go`:

- `Verse(v int) string` — returns the lyrics for a single verse (1-8)
- `Verses(start, end int) string` — returns verses from `start` to `end`, joined by double newlines
- `Song() string` — returns the complete song (all 8 verses)

## Acceptance Criteria

1. `Verse(v)` returns the correct text for each verse 1 through 8
2. `Verses(1, 3)` returns verses 1-3 joined by `"\n\n"`
3. `Song()` returns the full song (verses 1-8 joined by `"\n\n"`)
4. All tests in `food_chain_test.go` pass (`go test ./...`)
5. Code passes `go vet ./...`
6. Solution is algorithmic (not hardcoded lyrics)

## Key Constraints

- Package name must be `foodchain`
- Must match exact expected text from test file (punctuation, whitespace, newlines)
- The spider verse has special "that wriggled and jiggled and tickled inside her" chain text
- Verse 1 (fly) and verse 8 (horse) are special cases with no cumulative chain
