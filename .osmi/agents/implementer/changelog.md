# Implementer Changelog

## bottle-song: Implement Recite function

- Implemented `Recite(startBottles, takeDown int) []string` in `bottle_song.go`
- Added `numberToWord` map for int-to-word conversion (0-10)
- Added `bottleStr` helper for singular/plural "bottle"/"bottles"
- Added `verse` helper to generate a single 4-line verse
- Uses `Title` from the test file for capitalizing the first line of each verse
- All 7 tests pass
