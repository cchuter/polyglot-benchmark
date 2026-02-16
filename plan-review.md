# Plan Review: Beer Song

## Verdict: FAIL

The plan contains a critical bug in the `Verses()` function that will cause test failures for all multi-verse test cases.

---

## Detailed Findings

### 1. Verse() function -- PASS

The `Verse()` function is correct for all test cases:

| Test Case | Verse | Plan Output | Expected Output | Match? |
|-----------|-------|-------------|-----------------|--------|
| verse 8 (general) | `Verse(8)` | `"8 bottles of beer on the wall, 8 bottles of beer.\nTake one down and pass it around, 7 bottles of beer on the wall.\n"` | `verse8` | Yes |
| verse 3 (general) | `Verse(3)` | `"3 bottles of beer on the wall, 3 bottles of beer.\nTake one down and pass it around, 2 bottles of beer on the wall.\n"` | `verse3` | Yes |
| verse 2 (special) | `Verse(2)` | `"2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"` | `verse2` | Yes |
| verse 1 (special) | `Verse(1)` | `"1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"` | `verse1` | Yes |
| verse 0 (special) | `Verse(0)` | `"No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"` | `verse0` | Yes |
| invalid 104 | `Verse(104)` | `"", error` | `"", error` | Yes |

All string literals match the test expectations character-for-character. The switch cases (0, 1, 2, default) correctly cover all the distinct verse patterns. The error handling for out-of-range values is correct.

### 2. Verses() function -- FAIL (Critical Bug)

**The `strings.Join` approach produces different output than the reference's `bytes.Buffer` approach.** This will cause test failures for `Verses(8,6)`, `Verses(7,5)`, and indirectly `Song()`.

#### Root Cause

Each verse string returned by `Verse()` already ends with `\n`. The reference solution appends an *additional* `\n` after every verse (including the last one):

```go
// Reference (bytes.Buffer):
for i := start; i >= stop; i-- {
    v, _ := Verse(i)
    buff.WriteString(v)       // writes "...wall.\n"
    buff.WriteString("\n")    // writes "\n"
}
// Result: "verse8\n\nverse7\n\nverse6\n\n"
//                                       ^^ trailing \n after last verse
```

The plan uses `strings.Join`, which only inserts the separator *between* elements, not after the last one:

```go
// Plan (strings.Join):
parts = ["verse8\n", "verse7\n", "verse6\n"]
strings.Join(parts, "\n")
// Result: "verse8\n\nverse7\n\nverse6\n"
//                                      ^ NO trailing \n after last verse
```

#### Concrete Example: `Verses(8, 6)`

**Plan output:**
```
"8 bottles...wall.\n\n7 bottles...wall.\n\n6 bottles...wall.\n"
```
(347 characters, ends with single `\n`)

**Test expectation (`verses86`):**
```
"8 bottles...wall.\n\n7 bottles...wall.\n\n6 bottles...wall.\n\n"
```
(348 characters, ends with `\n\n`)

The plan's output is missing the final trailing `\n`. The test performs an exact string comparison (`actualVerse != tc.expectedVerse`), so this one-character difference will cause a failure.

#### Fix

Either:

**(a)** Switch to the `bytes.Buffer` approach from the reference:
```go
var buff bytes.Buffer
for i := start; i >= stop; i-- {
    v, _ := Verse(i)
    buff.WriteString(v)
    buff.WriteString("\n")
}
return buff.String(), nil
```

**(b)** Keep `strings.Join` but append a trailing `\n`:
```go
return strings.Join(parts, "\n") + "\n", nil
```

### 3. Song() function -- CONDITIONAL FAIL

`Song()` delegates to `Verses(99, 0)`. The `TestEntireSong` test compares `Song()` against `Verses(99, 0)` directly, so if `Verses` is internally consistent, this test passes regardless of the trailing newline issue. However, if the `Verses` bug is fixed to match the expected output, `Song()` will automatically be correct.

As currently planned, `Song()` is self-consistent but will produce output with the wrong trailing newline (missing the final `\n`). This does not cause the `TestEntireSong` test to fail because the test compares `Song()` against `Verses(99,0)` rather than against a hardcoded expected string. Still, fixing `Verses()` is necessary for the multi-verse tests.

### 4. Error Handling -- PASS

The plan's error validation in `Verses()` matches the reference exactly:
- `start < 0 || start > 99` returns an error (covers test case: `Verses(109, 5)`)
- `stop < 0 || stop > 99` returns an error (covers test case: `Verses(99, -20)`)
- `start < stop` returns an error (covers test case: `Verses(8, 14)`)

### 5. Edge Cases and Formatting -- PASS (aside from Verses bug)

- All verse string literals use `\n` line endings, not `\r\n`
- No trailing whitespace issues in verse strings
- Capitalization is correct ("No more" for verse 0, "no more" in verse 1's second line)
- Singular/plural ("bottle" vs "bottles") is correct
- "Take it down" (verse 1) vs "Take one down" (verses 2+) is correct

---

## Summary

| Function | Verdict | Notes |
|----------|---------|-------|
| `Verse()` | PASS | All 6 test cases will pass |
| `Verses()` | **FAIL** | Missing trailing `\n` due to `strings.Join` behavior |
| `Song()` | PASS* | Self-consistent test will pass, but output differs from expected format |

**Overall Verdict: FAIL** -- The plan requires a one-line fix in `Verses()` to append a trailing newline after the joined string (or switch to the `bytes.Buffer` approach from the reference). Without this fix, the `TestSeveralVerses` test cases for `Verses(8,6)` and `Verses(7,5)` will fail.
