# Test Results: crypto-square

## Test Run (`go test -v ./...`)

**Result: ALL PASS (19/19 tests)**

```
=== RUN   TestEncode
=== RUN   TestEncode/s#$%^&plunk
=== RUN   TestEncode/1,_2,_3_GO!
=== RUN   TestEncode/1234
=== RUN   TestEncode/123456789
=== RUN   TestEncode/123456789abc
=== RUN   TestEncode/Never_vex_thine_heart_with_idle_woes
=== RUN   TestEncode/ZOMG!_ZOMBIES!!!
=== RUN   TestEncode/Time_is_an_illusion._Lunchtime_doubly_so.
=== RUN   TestEncode/We_all_know_interspecies_romance_is_weird.
=== RUN   TestEncode/Madness,_and_then_illumination.
=== RUN   TestEncode/Vampires_are_people_too!
=== RUN   TestEncode/#00
=== RUN   TestEncode/1
=== RUN   TestEncode/12
=== RUN   TestEncode/12_3
=== RUN   TestEncode/12345678
=== RUN   TestEncode/123456789a
=== RUN   TestEncode/If_man_was_meant_to_stay_on_the_ground_god_would_have_given_us_roots
=== RUN   TestEncode/Have_a_nice_day._Feed_the_dog_&_chill_out!
--- PASS: TestEncode (0.00s)
    --- PASS: TestEncode/s#$%^&plunk (0.00s)
    --- PASS: TestEncode/1,_2,_3_GO! (0.00s)
    --- PASS: TestEncode/1234 (0.00s)
    --- PASS: TestEncode/123456789 (0.00s)
    --- PASS: TestEncode/123456789abc (0.00s)
    --- PASS: TestEncode/Never_vex_thine_heart_with_idle_woes (0.00s)
    --- PASS: TestEncode/ZOMG!_ZOMBIES!!! (0.00s)
    --- PASS: TestEncode/Time_is_an_illusion._Lunchtime_doubly_so. (0.00s)
    --- PASS: TestEncode/We_all_know_interspecies_romance_is_weird. (0.00s)
    --- PASS: TestEncode/Madness,_and_then_illumination. (0.00s)
    --- PASS: TestEncode/Vampires_are_people_too! (0.00s)
    --- PASS: TestEncode/#00 (0.00s)
    --- PASS: TestEncode/1 (0.00s)
    --- PASS: TestEncode/12 (0.00s)
    --- PASS: TestEncode/12_3 (0.00s)
    --- PASS: TestEncode/12345678 (0.00s)
    --- PASS: TestEncode/123456789a (0.00s)
    --- PASS: TestEncode/If_man_was_meant_to_stay_on_the_ground_god_would_have_given_us_roots (0.00s)
    --- PASS: TestEncode/Have_a_nice_day._Feed_the_dog_&_chill_out! (0.00s)
PASS
ok  	cryptosquare	(cached)
```

## Benchmark Run (`go test -bench=.`)

**Result: PASS**

```
goos: linux
goarch: amd64
pkg: cryptosquare
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkEncode-128     	   39056	     33085 ns/op
PASS
ok  	cryptosquare	1.614s
```

## Summary

- **Tests**: 19/19 PASS
- **Benchmark**: BenchmarkEncode-128: 39,056 iterations at 33,085 ns/op
- **Platform**: linux/amd64, AMD Ryzen Threadripper PRO 5995WX 64-Cores
