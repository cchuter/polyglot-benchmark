# Goal: polyglot-go-food-chain (Issue #125)

## Problem Statement

Implement the Go solution for the "Food Chain" exercise: generate the lyrics of the cumulative song "I Know an Old Lady Who Swallowed a Fly" algorithmically.

The solution must implement three functions in `go/exercises/practice/food-chain/food_chain.go`:

1. `Verse(v int) string` — Returns a single verse (1-8)
2. `Verses(start, end int) string` — Returns a range of verses joined by blank lines
3. `Song() string` — Returns the entire song (all 8 verses)

## Acceptance Criteria

1. `Verse(v)` returns the correct lyrics for each verse 1 through 8
2. `Verses(1, 3)` returns verses 1-3 joined by `"\n\n"`
3. `Song()` returns all 8 verses joined by `"\n\n"`
4. All tests in `food_chain_test.go` pass: `TestVerse`, `TestVerses`, `TestSong`, `BenchmarkSong`
5. Code compiles with `go build ./...`
6. Solution is algorithmic (not hardcoded lyrics)

## Key Constraints

- Package name must be `foodchain`
- File must be `food_chain.go`
- Must match exact expected output from the test file (whitespace-sensitive)
- The spider verse includes "that wriggled and jiggled and tickled inside her" in the cumulative chain
- Verse 8 (horse) is a special terminal verse: "She's dead, of course!"
- Verse 1 (fly) has no exclamation comment, just the "Perhaps she'll die" refrain
