# Challenger Review: bowling.go

## Verdict: NO ISSUES FOUND

The implementation is correct and handles all test cases properly.

## Detailed Analysis

### Roll Validation (Normal Frames 0-8)
- Negative pins: caught by `pins < 0` check (line 23) ✓
- Pins > 10: caught by `pins > 10` check (line 26) ✓
- Two rolls exceeding 10: caught by `prev+pins > 10` check (line 41) ✓
- Strike detection and frame advancement: correct (lines 32-34) ✓

### 10th Frame Validation
- **Strike first ball → ball 2**: `firstRoll==10` correctly bypasses the sum check and allows ball 2 ✓
- **Spare → ball 2**: `firstRoll+pins==10` correctly triggers ball 2 ✓
- **Open frame**: `done=true` after 2 balls, no bonus ✓
- **Strike+Strike → ball 2**: both strikes detected, no extra check needed (top-level ≤10 suffices since pins reset) ✓
- **Strike+non-strike → ball 2**: correctly validates `secondRoll+pins ≤ 10` (pins don't reset after non-strike second ball) ✓
- **"Second bonus cannot be strike if first isn't"**: With [10, 6] then roll 10: `secondRoll(6)+pins(10)=16 > 10` → error ✓
- **Spare bonus ball**: pins reset, 0-10 valid, caught by top-level check ✓

### Score Calculation (Frame-Walking)
- Strike: `10 + next two rolls`, rollIndex advances by 1 ✓
- Spare: `10 + next roll`, rollIndex advances by 2 ✓
- Open: `first + second`, rollIndex advances by 2 ✓
- 10th frame bonus rolls are naturally consumed as "next rolls" by the frame-walking algorithm ✓

### Verified Test Cases
All 21 score test cases and 10 roll validation test cases traced through the logic correctly:
- Perfect game (300) ✓
- Consecutive strikes bonus (81) ✓
- Last two strikes + bonus (31) ✓
- 10th frame spare/strike variants ✓
- All error cases (negative, >10, frame overflow, game-over) ✓

### Game Completion
- `done` flag only set when game is truly over ✓
- `Score()` errors when `!done` ✓
- `Roll()` errors when `done` ✓
- Incomplete game scenarios (unstarted, mid-game, missing bonus rolls) all correctly return errors ✓

### Code Quality
- Clean, idiomatic Go
- Proper error handling with `errors.New`
- Exported types and methods follow Go conventions
- Frame-walking scoring is the standard algorithm
