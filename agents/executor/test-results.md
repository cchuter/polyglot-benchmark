# Bowling Exercise - Test Results

## Full Test Output

```
=== RUN   TestRoll
=== RUN   TestRoll/rolls_cannot_score_negative_points
=== RUN   TestRoll/a_roll_cannot_score_more_than_10_points
=== RUN   TestRoll/two_rolls_in_a_frame_cannot_score_more_than_10_points
=== RUN   TestRoll/bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points
=== RUN   TestRoll/two_bonus_rolls_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points
=== RUN   TestRoll/the_second_bonus_rolls_after_a_strike_in_the_last_frame_cannot_be_a_strike_if_the_first_one_is_not_a_strike
=== RUN   TestRoll/second_bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points
=== RUN   TestRoll/cannot_roll_if_game_already_has_ten_frames
=== RUN   TestRoll/cannot_roll_after_bonus_roll_for_spare
=== RUN   TestRoll/cannot_roll_after_bonus_rolls_for_strike
--- PASS: TestRoll (0.00s)
    --- PASS: TestRoll/rolls_cannot_score_negative_points (0.00s)
    --- PASS: TestRoll/a_roll_cannot_score_more_than_10_points (0.00s)
    --- PASS: TestRoll/two_rolls_in_a_frame_cannot_score_more_than_10_points (0.00s)
    --- PASS: TestRoll/bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points (0.00s)
    --- PASS: TestRoll/two_bonus_rolls_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points (0.00s)
    --- PASS: TestRoll/the_second_bonus_rolls_after_a_strike_in_the_last_frame_cannot_be_a_strike_if_the_first_one_is_not_a_strike (0.00s)
    --- PASS: TestRoll/second_bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points (0.00s)
    --- PASS: TestRoll/cannot_roll_if_game_already_has_ten_frames (0.00s)
    --- PASS: TestRoll/cannot_roll_after_bonus_roll_for_spare (0.00s)
    --- PASS: TestRoll/cannot_roll_after_bonus_rolls_for_strike (0.00s)
=== RUN   TestScore
=== RUN   TestScore/should_be_able_to_score_a_game_with_all_zeros
=== RUN   TestScore/should_be_able_to_score_a_game_with_no_strikes_or_spares
=== RUN   TestScore/a_spare_followed_by_zeros_is_worth_ten_points
=== RUN   TestScore/points_scored_in_the_roll_after_a_spare_are_counted_twice
=== RUN   TestScore/consecutive_spares_each_get_a_one_roll_bonus
=== RUN   TestScore/a_spare_in_the_last_frame_gets_a_one_roll_bonus_that_is_counted_once
=== RUN   TestScore/a_strike_earns_ten_points_in_a_frame_with_a_single_roll
=== RUN   TestScore/points_scored_in_the_two_rolls_after_a_strike_are_counted_twice_as_a_bonus
=== RUN   TestScore/consecutive_strikes_each_get_the_two_roll_bonus
=== RUN   TestScore/a_strike_in_the_last_frame_gets_a_two_roll_bonus_that_is_counted_once
=== RUN   TestScore/rolling_a_spare_with_the_two_roll_bonus_does_not_get_a_bonus_roll
=== RUN   TestScore/strikes_with_the_two_roll_bonus_do_not_get_bonus_rolls
=== RUN   TestScore/last_two_strikes_followed_by_only_last_bonus_with_non_strike_points
=== RUN   TestScore/a_strike_with_the_one_roll_bonus_after_a_spare_in_the_last_frame_does_not_get_a_bonus
=== RUN   TestScore/all_strikes_is_a_perfect_game
=== RUN   TestScore/two_bonus_rolls_after_a_strike_in_the_last_frame_can_score_more_than_10_points_if_one_is_a_strike
=== RUN   TestScore/an_unstarted_game_cannot_be_scored
=== RUN   TestScore/an_incomplete_game_cannot_be_scored
=== RUN   TestScore/bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated
=== RUN   TestScore/both_bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated
=== RUN   TestScore/bonus_roll_for_a_spare_in_the_last_frame_must_be_rolled_before_score_can_be_calculated
--- PASS: TestScore (0.00s)
    --- PASS: TestScore/should_be_able_to_score_a_game_with_all_zeros (0.00s)
    --- PASS: TestScore/should_be_able_to_score_a_game_with_no_strikes_or_spares (0.00s)
    --- PASS: TestScore/a_spare_followed_by_zeros_is_worth_ten_points (0.00s)
    --- PASS: TestScore/points_scored_in_the_roll_after_a_spare_are_counted_twice (0.00s)
    --- PASS: TestScore/consecutive_spares_each_get_a_one_roll_bonus (0.00s)
    --- PASS: TestScore/a_spare_in_the_last_frame_gets_a_one_roll_bonus_that_is_counted_once (0.00s)
    --- PASS: TestScore/a_strike_earns_ten_points_in_a_frame_with_a_single_roll (0.00s)
    --- PASS: TestScore/points_scored_in_the_two_rolls_after_a_strike_are_counted_twice_as_a_bonus (0.00s)
    --- PASS: TestScore/consecutive_strikes_each_get_the_two_roll_bonus (0.00s)
    --- PASS: TestScore/a_strike_in_the_last_frame_gets_a_two_roll_bonus_that_is_counted_once (0.00s)
    --- PASS: TestScore/rolling_a_spare_with_the_two_roll_bonus_does_not_get_a_bonus_roll (0.00s)
    --- PASS: TestScore/strikes_with_the_two_roll_bonus_do_not_get_bonus_rolls (0.00s)
    --- PASS: TestScore/last_two_strikes_followed_by_only_last_bonus_with_non_strike_points (0.00s)
    --- PASS: TestScore/a_strike_with_the_one_roll_bonus_after_a_spare_in_the_last_frame_does_not_get_a_bonus (0.00s)
    --- PASS: TestScore/all_strikes_is_a_perfect_game (0.00s)
    --- PASS: TestScore/two_bonus_rolls_after_a_strike_in_the_last_frame_can_score_more_than_10_points_if_one_is_a_strike (0.00s)
    --- PASS: TestScore/an_unstarted_game_cannot_be_scored (0.00s)
    --- PASS: TestScore/an_incomplete_game_cannot_be_scored (0.00s)
    --- PASS: TestScore/bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated (0.00s)
    --- PASS: TestScore/both_bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated (0.00s)
    --- PASS: TestScore/bonus_roll_for_a_spare_in_the_last_frame_must_be_rolled_before_score_can_be_calculated (0.00s)
PASS
ok  	bowling	0.005s
```

## Summary

| Metric | Value |
|--------|-------|
| **Overall Result** | PASS |
| **Total Test Suites** | 2 (TestRoll, TestScore) |
| **Total Test Cases** | 31 |
| **Passed** | 31 |
| **Failed** | 0 |
| **Execution Time** | 0.005s |

### Breakdown

- **TestRoll** (10 subtests): All 10 PASSED - covers input validation (negative pins, too many pins, frame limits, bonus roll rules)
- **TestScore** (21 subtests): All 21 PASSED - covers scoring logic (zeros, no strikes/spares, spares, strikes, consecutive strikes/spares, perfect game, last frame bonus rules, incomplete game errors)
