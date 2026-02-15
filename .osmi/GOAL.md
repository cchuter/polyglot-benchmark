# Goal: polyglot-go-beer-song (Issue #15)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" exercise in Go as part of the polyglot benchmark suite. The exercise requires generating the lyrics to the classic song with correct handling of singular/plural forms and special verses.

The implementation already exists at `go/exercises/practice/beer-song/` with all tests passing. This issue validates the implementation meets acceptance criteria and closes issue #15 on a proper feature branch.

## Acceptance Criteria

1. **`Verse(n int) (string, error)`** — Returns a single verse for bottle count `n` (0-99)
   - Verse 0: "No more bottles..." / "Go to the store and buy some more, 99 bottles..."
   - Verse 1: "1 bottle..." (singular) / "Take it down..." / "no more bottles..."
   - Verse 2: "2 bottles..." / "Take one down..." / "1 bottle..." (singular next)
   - Verses 3-99: Standard plural form with decrementing count
   - Invalid input (n < 0 or n > 99): returns error

2. **`Verses(start, stop int) (string, error)`** — Returns verses from `start` down to `stop`
   - Each verse separated by a blank line
   - Validates start and stop are in range [0, 99]
   - Validates start >= stop
   - Returns error for invalid inputs

3. **`Song() string`** — Returns the complete song (all verses from 99 down to 0)

4. **All existing tests pass** (`go test -v ./...` in `go/exercises/practice/beer-song/`)

5. **Package name:** `beer`

6. **Module:** `beer` with `go 1.18`

## Key Constraints

- Must follow existing exercism exercise conventions in the repository
- No external dependencies
- Must handle edge cases: singular "bottle" vs plural "bottles", "Take it down" vs "Take one down", "No more" vs number
