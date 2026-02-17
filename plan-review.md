# Plan Review

## Review Method
Manual analysis of regex patterns against all 22 test cases (no codex agent available).

## Regex Analysis

### `vowel = ^([aeiou]|y[^aeiou]|xr)[a-z]*`

| Test | Input | Expected | Match? | Result |
|------|-------|----------|--------|--------|
| word beginning with a | apple | appleay | Yes, `a` matches `[aeiou]` | appleay ✓ |
| word beginning with e | ear | earay | Yes | earay ✓ |
| word beginning with i | igloo | iglooay | Yes | iglooay ✓ |
| word beginning with o | object | objectay | Yes | objectay ✓ |
| word beginning with u | under | underay | Yes | underay ✓ |
| vowel + qu | equal | equalay | Yes, `e` matches | equalay ✓ |
| word beginning with xr | xray | xrayay | Yes, `xr` matches | xrayay ✓ |
| word beginning with yt | yttria | yttriaay | Yes, `yt` → `y[^aeiou]` where `t` is not a vowel | yttriaay ✓ |

**Issue found**: The `yt` case relies on `y[^aeiou]` which matches `yt` correctly. But what about `yellow`? `ye` → `y` followed by `e` which IS a vowel, so `y[^aeiou]` does NOT match. Good — `yellow` should NOT be treated as vowel-start.

### `containsy = ^([^aeiou]+)y([a-z]*)`

| Test | Input | Expected | Match? |
|------|-------|----------|--------|
| y treated as vowel | rhythm | ythmrhay | `rh` matches `[^aeiou]+`, then `y` | ✓ |
| y as second letter | my | ymay | `m` matches `[^aeiou]+`, then `y` | ✓ |
| yellow | yellow | ellowyay | `y` matches `[^aeiou]+`... but then expects `y` next. `y` is consonant at start so `[^aeiou]+` matches `y`, then needs literal `y` at position 1 which is `e`. NO match. | Falls through to consonant rule ✓ |

Wait — `yellow`: `y` is not in `[aeiou]`, so `[^aeiou]+` greedily matches `y`, then regex needs literal `y` but finds `e`. No match. Falls to cons rule: `[^aeiou]+` matches `y`, remainder is `ellow`. Result: `ellow` + `y` + `ay` = `ellowyay` ✓

### `cons = ^([^aeiou]?qu|[^aeiou]+)([a-z]*)`

| Test | Input | Expected | Match? |
|------|-------|----------|--------|
| pig | pig | igpay | `[^aeiou]+` matches `p`, remainder `ig` | igpay ✓ |
| koala | koala | oalakay | `[^aeiou]+` matches `k`, remainder `oala` | oalakay ✓ |
| xenon | xenon | enonxay | `[^aeiou]+` matches `x`, remainder `enon` | enonxay ✓ |
| qat | qat | atqay | `[^aeiou]+` matches `q`, remainder `at` | atqay ✓ |
| chair | chair | airchay | `[^aeiou]+` matches `ch`, remainder `air` | airchay ✓ |
| queen | queen | eenquay | `[^aeiou]?qu` matches `qu` (with `[^aeiou]?` matching 0), remainder `een` | eenquay ✓ |
| square | square | aresquay | `[^aeiou]?qu` matches `squ`, remainder `are` | aresquay ✓ |
| therapy | therapy | erapythay | `[^aeiou]+` matches `th`, remainder `erapy` | erapythay ✓ |
| thrush | thrush | ushthray | `[^aeiou]+` matches `thr`, remainder `ush` | ushthray ✓ |
| school | school | oolschay | `[^aeiou]+` matches `sch`, remainder `ool` | oolschay ✓ |

**Wait — potential issue with `cons` regex**: `[^aeiou]?qu` only allows 0 or 1 consonant before `qu`. What if there were 2 consonants before `qu`? Not in the test cases but let's verify correctness of the alternation order.

The regex is `^([^aeiou]?qu|[^aeiou]+)`. For `square`: tries `[^aeiou]?qu` first — `s` matches `[^aeiou]?`, then `qu` matches. Captures `squ`. ✓

For `queen`: tries `[^aeiou]?qu` first — `[^aeiou]?` matches 0 chars, then `qu` matches. Captures `qu`. ✓

For `qat`: tries `[^aeiou]?qu` first — `[^aeiou]?` matches 0 chars, then `qu` needs `qu` but finds `qa`. Fails. Falls to `[^aeiou]+` which matches `q`. Captures `q`. ✓

### Multi-word phrase
| Test | Input | Expected |
|------|-------|----------|
| whole phrase | quick fast run | ickquay astfay unray |

`quick`: containsy? `[^aeiou]+` matches `q` (greedy: `qu` → `u` is vowel, so only `q`), then needs `y` but finds `u`. No match. Vowel? No. Cons? `[^aeiou]?qu` matches `qu`, remainder `ick`. Result: `ickquay` ✓
`fast`: cons matches `f`, remainder `ast`. `astfay` ✓
`run`: cons matches `r`, remainder `un`. `unray` ✓

## Verdict

**All 22 test cases verified against the regex patterns. No bugs found.**

The ordering of checks (containsy → vowel → cons) is correct:
1. `containsy` must be checked first to handle `y`-as-vowel before the generic consonant rule consumes it
2. `vowel` catches all vowel-initial words plus `xr`/`yt` special cases
3. `cons` handles remaining consonant-initial words including `qu` combinations

## Recommendation
Proceed with implementation as designed. No changes needed.
