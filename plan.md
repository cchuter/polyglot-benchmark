# Implementation Plan: polyglot-go-beer-song

## Overview

Single file modification: `go/exercises/practice/beer-song/beer_song.go`

The implementation follows the reference example pattern from `.meta/example.go` with three exported functions using standard library only (`fmt`, `bytes`).

## File to Modify

- `go/exercises/practice/beer-song/beer_song.go` — Replace stub with full implementation

## Imports Required

```go
import (
    "bytes"
    "fmt"
)
```

## Implementation Details

### Function 1: `Verse(n int) (string, error)`

Returns a single verse. Uses a switch statement to handle five cases:

1. **Invalid** (n < 0 or n > 99): Return error `fmt.Errorf("%d is not a valid verse", n)`
2. **n == 0**: Fixed string — "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
3. **n == 1**: Singular "bottle", "Take it down", "no more bottles"
4. **n == 2**: Plural "bottles" on first line, singular "1 bottle" on second line
5. **n >= 3** (default): `fmt.Sprintf` with n and n-1, both plural

Each verse string ends with `\n`.

### Function 2: `Verses(start, stop int) (string, error)`

Returns a range of verses from `start` down to `stop` (inclusive). Uses `bytes.Buffer` for efficient string concatenation.

Validation (in order):
1. start outside 0-99: error with `"start value[%d] is not a valid verse"`
2. stop outside 0-99: error with `"stop value[%d] is not a valid verse"`
3. start < stop: error with `"start value[%d] is less than stop value[%d]"`

Loop from start down to stop, calling `Verse(i)` for each. **After every verse (including the last), append an extra `\n`**. This means each verse is followed by a blank line, and the returned string ends with `\n\n`. This matches the test expectations in `verses86` and `verses75`.

### Function 3: `Song() string`

Signature: `func Song() string` (single return value, no error).
Internally calls `Verses(99, 0)` and discards the error with `_`.

## Architectural Decisions

- **Use `bytes.Buffer`** for `Verses()` string concatenation (efficient for many verses)
- **Use `switch` statement** in `Verse()` for clarity on special cases
- **Match error message formats** to reference solution (tests only check nil/non-nil, but consistency is good practice)

## Order of Changes

1. Write the complete `beer_song.go` with all three functions and imports
2. Run tests to verify
3. Commit
