package piglatin

import (
	"regexp"
	"strings"
)

var vowel = regexp.MustCompile(`^([aeiou]|y[^aeiou]|xr)[a-z]*`)
var containsy = regexp.MustCompile(`^([^aeiou]+)y([a-z]*)`)
var cons = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`)

func Word(s string) (result string) {
	if m := containsy.FindStringSubmatch(s); m != nil {
		return "y" + m[2] + m[1] + "ay"
	}
	if vowel.MatchString(s) {
		return s + "ay"
	}
	if m := cons.FindStringSubmatch(s); m != nil {
		return m[2] + m[1] + "ay"
	}
	return s
}

func Sentence(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		words[i] = Word(strings.ToLower(w))
	}
	return strings.Join(words, " ")
}
