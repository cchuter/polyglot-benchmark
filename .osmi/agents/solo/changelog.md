# Solo Agent Change Log

## Change 1: Implement Pig Latin Sentence function

**File**: `go/exercises/practice/pig-latin/pig_latin.go`
**Commit**: `972f088`

**What changed**:
- Added `Sentence(sentence string) string` — splits input on spaces, translates each word, joins result
- Added `translateWord(word string) string` — applies Pig Latin rules in priority order
- Added `isVowel(b byte) bool` — helper to check vowel characters

**Why**: Stub file only had package declaration. Implemented full translation logic per the four Pig Latin rules.

**Test results**: 22/22 tests pass. `go vet` clean.
