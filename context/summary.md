# Context Summary: bottle-song (Issue #26)

## Key Decisions

- Replaced `strings.Title()` with a custom `capitalize()` function to avoid deprecated API
- Used `strings.ToUpper(s[:1]) + s[1:]` which is safe for ASCII-only inputs (number words "one" through "ten")
- No external dependencies added — standard library only

## Files Modified

- `go/exercises/practice/bottle-song/bottle_song.go` — Added `capitalize` helper, replaced `strings.Title` calls

## Test Results

- 7/7 tests pass (`go test -v`)
- No staticcheck warnings (`staticcheck ./...`)
- No go vet warnings (`go vet ./...`)
- Build succeeds (`go build ./...`)

## Commit

- `6de25d9` — "fix: replace deprecated strings.Title with capitalize helper in bottle-song"
- Branch: `issue-26`
