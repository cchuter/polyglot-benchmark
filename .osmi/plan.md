# Implementation Plan: food-chain

## File to Modify

- `go/exercises/practice/food-chain/food_chain.go`

## Approach

The reference solution in `.meta/example.go` shows the idiomatic approach. We'll implement the same algorithmic pattern:

### Data Structure

Define a slice of structs containing each animal's name and its associated comment line. Use a `const` for the spider's "wriggled and jiggled and tickled inside her" phrase since it appears in both the spider's comment and the cumulative chain.

```go
var verse = []struct{ eaten, comment string }{
    {"", ""},  // index 0 unused
    {"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
    {"spider", "It wriggled and jiggled and tickled inside her.\n"},
    {"bird", "How absurd to swallow a bird!\n"},
    {"cat", "Imagine that, to swallow a cat!\n"},
    {"dog", "What a hog, to swallow a dog!\n"},
    {"goat", "Just opened her throat and swallowed a goat!\n"},
    {"cow", "I don't know how she swallowed a cow!\n"},
    {"horse", "She's dead, of course!"},
}
```

### `Verse(v int)` Logic

1. Start with `"I know an old lady who swallowed a {animal}.\n"`
2. Add the verse-specific comment
3. Special cases: verse 1 (fly) and verse 8 (horse) return immediately (no cumulative chain)
4. For verses 2-7: loop from `v` down to 2, building `"She swallowed the X to catch the Y"` lines
5. When reaching spider (v==3 in chain), append "that wriggled and jiggled and tickled inside her"
6. End with the fly refrain

### `Verses(start, end int)` Logic

Call `Verse()` for each verse in range, join with `"\n\n"`.

### `Song()` Logic

Return `Verses(1, 8)`.

## Order of Changes

1. Write the complete `food_chain.go` implementation
2. Run tests to verify correctness
