**Findings (ordered by severity)**

1. `math.MinInt` overflow breaks negative amount formatting  
`go/exercises/practice/ledger/ledger.go:86` and `go/exercises/practice/ledger/ledger.go:88` negate `cents` with `abs = -abs`. For `Change == math.MinInt`, this overflows and stays negative, yielding malformed output (example I reproduced: `($-92,233,720,368,547,758.-8)`).  
Impact: incorrect formatting for an edge negative amount, including bad decimal/thousand formatting.

2. Overlong UTF-8 descriptions can be truncated mid-rune  
`go/exercises/practice/ledger/ledger.go:77` to `go/exercises/practice/ledger/ledger.go:80` uses byte-based slicing (`desc[:22]`). With multibyte characters, truncation can split a rune and emit invalid UTF-8 bytes (I reproduced output containing `\xc3...`).  
Impact: corrupted text and potential column misalignment for non-ASCII overlong descriptions.

No other issues found against the current test expectations:
- Exact formatting/alignment/separators/trailing spaces: matches tests.
- Empty entries, zero/normal negative amounts, thousand separators: matches tests.
- Input immutability: preserved (copy + sort on copied slice).
- Sort order (date, description, change): correct.
- Error handling for invalid currency/locale/date: returns errors as expected.

I also ran `go test` in `go/exercises/practice/ledger`; all existing tests pass.