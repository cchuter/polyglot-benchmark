# Bowling Exercise - Test Results

## Summary

- **Status**: ALL TESTS PASS
- **Total Tests**: 31 (10 Roll tests + 21 Score tests)
- **Passed**: 31
- **Failed**: 0
- **Package**: `bowling`

## Command

```
cd /tmp/osmi-bench-polyglot-go-bowling-177692672/go/exercises/practice/bowling && go test -v ./...
```

## Test Output

### TestRoll (10/10 PASS)

| Test | Result |
|------|--------|
| rolls_cannot_score_negative_points | PASS |
| a_roll_cannot_score_more_than_10_points | PASS |
| two_rolls_in_a_frame_cannot_score_more_than_10_points | PASS |
| bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points | PASS |
| two_bonus_rolls_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points | PASS |
| the_second_bonus_rolls_after_a_strike_in_the_last_frame_cannot_be_a_strike_if_the_first_one_is_not_a_strike | PASS |
| second_bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points | PASS |
| cannot_roll_if_game_already_has_ten_frames | PASS |
| cannot_roll_after_bonus_roll_for_spare | PASS |
| cannot_roll_after_bonus_rolls_for_strike | PASS |

### TestScore (21/21 PASS)

| Test | Result |
|------|--------|
| should_be_able_to_score_a_game_with_all_zeros | PASS |
| should_be_able_to_score_a_game_with_no_strikes_or_spares | PASS |
| a_spare_followed_by_zeros_is_worth_ten_points | PASS |
| points_scored_in_the_roll_after_a_spare_are_counted_twice | PASS |
| consecutive_spares_each_get_a_one_roll_bonus | PASS |
| a_spare_in_the_last_frame_gets_a_one_roll_bonus_that_is_counted_once | PASS |
| a_strike_earns_ten_points_in_a_frame_with_a_single_roll | PASS |
| points_scored_in_the_two_rolls_after_a_strike_are_counted_twice_as_a_bonus | PASS |
| consecutive_strikes_each_get_the_two_roll_bonus | PASS |
| a_strike_in_the_last_frame_gets_a_two_roll_bonus_that_is_counted_once | PASS |
| rolling_a_spare_with_the_two_roll_bonus_does_not_get_a_bonus_roll | PASS |
| strikes_with_the_two_roll_bonus_do_not_get_bonus_rolls | PASS |
| last_two_strikes_followed_by_only_last_bonus_with_non_strike_points | PASS |
| a_strike_with_the_one_roll_bonus_after_a_spare_in_the_last_frame_does_not_get_a_bonus | PASS |
| all_strikes_is_a_perfect_game | PASS |
| two_bonus_rolls_after_a_strike_in_the_last_frame_can_score_more_than_10_points_if_one_is_a_strike | PASS |
| an_unstarted_game_cannot_be_scored | PASS |
| an_incomplete_game_cannot_be_scored | PASS |
| bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated | PASS |
| both_bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated | PASS |
| bonus_roll_for_a_spare_in_the_last_frame_must_be_rolled_before_score_can_be_calculated | PASS |
