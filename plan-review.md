# Plan Review: Book Store -- Branch 1 (Greedy + Adjustment)

## Reviewer: codex
## Date: 2026-02-16

---

## 1. Correctness: Manual Walkthrough of All 18 Test Cases

### Discount table used throughout

| Group size | Unit price | Discount | Group cost |
|-----------|-----------|---------|-----------|
| 1 | 800 | 0% | 800 |
| 2 | 1600 | 5% | 1520 |
| 3 | 2400 | 10% | 2160 |
| 4 | 3200 | 20% | 2560 |
| 5 | 4000 | 25% | 3000 |

Key observation for adjustment: `cost(5) + cost(3) = 3000 + 2160 = 5160`, while `2 * cost(4) = 5120`. So each 5+3 pair saves 40 cents when converted to 4+4.

### Simple cases (1-7): trivially correct

| # | Description | Frequencies (desc) | Greedy groups | After adjust | Cost | Expected | OK? |
|---|------------|-------------------|--------------|-------------|------|----------|-----|
| 1 | Single book | [1] | [1] | [1] | 800 | 800 | YES |
| 2 | Two same | [2] | [1,1] | [1,1] | 1600 | 1600 | YES |
| 3 | Empty | [] | [] | [] | 0 | 0 | YES |
| 4 | Two different | [1,1] | [2] | [2] | 1520 | 1520 | YES |
| 5 | Three different | [1,1,1] | [3] | [3] | 2160 | 2160 | YES |
| 6 | Four different | [1,1,1,1] | [4] | [4] | 2560 | 2560 | YES |
| 7 | Five different | [1,1,1,1,1] | [5] | [5] | 3000 | 3000 | YES |

### Case 8: "Two groups of four is cheaper than group of five plus group of three"
- Basket: [1,1,2,2,3,3,4,5]
- Frequencies: {1:2, 2:2, 3:2, 4:1, 5:1} -> sorted desc: [2,2,2,1,1]
- Greedy layer peeling:
  - Layer 1: 5 titles have freq >= 1 -> group of 5. Remaining: [1,1,1,0,0]
  - Layer 2: 3 titles have freq >= 2 -> group of 3. Remaining: [0,0,0,0,0]
- Greedy groups: one 5, one 3
- Adjustment: min(1,1) = 1 pair. Convert to two 4s.
- Cost: 2 * 2560 = 5120
- Expected: 5120. **YES**

### Case 9: "Two groups of four is cheaper than groups of five and three"
- Basket: [1,1,2,3,4,4,5,5]
- Frequencies: {1:2, 2:1, 3:1, 4:2, 5:2} -> sorted desc: [2,2,2,1,1]
- Same frequency profile as Case 8. Same result: **5120. YES**

### Case 10: "Group of four plus group of two is cheaper than two groups of three"
- Basket: [1,1,2,2,3,4]
- Frequencies: {1:2, 2:2, 3:1, 4:1} -> sorted desc: [2,2,1,1]
- Greedy layer peeling:
  - Layer 1: 4 titles have freq >= 1 -> group of 4. Remaining: [1,1,0,0]
  - Layer 2: 2 titles have freq >= 2 -> group of 2. Remaining: [0,0,0,0]
- Greedy groups: one 4, one 2
- No 5+3 pairs to adjust.
- Cost: 2560 + 1520 = 4080
- Expected: 4080. **YES**

Note: This test verifies that greedy naturally produces 4+2 rather than 3+3. Since greedy takes the largest group first, it produces [4,2] not [3,3]. The greedy approach is naturally correct here without any adjustment needed. Let me double-check: `cost(3) + cost(3) = 2160 + 2160 = 4320`, which is more expensive. The greedy approach correctly avoids this.

### Case 11: "Two each of first four books and one copy each of rest"
- Basket: [1,1,2,2,3,3,4,4,5]
- Frequencies: {1:2, 2:2, 3:2, 4:2, 5:1} -> sorted desc: [2,2,2,2,1]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [1,1,1,1,0]
  - Layer 2: 4 titles -> group of 4. Remaining: [0,0,0,0,0]
- Groups: one 5, one 4
- No 5+3 pair to adjust.
- Cost: 3000 + 2560 = 5560
- Expected: 5560. **YES**

