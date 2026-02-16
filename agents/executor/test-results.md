# Executor Test Results

## Command
```
cd go/exercises/practice/alphametics && go test -v -count=1 -timeout 10m ./...
```

## Output
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
--- PASS: TestSolve (0.00s)
    --- PASS: TestSolve/puzzle_with_three_letters (0.00s)
    --- PASS: TestSolve/solution_must_have_unique_value_for_each_letter (0.00s)
    --- PASS: TestSolve/leading_zero_solution_is_invalid (0.00s)
    --- PASS: TestSolve/puzzle_with_two_digits_final_carry (0.00s)
    --- PASS: TestSolve/puzzle_with_four_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_six_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_seven_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_eight_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_ten_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_ten_letters_and_199_addends (0.00s)
PASS
ok  	alphametics	0.005s
```

## Summary
- **10/10 tests PASS**
- Total execution time: 0.005s
- No compilation errors
- No runtime errors
