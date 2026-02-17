package piglatin

import (
	"regexp"
	"strings"
)

var vowel = regexp.MustCompile(`^([aeiou]|y[^aeiou]|xr)[a-z]*`)
var cons = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`)
var containsy = regexp.MustCompile(`^([^aeiou]+)y([a-z]*)`)

// Sentence translates a sentence from English to Pig Latin.
func Sentence(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = Word(strings.ToLower(w))
	}
	return strings.Join(words, " ")
}

// Word translates a single word from English to Pig Latin.
func Word(s string) string {
	// Rule 4: consonant cluster followed by y
	if containsy.MatchString(s) {
		pos := containsy.FindStringSubmatchIndex(s)
		return s[pos[3]:] + s[:pos[3]] + "ay"
	}
	// Rule 1: starts with vowel, xr, or yt
	if vowel.MatchString(s) {
		return s + "ay"
	}
	// Rules 2 & 3: consonant cluster (with optional qu)
	if x := cons.FindStringSubmatchIndex(s); x != nil {
		return s[x[3]:] + s[:x[3]] + "ay"
	}
	return s
}
