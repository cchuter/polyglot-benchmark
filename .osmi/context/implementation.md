# Context: Pig Latin Implementation

## Solution file
`go/exercises/practice/pig-latin/pig_latin.go`

## Approach
Simple imperative loop scanning the consonant cluster from left to right. Rules applied in priority order within the loop:
1. Rule 1 checked before the loop (vowel start, "xr", "yt" prefix)
2. Rule 3: `q` followed by `u` → include both in moved cluster
3. Rule 4: `y` at position > 0 → split before `y`
4. Rule 2: any vowel → split at vowel

## Test results
22/22 PASS, go vet clean

## Branch
`issue-273` pushed to origin
