# Test Results: dnd-character

## Test Run (`go test -v ./...`)

**Result: ALL PASS**

```
=== RUN   TestModifier
=== RUN   TestModifier/ability_modifier_for_score_3_is_-4
=== RUN   TestModifier/ability_modifier_for_score_4_is_-3
=== RUN   TestModifier/ability_modifier_for_score_5_is_-3
=== RUN   TestModifier/ability_modifier_for_score_6_is_-2
=== RUN   TestModifier/ability_modifier_for_score_7_is_-2
=== RUN   TestModifier/ability_modifier_for_score_8_is_-1
=== RUN   TestModifier/ability_modifier_for_score_9_is_-1
=== RUN   TestModifier/ability_modifier_for_score_10_is_0
=== RUN   TestModifier/ability_modifier_for_score_11_is_0
=== RUN   TestModifier/ability_modifier_for_score_12_is_+1
=== RUN   TestModifier/ability_modifier_for_score_13_is_+1
=== RUN   TestModifier/ability_modifier_for_score_14_is_+2
=== RUN   TestModifier/ability_modifier_for_score_15_is_+2
=== RUN   TestModifier/ability_modifier_for_score_16_is_+3
=== RUN   TestModifier/ability_modifier_for_score_17_is_+3
=== RUN   TestModifier/ability_modifier_for_score_18_is_+4
--- PASS: TestModifier (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_3_is_-4 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_4_is_-3 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_5_is_-3 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_6_is_-2 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_7_is_-2 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_8_is_-1 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_9_is_-1 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_10_is_0 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_11_is_0 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_12_is_+1 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_13_is_+1 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_14_is_+2 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_15_is_+2 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_16_is_+3 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_17_is_+3 (0.00s)
    --- PASS: TestModifier/ability_modifier_for_score_18_is_+4 (0.00s)
=== RUN   TestAbility
=== RUN   TestAbility/should_generate_ability_score_within_accepted_range
--- PASS: TestAbility (0.00s)
    --- PASS: TestAbility/should_generate_ability_score_within_accepted_range (0.00s)
=== RUN   TestGenerateCharacter
=== RUN   TestGenerateCharacter/should_generate_a_character_with_random_ability_scores
--- PASS: TestGenerateCharacter (0.00s)
    --- PASS: TestGenerateCharacter/should_generate_a_character_with_random_ability_scores (0.00s)
PASS
ok  	dnd-character	0.009s
```

## Summary

| Test Suite         | Subtests | Result |
|--------------------|----------|--------|
| TestModifier       | 16       | PASS   |
| TestAbility        | 1        | PASS   |
| TestGenerateCharacter | 1     | PASS   |
| **Total**          | **18**   | **ALL PASS** |

## Benchmark Run (`go test -bench=. -benchtime=1s ./...`)

```
goos: linux
goarch: amd64
pkg: dnd-character
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkModifier-128      	1000000000	         0.6481 ns/op
BenchmarkAbility-128       	34188122	        33.05 ns/op
BenchmarkCharacter-128     	 6010840	       197.5 ns/op
PASS
ok  	dnd-character	4.954s
```

| Benchmark            | Iterations     | ns/op   |
|----------------------|----------------|---------|
| BenchmarkModifier    | 1,000,000,000  | 0.65    |
| BenchmarkAbility     | 34,188,122     | 33.05   |
| BenchmarkCharacter   | 6,010,840      | 197.50  |
