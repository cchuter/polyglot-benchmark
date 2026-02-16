# Implementation Plan: Beer Song (Issue #82)

## Overview

Implement three exported functions (`Verse`, `Verses`, `Song`) in `go/exercises/practice/beer-song/beer_song.go` to generate the lyrics for "99 Bottles of Beer on the Wall".

## File to Modify

**`go/exercises/practice/beer-song/beer_song.go`** — Replace the stub (currently only `package beer`) with the full implementation.

No other files need to be created or modified.

## Architectural Decisions

1. **Single file implementation** — All three functions in one file, following the pattern of other exercises in this repo.
2. **Switch-based verse generation** — Use a `switch` statement in `Verse()` with cases for n=0, n=1, n=2, and default (n>=3). This balances readability with correctness.
3. **`bytes.Buffer` for multi-verse concatenation** — Use `bytes.Buffer` in `Verses()` for efficient string building.
4. **`fmt` for formatting and errors** — Use `fmt.Sprintf` for the generic verse template and `fmt.Errorf` for error messages.

## Implementation Details

### Imports

```go
import (
    "bytes"
    "fmt"
)
```

### `Verse(n int) (string, error)`

- Validate: `n < 0 || n > 99` → return error
- Switch on n:
  - `n == 0`: Return the "No more bottles" verse (hardcoded string)
  - `n == 1`: Return the "1 bottle" verse (singular, "Take it down", "no more")
  - `n == 2`: Return the "2 bottles" verse (next bottle is singular "1 bottle")
  - `default` (n >= 3): Use `fmt.Sprintf` with `n` and `n-1`

### `Verses(start, stop int) (string, error)`

- Validate start range (0-99), stop range (0-99), start >= stop
- Loop from start down to stop, calling `Verse(i)` for each
- Append each verse to a `bytes.Buffer`, followed by `"\n"` (blank line separator)
- Return `buff.String()`

### `Song() string`

- Call `Verses(99, 0)`, ignore error (guaranteed valid), return result

## Ordering of Changes

1. Write the complete implementation in `beer_song.go`
2. Run `go test` to verify all tests pass
3. Run `go test -bench=.` to verify benchmarks work
4. Commit the change

## Risks and Mitigations

- **String format mismatch**: The tests compare exact strings. Mitigation: carefully match the expected format from test constants and the reference implementation.
- **Trailing newline handling**: Each verse ends with `\n`, and verses are separated by an additional `\n` (blank line). The `Verses` function adds `\n` after each verse including the last, producing a trailing blank line. This matches the test expectations (test constants like `verses86` end with `\n\n`).
