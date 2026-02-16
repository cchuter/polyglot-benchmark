# Changes

## Implemented `Recite` function in `bottle_song.go`

- Added `Recite(startBottles, takeDown int) []string` function that generates verses of the "Ten Green Bottles" song
- Added `numberToWord` map for converting integers 1-10 to English words
- Added `verse(n int) []string` helper with three cases: singular (n==1), two bottles (n==2), and generic plural (n>=3)
- All 7 tests pass
