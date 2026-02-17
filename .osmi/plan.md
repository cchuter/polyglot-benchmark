# Design Plan: Bowling Game Scoring

## Branch 1: Array-Based Roll Tracking (Minimal, Direct)

Store all rolls in a fixed-size array. Track the number of rolls and completed frames. On each `Roll()`, validate the pin count against game state. On `Score()`, walk through rolls frame-by-frame, applying strike/spare bonuses by looking ahead in the array.

**Files:** Only `bowling.go`

**Approach:**
- `Game` struct with `rolls [21]int`, `nRolls int`, `nFrames int`, `rFrameStart int`
- `Roll()` validates pins, records the roll, and advances frame tracking
- `Score()` iterates frames 0–9, summing base + bonus using lookahead
- Helper methods: `isStrike`, `isSpare`, `rawFrameScore`, `strikeBonus`, `spareBonus`

**Evaluation:**
- Feasibility: High — matches the reference solution pattern exactly
- Risk: Low — well-understood algorithm, proven by `.meta/example.go`
- Alignment: Fully satisfies all acceptance criteria
- Complexity: Single file, ~130 lines

## Branch 2: Frame-Object Model (Extensible)

Model each frame as a struct with its own state (rolls, type, bonus tracking). The `Game` holds a slice of `Frame` objects and delegates roll recording and scoring to each frame.

**Files:** Only `bowling.go`

**Approach:**
- `Frame` struct with `rolls []int`, `frameType` enum, `complete bool`
- `Game` holds `frames [10]Frame` and routes rolls to the current frame
- Each frame determines its own completeness and score

**Evaluation:**
- Feasibility: High — achievable with only stdlib
- Risk: Medium — more complex validation logic spread across types, higher chance of edge case bugs in 10th frame handling
- Alignment: Satisfies criteria but more code to verify
- Complexity: Single file but ~200+ lines, more abstractions

## Branch 3: State Machine (Roll-by-Roll)

Model the game as a finite state machine with states like `FirstRoll`, `SecondRoll`, `BonusRoll1`, `BonusRoll2`, `GameOver`. Each state defines valid transitions and pin count constraints.

**Files:** Only `bowling.go`

**Approach:**
- Enum-like state constants
- `Roll()` switches on current state and transitions
- Accumulate rolls in array, score at end using standard frame walk

**Evaluation:**
- Feasibility: High — state machines are well-suited to bowling
- Risk: Medium — many state transitions to get right, especially for 10th frame
- Alignment: Satisfies criteria
- Complexity: Single file, ~150 lines, but more complex control flow

## Selected Plan

**Branch 1: Array-Based Roll Tracking** is selected.

**Rationale:** This approach is the simplest, most direct, and matches the proven reference implementation. It has the lowest risk since the algorithm is validated by the existing `.meta/example.go`. It uses minimal abstraction — just a struct with an array and counters — making it easy to verify against all test cases. The other branches add complexity without benefit for this well-scoped exercise.

### Detailed Implementation Plan

**File to modify:** `go/exercises/practice/bowling/bowling.go`

**Step 1: Define error variables**
```go
var (
    ErrNegativeRollIsInvalid        = errors.New("Negative roll is invalid")
    ErrPinCountExceedsPinsOnTheLane = errors.New("Pin count exceeds pins on the lane")
    ErrPrematureScore               = errors.New("Score cannot be taken until the end of the game")
    ErrCannotRollAfterGameOver      = errors.New("Cannot roll after game is over")
)
```

**Step 2: Define constants**
```go
const (
    pinsPerFrame      = 10
    framesPerGame     = 10
    maxRollsPerFrame  = 2
    maxRollsLastFrame = 3
    maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame
)
```

**Step 3: Define Game struct**
```go
type Game struct {
    rolls       [maxRolls]int
    nRolls      int
    nFrames     int
    rFrameStart int
}
```

**Step 4: Implement NewGame()**
Returns a zero-valued `*Game`.

**Step 5: Implement Roll()**
- Validate: pins < 0 → error, pins > 10 → error, game over → error
- Record the roll
- Handle strikes in frames 1–9 (complete frame immediately)
- Handle normal frames (2 rolls, validate total ≤ 10)
- Handle 10th frame special cases (up to 3 rolls, pin reset rules)

**Step 6: Implement Score()**
- If game not complete, return error
- Walk frames 0–9, sum score with strike/spare bonuses via lookahead

**Step 7: Implement helper methods**
`rollsThisFrame`, `completeTheFrame`, `completedFrames`, `isStrike`, `isSpare`, `rawFrameScore`, `strikeBonus`, `spareBonus`

**Step 8: Run tests** with `go test ./...` in the bowling directory
