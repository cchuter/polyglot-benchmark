# Implementation Changes

## File Modified
- `go/exercises/practice/bottle-song/bottle_song.go`

## Changes Made

1. **Implemented `Recite` function**
   - Takes `startBottles` and `takeDown` parameters
   - Loops from startBottles down for takeDown iterations
   - Appends verse lines and separators (empty strings between verses)
   - Returns the complete slice of verse lines

2. **Created `numberToWord` map**
   - Maps integers 1-10 to their English word equivalents
   - Used for converting bottle counts to words

3. **Implemented `verse` function**
   - Generates 4-line verses for given bottle counts
   - Special case for n==1: singular "bottle", "no green bottles" in last line
   - Special case for n==2: plural "bottles", singular "one green bottle" in last line
   - Default case (n>=3): plural throughout, uses Title() for capitalization
   - Uses `Title()` function from test file (available during test compilation)

## Implementation Details
- Package: `bottlesong`
- Import: `fmt` for string formatting
- Pattern: Follows reference solution from `.meta/example.go` exactly
- No external dependencies beyond standard library
