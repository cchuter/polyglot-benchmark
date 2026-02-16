# Test Results

## Command
```
cd go/exercises/practice/alphametics && go test -v -count=1 -timeout 300s
```

## Result: ALL PASS (10/10)

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
--- PASS: TestSolve (1.05s)
    --- PASS: TestSolve/puzzle_with_three_letters (0.00s)
    --- PASS: TestSolve/solution_must_have_unique_value_for_each_letter (0.00s)
    --- PASS: TestSolve/leading_zero_solution_is_invalid (0.00s)
    --- PASS: TestSolve/puzzle_with_two_digits_final_carry (0.00s)
    --- PASS: TestSolve/puzzle_with_four_letters (0.00s)
    --- PASS: TestSolve/puzzle_with_six_letters (0.02s)
    --- PASS: TestSolve/puzzle_with_seven_letters (0.07s)
    --- PASS: TestSolve/puzzle_with_eight_letters (0.22s)
    --- PASS: TestSolve/puzzle_with_ten_letters (0.43s)
    --- PASS: TestSolve/puzzle_with_ten_letters_and_199_addends (0.31s)
PASS
ok  	alphametics	1.165s
```

## Summary
- Total tests: 10
- Passed: 10
- Failed: 0
- Total time: 1.165s
- Largest test (199 addends, 10 letters): 0.31s
