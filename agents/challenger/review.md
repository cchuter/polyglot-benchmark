# Challenger Review: Kindergarten Garden Implementation

## Verdict: PASS — No blocking issues found

## Type & API Signatures

- `Garden` is `map[string][]string` — works with test expectations (`*Garden` pointer, map-like semantics). **OK**
- `NewGarden(diagram string, children []string) (*Garden, error)` — matches tests. **OK**
- `Plants(child string) ([]string, bool)` with pointer receiver — matches tests. **OK**

## Correctness Trace

### Single student (`"\nRC\nGG"`, Alice)
- rows = ["", "RC", "GG"]
- Row "RC": Alice gets row[0]='R'→radishes, row[1]='C'→clover
- Row "GG": Alice gets row[0]='G'→grass, row[1]='G'→grass
- Result: ["radishes", "clover", "grass", "grass"] — **matches test**

### Two students (`"\nVVCG\nVVRC"`, Alice/Bob)
- Row "VVCG": Alice→V,V; Bob→C,G
- Row "VVRC": Alice→V,V; Bob→R,C
- Bob: ["clover", "grass", "radishes", "clover"] — **matches test**

### Names out of order (test6: Samantha, Patricia, Xander, Roger)
- Sorted: [Patricia, Roger, Samantha, Xander]
- Row "VCRRGVRG": Patricia→V,C; Roger→R,R; Samantha→G,V; Xander→R,G
- Row "RVGCCGCV": Patricia→R,V; Roger→G,C; Samantha→C,G; Xander→C,V
- Patricia: ["violets","clover","radishes","violets"] — **matches test**
- Roger: ["radishes","radishes","grass","clover"] — **matches test**
- Samantha: ["grass","violets","clover","grass"] — **matches test**
- Xander: ["radishes","grass","clover","violets"] — **matches test**

### Full garden (test5, 12 children) — spot-checked Alice and Bob
- Alice (nx=0): row1[0,1]→V,R; row2[0,1]→V,R → ["violets","radishes","violets","radishes"] — **matches**
- Bob (nx=1): row1[2,3]→C,G; row2[2,3]→C,C → ["clover","grass","clover","clover"] — **matches**

### Two gardens independence (TestTwoGardens)
- g1 children: [Alice, Bob, Charlie, Dan] → Bob is at index 1
- g2 children: [Bob, Charlie, Dan, Erin] → Bob is at index 0
- g1 Bob: row[2,3] = R,R / G,C → ["radishes","radishes","grass","clover"] — **matches**
- g2 Bob: row[0,1] = V,C / R,V → ["violets","clover","radishes","violets"] — **matches**

### Lookup invalid name
- `Plants("Bob")` on a garden that only has Alice → map lookup returns `ok=false`. **OK**

## Error Case Handling

| Test Case | Diagram | Trigger | Caught? |
|-----------|---------|---------|---------|
| Wrong format | `"RC\nGG"` | `len(rows)!=3` (rows=["RC","GG"], len=2) | **YES** |
| Mismatched rows | `"\nRCCC\nGG"` | `len(rows[1])!=len(rows[2])` (4!=2) | **YES** |
| Odd cups | `"\nRCC\nGGC"` | `len(rows[1])!=2*len(children)` (3!=2) | **YES** |
| Duplicate name | `"\nRCCC\nGGCC"` children=["Alice","Alice"] | `len(g)!=len(alpha)` (1!=2) | **YES** |
| Invalid codes | `"\nrc\ngg"` | `default` in switch | **YES** |

## Safety: Input Children Slice Not Modified

```go
alpha := append([]string{}, children...)
sort.Strings(alpha)
```

This creates a **new backing array** via `append([]string{}, children...)`, then sorts the copy. The original `children` slice is never touched. **SAFE** — TestNamesNotModified will pass.

## Independence: No Package Globals

Each `NewGarden` call creates a fresh `Garden{}` map local to the function. No package-level variables are used for state. **SAFE** — TestTwoGardens will pass.

## Loop Correctness: `for cx := range rows[1:]`

The innermost loop `for cx := range rows[1:]` is an idiom to generate indices 0 and 1. Since we validated `len(rows) == 3`, `rows[1:]` always has exactly 2 elements, so `cx` is always 0 then 1. Combined with `row[2*nx+cx]`, this correctly accesses the two cups per child in each row.

This is functionally equivalent to `for cx := 0; cx < 2; cx++` — slightly less readable but correct.

## Minor Style Observations (non-blocking)

1. The `for cx := range rows[1:]` idiom is clever but could be clearer as `for cx := 0; cx < 2; cx++`. Not a bug.
2. The error messages are descriptive and reasonable.
3. No unnecessary allocations — `make([]string, 0, 4)` preallocates capacity correctly (each child gets exactly 4 plants).

## Conclusion

The implementation is **correct, safe, and complete**. All test cases should pass, including:
- TestGarden (all 11 sub-tests)
- TestNamesNotModified
- TestTwoGardens
- Both benchmarks

**No changes required.**