### Case 12: "Two copies of each book"
- Basket: [1,1,2,2,3,3,4,4,5,5]
- Frequencies: [2,2,2,2,2]
- Greedy:
  - Layer 1: 5 titles -> group of 5
  - Layer 2: 5 titles -> group of 5
- Groups: two 5s
- No 3s to adjust with.
- Cost: 2 * 3000 = 6000
- Expected: 6000. **YES**

### Case 13: "Three copies of first book and two each of remaining"
- Basket: [1,1,1,2,2,3,3,4,4,5,5]
- Frequencies: {1:3, 2:2, 3:2, 4:2, 5:2} -> sorted desc: [3,2,2,2,2]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [2,1,1,1,1]
  - Layer 2: 5 titles -> group of 5. Remaining: [1,0,0,0,0]
  - Layer 3: 1 title -> group of 1. Remaining: [0,0,0,0,0]
- Groups: two 5s, one 1
- No 3s. No adjustment.
- Cost: 2 * 3000 + 800 = 6800
- Expected: 6800. **YES**

### Case 14: "Three each of first two books and two each of remaining books"
- Basket: [1,1,1,2,2,2,3,3,4,4,5,5]
- Frequencies: {1:3, 2:3, 3:2, 4:2, 5:2} -> sorted desc: [3,3,2,2,2]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [2,2,1,1,1]
  - Layer 2: 5 titles -> group of 5. Remaining: [1,1,0,0,0]
  - Layer 3: 2 titles -> group of 2. Remaining: [0,0,0,0,0]
- Groups: two 5s, one 2
- No 3s. No adjustment.
- Cost: 2 * 3000 + 1520 = 7520
- Expected: 7520. **YES**

### Case 15: "Four groups of four are cheaper than two groups each of five and three" (CRITICAL)
- Basket: [1,1,2,2,3,3,4,5,1,1,2,2,3,3,4,5]
- Frequencies: {1:4, 2:4, 3:4, 4:2, 5:2} -> sorted desc: [4,4,4,2,2]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [3,3,3,1,1]
  - Layer 2: 5 titles -> group of 5. Remaining: [2,2,2,0,0]
  - Layer 3: 3 titles -> group of 3. Remaining: [1,1,1,0,0]
  - Layer 4: 3 titles -> group of 3. Remaining: [0,0,0,0,0]
- Groups: two 5s, two 3s
- Adjustment: min(2,2) = 2 pairs. Convert to four 4s.
- Cost: 4 * 2560 = 10240
- Expected: 10240. **YES**

### Case 16: "Groups of four created properly even with more groups of three than five" (CRITICAL)
- Basket: [1,1,1,1,1,1, 2,2,2,2,2,2, 3,3,3,3,3,3, 4,4, 5,5]
- Frequencies: {1:6, 2:6, 3:6, 4:2, 5:2} -> sorted desc: [6,6,6,2,2]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [5,5,5,1,1]
  - Layer 2: 5 titles -> group of 5. Remaining: [4,4,4,0,0]
  - Layer 3: 3 titles -> group of 3. Remaining: [3,3,3,0,0]
  - Layer 4: 3 titles -> group of 3. Remaining: [2,2,2,0,0]
  - Layer 5: 3 titles -> group of 3. Remaining: [1,1,1,0,0]
  - Layer 6: 3 titles -> group of 3. Remaining: [0,0,0,0,0]
- Groups: two 5s, four 3s
- Adjustment: min(2,4) = 2 pairs. groups[5]=0, groups[3]=2, groups[4]=4.
- Cost: 4 * 2560 + 2 * 2160 = 10240 + 4320 = 14560
- Expected: 14560. **YES**

### Case 17: "One group of one and four is cheaper than one group of two and three"
- Basket: [1,1,2,3,4]
- Frequencies: {1:2, 2:1, 3:1, 4:1} -> sorted desc: [2,1,1,1]
- Greedy:
  - Layer 1: 4 titles -> group of 4. Remaining: [1,0,0,0]
  - Layer 2: 1 title -> group of 1. Remaining: [0,0,0,0]
- Groups: one 4, one 1
- No 5+3 pairs.
- Cost: 2560 + 800 = 3360
- Expected: 3360. **YES**

Note: Alternative would be 3+2 = 2160 + 1520 = 3680, which is more expensive. Greedy correctly picks the larger group first.

