# Implementation Plan: Beer Song

## Branch 1: Switch-based approach (minimal, matches example)

Directly follow the reference solution pattern from `.meta/example.go`. Use a switch statement in `Verse()` to handle the four cases (0, 1, 2, default). Use `bytes.Buffer` in `Verses()` to concatenate. `Song()` delegates to `Verses(99, 0)`.

### Files to modify
- `go/exercises/practice/beer-song/beer_song.go` — implement all three functions

### Approach
1. Import `bytes` and `fmt`
2. `Verse(n)`: switch on n with cases for 0, 1, 2, and default (3-99). Return error for n<0 or n>99
3. `Verses(start, stop)`: validate inputs, loop from start down to stop, call `Verse()` for each, join with `\n`
4. `Song()`: call `Verses(99, 0)`, return result

### Evaluation
- **Feasibility**: High — proven pattern from the reference solution
- **Risk**: Very low — straightforward switch logic
- **Alignment**: Fully satisfies all acceptance criteria
- **Complexity**: ~50 lines, single file, minimal imports

## Branch 2: fmt.Sprintf template approach (extensible)

Use a helper function that takes a verse number and produces the verse using `fmt.Sprintf` with conditional template selection. Separate the concerns of formatting from iteration.

### Files to modify
- `go/exercises/practice/beer-song/beer_song.go` — implement all three functions plus a helper

### Approach
1. Create helper `bottleStr(n)` returning "N bottles", "1 bottle", "no more bottles"
2. Create helper `actionStr(n)` returning the appropriate action line
3. `Verse(n)`: compose using helpers with Sprintf
4. `Verses(start, stop)`: same loop pattern
5. `Song()`: delegates to Verses

### Evaluation
- **Feasibility**: High — but more abstraction than needed
- **Risk**: Low, but more places for subtle bugs in string formatting
- **Alignment**: Fully satisfies acceptance criteria
- **Complexity**: ~70 lines, more functions, more potential for string mismatch

## Branch 3: strings.Builder with precomputed map (performance)

Precompute all 100 verses into a map at init time. Verse/Verses/Song just do lookups.

### Files to modify
- `go/exercises/practice/beer-song/beer_song.go` — implement with init() and map

### Approach
1. `var verseCache map[int]string` populated in `init()`
2. `Verse(n)`: map lookup
3. `Verses(start, stop)`: concatenate from map with strings.Builder
4. `Song()`: delegates to Verses

### Evaluation
- **Feasibility**: High — but over-engineered for this problem
- **Risk**: Medium — init() adds complexity, harder to debug
- **Alignment**: Satisfies criteria but adds unnecessary complexity
- **Complexity**: ~60 lines, uses init(), map allocation

## Selected Plan

**Branch 1 (Switch-based approach)** is selected.

### Rationale
- It is the simplest, most readable approach
- It directly matches the proven reference solution pattern
- Minimal code, minimal risk of string formatting bugs
- No unnecessary abstractions or init-time computation
- The problem domain (4 distinct verse types) maps naturally to a switch statement

### Detailed Implementation Plan

**File**: `go/exercises/practice/beer-song/beer_song.go`

**Step 1**: Add imports for `bytes` and `fmt`

**Step 2**: Implement `Verse(n int) (string, error)`:
```go
func Verse(n int) (string, error) {
    switch {
    case n < 0 || n > 99:
        return "", fmt.Errorf("%d is not a valid verse", n)
    case n == 0:
        return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
    case n == 1:
        return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
    case n == 2:
        return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
    default:
        return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
    }
}
```

**Step 3**: Implement `Verses(start, stop int) (string, error)`:
```go
func Verses(start, stop int) (string, error) {
    switch {
    case start < 0 || start > 99:
        return "", fmt.Errorf("start value [%d] is not valid", start)
    case stop < 0 || stop > 99:
        return "", fmt.Errorf("stop value [%d] is not valid", stop)
    case start < stop:
        return "", fmt.Errorf("start [%d] is less than stop [%d]", start, stop)
    }
    var buf bytes.Buffer
    for i := start; i >= stop; i-- {
        v, _ := Verse(i)
        buf.WriteString(v)
        buf.WriteString("\n")
    }
    return buf.String(), nil
}
```

**Step 4**: Implement `Song() string`:
```go
func Song() string {
    result, _ := Verses(99, 0)
    return result
}
```

**Step 5**: Run `go test` to verify all tests pass.
