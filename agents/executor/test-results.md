# Forth Exercise Test Results

## Build Status
**PASS** - Build succeeded with no compilation errors.

## Test Results
- **Total tests**: 46
- **Passed**: 46
- **Failed**: 0
- **Result**: ALL PASS

## Full Test Output

```
=== RUN   TestForth
=== RUN   TestForth/numbers_just_get_pushed_onto_the_stack
=== RUN   TestForth/pushes_negative_numbers_onto_the_stack
=== RUN   TestForth/can_add_two_numbers
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack
=== RUN   TestForth/can_subtract_two_numbers
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#01
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack#01
=== RUN   TestForth/can_multiply_two_numbers
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#02
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack#02
=== RUN   TestForth/can_divide_two_numbers
=== RUN   TestForth/performs_integer_division
=== RUN   TestForth/errors_if_dividing_by_zero
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#03
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack#03
=== RUN   TestForth/addition_and_subtraction
=== RUN   TestForth/multiplication_and_division
=== RUN   TestForth/copies_a_value_on_the_stack
=== RUN   TestForth/copies_the_top_value_on_the_stack
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#04
=== RUN   TestForth/removes_the_top_value_on_the_stack_if_it_is_the_only_one
=== RUN   TestForth/removes_the_top_value_on_the_stack_if_it_is_not_the_only_one
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#05
=== RUN   TestForth/swaps_the_top_two_values_on_the_stack_if_they_are_the_only_ones
=== RUN   TestForth/swaps_the_top_two_values_on_the_stack_if_they_are_not_the_only_ones
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#06
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack#04
=== RUN   TestForth/copies_the_second_element_if_there_are_only_two
=== RUN   TestForth/copies_the_second_element_if_there_are_more_than_two
=== RUN   TestForth/errors_if_there_is_nothing_on_the_stack#07
=== RUN   TestForth/errors_if_there_is_only_one_value_on_the_stack#05
=== RUN   TestForth/can_consist_of_built-in_words
=== RUN   TestForth/execute_in_the_right_order
=== RUN   TestForth/can_override_other_user-defined_words
=== RUN   TestForth/can_override_built-in_words
=== RUN   TestForth/can_override_built-in_operators
=== RUN   TestForth/can_use_different_words_with_the_same_name
=== RUN   TestForth/can_define_word_that_uses_word_with_the_same_name
=== RUN   TestForth/cannot_redefine_non-negative_numbers
=== RUN   TestForth/cannot_redefine_negative_numbers
=== RUN   TestForth/errors_if_executing_a_non-existent_word
=== RUN   TestForth/DUP_is_case-insensitive
=== RUN   TestForth/DROP_is_case-insensitive
=== RUN   TestForth/SWAP_is_case-insensitive
=== RUN   TestForth/OVER_is_case-insensitive
=== RUN   TestForth/user-defined_words_are_case-insensitive
=== RUN   TestForth/definitions_are_case-insensitive
--- PASS: TestForth (0.00s)
    --- PASS: TestForth/numbers_just_get_pushed_onto_the_stack (0.00s)
    --- PASS: TestForth/pushes_negative_numbers_onto_the_stack (0.00s)
    --- PASS: TestForth/can_add_two_numbers (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack (0.00s)
    --- PASS: TestForth/can_subtract_two_numbers (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#01 (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack#01 (0.00s)
    --- PASS: TestForth/can_multiply_two_numbers (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#02 (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack#02 (0.00s)
    --- PASS: TestForth/can_divide_two_numbers (0.00s)
    --- PASS: TestForth/performs_integer_division (0.00s)
    --- PASS: TestForth/errors_if_dividing_by_zero (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#03 (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack#03 (0.00s)
    --- PASS: TestForth/addition_and_subtraction (0.00s)
    --- PASS: TestForth/multiplication_and_division (0.00s)
    --- PASS: TestForth/copies_a_value_on_the_stack (0.00s)
    --- PASS: TestForth/copies_the_top_value_on_the_stack (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#04 (0.00s)
    --- PASS: TestForth/removes_the_top_value_on_the_stack_if_it_is_the_only_one (0.00s)
    --- PASS: TestForth/removes_the_top_value_on_the_stack_if_it_is_not_the_only_one (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#05 (0.00s)
    --- PASS: TestForth/swaps_the_top_two_values_on_the_stack_if_they_are_the_only_ones (0.00s)
    --- PASS: TestForth/swaps_the_top_two_values_on_the_stack_if_they_are_not_the_only_ones (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#06 (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack#04 (0.00s)
    --- PASS: TestForth/copies_the_second_element_if_there_are_only_two (0.00s)
    --- PASS: TestForth/copies_the_second_element_if_there_are_more_than_two (0.00s)
    --- PASS: TestForth/errors_if_there_is_nothing_on_the_stack#07 (0.00s)
    --- PASS: TestForth/errors_if_there_is_only_one_value_on_the_stack#05 (0.00s)
    --- PASS: TestForth/can_consist_of_built-in_words (0.00s)
    --- PASS: TestForth/execute_in_the_right_order (0.00s)
    --- PASS: TestForth/can_override_other_user-defined_words (0.00s)
    --- PASS: TestForth/can_override_built-in_words (0.00s)
    --- PASS: TestForth/can_override_built-in_operators (0.00s)
    --- PASS: TestForth/can_use_different_words_with_the_same_name (0.00s)
    --- PASS: TestForth/can_define_word_that_uses_word_with_the_same_name (0.00s)
    --- PASS: TestForth/cannot_redefine_non-negative_numbers (0.00s)
    --- PASS: TestForth/cannot_redefine_negative_numbers (0.00s)
    --- PASS: TestForth/errors_if_executing_a_non-existent_word (0.00s)
    --- PASS: TestForth/DUP_is_case-insensitive (0.00s)
    --- PASS: TestForth/DROP_is_case-insensitive (0.00s)
    --- PASS: TestForth/SWAP_is_case-insensitive (0.00s)
    --- PASS: TestForth/OVER_is_case-insensitive (0.00s)
    --- PASS: TestForth/user-defined_words_are_case-insensitive (0.00s)
    --- PASS: TestForth/definitions_are_case-insensitive (0.00s)
PASS
ok  	forth	0.005s
```

## Compilation Errors
None.
