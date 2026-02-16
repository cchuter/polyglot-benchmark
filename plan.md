# Implementation Plan: Bowling Scorer

## File to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file that needs changes

## Architectural Approach

Use the reference solution from `.meta/example.go` as the basis. The approach tracks rolls in a fixed-size array and uses frame completion counting to manage game state.

### Data Structure

```go
type Game struct {
    rolls       [21]int  // max 21 rolls in a game (9 frames × 2 + 3 for 10th)
    nRolls      int      // number of rolls recorded so far
    nFrames     int      // number of completed frames
    rFrameStart int      // index of the first roll in the current frame
}
```

### Key Design Decisions

1. **Roll-based storage**: Store each roll individually rather than frame-based. This simplifies bonus calculation since strikes/spares look ahead at subsequent rolls regardless of frame boundaries.

2. **Frame tracking during Roll()**: Track completed frames as rolls come in, rather than reconstructing frames during Score(). This enables immediate validation (e.g., "game is over" checks).

3. **Validation in Roll()**: All validation happens at roll time:
   - Negative pin count → error
   - Pin count > 10 → error
   - Game already over (10 frames completed) → error
   - Two rolls in a frame exceeding 10 pins → error (with 10th frame exceptions)
   - 10th frame bonus roll validation (complex rules around when 10+ is valid)

4. **Score calculation**: Walk through frames using frame start indices. For each frame:
   - Strike: 10 + next two rolls
   - Spare: 10 + next one roll
   - Open: sum of two rolls

5. **Error types**: Define package-level error variables for the four error cases.

### Constants

```go
pinsPerFrame      = 10
framesPerGame     = 10
maxRollsPerFrame  = 2
maxRollsLastFrame = 3
maxRolls          = (maxRollsPerFrame * (framesPerGame - 1)) + maxRollsLastFrame  // = 21
```

### Roll() Logic

1. Check pins < 0 → ErrNegativeRollIsInvalid
2. Check pins > 10 → ErrPinCountExceedsPinsOnTheLane
3. Check completed frames == 10 → ErrCannotRollAfterGameOver
4. Record the roll
5. Handle strike in frames 1-9: complete frame immediately
6. Handle 2nd roll in frames 1-9: validate sum ≤ 10, complete frame
7. Handle 10th frame special cases:
   - After 2 rolls: if score < 10, complete frame (no bonus rolls)
   - After 3 rolls: validate bonus roll constraints, complete frame

### Score() Logic

1. Check completed frames == 10, else return ErrPrematureScore
2. Walk through frames computing score with bonus lookups

## Implementation Order

1. Write the complete `bowling.go` with all types, constants, errors, and methods
2. Run tests to verify

## Risk Assessment

- The 10th frame validation is the most complex part — multiple edge cases around when bonus rolls can exceed the normal pin limit
- The reference solution handles all test cases, so following its logic closely minimizes risk
