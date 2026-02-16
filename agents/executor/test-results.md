# Pig Latin - Build & Test Results

## Build Status: PASS

```
$ go build ./...
(no errors)
```

## Go Vet Status: PASS

```
$ go vet ./...
(no issues)
```

## Test Results: ALL 22 TESTS PASSED

```
$ go test -v ./...
=== RUN   TestPigLatin
=== RUN   TestPigLatin/word_beginning_with_a
=== RUN   TestPigLatin/word_beginning_with_e
=== RUN   TestPigLatin/word_beginning_with_i
=== RUN   TestPigLatin/word_beginning_with_o
=== RUN   TestPigLatin/word_beginning_with_u
=== RUN   TestPigLatin/word_beginning_with_a_vowel_and_followed_by_a_qu
=== RUN   TestPigLatin/word_beginning_with_p
=== RUN   TestPigLatin/word_beginning_with_k
=== RUN   TestPigLatin/word_beginning_with_x
=== RUN   TestPigLatin/word_beginning_with_q_without_a_following_u
=== RUN   TestPigLatin/word_beginning_with_ch
=== RUN   TestPigLatin/word_beginning_with_qu
=== RUN   TestPigLatin/word_beginning_with_qu_and_a_preceding_consonant
=== RUN   TestPigLatin/word_beginning_with_th
=== RUN   TestPigLatin/word_beginning_with_thr
=== RUN   TestPigLatin/word_beginning_with_sch
=== RUN   TestPigLatin/word_beginning_with_yt
=== RUN   TestPigLatin/word_beginning_with_xr
=== RUN   TestPigLatin/y_is_treated_like_a_consonant_at_the_beginning_of_a_word
=== RUN   TestPigLatin/y_is_treated_like_a_vowel_at_the_end_of_a_consonant_cluster
=== RUN   TestPigLatin/y_as_second_letter_in_two_letter_word
=== RUN   TestPigLatin/a_whole_phrase
--- PASS: TestPigLatin (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_a (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_e (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_i (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_o (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_u (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_a_vowel_and_followed_by_a_qu (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_p (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_k (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_x (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_q_without_a_following_u (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_ch (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_qu (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_qu_and_a_preceding_consonant (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_th (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_thr (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_sch (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_yt (0.00s)
    --- PASS: TestPigLatin/word_beginning_with_xr (0.00s)
    --- PASS: TestPigLatin/y_is_treated_like_a_consonant_at_the_beginning_of_a_word (0.00s)
    --- PASS: TestPigLatin/y_is_treated_like_a_vowel_at_the_end_of_a_consonant_cluster (0.00s)
    --- PASS: TestPigLatin/y_as_second_letter_in_two_letter_word (0.00s)
    --- PASS: TestPigLatin/a_whole_phrase (0.00s)
PASS
ok      piglatin        (cached)
```

## Summary

| Check     | Status |
|-----------|--------|
| Build     | PASS   |
| Vet       | PASS   |
| Tests     | 22/22 PASSED |
| Failures  | 0      |
