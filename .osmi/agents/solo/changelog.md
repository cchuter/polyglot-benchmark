# Solo Agent Change Log

## Change 1: Implement bowling.go

**File:** `go/exercises/practice/bowling/bowling.go`
**Action:** Replaced stub with full implementation
**Commit:** 20fcf16

### What was done
- Added error variables for all validation cases
- Defined constants for bowling rules (pins per frame, frames per game, max rolls)
- Implemented `Game` struct with array-based roll tracking
- Implemented `NewGame()` constructor
- Implemented `Roll(pins int) error` with full validation:
  - Negative pins check
  - Pin count > 10 check
  - Game-over check
  - Frame total > 10 check
  - 10th frame bonus roll validation
- Implemented `Score() (int, error)` with frame-by-frame scoring including strike/spare bonuses
- Added helper methods for frame state tracking

### Test Results
- All 10 Roll test cases: PASS
- All 21 Score test cases: PASS
- `go vet`: clean
