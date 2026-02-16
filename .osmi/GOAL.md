# Goal: polyglot-go-beer-song (Issue #149)

## Problem Statement

Implement the "99 Bottles of Beer on the Wall" song generator in Go. The solution must provide three functions in the `beer` package:

- `Verse(n int) (string, error)` — Returns a single verse for bottle number `n`
- `Verses(start, stop int) (string, error)` — Returns verses from `start` down to `stop`
- `Song() string` — Returns the complete song (verses 99 through 0)

## Key Rules

Not all verses are identical. Special cases:
- **Verse 0**: "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
- **Verse 1**: Uses "bottle" (singular), "Take it down" (not "Take one down")
- **Verse 2**: Next verse uses "1 bottle" (singular)
- **Verses 3-99**: Standard plural "bottles", "Take one down and pass it around"

## Acceptance Criteria

1. `Verse(n)` returns the correct verse string for any valid n (0-99)
2. `Verse(n)` returns an error for invalid n (< 0 or > 99)
3. `Verses(start, stop)` returns concatenated verses separated by blank lines
4. `Verses(start, stop)` returns errors for invalid ranges
5. `Song()` returns the complete song (equivalent to `Verses(99, 0)`)
6. All tests in `beer_song_test.go` pass
7. `go vet` reports no issues

## Constraints

- Package name must be `beer`
- Module is `beer` with Go 1.18
- Solution goes in `beer_song.go`
- Do not modify test files
