# Test Results: kindergarten-garden

## Build
```
$ go build ./...
```
**Result: SUCCESS** (no errors)

## Tests (go test -v ./...)
```
=== RUN   TestGarden
=== RUN   TestGarden/garden_with_single_student
=== RUN   TestGarden/different_garden_with_single_student
=== RUN   TestGarden/garden_with_two_students
=== RUN   TestGarden/garden_with_three_students
=== RUN   TestGarden/full_garden
=== RUN   TestGarden/names_out_of_order
=== RUN   TestGarden/lookup_invalid_name
=== RUN   TestGarden/wrong_diagram_format
=== RUN   TestGarden/mismatched_rows
=== RUN   TestGarden/odd_number_of_cups
=== RUN   TestGarden/duplicate_name
=== RUN   TestGarden/invalid_cup_codes
--- PASS: TestGarden (0.00s)
    --- PASS: TestGarden/garden_with_single_student (0.00s)
    --- PASS: TestGarden/different_garden_with_single_student (0.00s)
    --- PASS: TestGarden/garden_with_two_students (0.00s)
    --- PASS: TestGarden/garden_with_three_students (0.00s)
    --- PASS: TestGarden/full_garden (0.00s)
    --- PASS: TestGarden/names_out_of_order (0.00s)
    --- PASS: TestGarden/lookup_invalid_name (0.00s)
    --- PASS: TestGarden/wrong_diagram_format (0.00s)
    --- PASS: TestGarden/mismatched_rows (0.00s)
    --- PASS: TestGarden/odd_number_of_cups (0.00s)
    --- PASS: TestGarden/duplicate_name (0.00s)
    --- PASS: TestGarden/invalid_cup_codes (0.00s)
=== RUN   TestNamesNotModified
--- PASS: TestNamesNotModified (0.00s)
=== RUN   TestTwoGardens
--- PASS: TestTwoGardens (0.00s)
PASS
ok  	kindergarten	0.005s
```
**Result: ALL 15 TESTS PASSED**

## Benchmarks (go test -bench=. -short ./...)
```
PASS
ok  	kindergarten	0.010s
```
**Result: PASS** (no benchmarks defined in test suite)

## Summary
- Build: PASS
- Tests: 15/15 PASSED (0 failures)
- Benchmarks: PASS
