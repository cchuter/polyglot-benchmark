# Alphametics Implementation Review

**Reviewer**: Challenger agent
**File reviewed**: `go/exercises/practice/alphametics/alphametics.go`
**Reference**: `.meta/example.go`
**Status**: Review complete — **PASS**

---

## Summary

The implementation is based on the reference solution with one significant improvement: the leading-zero check now validates ALL words, not just the answer. The brute-force permutation approach is correct for all 10 test cases.

---

## 1. Correctness: All 10 Test Cases

| # | Description | Verdict | Notes |
|---|-------------|---------|-------|
| 1 | Puzzle with three letters (`I + BB == ILL`) | PASS | B=9, I=1, L=0; L is not a leading digit |
| 2 | Unique value constraint (`A == B`) | PASS | Permutations always assign different digits → no solution |
| 3 | Leading zero invalid (`ACA + DD == BD`) | PASS | `isPuzzleSolution` returns false at column 2 (answer has no digit there); correctly returns "no solution" |
| 4 | Two-digit carry (`A+A+...+B == BCC`) | PASS | A=9, B=1, C=0; C is not a leading digit |
| 5 | Four letters (`AS + A == MOM`) | PASS | All leading digits non-zero |
| 6 | Six letters (`NO + NO + TOO == LATE`) | PASS | A=0 is not a leading digit |
| 7 | Seven letters (`HE + SEES + THE == LIGHT`) | PASS | All leading digits non-zero |
| 8 | Eight letters (`SEND + MORE == MONEY`) | PASS | S=9, M=1 |
| 9 | Ten letters (OFFENSE puzzle) | PASS | Full digit allocation |
| 10 | Ten letters, 199 addends (FORTRESSES) | PASS | Correct despite performance cost |

---

## 2. Key Improvement Over Reference Solution: Leading-Zero Check

**Reference** (`example.go` lines 79-83) — only checks the answer word:
```go
r := p.vDigits[len(p.vDigits)-1][p.maxDigits-1]
if p.letterValues[r-1] == 0 {
    continue
}
```

**Implementation** (`alphametics.go` lines 75-94) — checks ALL words:
```go
valid := true
for _, word := range p.vDigits {
    for d := len(word) - 1; d >= 0; d-- {
        if word[d] != 0 {
            if d > 0 && p.letterValues[word[d]-1] == 0 {
                valid = false
            }
            break
        }
    }
    if !valid {
        break
    }
}
```

**Analysis**: This correctly:
- Iterates ALL words (addends and answer), not just the answer
- Finds the actual leading digit by skipping padding zeros (highest non-zero position)
- Only flags leading zeros on multi-digit words (`d > 0`), allowing single-letter words like "A" or "I" to be assigned digit 0
- Short-circuits on first violation

This eliminates the latent bug in the reference solution where addend leading zeros were not validated.

---

## 3. Other Differences from Reference

| Aspect | Reference | Implementation | Assessment |
|--------|-----------|----------------|------------|
| Leading-zero check | Answer only | All words | **Improvement** |
| `nLetters` init | Implicit zero | Explicit `p.nLetters = 0` | Equivalent (struct zero-values) |
| Perm variable name | `p` | `p2` | Avoids shadowing; minor improvement |
| Comments | More verbose | Leaner | Cosmetic |

---

## 4. Edge Case Analysis

- **Leading zeros on addends**: Fixed (checks all words). Single-letter words correctly exempted via `d > 0` guard.
- **Unique digit constraint ("A == B")**: Permutations always assign distinct values to distinct letters → no solution found correctly.
- **No-solution cases**: Exhaustive search returns `errors.New("no solution")` after all permutations checked.
- **Invalid input (non-uppercase)**: `parsePuzzle` returns nil → `Solve` returns "invalid puzzle" error.
- **Answer shorter than longest addend**: `isPuzzleSolution` returns false at high columns where answer has no digit (r=0 check on line 121).

---

## 5. Permutation Generator

Faithful port of Python `itertools.permutations`. Verified:
- Capacity calculation: P(n, r) = n! / (n-r)! — correct
- Index/cycle manipulation matches the Python algorithm
- Each permutation is correctly copied into a new slice before appending

**Performance note**: Generates all P(10, r) permutations upfront in memory. For r=10: ~3.6M permutations × ~104 bytes ≈ 288MB. Acceptable for exercise context.

---

## 6. Verdict

**PASS** — The implementation is correct for all 10 test cases and improves upon the reference solution's leading-zero validation. No bugs found. Code is clean, well-structured Go. Ready for testing.
