# Test Results: palindrome-products

## `go build ./...`

```
(no output)
```

**Exit code: 0 (PASS)**

## `go test -v ./...`

```
=== RUN   TestPalindromeProducts
=== RUN   TestPalindromeProducts/valid_limits_1-9
=== RUN   TestPalindromeProducts/valid_limits_10-99
=== RUN   TestPalindromeProducts/valid_limits_100-999
=== RUN   TestPalindromeProducts/no_palindromes
=== RUN   TestPalindromeProducts/fmin_>_fmax
--- PASS: TestPalindromeProducts (0.03s)
    --- PASS: TestPalindromeProducts/valid_limits_1-9 (0.00s)
    --- PASS: TestPalindromeProducts/valid_limits_10-99 (0.00s)
    --- PASS: TestPalindromeProducts/valid_limits_100-999 (0.03s)
    --- PASS: TestPalindromeProducts/no_palindromes (0.00s)
    --- PASS: TestPalindromeProducts/fmin_>_fmax (0.00s)
PASS
ok  	palindrome	(cached)
```

**Exit code: 0 (PASS)**

## `go test -bench=. -short`

```
PASS
ok  	palindrome	0.022s
```

**Exit code: 0 (PASS)**

## `go vet ./...`

```
(no output)
```

**Exit code: 0 (PASS)**

## Summary

| Check       | Result |
|-------------|--------|
| Build       | PASS   |
| Tests (5/5) | PASS   |
| Benchmarks  | PASS   |
| Vet         | PASS   |
