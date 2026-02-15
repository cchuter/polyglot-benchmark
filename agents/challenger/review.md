# Challenger Review: book-store Cost function

## Verdict: PASS

All 18 test cases pass. No correctness bugs found. Implementation faithfully follows the plan.

---

## 1. Algorithm Correctness

**Status: Correct**

The implementation uses frequency-based dynamic programming with memoization, exactly as specified in the plan. It correctly explores ALL possible group sizes at each step (line 36: `for gs := 1; gs <= distinct; gs++`), not just the greedy maximum. This is critical because greedy fails on cases like `[1,1,2,2,3,3,4,5]` where 4+4 grouping ($51.20) beats 5+3 ($51.60).

The memoization key is the sorted frequency tuple `[5]int`, which correctly canonicalizes equivalent states. The re-sort after each decrement (line 42) ensures the canonical form is maintained.

## 2. Edge Cases

**Status: All handled correctly**

| Case | Input | Behavior | Result |
|------|-------|----------|--------|
| Empty basket | `[]int{}` | freq = [0,0,0,0,0], `freq[0]==0` returns 0 | 0 (correct) |
| Single book | `[]int{1}` | freq = [1,0,0,0,0], one group of 1 | 800 (correct) |
| All same books | `[]int{2,2}` | freq = [2,0,0,0,0], distinct=1, only gs=1 possible | 1600 (correct) |
| Nil input | `nil` | `range nil` is no-op in Go, returns 0 | 0 (correct, untested) |

**Note:** Invalid book numbers (0, negative, >5) would cause an index-out-of-range panic at `freq[b-1]++` (line 13). This is acceptable since the exercise specification constrains books to 1-5, and no validation is expected.

## 3. Discount Math

**Status: No truncation issues**

`groupCost(n) = n * 800 * (100 - discounts[n]) / 100`

Verification:
- n=1: 1 x 800 x 100 = 80,000 / 100 = **800**
- n=2: 2 x 800 x 95 = 152,000 / 100 = **1,520**
- n=3: 3 x 800 x 90 = 216,000 / 100 = **2,160**
- n=4: 4 x 800 x 80 = 256,000 / 100 = **2,560**
- n=5: 5 x 800 x 75 = 300,000 / 100 = **3,000**

All intermediate values are cleanly divisible by 100. No integer truncation occurs. The discount table `[6]int{0, 0, 5, 10, 20, 25}` correctly maps group sizes 1-5 (index 0 unused).

## 4. Code Quality

**Status: Clean**

- `go vet`: No issues
- `gofmt`: No formatting differences
- No unused imports (`math` needed for `math.MaxInt32`, `sort` for sorting)
- Proper Go conventions: unexported helper functions, package-level var for constants
- Fixed-size `[5]int` arrays for frequencies -- efficient and avoids heap allocation
- Memoization map created per `Cost` call -- no state leakage between invocations

## 5. Plan Adherence

**Status: Faithful implementation**

The code matches the plan's algorithm pseudocode exactly:
- Frequency counting, sort descending, recursive search with memoization
- Discount table values match
- Function signatures match (`Cost`, `minCost`, `groupCost`)
- File modified: `book_store.go` only

## 6. Minor Observations (non-blocking)

1. **`math.MaxInt32` vs `math.MaxInt`**: The code uses `MaxInt32`. Since Go 1.17+, `math.MaxInt` is available and more idiomatic. Not a bug -- `MaxInt32` is sufficient for the cost domain.

2. **Redundant initial sort**: Line 15 sorts freq before passing to `minCost`, but `minCost` re-sorts after each decrement (line 42). The initial sort is still necessary for the first memoization lookup, so it's correct.

3. **No input validation**: Out-of-range book numbers would panic. Acceptable for exercism.

---

## Test Results

```
--- PASS: TestCost (0.00s)    [18/18 subtests passed]
PASS
ok  bookstore  0.005s
```
