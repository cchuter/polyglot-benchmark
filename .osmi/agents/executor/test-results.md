# Beer Song - Build & Test Results

## 1. `go build ./...`

**Exit Code:** 0

**Output:**
_(no output - build succeeded)_

---

## 2. `go test -v ./...`

**Exit Code:** 0

**Output:**
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

---

## 3. `go vet ./...`

**Exit Code:** 0

**Output:**
_(no output - vet passed with no issues)_

---

## 4. `go test -bench=. -benchmem ./...`

**Exit Code:** 0

**Output:**
```
goos: linux
goarch: amd64
pkg: beer
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkSeveralVerses-128     	  251890	      5312 ns/op	    3486 B/op	      21 allocs/op
BenchmarkEntireSong-128        	   18686	     64883 ns/op	   57709 B/op	     106 allocs/op
PASS
ok  	beer	3.265s
```

---

## Summary

| Command | Exit Code | Result |
|---------|-----------|--------|
| `go build ./...` | 0 | SUCCESS |
| `go test -v ./...` | 0 | ALL 12 TESTS PASSED |
| `go vet ./...` | 0 | NO ISSUES |
| `go test -bench=. -benchmem ./...` | 0 | BENCHMARKS PASSED |

**Total tests:** 12 (3 test suites: TestBottlesVerse with 6 subtests, TestSeveralVerses with 5 subtests, TestEntireSong)
**Passed:** 12
**Failed:** 0
**Errors:** 0
