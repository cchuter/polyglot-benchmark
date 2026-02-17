package piglatin

import "strings"

// Sentence translates a sentence from English to Pig Latin.
func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	for i, w := range words {
		words[i] = translateWord(w)
	}
	return strings.Join(words, " ")
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func translateWord(word string) string {
	// Rule 1: starts with vowel, "xr", or "yt"
	if isVowel(word[0]) || (len(word) >= 2 && (word[:2] == "xr" || word[:2] == "yt")) {
		return word + "ay"
	}

	// Scan consonant cluster to find split point
	for i := 0; i < len(word); i++ {
		// Rule 3: "qu" — include both in prefix
		if word[i] == 'q' && i+1 < len(word) && word[i+1] == 'u' {
			return word[i+2:] + word[:i+2] + "ay"
		}
		// Rule 4: "y" after consonants acts as vowel
		if word[i] == 'y' && i > 0 {
			return word[i:] + word[:i] + "ay"
		}
		// Rule 2: hit a vowel — split here
		if isVowel(word[i]) {
			return word[i:] + word[:i] + "ay"
		}
	}
	// Fallback for all-consonant words
	return word + "ay"
}
