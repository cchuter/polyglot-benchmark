# Review: Book Store Implementation

## Verdict: PASS

All 18 test cases pass. The implementation is correct, clean, and follows the plan exactly.

## Algorithm Correctness

The recursive brute-force approach correctly finds optimal groupings by trying all possible group sizes (1 to len(counts)) at each step and returning the minimum cost. This guarantees the globally optimal solution.

### Manual Traces

**Trace 1: Empty basket `[]int{}`**
- freq = empty → counts = [] → minCost([]) → len == 0 → return 0
- Expected: 0 ✓

**Trace 2: Single book `[]int{1}`**
- counts = [1] → minCost([1]) → groupSize=1 → groupCost(1)=800 + minCost([])=0 → 800
- Expected: 800 ✓

**Trace 3: Key case `[1,1,2,2,3,3,4,5]` → two groups of 4 vs 5+3**
- counts = [2, 2, 2, 1, 1]
- groupSize=5: [1,1,1,0,0] → [1,1,1] → groupCost(5)=3000 + minCost([1,1,1])=2160 → **5160**
- groupSize=4: [1,1,1,0,1] → [1,1,1,1] → groupCost(4)=2560 + minCost([1,1,1,1])=2560 → **5120** ← minimum
- Expected: 5120 ✓ (correctly prefers 4+4 over 5+3)

## Edge Cases

| Case | Input | Result | Status |
|------|-------|--------|--------|
| Empty basket | `[]int{}` | 0 | ✓ |
| Single book | `[]int{1}` | 800 | ✓ |
| All same books | `[]int{2, 2}` | 1600 | ✓ |
| Two groups of 4 vs 5+3 | `[1,1,2,2,3,3,4,5]` | 5120 | ✓ |
| Complex grouping | `[1,2,2,3,3,3,4,4,4,4,5,5,5,5,5]` | 10000 | ✓ |

## Code Quality

- **Go conventions**: Clean idiomatic Go. Proper use of `sort` package, slices, maps.
- **Imports**: Only `"sort"` — necessary and sufficient.
- **Package name**: `bookstore` — correct.
- **Integer arithmetic**: All calculations use integer math (no floats). `bookPrice * size * (100 - discounts[size]) / 100` is safe.
- **MaxInt idiom**: `int(^uint(0) >> 1)` is the standard Go way to express MaxInt.

## Bounds Safety

- `discounts` array has 6 elements (indices 0-5). `groupSize` ranges from 1 to `len(counts)`, which is at most 5 (only 5 distinct book titles). No out-of-bounds risk.
- `counts[i]--` only for i < groupSize ≤ len(counts). Safe.
- No integer overflow concern: max single groupCost is 4000 (5 books × 800 × 75/100). Total cost is bounded by number of books × 800.

## Adherence to Plan

The implementation is an exact match of the code structure in `.osmi/plan.md`. Every function, constant, and line matches the plan.

## Minor Note: Performance

The last test case (`[1,1,1,1,1,1,2,2,2,2,2,2,3,3,3,3,3,3,4,4,5,5]`, 22 books) takes ~0.64s. The recursive approach without memoization is exponential in the worst case. For the bounded inputs in this exercise (max 5 distinct books, reasonable quantities), this is acceptable. Adding memoization on sorted count tuples would make it O(states) but is not needed to pass the tests.

## Security

No security concerns — pure computation, no I/O, no user-facing strings, no external dependencies.

## Test Results

```
PASS
ok  bookstore  0.667s
```

All 18/18 test cases pass.
