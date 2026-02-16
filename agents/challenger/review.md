# Challenger Review: ParseOctal Implementation

## Verdict: PASS

The implementation is correct, follows the plan exactly, and should pass all test cases.

## Algorithm Correctness

The octal-to-decimal conversion algorithm is correct:

- `num = num<<3 + int64(digit-'0')` works because in Go, `<<` (precedence group 5) binds tighter than `+` (precedence group 4), so this evaluates as `(num << 3) + int64(digit-'0')`, which is equivalent to `num*8 + digitValue`.

### Manual verification against test cases:

| Input       | Expected | Trace                                                | Result |
|-------------|----------|------------------------------------------------------|--------|
| "1"         | 1        | 0<<3+1 = 1                                          | PASS   |
| "10"        | 8        | 0<<3+1=1, 1<<3+0=8                                  | PASS   |
| "1234567"   | 342391   | 0->1->10->83->668->5349->42798->342391               | PASS   |
| "carrot"    | error    | 'c' > '7', returns error                            | PASS   |
| "35682"     | error    | '3','5','6' ok, '8' > '7', returns error             | PASS   |

## Error Handling

- Invalid digits (8, 9) are correctly rejected: `digit > '7'` catches them.
- Non-digit characters (letters, symbols) are caught: either `digit < '0'` or `digit > '7'`.
- Error message uses `fmt.Errorf` with the offending rune, which is informative.
- Returns `(0, error)` on invalid input, matching expected behavior.

## Edge Cases

| Case               | Behavior                    | Assessment |
|--------------------|-----------------------------|------------|
| Empty string `""`  | Returns `(0, nil)`          | Acceptable (not tested) |
| Digits 8-9         | Returns error               | Correct    |
| Non-digit chars    | Returns error               | Correct    |
| Non-ASCII (e.g. emoji) | Returns error (rune > '7') | Correct    |
| Leading zeros      | Handled correctly (0<<3=0)  | Correct    |

**Note:** Empty string returning `(0, nil)` is a design choice. One could argue it should be an error, but the test suite does not test this case, and returning 0 for an empty octal string is a reasonable interpretation (no digits = value 0).

## No Built-in Conversion Functions

Confirmed: No `strconv` or other built-in conversion functions are used. Only `fmt` is imported, solely for `fmt.Errorf`.

## Adherence to Plan

The implementation matches the plan's specified code exactly, including:
- Bit shifting (`<<3`) instead of multiplication
- `fmt.Errorf` for error creation
- Range over string for rune iteration
- Early return on invalid input

## Potential Concerns (informational, not blocking)

1. **No int64 overflow protection** for extremely large octal strings. Not required by tests.
2. **Empty string** returns `(0, nil)` rather than an error. Not tested, so acceptable.

## Conclusion

The implementation is clean, correct, and minimal. It should pass all 5 test cases and the benchmark. No changes needed.
