# Hexadecimal Exercise - Test Results

## Build

```
$ go build ./...
```

**Result:** SUCCESS (no errors)

## Unit Tests

```
$ go test -v ./...
=== RUN   TestParseHex
--- PASS: TestParseHex (0.00s)
=== RUN   TestHandleErrors
--- PASS: TestHandleErrors (0.00s)
PASS
ok  	hexadecimal	0.003s
```

**Result:** ALL TESTS PASSED (2/2)

## Benchmarks

```
$ go test -bench=. -benchtime=1s ./...
goos: linux
goarch: amd64
pkg: hexadecimal
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkParseHex-128     	 2741164	       456.4 ns/op
PASS
ok  	hexadecimal	1.706s
```

**Result:** BENCHMARK PASSED

## Summary

| Check       | Status |
|-------------|--------|
| Build       | PASS   |
| TestParseHex | PASS  |
| TestHandleErrors | PASS |
| BenchmarkParseHex | PASS (456.4 ns/op) |
