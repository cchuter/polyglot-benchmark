# Beer Song - Build & Test Results

## Build

```
$ go build ./...
```

Build succeeded with no errors.

## Tests

```
$ go test -v ./...
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
ok  	beer	0.005s
```

## Summary

- Build: PASS
- Tests: ALL PASS (12/12 test cases)
  - TestBottlesVerse: 6/6 passed
  - TestSeveralVerses: 5/5 passed
  - TestEntireSong: 1/1 passed
