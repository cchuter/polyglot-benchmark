# Octal Exercise Test Results

## Command 1: `go test -v ./...`

```
=== RUN   TestParseOctal
--- PASS: TestParseOctal (0.00s)
PASS
ok  	octal	0.003s
```

**Exit code: 0**

## Command 2: `go test -bench=. ./...`

```
goos: linux
goarch: amd64
pkg: octal
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkParseOctal-128     	 2584717	       470.1 ns/op
PASS
ok  	octal	1.693s
```

**Exit code: 0**

## Summary

- All tests: **PASSED**
- Benchmark: **PASSED** (470.1 ns/op)
