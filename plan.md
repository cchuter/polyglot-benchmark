# Implementation Plan: Bowling Game Scorer

## File to Modify

- `go/exercises/practice/bowling/bowling.go` — the only file that needs changes

## Data Model

```go
type Game struct {
    rolls       []int // all rolls recorded so far
    currentRoll int   // index into rolls (same as len(rolls))
    frame       int   // current frame (0-indexed, 0-9)
    rollInFrame int   // 0 = first roll in frame, 1 = second, 2 = bonus (10th frame only)
    done        bool  // true when game is complete
}
```

### Design Decision: Roll-based storage with frame tracking

Store all rolls in a flat slice and track frame/roll position during `Roll()` for validation. During `Score()`, walk the rolls slice frame-by-frame to compute the score with bonuses. This approach cleanly separates validation (during Roll) from scoring (during Score).

## Roll() Logic

1. **Game over check**: if `done`, return error
2. **Pin range check**: if `pins < 0 || pins > 10`, return error
3. **Frame pin total check** (varies by frame and roll position):
   - Frames 0-8: `rollInFrame == 1` → `rolls[len(rolls)-1] + pins > 10` → error
   - Frame 9 (10th frame):
     - `rollInFrame == 1`: if previous roll was a strike, pins can be 0-10; otherwise `prev + pins > 10` → error
     - `rollInFrame == 2`: if first two rolls sum includes a strike: if roll[0]=10 and roll[1]=10, pins can be 0-10; if roll[0]=10 and roll[1]<10, then `roll[1]+pins <= 10`; if spare (roll[0]+roll[1]=10), pins can be 0-10
4. **Record the roll**: append to `rolls`
5. **Advance frame state**:
   - Frames 0-8: strike (pins==10, rollInFrame==0) → advance frame; rollInFrame==1 → advance frame
   - Frame 9: after 2 rolls if no strike/spare → done; after 3 rolls → done; after 2 rolls if open frame → done

## Score() Logic

1. If `!done`, return error
2. Walk through rolls frame by frame:
   - For each frame 0-8:
     - Strike: score += rolls[i] + rolls[i+1] + rolls[i+2]; i += 1
     - Spare: score += 10 + rolls[i+2]; i += 2
     - Open: score += rolls[i] + rolls[i+1]; i += 2
   - Frame 9: sum remaining rolls (no bonus beyond what's there)

## Validation Edge Cases for 10th Frame

The 10th frame has complex validation:

- After a strike (first roll = 10):
  - Second roll can be 0-10
  - Third roll: if second was also a strike (10), can be 0-10; otherwise third + second <= 10
- After a spare (first + second = 10):
  - Third roll can be 0-10
- Open frame (first + second < 10):
  - No third roll, game is done after second

## Approach and Ordering

1. Define the `Game` struct and `NewGame()` constructor
2. Implement `Roll()` with all validation logic
3. Implement `Score()` with frame-by-frame walk
4. Test with `go test ./...` in the bowling directory

## Error Handling

Use `fmt.Errorf()` or `errors.New()` for error messages. The tests only check for `err != nil`, not specific messages.
