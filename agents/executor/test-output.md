# Executor: Build and Test Results

## Command
```
cd go/exercises/practice/bowling && go test -v ./...
```

## Result: ALL PASS

### TestRoll (10/10 passed)
- rolls_cannot_score_negative_points ✓
- a_roll_cannot_score_more_than_10_points ✓
- two_rolls_in_a_frame_cannot_score_more_than_10_points ✓
- bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points ✓
- two_bonus_rolls_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points ✓
- the_second_bonus_rolls_after_a_strike_in_the_last_frame_cannot_be_a_strike_if_the_first_one_is_not_a_strike ✓
- second_bonus_roll_after_a_strike_in_the_last_frame_cannot_score_more_than_10_points ✓
- cannot_roll_if_game_already_has_ten_frames ✓
- cannot_roll_after_bonus_roll_for_spare ✓
- cannot_roll_after_bonus_rolls_for_strike ✓

### TestScore (21/21 passed)
- should_be_able_to_score_a_game_with_all_zeros ✓
- should_be_able_to_score_a_game_with_no_strikes_or_spares ✓
- a_spare_followed_by_zeros_is_worth_ten_points ✓
- points_scored_in_the_roll_after_a_spare_are_counted_twice ✓
- consecutive_spares_each_get_a_one_roll_bonus ✓
- a_spare_in_the_last_frame_gets_a_one_roll_bonus_that_is_counted_once ✓
- a_strike_earns_ten_points_in_a_frame_with_a_single_roll ✓
- points_scored_in_the_two_rolls_after_a_strike_are_counted_twice_as_a_bonus ✓
- consecutive_strikes_each_get_the_two_roll_bonus ✓
- a_strike_in_the_last_frame_gets_a_two_roll_bonus_that_is_counted_once ✓
- rolling_a_spare_with_the_two_roll_bonus_does_not_get_a_bonus_roll ✓
- strikes_with_the_two_roll_bonus_do_not_get_bonus_rolls ✓
- last_two_strikes_followed_by_only_last_bonus_with_non_strike_points ✓
- a_strike_with_the_one_roll_bonus_after_a_spare_in_the_last_frame_does_not_get_a_bonus ✓
- all_strikes_is_a_perfect_game ✓
- two_bonus_rolls_after_a_strike_in_the_last_frame_can_score_more_than_10_points_if_one_is_a_strike ✓
- an_unstarted_game_cannot_be_scored ✓
- an_incomplete_game_cannot_be_scored ✓
- bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated ✓
- both_bonus_rolls_for_a_strike_in_the_last_frame_must_be_rolled_before_score_can_be_calculated ✓
- bonus_roll_for_a_spare_in_the_last_frame_must_be_rolled_before_score_can_be_calculated ✓

## Summary
31/31 tests passed. Build successful. No errors.