### Case 18: "One group of one and two plus three groups of four is cheaper than one group of each size" (CRITICAL)
- Basket: [1, 2,2, 3,3,3, 4,4,4,4, 5,5,5,5,5]
- Frequencies: {1:1, 2:2, 3:3, 4:4, 5:5} -> sorted desc: [5,4,3,2,1]
- Greedy:
  - Layer 1: 5 titles -> group of 5. Remaining: [4,3,2,1,0]
  - Layer 2: 4 titles -> group of 4. Remaining: [3,2,1,0,0]
  - Layer 3: 3 titles -> group of 3. Remaining: [2,1,0,0,0]
  - Layer 4: 2 titles -> group of 2. Remaining: [1,0,0,0,0]
  - Layer 5: 1 title -> group of 1. Remaining: [0,0,0,0,0]
- Groups: one 5, one 4, one 3, one 2, one 1
- Adjustment: min(1,1) = 1 pair. groups[5]=0, groups[3]=0, groups[4]=3.
- Cost: 3 * 2560 + 1520 + 800 = 7680 + 1520 + 800 = 10000
- Expected: 10000. **YES**

Verification of alternative (no adjustment): 3000 + 2560 + 2160 + 1520 + 800 = 10040. So the adjustment saves 40.

**Result: All 18 test cases produce the correct answer.**

---

## 2. Edge Cases

| Edge case | Handling in plan | Verdict |
|-----------|-----------------|---------|
| Empty basket (`[]`) | Step 2 says "If books is empty, return 0" | Correct (test case 3) |
| Single book | Frequencies [1], one group of 1, cost = 800 | Correct (test case 1) |
| All same books (e.g., [2,2]) | Frequencies [2], two groups of 1, cost = 1600 | Correct (test case 2) |
| All five distinct books once | Frequencies [1,1,1,1,1], one group of 5 | Correct (test case 7) |
| Large basket with uneven distribution | Test case 16 and 18 cover this | Correct (verified above) |

No missing edge cases identified.

---

## 3. Algorithm Soundness: Is 5+3 -> 4+4 the Only Needed Adjustment?

### Mathematical analysis

The per-book cost for each group size:
- Size 1: 800 / 1 = 800.0 per book
- Size 2: 1520 / 2 = 760.0 per book
- Size 3: 2160 / 3 = 720.0 per book
- Size 4: 2560 / 4 = 640.0 per book
- Size 5: 3000 / 5 = 600.0 per book

The greedy algorithm (largest groups first) is driven by the fact that larger groups have lower per-book cost. The only situation where this fails is when combining groups can yield a cheaper result. Let me check all possible pair conversions:

| Conversion | Before cost | After cost | Savings |
|-----------|------------|------------|---------|
| 5+1 -> 3+3 | 3800 | 4320 | -520 (worse) |
| 5+1 -> 4+2 | 3800 | 4080 | -280 (worse) |
| 5+2 -> 4+3 | 4520 | 4720 | -200 (worse) |
| 5+3 -> 4+4 | 5160 | 5120 | +40 (better!) |
| 4+1 -> 3+2 | 3360 | 3680 | -320 (worse) |
| 4+2 -> 3+3 | 4080 | 4320 | -240 (worse) |
| 3+1 -> 2+2 | 2960 | 3040 | -80 (worse) |
| 2+1 -> 1+1+1 | 2320 | 2400 | -80 (worse) |

The ONLY profitable conversion is 5+3 -> 4+4 (saving 40 per pair).

Now, could there be multi-step conversions that help? For example:
- 5+5+1+1 -> could rearranging three or more groups help?
  - 5+5+1+1 = 7600. vs 4+4+2+2 = 8160. No.
  - 5+5+1+1 = 7600. vs 5+4+2+1 = 7880. No.
  - All alternatives are worse.
- 5+3+3 -> 4+4+3 (one conversion) = saves 40. Could we do 4+4+3 -> 4+3+4? Same thing. No further improvement.

The key insight: once all 5+3 pairs are eliminated, no further conversions are beneficial. This is because:
1. Only the 5+3 -> 4+4 conversion is profitable (shown above).
2. After conversion, the remaining groups contain no 5+3 pairs, so no further improvement is possible.
3. No multi-group conversions involving more than two groups produce savings beyond what pairwise 5+3 conversion achieves.

**Verdict: The 5+3 -> 4+4 adjustment is provably the only correction needed.**

---

## 4. Implementation Concerns

