# Context Summary

## Issue
#323 — polyglot-go-bottle-song: Implement Ten Green Bottles song

## Status
Complete. All tests pass.

## Files Modified
- `go/exercises/practice/bottle-song/bottle_song.go` — Added `Recite` function with helpers

## Branch
`issue-323` — pushed to origin

## Key Details
- Function signature: `Recite(startBottles int, takeDown int) []string`
- Number words stored in a slice indexed 0-10
- Singular "bottle" when count == 1, "bottles" otherwise
- Verses separated by empty string elements in the returned slice
- First two lines of each verse use capitalized number word; last line lowercase
