# Book-Store Exercise - Test Results

**Date:** 2026-02-15
**Exercise:** book-store (Go)
**Path:** go/exercises/practice/book-store

---

## 1. Build Output

```
$ go build ./...
```

**Result:** SUCCESS (no errors)

---

## 2. Test Output

```
$ go test -v ./...

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
ok  	bookstore	(cached)
```

**Result:** ALL 18 TESTS PASSED (18/18)

| # | Test Case | Result |
|---|-----------|--------|
| 1 | Only_a_single_book | PASS |
| 2 | Two_of_the_same_book | PASS |
| 3 | Empty_basket | PASS |
| 4 | Two_different_books | PASS |
| 5 | Three_different_books | PASS |
| 6 | Four_different_books | PASS |
| 7 | Five_different_books | PASS |
| 8 | Two_groups_of_four_is_cheaper_than_group_of_five_plus_group_of_three | PASS |
| 9 | Two_groups_of_four_is_cheaper_than_groups_of_five_and_three | PASS |
| 10 | Group_of_four_plus_group_of_two_is_cheaper_than_two_groups_of_three | PASS |
| 11 | Two_each_of_first_four_books_and_one_copy_each_of_rest | PASS |
| 12 | Two_copies_of_each_book | PASS |
| 13 | Three_copies_of_first_book_and_two_each_of_remaining | PASS |
| 14 | Three_each_of_first_two_books_and_two_each_of_remaining_books | PASS |
| 15 | Four_groups_of_four_are_cheaper_than_two_groups_each_of_five_and_three | PASS |
| 16 | Check_that_groups_of_four_are_created_properly_even_when_there_are_more_groups_of_three_than_groups_of_five | PASS |
| 17 | One_group_of_one_and_four_is_cheaper_than_one_group_of_two_and_three | PASS |
| 18 | One_group_of_one_and_two_plus_three_groups_of_four_is_cheaper_than_one_group_of_each_size | PASS |

---

## 3. Vet Output

```
$ go vet ./...
```

**Result:** SUCCESS (no issues found)

---

## 4. Benchmark Results

```
$ go test -bench=. -benchtime=1s ./...

goos: linux
goarch: amd64
pkg: bookstore
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkCost-128     	    4066	    289097 ns/op
PASS
ok  	bookstore	1.218s
```

**Result:** BenchmarkCost-128: 4,066 iterations, 289,097 ns/op (~289 us/op)

---

## Summary

| Check | Result |
|-------|--------|
| Build | SUCCESS |
| Tests | 18/18 PASSED |
| Vet | SUCCESS (no issues) |
| Benchmark | 289 us/op (4,066 iterations) |

**Overall: ALL CHECKS PASSED**