### 4.1 Greedy group formation (layer peeling)

The plan describes "peeling layers off the frequency histogram." This is correct. With frequencies sorted in descending order, each layer's width equals the count of frequencies at or above that layer. Concretely:

```
sorted_freqs = [f1, f2, ..., fn] (descending)
Layer k: group size = number of fi >= k
```

An equivalent and simpler implementation: after sorting frequencies descending, the group sizes are the differences in a staircase. Specifically, for sorted frequencies `[f1, f2, ..., fn]`:
- Group at layer k has size = count of i where fi >= k

This can be computed by iterating through the sorted frequencies and building groups. The plan's Step 2 item 4 mentions this correctly.

**Potential issue**: The plan says "iterate frequencies from high to low" which could be ambiguous. It should iterate layer by layer (level 1, 2, 3...) up to the maximum frequency. However, a cleaner equivalent is: for each index i in the sorted (desc) array, the contribution to groups is that position i is included in layers 1 through f[i]. This means groups can be computed by looking at consecutive frequency differences.

Actually, the simplest correct implementation is:

```go
// sorted desc: freqs = [5, 4, 3, 2, 1]
// layer 1: all 5 have freq >= 1, group size = 5
// layer 2: first 4 have freq >= 2, group size = 4
// etc.
```

This is equivalent to: the group sizes are exactly the sorted frequencies read as a "transposed" histogram. For `n` distinct titles with sorted frequencies `[f1 >= f2 >= ... >= fn]`, the group sizes are: `f1` groups total, where group at layer `k` has size = number of `fi >= k`.

The plan gets this right conceptually. Implementation should be straightforward.

### 4.2 Integer arithmetic

The plan uses integer arithmetic throughout (cents, integer percentages). This is correct and avoids floating-point issues. The discount computation `normalPrice * discountPercent / 100` works correctly in integer arithmetic for the given discount percentages (0, 5, 10, 20, 25) since all resulting products are divisible by 100.

### 4.3 Book title range

The test cases use book titles 1-5. The plan uses a map or array for counting. Either works. If using an array, it should be indexed 1-5 (or 0-4 with offset). No issue here.

### 4.4 The `min` function

The plan uses `min(groups[5], groups[3])` in the adjustment step. Go 1.21+ has a built-in `min` function. If targeting an older Go version, this would need to be implemented manually. Check the `go.mod` for the Go version.

### 4.5 Complexity

The plan claims O(n) time. More precisely, it is O(n + m * max_freq) where m is the number of distinct titles (at most 5) and max_freq is the highest frequency. Since m <= 5, this is effectively O(n). Correct.

### 4.6 Discount table indexing

The plan defines discounts as `[0, 0, 5, 10, 20, 25]` indexed by group size (0-5). The reference example uses `[0, 5, 10, 20, 25]` indexed 0-4 (with `discountTiers[groupSize-1]`). Either approach works, but the plan's zero-indexed-by-size approach is cleaner and less error-prone.

### 4.7 Missing concern: nil vs empty slice

The plan handles `len(books) == 0` returning 0. In Go, `len(nil) == 0`, so a nil input is also handled. Good.

---

## 5. Minor Suggestions

1. **Clarity on layer peeling implementation**: The detailed implementation plan could be more explicit about how to compute group sizes from sorted frequencies. A concrete code sketch would reduce implementation ambiguity. For example:

   ```
   Sort frequencies descending.
   prevWidth = 0
   For layer from maxFreq down to 1:
       width = count of freqs >= layer
       if width != prevWidth, record group
   ```

   Or more simply: the transpose approach where the group at each layer has size equal to the number of nonzero frequencies at that layer.

2. **Consider using an array `[6]int` for frequency counts** (indexed 1-5) rather than a map, since book titles are known to be in range 1-5. This is simpler and faster.

3. **Consider using an array `[6]int` for group counts** (indexed by group size 1-5) as the plan already describes in Step 2 item 5.

---

## 6. Verdict

**APPROVE**

The Greedy + Adjustment approach (Branch 1) is correct, efficient, and well-suited for this problem. All 18 test cases have been manually verified to produce the correct output. The 5+3 -> 4+4 adjustment is mathematically proven to be the only correction needed for this discount structure. The implementation plan is sound with no significant concerns. The minor suggestions above are optional improvements for clarity but do not affect correctness.

The plan is ready for implementation.
