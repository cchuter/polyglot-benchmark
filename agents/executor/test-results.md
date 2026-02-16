# Beer Song - Test Results

## 1. `go vet ./...`
**Result: PASS** (no output, no issues found)

## 2. `go build ./...`
**Result: PASS** (compiles successfully, no errors)

## 3. `go test -v ./...`
**Result: ALL TESTS PASS**

```
=== RUN   TestBottlesVerse
=== RUN   TestBottlesVerse/a_typical_verse
=== RUN   TestBottlesVerse/another_typical_verse
=== RUN   TestBottlesVerse/verse_2
=== RUN   TestBottlesVerse/verse_1
=== RUN   TestBottlesVerse/verse_0
=== RUN   TestBottlesVerse/invalid_verse
--- PASS: TestBottlesVerse (0.00s)
    --- PASS: TestBottlesVerse/a_typical_verse (0.00s)
    --- PASS: TestBottlesVerse/another_typical_verse (0.00s)
    --- PASS: TestBottlesVerse/verse_2 (0.00s)
    --- PASS: TestBottlesVerse/verse_1 (0.00s)
    --- PASS: TestBottlesVerse/verse_0 (0.00s)
    --- PASS: TestBottlesVerse/invalid_verse (0.00s)
=== RUN   TestSeveralVerses
=== RUN   TestSeveralVerses/multiple_verses
=== RUN   TestSeveralVerses/a_different_set_of_verses
=== RUN   TestSeveralVerses/invalid_start
=== RUN   TestSeveralVerses/invalid_stop
=== RUN   TestSeveralVerses/start_less_than_stop
--- PASS: TestSeveralVerses (0.00s)
    --- PASS: TestSeveralVerses/multiple_verses (0.00s)
    --- PASS: TestSeveralVerses/a_different_set_of_verses (0.00s)
    --- PASS: TestSeveralVerses/invalid_start (0.00s)
    --- PASS: TestSeveralVerses/invalid_stop (0.00s)
    --- PASS: TestSeveralVerses/start_less_than_stop (0.00s)
=== RUN   TestEntireSong
--- PASS: TestEntireSong (0.00s)
PASS
ok  	beer	(cached)
```

## 4. `go test -bench=. -benchtime=1s ./...`
**Result: BENCHMARKS RUN SUCCESSFULLY**

```
goos: linux
goarch: amd64
pkg: beer
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkSeveralVerses-128     	  267333	      4995 ns/op
BenchmarkEntireSong-128        	   18856	     63792 ns/op
PASS
ok  	beer	3.245s
```

## Summary

| Check              | Status |
|--------------------|--------|
| Build passes       | YES    |
| All tests pass     | YES (12/12 test cases across 3 test functions) |
| Benchmarks run     | YES    |
| Warnings or errors | NONE   |
