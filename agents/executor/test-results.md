# Test Results: beer-song (Go)

## 1. Static Analysis — `go vet ./...`

```
(no output — all clear)
```

**Exit code:** 0

---

## 2. Tests — `go test -v ./...`

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

**Exit code:** 0
**Result:** ALL 12 TESTS PASSED

---

## 3. Benchmarks — `go test -bench=. -benchtime=1s`

```
goos: linux
goarch: amd64
pkg: beer
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkSeveralVerses-128     	  238098	      5290 ns/op
BenchmarkEntireSong-128        	   18039	     66084 ns/op
PASS
ok  	beer	3.193s
```

**Exit code:** 0

---

## Summary

| Check           | Status | Details                          |
|-----------------|--------|----------------------------------|
| `go vet`        | PASS   | No issues found                  |
| `go test -v`    | PASS   | 12/12 tests passed               |
| Benchmarks      | PASS   | SeveralVerses: 5290 ns/op, EntireSong: 66084 ns/op |
