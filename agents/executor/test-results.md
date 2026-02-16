# Beer Song - Test Results

## Unit Tests (`go test -v ./...`)

**Result: ALL PASS**

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

## Benchmarks (`go test -bench=. ./...`)

**Result: PASS**

```
goos: linux
goarch: amd64
pkg: beer
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkSeveralVerses-128     	  275383	      5064 ns/op
BenchmarkEntireSong-128        	   18788	     63588 ns/op
PASS
ok  	beer	3.294s
```

## Summary

| Test Suite | Tests | Passed | Failed |
|---|---|---|---|
| TestBottlesVerse | 6 | 6 | 0 |
| TestSeveralVerses | 5 | 5 | 0 |
| TestEntireSong | 1 | 1 | 0 |
| **Total** | **12** | **12** | **0** |

| Benchmark | Iterations | ns/op |
|---|---|---|
| BenchmarkSeveralVerses-128 | 275,383 | 5,064 |
| BenchmarkEntireSong-128 | 18,788 | 63,588 |
