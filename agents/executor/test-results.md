# Test Results

## Test Results

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
--- PASS: TestGenerateCharacter (0.01s)
    --- PASS: TestGenerateCharacter/should_generate_a_character_with_random_ability_scores (0.01s)
PASS
ok  	dnd-character	0.010s
```

**Status: ALL TESTS PASS (3/3 test suites, 18 subtests)**

## Benchmark Results

```
goos: linux
goarch: amd64
pkg: dnd-character
cpu: AMD Ryzen Threadripper PRO 5995WX 64-Cores
BenchmarkModifier-128      	1000000000	         0.6925 ns/op
BenchmarkAbility-128       	 5798413	       210.6 ns/op
BenchmarkCharacter-128     	  825864	      1275 ns/op
PASS
ok  	dnd-character	3.308s
```

**Benchmark Summary:**
- `Modifier`: 0.69 ns/op (extremely fast, pure math)
- `Ability`: 210.6 ns/op (dice rolling with randomness)
- `Character`: 1275 ns/op (full character generation)
