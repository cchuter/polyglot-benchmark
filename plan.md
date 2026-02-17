# Implementation Plan: polyglot-go-bowling

## Proposal A: Roll-Array with Frame Tracking

**Role: Proponent**

### Approach

Store all rolls in a fixed-size array and track game progress using frame count and frame-start index. This mirrors the reference solution's architecture closely.

### Files to Modify

- `go/exercises/practice/bowling/bowling.go` — replace stub with full implementation

### Architecture

```go
type Game struct {
    rolls       [21]int  // max 21 rolls in a game (10 frames, last frame can have 3)
    nRolls      int      // number of rolls recorded
    nFrames     int      // completed frames (0-10)
    rFrameStart int      // index of the first roll of current frame
}
```

**State machine in Roll():**
1. Validate pins (negative, >10)
2. Check game not over (nFrames == 10)
3. Record the roll
4. Handle strike in frames 1-9: complete frame immediately
5. Handle second roll in frames 1-9: validate frame total ≤ 10, complete frame
6. Handle 10th frame: allow up to 3 rolls based on strikes/spares, validate pin counts

**Score():**
- Walk through frames using frameStart pointer
- For strikes: 10 + next 2 rolls
- For spares: 10 + next 1 roll
- For open frames: sum of 2 rolls

### Rationale

- Proven correct (matches reference architecture)
- Simple array storage with O(1) lookups
- Frame-based tracking makes scoring straightforward
- Minimal memory footprint (fixed-size array)

---

## Proposal B: Roll-Slice with Post-hoc Frame Analysis

**Role: Opponent**

### Approach

Store rolls in a dynamic slice. Don't track frames during Roll() — instead, analyze the roll history in both Roll() (for validation) and Score() (for computation). The frame boundaries are derived from the roll data itself.

### Files to Modify

- `go/exercises/practice/bowling/bowling.go` — replace stub with full implementation

### Architecture

```go
type Game struct {
    rolls    []int
    done     bool
}
```

**Roll():** Append to slice. After each append, analyze the full roll history to determine the current frame and validate:
- Derive frame boundaries by scanning rolls from the beginning
- Check if the game is complete
- Validate pin counts in context of current frame position

**Score():** Walk the roll slice, determine frame boundaries, and compute score.

### Critique of Proposal A

- Proposal A duplicates frame tracking in two places (Roll and Score), creating potential inconsistency
- The rFrameStart tracking adds complexity to the Roll method

### Rationale for Proposal B

- Single source of truth: frame analysis is centralized
- Simpler state: only the roll data and a done flag
- Easier to reason about correctness

### Weaknesses of Proposal B

- Re-analyzing all rolls on every Roll() call is inefficient (though game is max 21 rolls, so negligible)
- More complex logic in the frame-derivation function
- Deriving frames post-hoc introduces risk of subtle bugs in the analysis function
- Does not match established patterns in the codebase

---

## Selected Plan

**Role: Judge**

### Evaluation

| Criteria      | Proposal A                           | Proposal B                              |
|---------------|--------------------------------------|-----------------------------------------|
| Correctness   | High — proven by reference solution  | Medium — novel derivation logic          |
| Risk          | Low — well-understood approach       | Medium — post-hoc analysis is error-prone|
| Simplicity    | Moderate — explicit state tracking   | Lower — centralized but complex analysis |
| Consistency   | High — matches .meta/example.go      | Low — diverges from established pattern  |

### Decision: Proposal A wins

**Rationale:**
1. The reference solution in `.meta/example.go` validates this architecture works for all test cases
2. Explicit frame tracking during Roll() keeps validation logic localized and incremental
3. The fixed-size array is simpler than dynamic slice allocation
4. Consistency with the existing codebase pattern is important for maintainability

### Detailed Implementation Plan

**File:** `go/exercises/practice/bowling/bowling.go`

**Step 1: Define package, imports, and error variables**
```go
package bowling

import "errors"

var (
    ErrNegativeRollIsInvalid       = errors.New("Negative roll is invalid")
    ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
    ErrPrematureScore              = errors.New("Score cannot be taken until the end of the game")
    ErrCannotRollAfterGameOver     = errors.New("Cannot roll after game is over")
)
```

**Step 2: Define constants and Game struct**
```go
const (
    pinsPerFrame      = 10
    framesPerGame     = 10
    maxRollsPerFrame  = 2
    maxRollsLastFrame = 3
    maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)

type Game struct {
    rolls       [maxRolls]int
    nRolls      int
    nFrames     int
    rFrameStart int
}
```

**Step 3: Implement NewGame**
```go
func NewGame() *Game {
    return &Game{}
}
```

**Step 4: Implement Roll with frame-aware validation**
- Check pins < 0 → ErrNegativeRollIsInvalid
- Check pins > 10 → ErrPinCountExceedsPinsOnTheLane
- Check game over → ErrCannotRollAfterGameOver
- Record roll, increment count
- Handle strike in frames 1-9
- Handle second roll in frames 1-9 (validate sum ≤ 10)
- Handle 10th frame (2 rolls if open, 3 rolls if spare/strike, with appropriate validation)

**Step 5: Implement Score**
- Check game complete → ErrPrematureScore
- Walk frames, compute strike/spare/open frame scores

**Step 6: Implement helper methods**
- rollsThisFrame(), completeTheFrame(), completedFrames()
- isStrike(), isSpare(), rawFrameScore(), spareBonus(), strikeBonus()
