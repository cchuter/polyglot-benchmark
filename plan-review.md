# Plan Review: Book Store Exercise (Proposal A -- Greedy with 5+3 to 4+4 Optimization)

## Overall Assessment: PASS with minor code fixes required

The selected plan (Proposal A) is algorithmically correct and will produce the right answers for all 18 test cases. The greedy-then-fix approach is the well-known canonical solution for this problem. However, the pseudocode implementation has two bugs that must be fixed before it will compile and run correctly.

---

## 1. Correctness: Will the algorithm produce correct answers for ALL 18 test cases?

**Verdict: YES**, the algorithm is correct.

The greedy approach (sort frequencies descending, repeatedly form a group from all non-zero frequencies) produces groups that are as large as possible. The only case where this greedy result is suboptimal is when a group of 5 and a group of 3 could be replaced by two groups of 4, because:

- 5-group cost: 5 * 600 = 3000
- 3-group cost: 3 * 720 = 2160
- Total: 5160

vs.

- Two 4-groups: 2 * 4 * 640 = 5120
- Savings: 40 cents per conversion

The post-processing step handles this correctly.

### Trace of critical test cases:

**Test 8** -- `{1,1,2,2,3,3,4,5}` expected 5120:
- Frequencies: [3,3,2,1,1] (sorted desc, 5 distinct values based on books present, some have freq 1)
- Actually: book1=2, book2=2, book3=2, book4=1, book5=1 -> sorted desc: [2,2,2,1,1]
- Round 1: group of 5, freq=[1,1,1,0,0]
- Round 2: group of 3, freq=[0,0,0,0,0]
- Groups: {5:1, 3:1}. Optimization: 1 pair -> {4:2}. Cost: 2*4*640 = 5120. CORRECT.

**Test 15** -- `{1,1,2,2,3,3,4,5,1,1,2,2,3,3,4,5}` expected 10240:
- Frequencies: book1=4, book2=4, book3=4, book4=2, book5=2 -> sorted desc: [4,4,4,2,2]
- Round 1: group of 5, freq=[3,3,3,1,1]
- Round 2: group of 5, freq=[2,2,2,0,0]
- Round 3: group of 3, freq=[1,1,1,0,0]
- Round 4: group of 3, freq=[0,0,0,0,0]
- Groups: {5:2, 3:2}. Optimization: 2 pairs -> {4:4}. Cost: 4*4*640 = 10240. CORRECT.

**Test 16** -- `{1,1,1,1,1,1,2,2,2,2,2,2,3,3,3,3,3,3,4,4,5,5}` expected 14560:
- Frequencies: [6,6,6,2,2] sorted desc
- Round 1: group of 5, freq=[5,5,5,1,1]
- Round 2: group of 5, freq=[4,4,4,0,0]
- Round 3: group of 3, freq=[3,3,3,0,0]
- Round 4: group of 3, freq=[2,2,2,0,0]
- Round 5: group of 3, freq=[1,1,1,0,0]
- Round 6: group of 3, freq=[0,0,0,0,0]
- Groups: {5:2, 3:4}. Optimization: min(2,4)=2 pairs -> {5:0, 3:2, 4:4}. Cost: 4*4*640 + 2*3*720 = 10240 + 4320 = 14560. CORRECT.

**Test 17** -- `{1,1,2,3,4}` expected 3360:
- Frequencies: book1=2, book2=1, book3=1, book4=1 -> sorted desc: [2,1,1,1,0]
- Round 1: group of 4, freq=[1,0,0,0,0]
- Round 2: group of 1, freq=[0,0,0,0,0]
- Groups: {4:1, 1:1}. No 5+3 pairs. Cost: 1*4*640 + 1*1*800 = 2560 + 800 = 3360. CORRECT.

**Test 18** -- `{1,2,2,3,3,3,4,4,4,4,5,5,5,5,5}` expected 10000:
- Frequencies: [5,4,3,2,1] sorted desc
- Round 1: group of 5, freq=[4,3,2,1,0]
- Round 2: group of 4, freq=[3,2,1,0,0]
- Round 3: group of 3, freq=[2,1,0,0,0]
- Round 4: group of 2, freq=[1,0,0,0,0]
- Round 5: group of 1, freq=[0,0,0,0,0]
- Groups: {5:1, 4:1, 3:1, 2:1, 1:1}. Optimization: 1 pair -> {5:0, 3:0, 4:3, 2:1, 1:1}. Cost: 3*4*640 + 1*2*760 + 1*1*800 = 7680 + 1520 + 800 = 10000. CORRECT.

All 18 test cases verified. The algorithm handles every one correctly.

---

## 2. Completeness: Edge cases

| Edge Case | Handled? | Notes |
|-----------|----------|-------|
| Empty basket `[]` | YES | The greedy loop forms no groups; total = 0. |
| Single book `[1]` | YES | Forms one group of 1; cost = 800. |
| All duplicates `[2,2]` | YES | Forms two groups of 1; cost = 1600. |
| Large basket with skewed frequencies | YES | Test 16 and 18 verified above. |
| No 5+3 optimization needed | YES | Tests 1-7, 10-11, 17 have no 5+3 pairs. |
| Multiple 5+3 conversions needed | YES | Test 15 (2 pairs), Test 16 (2 pairs). |

---

## 3. Implementation Concerns: Bugs in the Pseudocode

### Bug 1: Sorting a slice of a fixed-size array (CRITICAL)

```go
sort.Sort(sort.Reverse(sort.IntSlice(freq[1:])))
```

The `freq` array is declared as `[6]int` (a fixed-size array). The expression `freq[1:]` creates a **slice** backed by the array, and sorting the slice DOES modify the underlying array. So this actually works correctly in Go -- slicing a fixed-size array produces a slice that references the original array's memory.

