# Context Summary: Issue #197 — Bowling Game Scorer

## Status: DONE

## Branch: `issue-197`
## Commit: `feat: implement bowling game scorer`

## Key Facts
- Exercise: Exercism bowling problem adapted for polyglot benchmark
- Language: Go 1.18+
- Solution file: `go/exercises/practice/bowling/bowling.go`
- Test files (read-only): `bowling_test.go`, `cases_test.go`
- Tests: 31 total (10 roll validation + 21 score calculation), all passing
- Architecture: Flat roll array with frame-tracking state machine
- Dependencies: Go stdlib only (`errors` package)

## Decisions
- Used reference solution architecture from `.meta/example.go`
- Fixed-size array `[21]int` for roll storage (max possible rolls in bowling)
- Frame completion tracked via `nFrames` counter and `rFrameStart` index
- 10th frame special-cased within `Roll()` for bonus roll validation
