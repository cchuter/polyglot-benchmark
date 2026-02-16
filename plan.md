# Implementation Plan: polyglot-go-counter

## Overview

Write a comprehensive test suite in `counter_test.go` that validates the `Counter` interface. The tests must pass for Impl4 (correct) and fail for Impl1, Impl2, and Impl3 (each buggy in a different way).

## File to Modify

- `go/exercises/practice/counter/counter_test.go` — Replace stub with full test suite

## File That Stays As-Is

- `go/exercises/practice/counter/counter.go` — Package declaration only (no solution code needed; this is a test-design exercise)

## Test Design

### Test Functions

Each test function uses `makeCounter()` and calls `AddString()` zero or more times, then asserts `Lines()`, `Letters()`, and `Characters()`.

#### 1. `TestNoAddString`
- Create counter, don't call AddString
- Expect: Lines=0, Letters=0, Characters=0
- Catches: nothing (all impls handle this)
- Purpose: baseline sanity check

#### 2. `TestEmptyString`
- AddString("")
- Expect: Lines=0, Letters=0, Characters=0
- Catches: nothing (all impls handle this)
- Purpose: edge case

#### 3. `TestSimpleASCIINoNewline`
- AddString("hello")
- Expect: Lines=1, Letters=5, Characters=5
- **Catches Impl1**: Impl1 returns Lines=0 (no `\n` found, doesn't add 1)

#### 4. `TestASCIIWithNewlineInMiddle`
- AddString("Hello\nworld!")
- Expect: Lines=2, Letters=10, Characters=12
- **Catches Impl1**: returns Lines=1 (counts only the `\n`)

#### 5. `TestStringEndingWithNewline`
- AddString("hello\n")
- Expect: Lines=1, Letters=5, Characters=6
- Purpose: verify trailing newline handling (all correct impls agree: 1 line)

#### 6. `TestUnicodeLetters`
- AddString("здравствуй, мир\n")
- Expect: Lines=1, Letters=13, Characters=16
- **Catches Impl2**: Impl2 counts 0 letters (ASCII-only check)
- **Catches Impl3**: Impl3 iterates bytes, gets wrong character count (30 bytes vs 16 runes) and wrong letter count

#### 7. `TestMultipleAddStrings`
- AddString("hello\n"), then AddString("world")
- Expect: Lines=2, Letters=10, Characters=11
- Purpose: verify accumulation works across multiple calls
- **Catches Impl1**: returns Lines=1
- Note: "hello\n" = 6 chars, "world" = 5 chars, total = 11

#### 8. `TestOnlyNewlines`
- AddString("\n\n\n")
- Expect: Lines=3, Letters=0, Characters=3
- Purpose: edge case with only newlines

#### 9. `TestMixedContent`
- AddString("abc 123!@#\ndef")
- Expect: Lines=2, Letters=6, Characters=14
- Purpose: non-letter ASCII characters, verify letters vs characters distinction

### Helper Pattern

Use a helper function `assertCounts(t, counter, expectedLines, expectedLetters, expectedChars)` to reduce boilerplate and provide clear error messages.

## Approach and Ordering

1. Write `counter_test.go` with all test functions
2. Verify: `COUNTER_IMPL=4 go test` passes
3. Verify: `COUNTER_IMPL=1 go test` fails (line counting)
4. Verify: `COUNTER_IMPL=2 go test` fails (Unicode letters)
5. Verify: `COUNTER_IMPL=3 go test` fails (byte iteration)

## Rationale

- Tests are modeled after the example in `.meta/example.go` but with correct expected values
- The helper function follows Go testing idioms (using `t.Helper()`)
- Each buggy impl is caught by at least one test:
  - Impl1 caught by tests 3, 4, 7 (line counting with no trailing newline)
  - Impl2 caught by test 6 (Unicode letters)
  - Impl3 caught by test 6 (byte iteration on multi-byte chars)
