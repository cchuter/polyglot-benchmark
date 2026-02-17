# Implementation Plan: polyglot-go-bowling

## Proposal A: Roll-tracking with frame-state machine

**Role: Proponent**

### Approach

Use the reference solution from `.meta/example.go` as the architectural guide. Store all rolls in a fixed-size array and track frame completion with counters. The `Roll()` method validates and records rolls while tracking frame boundaries. The `Score()` method walks the rolls array frame-by-frame, applying strike/spare bonuses.

### Files to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file to implement.

### Architecture

- **Game struct**: Fixed-size array `[21]int` for rolls, counters for total rolls, completed frames, and current frame start index.
- **Constants**: `pinsPerFrame=10`, `framesPerGame=10`, `maxRollsPerFrame=2`, `maxRollsLastFrame=3`, `maxRolls=21`.
- **Roll()**: Validates pins (negative, >10, game over), records the roll, then determines if the frame is complete based on:
  - Strike in frames 1-9: immediate frame completion.
  - Two rolls in frames 1-9: validate sum ≤ 10, complete frame.
  - 10th frame: up to 3 rolls with special validation for bonus balls after strike/spare.
- **Score()**: Iterates 10 frames, advancing a `frameStart` pointer. Adds strike bonus (next 2 rolls), spare bonus (next 1 roll), or raw score.
- **Helper methods**: `rollsThisFrame()`, `completeTheFrame()`, `completedFrames()`, `isStrike()`, `isSpare()`, `rawFrameScore()`, `spareBonus()`, `strikeBonus()`.

### Rationale

This is the canonical Exercism approach. It directly mirrors the reference solution, which is known to pass all tests. The frame-state tracking is straightforward and the scoring walk is clean.

---

## Proposal B: Append-only roll slice with computed scoring

**Role: Opponent**

### Approach

Instead of a fixed array with explicit frame tracking, use a dynamic slice to store rolls and compute frame boundaries lazily during both `Roll()` validation and `Score()`. Track only the current frame number and whether we're on the first or second ball of a frame.

### Files to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file to implement.

### Architecture

- **Game struct**: `rolls []int`, `frame int`, `firstBallInFrame bool`, `done bool`.
- **Roll()**: Validate pins, determine if this is a fill ball (frame 10), apply appropriate validation, append to slice, advance frame state.
- **Score()**: Walk the slice using a roll index, computing frame boundaries on the fly based on whether each frame starts with a strike.

### Critique of Proposal A

Proposal A uses a fixed-size array which is slightly over-engineered for a problem that never exceeds 21 rolls. The frame-start-index tracking in Proposal A is redundant information that could be computed. However, these are minor concerns.

### Rationale

A dynamic slice is more idiomatic Go in some contexts. The lazy computation avoids maintaining frame-start as mutable state.

### Weaknesses

The lazy approach makes 10th-frame validation harder — you need to reconstruct frame state during `Roll()` anyway, which defeats the "lazy" benefit. The reference solution's explicit tracking is actually simpler for the 10th frame edge cases.

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criterion    | Proposal A                         | Proposal B                          |
|-------------|-------------------------------------|--------------------------------------|
| Correctness | Proven (matches reference solution) | Likely correct but unproven          |
| Risk        | Very low (known-good pattern)       | Medium (10th frame edge cases)       |
| Simplicity  | Explicit state, easy to debug       | Seems simpler but hidden complexity  |
| Consistency | Matches `.meta/example.go` exactly  | Diverges from project conventions    |

### Decision

**Proposal A wins.** The reference solution in `.meta/example.go` is proven correct against the test suite. The explicit frame-tracking approach handles the complex 10th-frame validation cleanly. There is no benefit to deviating from a known-good pattern.

### Final Implementation Plan

**File**: `go/exercises/practice/bowling/bowling.go`

**Step 1**: Define error variables and constants.
```go
var (
    ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
    ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
    ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
    ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)

const (
    pinsPerFrame      = 10
    framesPerGame     = 10
    maxRollsPerFrame  = 2
    maxRollsLastFrame = 3
    maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)
```

**Step 2**: Define the `Game` struct with a fixed-size rolls array and frame-tracking counters.
```go
type Game struct {
    rolls       [maxRolls]int
    nRolls      int
    nFrames     int
    rFrameStart int
}
```

**Step 3**: Implement `NewGame()` returning a zero-valued `*Game`.

**Step 4**: Implement `Roll(pins int) error` with:
- Pin validation (negative, >10)
- Game-over check
- Roll recording
- Strike detection for frames 1-9
- Normal frame completion (2 rolls, sum ≤ 10)
- 10th frame logic: 2 rolls < 10 → done; otherwise up to 3 rolls with bonus validation

**Step 5**: Implement `Score() (int, error)` with:
- Incomplete game check
- Frame-by-frame scoring walk with strike/spare bonus application

**Step 6**: Implement helper methods as one-liners for readability.

**Step 7**: Run `go test ./...` and `go vet ./...` to verify.
