# Palindrome Products - Build & Test Results

## go vet

```
$ cd go/exercises/practice/palindrome-products && go vet ./...
```

**Result:** PASS (no issues found)

## go test -v

```
$ go test -v ./...
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
ok  	palindrome	0.034s
```

**Result:** ALL 5 TESTS PASSED

## Summary

| Check       | Status |
|-------------|--------|
| go vet      | PASS   |
| go test     | 5/5 PASS |
