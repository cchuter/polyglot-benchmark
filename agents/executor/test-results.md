# Kindergarten Garden - Test Results

## Build

```
$ go build ./...
```

Build completed successfully with no errors or warnings.

## Tests

```
$ go test -v ./...
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
ok  	kindergarten	0.004s
```

## Benchmarks

```
$ go test -bench=. -benchtime=1s
goos: linux
goarch: amd64
pkg: kindergarten
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkNewGarden-128         	   89662	     13466 ns/op
BenchmarkGarden_Plants-128     	10939570	       106.4 ns/op
PASS
ok  	kindergarten	3.482s
```

## Summary

| Metric | Result |
|--------|--------|
| Build | PASS |
| Tests | 14/14 PASS, 0 FAIL |
| Benchmarks | 2/2 PASS |
| BenchmarkNewGarden | 13,466 ns/op |
| BenchmarkGarden_Plants | 106.4 ns/op |
