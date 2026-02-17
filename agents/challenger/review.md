# Code Review: book_store.go

## Verdict: PASS — No issues found

The implementation is correct, clean, and follows the plan.

---

## 1. Correctness: Greedy grouping + 5+3→4+4 optimization

The algorithm correctly:
- Counts frequencies of each book into a 5-element slice (0-indexed)
- Sorts frequencies descending, then greedily forms groups by taking one copy from each non-zero frequency
- Counts group sizes in `groupCount[size]`
- Applies the 5+3→4+4 optimization post-hoc

This matches the "Proposal A — Greedy with 5→3 to 4→4 Optimization" from the plan. **PASS**

## 2. Edge cases

- **Empty basket (`[]int{}`):** The `range books` loop is skipped. `freq` stays all zeros. In the greedy loop, `freq[0] == 0` on the first check → `break`. Total = 0. **PASS**
- **Single book (`[]int{1}`):** freq = [1,0,0,0,0]. One group of size 1. Total = 1×1×800 = 800. **PASS**
- **All duplicates (`[]int{2,2}`):** freq = [0,2,0,0,0]. Sorted → [2,0,0,0,0]. Two iterations produce two groups of size 1. Total = 2×1×800 = 1600. **PASS**

## 3. Go 1.18 compatibility

- **Custom `minInt` function** used instead of built-in `min` (Go 1.21+). **PASS**
- **`sort.Sort(sort.Reverse(sort.IntSlice(freq)))`** — uses standard library APIs available since Go 1.0. **PASS**
- No generics, no `slices` package, no Go 1.21+ features. **PASS**

## 4. Sort usage

`freq` is `make([]int, 5)` — a slice. `sort.IntSlice(freq)` wraps the slice header, so `sort.Sort` modifies the underlying array data in-place. The sorted result is visible via the original `freq` variable. **PASS**

## 5. Index correctness (books 1-5, array 0-4)

- `freq := make([]int, 5)` — indices 0-4
- `freq[b-1]++` — maps book 1→index 0, book 5→index 4
- Greedy loop iterates `i := 0; i < 5` — correctly covers all 5 positions
- No off-by-one errors. **PASS**

## 6. Discount table correctness

```
price := [6]int{0, 800, 760, 720, 640, 600}
```

Verification (800 cents base, discounts 0%/5%/10%/20%/25%):
| Size | Discount | Price = 800 × (1 - discount) | Table value |
|------|----------|------------------------------|-------------|
| 1    | 0%       | 800                          | 800 ✓       |
| 2    | 5%       | 760                          | 760 ✓       |
| 3    | 10%      | 720                          | 720 ✓       |
| 4    | 20%      | 640                          | 640 ✓       |
| 5    | 25%      | 600                          | 600 ✓       |

**PASS**

## 7. Total calculation

```go
total += groupCount[size] * size * price[size]
```

For `groupCount[s]` groups of size `s`, each group costs `s × price[s]` (s books at the discounted per-book price). Multiplied by count. **PASS**

## 8. Trace: "Two groups of four is cheaper than group of five plus group of three"

**Input:** basket = [1,1,2,2,3,3,4,5], expected = 5120

**Frequency:** freq = [2,2,2,1,1]

**Greedy iteration 1:** sort desc → [2,2,2,1,1]. All 5 slots > 0 → size=5. freq→[1,1,1,0,0]. groupCount[5]=1

**Greedy iteration 2:** sort desc → [1,1,1,0,0]. 3 slots > 0 → size=3. freq→[0,0,0,0,0]. groupCount[3]=1

**Greedy iteration 3:** freq[0]=0 → break

**After greedy:** groupCount[5]=1, groupCount[3]=1

**Optimization:** swaps = minInt(1,1) = 1. groupCount[5]→0, groupCount[3]→0, groupCount[4]→2

**Total:** 2 × 4 × 640 = **5120** ✓

## 9. Additional trace: Complex case with more 3-groups than 5-groups

**Input:** basket = [1,1,1,1,1,1, 2,2,2,2,2,2, 3,3,3,3,3,3, 4,4, 5,5], expected = 14560

**Frequency:** freq = [6,6,6,2,2]

**Greedy:** 2 groups of 5, then 4 groups of 3. groupCount[5]=2, groupCount[3]=4

**Optimization:** swaps = minInt(2,4) = 2. groupCount[5]→0, groupCount[3]→2, groupCount[4]→4

**Total:** 2×3×720 + 4×4×640 = 4320 + 10240 = **14560** ✓

## 10. Differences from plan pseudo-code (all acceptable)

| Aspect | Plan | Implementation | Assessment |
|--------|------|----------------|------------|
| Indexing | 1-based (`freq[b]`) | 0-based (`freq[b-1]`) | Equivalent |
| Price | Computed from discount% | Pre-computed table | Equivalent, cleaner |
| Group counting | Accumulate slice, then count | Direct `groupCount[size]++` | More efficient |
| `min` | Built-in `min` | Custom `minInt` | Go 1.18 compatible |

---

## Conclusion

The implementation is correct, efficient, Go 1.18 compatible, and follows the selected plan (Proposal A). All edge cases are handled. The 5+3→4+4 optimization is correctly applied. Ready for testing.