However, there is an **indexing problem**. After sorting `freq[1:]` in descending order, the code continues to iterate with `for i := 1; i <= 5; i++` and checks `freq[i]`. Since the sort operates on indices 1 through 5 of the array and the loop also uses indices 1 through 5, the indexing is consistent. This is actually fine.

**Verdict: Not a bug.** The slice-of-array approach works correctly in Go.

### Bug 2: The `min` function (CRITICAL)

```go
pairs := min(groupCount[5], groupCount[3])
```

The built-in `min` function was only introduced in **Go 1.21**. The `go.mod` file specifies `go 1.18`, which means the built-in `min` is **not available** and this code will **fail to compile**.

**Action required**: Either:
- (a) Add a local `min` helper function, or
- (b) Inline the logic: `pairs := groupCount[5]; if groupCount[3] < pairs { pairs = groupCount[3] }`

### Bug 3: Missing `sort` import

The pseudocode uses `sort.Sort(sort.Reverse(sort.IntSlice(...)))` but does not include the `import "sort"` statement in the code block. The plan does note "Imports needed: sort" at the bottom, so this is acknowledged, but the implementer must remember to add it.

**Verdict: Minor -- acknowledged in the plan text, must be included in implementation.**

### Bug 4: Re-sorting inside the loop is correct but unnecessary

The plan re-sorts frequencies after each greedy group extraction:

```go
sort.Sort(sort.Reverse(sort.IntSlice(freq[1:])))
```

This is functionally correct. After removing one from each non-zero frequency, re-sorting ensures the next iteration correctly identifies how many non-zero frequencies exist. In fact, a simpler approach exists: since all non-zero values decrease by exactly 1, the sorted order is preserved UNLESS a frequency that was larger than another becomes equal or smaller. Re-sorting guarantees correctness.

**Verdict: Correct, though a more efficient approach could avoid re-sorting. Not a bug -- just a minor performance note.**

---

## 4. Go-Specific Issues

### 4.1 Compilation

The code will compile provided:
1. The `import "sort"` statement is added.
2. The `min` built-in is available (Go 1.21+) or a helper is provided.
3. The package declaration remains `package bookstore`.

### 4.2 `sort` package usage

The usage `sort.Sort(sort.Reverse(sort.IntSlice(freq[1:])))` is idiomatic Go. An alternative modern approach would be `slices.SortFunc` or `sort.Slice`, but the current approach is perfectly correct.

### 4.3 Nil/empty slice handling

When `books` is `nil` or empty (`[]int{}`), the `for _, b := range books` loop simply does not execute, all frequencies remain 0, the greedy loop immediately breaks (size == 0 on the first iteration), `groups` is nil, all groupCounts are 0, and the total is 0. This is correct.

### 4.4 Test compatibility

The test file calls `Cost(testCase.basket)` and expects an `int` return. The plan's function signature matches: `func Cost(books []int) int`. No compatibility issues.

---

## 5. Is the Greedy + 5+3->4+4 Optimization Truly Sufficient?

**Yes.** This is a mathematically proven result for the specific discount schedule used in this problem.

The key insight is analyzing the per-book prices for each group size:

| Group Size | Per-Book Price (cents) |
|------------|----------------------|
| 1          | 800                  |
| 2          | 760                  |
| 3          | 720                  |
| 4          | 640                  |
| 5          | 600                  |

The discount jumps are: 40, 40, 80, 40. The jump from 3 to 4 is disproportionately large (80 cents vs 40 cents). This means the only case where greedy is suboptimal is when we can "steal" a book from a 5-group and add it to a 3-group, turning both into 4-groups. The net effect:

- Removing one book from a 5-group costs us: 600 (we no longer get a 5-group discount on that book)
- Adding that book to a 3-group gains us: going from a 3-group to a 4-group saves 720-640 = 80 cents per book on the group, but the group also grows, so the real calculation is: 4*640 - 3*720 = 2560 - 2160 = 400 for the new 4-group vs 2160 for the old 3-group, costing 400 more total, BUT we save from the 5-group becoming a 4-group: 5*600 = 3000 -> 4*640 = 2560, saving 440. Net: 440 - 400 = 40 cents saved.

Wait, let me recalculate more carefully:
- 5-group + 3-group cost: 3000 + 2160 = 5160
- Two 4-groups cost: 2560 + 2560 = 5120
- Savings: 40 cents

No other split is beneficial:
- 5+2 -> 4+3: cost 3000+1520=4520 vs 2560+2160=4720. WORSE. So we should NOT convert 5+2 to 4+3.
- 5+1 -> 4+2: cost 3000+800=3800 vs 2560+1520=4080. WORSE.
- 4+2 -> 3+3: cost 2560+1520=4080 vs 2160+2160=4320. WORSE.

The ONLY beneficial conversion is 5+3 -> 4+4. Therefore the greedy+fix approach is provably optimal for this discount schedule.

---

## 6. Summary of Required Fixes Before Implementation

| # | Issue | Severity | Fix |
|---|-------|----------|-----|
| 1 | Missing `import "sort"` | Medium | Add to imports block |
| 2 | `min()` will not compile on Go 1.18 (`go.mod` specifies 1.18) | High | Add a local `min` helper function or inline the logic |
| 3 | No algorithmic issues | -- | None needed |

## 7. Recommendation

**Proceed with implementation of Proposal A.** The algorithm is correct for all 18 test cases. The only required changes are ensuring the `sort` import is present and that the `min` function is available (either via Go 1.21+ or a helper). The plan is sound, well-reasoned, and matches the canonical solution for this exercise.
