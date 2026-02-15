# Changelog — implementer

## 2026-02-15

### Verified alphametics implementation

- Verified existing implementation at `go/exercises/practice/alphametics/alphametics.go`
- `go vet ./...` passed with no issues
- `gofmt -l` passed — code is properly formatted
- No code changes needed; implementation already complete and correct
- Implementation uses permutation-based brute-force solver with leading-zero validation

## 2026-02-14

### Implemented bottle-song Recite function

- Replaced stub in `go/exercises/practice/bottle-song/bottle_song.go` with full implementation
- Added `numberToWord` map (1–10) for integer-to-English-word conversion
- Added private `verse(n int) []string` helper with three cases:
  - `n == 1`: singular bottle, "no green bottles" ending
  - `n == 2`: plural bottles, singular "one green bottle" ending
  - default: plural bottles using `fmt.Sprintf` and `Title()` for capitalization
- Implemented `Recite(startBottles, takeDown int) []string` with verse iteration and blank-line separators
- All 7 test cases pass
