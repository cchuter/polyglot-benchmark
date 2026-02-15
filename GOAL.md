# Goal: Fix bottle-song Go exercise (Issue #26)

## Problem Statement

The `bottle-song` Go exercise implements the "Ten Green Bottles" children's song. The existing implementation in `go/exercises/practice/bottle-song/bottle_song.go` uses `strings.Title()` which has been deprecated since Go 1.18. The `staticcheck` linter flags this as SA1019. The implementation needs to be updated to avoid the deprecated function while keeping all tests passing.

## Acceptance Criteria

1. All 7 test cases in `bottle_song_test.go` pass (`go test -v` shows PASS for every case)
2. `staticcheck ./...` produces no warnings (no SA1019 deprecation warnings)
3. `go vet ./...` produces no warnings
4. The `Recite(startBottles, takeDown int) []string` function signature is preserved
5. The output matches the expected lyrics exactly (capitalized number words, correct singular/plural "bottle"/"bottles", correct "no green bottles" for zero)
6. No external dependencies added (go.mod stays at `go 1.18` with no require statements)

## Key Constraints

- Cannot use `golang.org/x/text/cases` (no external dependencies allowed)
- Must replace `strings.Title` with a manual capitalization approach
- Test files (`bottle_song_test.go`, `cases_test.go`) are auto-generated and must not be modified
- The package name must remain `bottlesong`
