# Connect Exercise - Test Results

## go test -v ./...

```
=== RUN   TestResultOf
=== RUN   TestResultOf/an_empty_board_has_no_winner
=== RUN   TestResultOf/X_can_win_on_a_1x1_board
=== RUN   TestResultOf/O_can_win_on_a_1x1_board
=== RUN   TestResultOf/only_edges_does_not_make_a_winner
=== RUN   TestResultOf/illegal_diagonal_does_not_make_a_winner
=== RUN   TestResultOf/nobody_wins_crossing_adjacent_angles
=== RUN   TestResultOf/X_wins_crossing_from_left_to_right
=== RUN   TestResultOf/O_wins_crossing_from_top_to_bottom
=== RUN   TestResultOf/X_wins_using_a_convoluted_path
=== RUN   TestResultOf/X_wins_using_a_spiral_path
--- PASS: TestResultOf (0.00s)
    --- PASS: TestResultOf/an_empty_board_has_no_winner (0.00s)
    --- PASS: TestResultOf/X_can_win_on_a_1x1_board (0.00s)
    --- PASS: TestResultOf/O_can_win_on_a_1x1_board (0.00s)
    --- PASS: TestResultOf/only_edges_does_not_make_a_winner (0.00s)
    --- PASS: TestResultOf/illegal_diagonal_does_not_make_a_winner (0.00s)
    --- PASS: TestResultOf/nobody_wins_crossing_adjacent_angles (0.00s)
    --- PASS: TestResultOf/X_wins_crossing_from_left_to_right (0.00s)
    --- PASS: TestResultOf/O_wins_crossing_from_top_to_bottom (0.00s)
    --- PASS: TestResultOf/X_wins_using_a_convoluted_path (0.00s)
    --- PASS: TestResultOf/X_wins_using_a_spiral_path (0.00s)
PASS
ok  	connect	(cached)
```

**Result: PASS** - 10/10 subtests passed

## go test -bench=. -benchtime=1s ./...

```
goos: linux
goarch: amd64
pkg: connect
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkResultOf-128     	   36735	     35760 ns/op
PASS
ok  	connect	1.660s
```

**Result: PASS** - Benchmark: 36,735 iterations, 35,760 ns/op

## go vet ./...

```
(no output - clean)
```

**Result: PASS** - No issues found

## Summary

| Check | Status |
|-------|--------|
| Unit Tests (10 subtests) | PASS |
| Benchmarks | PASS (35,760 ns/op) |
| go vet | PASS (clean) |

All checks passed successfully.
