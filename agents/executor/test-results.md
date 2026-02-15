# Alphametics - Build & Test Results

## go vet
```
$ go vet ./...
(no output - passed cleanly)
```
**Result: PASS**

## go test -v -count=1 ./...
```
=== RUN   TestSolve
=== RUN   TestSolve/puzzle_with_three_letters
=== RUN   TestSolve/solution_must_have_unique_value_for_each_letter
=== RUN   TestSolve/leading_zero_solution_is_invalid
=== RUN   TestSolve/puzzle_with_two_digits_final_carry
=== RUN   TestSolve/puzzle_with_four_letters
=== RUN   TestSolve/puzzle_with_six_letters
=== RUN   TestSolve/puzzle_with_seven_letters
=== RUN   TestSolve/puzzle_with_eight_letters
=== RUN   TestSolve/puzzle_with_ten_letters
=== RUN   TestSolve/puzzle_with_ten_letters_and_199_addends
--- PASS: TestSolve (1.03s)
    --- PASS: TestSolve/puzzle_with_three_letters (0.00s)
    --- PASS: TestSolve/solution_must_have_unique_value_for_each_letter (0.00s)
    --- PASS: TestSolve/leading_zero_solution_is_invalid (0.00s)
    --- PASS: TestSolve/puzzle_with_two_digits_final_carry (0.00s)
    --- PASS: TestSolve/puzzle_with_four_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_six_letters (0.02s)
    --- PASS: TestSolve/puzzle_with_seven_letters (0.06s)
    --- PASS: TestSolve/puzzle_with_eight_letters (0.19s)
    --- PASS: TestSolve/puzzle_with_ten_letters (0.39s)
    --- PASS: TestSolve/puzzle_with_ten_letters_and_199_addends (0.38s)
PASS
ok  	alphametics	1.137s
```
**Result: PASS (10/10 tests passed in 1.137s)**

## go test -bench=. -benchtime=1x -short ./...
```
PASS
ok  	alphametics	1.237s
```
**Result: PASS**

## Summary
- **go vet**: PASS
- **Tests**: 10/10 PASS (1.137s)
- **Benchmarks**: PASS (1.237s)
- **Overall**: ALL PASSING
