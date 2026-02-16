package piglatin

import "strings"

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		words[i] = translateWord(w)
	}
	return strings.Join(words, " ")
}

func translateWord(word string) string {
	// Rule 1: starts with vowel, "xr", or "yt"
	if isVowel(word[0]) || (len(word) >= 2 && (word[:2] == "xr" || word[:2] == "yt")) {
		return word + "ay"
	}

	// Scan consonant cluster for Rules 2/3/4
	for i := 0; i < len(word); i++ {
		// Rule 3: "qu" — include the "qu" in moved prefix
		if word[i] == 'q' && i+1 < len(word) && word[i+1] == 'u' {
			return word[i+2:] + word[:i+2] + "ay"
		}
		// Rule 4: "y" after consonants
		if word[i] == 'y' && i > 0 {
			return word[i:] + word[:i] + "ay"
		}
		// Vowel found — Rule 2: split here
		if isVowel(word[i]) {
			return word[i:] + word[:i] + "ay"
		}
	}
	return word + "ay"
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
