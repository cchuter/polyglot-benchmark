# Implementation Plan: Beer Song

## File to Modify

`go/exercises/practice/beer-song/beer_song.go`

## Approach

Implement the solution based on the reference example in `.meta/example.go`, which uses a switch/case pattern in `Verse()` to handle the grammatical edge cases. This is the canonical Exercism approach for this exercise.

### Function Implementations

#### 1. `Verse(n int) (string, error)`

Uses a switch statement with 4 cases:
- **n < 0 or n > 99**: Return error
- **n == 0**: "No more bottles..." / "Go to the store..."
- **n == 1**: "1 bottle..." / "Take it down..." / "no more bottles..."
- **n == 2**: "2 bottles..." / "Take one down..." / "1 bottle..."
- **default (3–99)**: "N bottles..." / "Take one down..." / "N-1 bottles..."

#### 2. `Verses(start, stop int) (string, error)`

- Validate: start and stop both in [0, 99], and start >= stop
- Loop from start down to stop, calling `Verse(i)` for each
- Separate verses with a blank line (`\n`)
- Use `bytes.Buffer` for efficient string building

#### 3. `Song() string`

- Simply calls `Verses(99, 0)` and returns the result

### Imports Required

- `bytes` — for efficient string concatenation in `Verses()`
- `fmt` — for `Sprintf` and `Errorf`

## Order of Changes

1. Replace the stub `beer_song.go` with the full implementation
2. Run tests to verify all pass

## Rationale

The reference solution in `.meta/example.go` is clean, idiomatic Go, and directly matches the test expectations. There's no need to deviate from this proven approach.
