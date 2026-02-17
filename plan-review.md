# Plan Review

## Review Method
Manual trace of the Selected Plan implementation against all 22 test cases.

## Implementation Under Review (Selected Plan)

```go
func translateWord(word string) string {
    if vowels[word[0]] || specials[word[:2]] {
        return word + "ay"
    }
    for pos := 1; pos < len(word); pos++ {
        letter := word[pos]
        if vowelsY[letter] {
            if letter == 'u' && word[pos-1] == 'q' {
                pos++
            }
            return word[pos:] + word[:pos] + "ay"
        }
    }
    return word + "ay"
}
```

## Bug Found: `specials[word[:2]]` panics on single-char words

**Severity: HIGH**

If `word` is only 1 character long (e.g., `"a"`), `word[:2]` will panic with an index out of range error. Although none of the current 22 test cases include a single-character word that starts with a consonant, the word `"a"` starts with a vowel so the `vowels[word[0]]` check would short-circuit. However, a single consonant word like `"I"` (if lowercased to `"i"`) would also be caught by the vowel check. But a hypothetical single-char consonant word would panic.

**Analysis of current test cases:** All test words are >= 2 chars. The `vowels[word[0]]` check short-circuits for vowel-starting words. For consonant-starting words, all are >= 2 chars. So this won't trigger for the existing tests but is fragile.

**Recommendation:** Add a length guard: `if vowels[word[0]] || (len(word) >= 2 && specials[word[:2]])`

## Test Case Trace

| # | Input | Expected | Rule | Trace Result | Pass? |
|---|-------|----------|------|-------------|-------|
| 1 | "apple" | "appleay" | R1 | vowels['a']=true → "apple"+"ay" | YES |
| 2 | "ear" | "earay" | R1 | vowels['e']=true → "ear"+"ay" | YES |
| 3 | "igloo" | "iglooay" | R1 | vowels['i']=true → "igloo"+"ay" | YES |
| 4 | "object" | "objectay" | R1 | vowels['o']=true → "object"+"ay" | YES |
| 5 | "under" | "underay" | R1 | vowels['u']=true → "under"+"ay" | YES |
| 6 | "equal" | "equalay" | R1 | vowels['e']=true → "equal"+"ay" | YES |
| 7 | "pig" | "igpay" | R2 | pos=1: 'i' in vowelsY → "ig"+"p"+"ay" = "igpay" | YES |
| 8 | "koala" | "oalakay" | R2 | pos=1: 'o' in vowelsY → "oala"+"k"+"ay" = "oalakay" | YES |
| 9 | "xenon" | "enonxay" | R2 | specials["xe"]=false, pos=1: 'e' in vowelsY → "enon"+"x"+"ay" = "enonxay" | YES |
| 10 | "qat" | "atqay" | R2 | pos=1: 'a' in vowelsY, letter='a', not 'u' → "at"+"q"+"ay" = "atqay" | YES |
| 11 | "chair" | "airchay" | R2 | pos=1: 'h' not in vowelsY, pos=2: 'a' in vowelsY → "air"+"ch"+"ay" = "airchay" | YES |
| 12 | "queen" | "eenquay" | R3 | pos=1: 'u' in vowelsY, letter='u' && word[0]='q' → pos becomes 2 → "een"+"qu"+"ay" = "eenquay" | YES |
| 13 | "square" | "aresquay" | R3 | pos=1: 'q' not in vowelsY, pos=2: 'u' in vowelsY, letter='u' && word[1]='q' → pos becomes 3 → "are"+"squ"+"ay" = "aresquay" | YES |
| 14 | "therapy" | "erapythay" | R2 | pos=1: 'h' not in vowelsY, pos=2: 'e' in vowelsY → "erapy"+"th"+"ay" = "erapythay" | YES |
| 15 | "thrush" | "ushthray" | R2 | pos=1: 'h' not, pos=2: 'r' not, pos=3: 'u' in vowelsY → "ush"+"thr"+"ay" = "ushthray" | YES |
| 16 | "school" | "oolschay" | R2 | pos=1: 'c' not, pos=2: 'h' not, pos=3: 'o' in vowelsY → "ool"+"sch"+"ay" = "oolschay" | YES |
| 17 | "yttria" | "yttriaay" | R1 | specials["yt"]=true → "yttria"+"ay" = "yttriaay" | YES |
| 18 | "xray" | "xrayay" | R1 | specials["xr"]=true → "xray"+"ay" = "xrayay" | YES |
| 19 | "yellow" | "ellowyay" | R2 | pos=1: 'e' in vowelsY → "ellow"+"y"+"ay" = "ellowyay" | YES |
| 20 | "rhythm" | "ythmrhay" | R4 | pos=1: 'h' not, pos=2: 'y' in vowelsY → "ythm"+"rh"+"ay" = "ythmrhay" | YES |
| 21 | "my" | "ymay" | R4 | pos=1: 'y' in vowelsY → "y"+"m"+"ay" = "ymay" | YES |
| 22 | "quick fast run" | "ickquay astfay unray" | Mixed | "quick"→"ickquay", "fast"→"astfay", "run"→"unray" | YES |

## Detailed Trace for Tricky Cases

### "queen" (test 12)
- word[0]='q', not a vowel. specials["qu"]=false.
- Loop: pos=1, letter='u', vowelsY['u']=true
- letter=='u' && word[0]=='q' → pos++ → pos=2
- return "een" + "qu" + "ay" = "eenquay" ✓

### "square" (test 13)
- word[0]='s', not a vowel. specials["sq"]=false.
- Loop: pos=1, letter='q', vowelsY['q']=false
- pos=2, letter='u', vowelsY['u']=true
- letter=='u' && word[1]=='q' → pos++ → pos=3
- return "are" + "squ" + "ay" = "aresquay" ✓

### "rhythm" (test 20)
- word[0]='r', not a vowel. specials["rh"]=false.
- Loop: pos=1, letter='h', vowelsY['h']=false
- pos=2, letter='y', vowelsY['y']=true
- letter != 'u', no qu adjustment
- return "ythm" + "rh" + "ay" = "ythmrhay" ✓

### "yellow" (test 19)
- word[0]='y', not a vowel (correct! y at start is consonant). specials["ye"]=false.
- Loop: pos=1, letter='e', vowelsY['e']=true
- return "ellow" + "y" + "ay" = "ellowyay" ✓

## Verdict

**The Selected Plan implementation will pass all 22 test cases.**

### Recommended Fix
Add a length guard for the `specials` lookup to prevent potential panics on short words:
```go
if vowels[word[0]] || (len(word) >= 2 && specials[word[:2]]) {
```

This is a defensive improvement. Without it, the code still passes all existing tests, but it's fragile.

## Overall Assessment: APPROVED with minor fix
