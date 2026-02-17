# Test Results

## Command
go test -v ./...

## Output
```
=== RUN   TestCost
=== RUN   TestCost/Only_a_single_book
=== RUN   TestCost/Two_of_the_same_book
=== RUN   TestCost/Empty_basket
=== RUN   TestCost/Two_different_books
=== RUN   TestCost/Three_different_books
=== RUN   TestCost/Four_different_books
=== RUN   TestCost/Five_different_books
=== RUN   TestCost/Two_groups_of_four_is_cheaper_than_group_of_five_plus_group_of_three
=== RUN   TestCost/Two_groups_of_four_is_cheaper_than_groups_of_five_and_three
=== RUN   TestCost/Group_of_four_plus_group_of_two_is_cheaper_than_two_groups_of_three
=== RUN   TestCost/Two_each_of_first_four_books_and_one_copy_each_of_rest
=== RUN   TestCost/Two_copies_of_each_book
=== RUN   TestCost/Three_copies_of_first_book_and_two_each_of_remaining
=== RUN   TestCost/Three_each_of_first_two_books_and_two_each_of_remaining_books
=== RUN   TestCost/Four_groups_of_four_are_cheaper_than_two_groups_each_of_five_and_three
=== RUN   TestCost/Check_that_groups_of_four_are_created_properly_even_when_there_are_more_groups_of_three_than_groups_of_five
=== RUN   TestCost/One_group_of_one_and_four_is_cheaper_than_one_group_of_two_and_three
=== RUN   TestCost/One_group_of_one_and_two_plus_three_groups_of_four_is_cheaper_than_one_group_of_each_size
--- PASS: TestCost (0.00s)
    --- PASS: TestCost/Only_a_single_book (0.00s)
    --- PASS: TestCost/Two_of_the_same_book (0.00s)
    --- PASS: TestCost/Empty_basket (0.00s)
    --- PASS: TestCost/Two_different_books (0.00s)
    --- PASS: TestCost/Three_different_books (0.00s)
    --- PASS: TestCost/Four_different_books (0.00s)
    --- PASS: TestCost/Five_different_books (0.00s)
    --- PASS: TestCost/Two_groups_of_four_is_cheaper_than_group_of_five_plus_group_of_three (0.00s)
    --- PASS: TestCost/Two_groups_of_four_is_cheaper_than_groups_of_five_and_three (0.00s)
    --- PASS: TestCost/Group_of_four_plus_group_of_two_is_cheaper_than_two_groups_of_three (0.00s)
    --- PASS: TestCost/Two_each_of_first_four_books_and_one_copy_each_of_rest (0.00s)
    --- PASS: TestCost/Two_copies_of_each_book (0.00s)
    --- PASS: TestCost/Three_copies_of_first_book_and_two_each_of_remaining (0.00s)
    --- PASS: TestCost/Three_each_of_first_two_books_and_two_each_of_remaining_books (0.00s)
    --- PASS: TestCost/Four_groups_of_four_are_cheaper_than_two_groups_each_of_five_and_three (0.00s)
    --- PASS: TestCost/Check_that_groups_of_four_are_created_properly_even_when_there_are_more_groups_of_three_than_groups_of_five (0.00s)
    --- PASS: TestCost/One_group_of_one_and_four_is_cheaper_than_one_group_of_two_and_three (0.00s)
    --- PASS: TestCost/One_group_of_one_and_two_plus_three_groups_of_four_is_cheaper_than_one_group_of_each_size (0.00s)
PASS
ok  	bookstore	0.005s
```

## Summary
- Total tests: 18
- Passed: 18
- Failed: 0
- Build errors: no
