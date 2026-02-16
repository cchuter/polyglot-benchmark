# Test Results: food-chain

## `go vet ./...`

```
(no output — clean)
```

**Result: PASS**

## `go test -v -count=1 ./...`

```
=== RUN   TestVerse
=== RUN   TestVerse/verse_1
=== RUN   TestVerse/verse_2
=== RUN   TestVerse/verse_3
=== RUN   TestVerse/verse_4
=== RUN   TestVerse/verse_5
=== RUN   TestVerse/verse_6
=== RUN   TestVerse/verse_7
=== RUN   TestVerse/verse_8
--- PASS: TestVerse (0.00s)
    --- PASS: TestVerse/verse_1 (0.00s)
    --- PASS: TestVerse/verse_2 (0.00s)
    --- PASS: TestVerse/verse_3 (0.00s)
    --- PASS: TestVerse/verse_4 (0.00s)
    --- PASS: TestVerse/verse_5 (0.00s)
    --- PASS: TestVerse/verse_6 (0.00s)
    --- PASS: TestVerse/verse_7 (0.00s)
    --- PASS: TestVerse/verse_8 (0.00s)
=== RUN   TestVerses
--- PASS: TestVerses (0.00s)
=== RUN   TestSong
--- PASS: TestSong (0.00s)
PASS
ok  	foodchain	0.003s
```

**Result: PASS** — 3/3 test functions passed (TestVerse with 8 subtests, TestVerses, TestSong)

## `go test -bench=. -benchtime=1s ./...`

```
goos: linux
goarch: amd64
pkg: foodchain
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkSong-128     	   76645	     17113 ns/op
PASS
ok  	foodchain	1.480s
```

**Result: PASS** — BenchmarkSong: 76,645 iterations, 17,113 ns/op
