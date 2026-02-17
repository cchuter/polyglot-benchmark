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

func translateWord(word string) string {
	// Rule 1: starts with vowel, "xr", or "yt"
	if isVowel(word[0]) || strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") {
		return word + "ay"
	}

	// Scan consonant cluster to find the split point
	for i := 0; i < len(word); i++ {
		// Rule 3: consonant(s) + "qu"
		if word[i] == 'q' && i+1 < len(word) && word[i+1] == 'u' {
			return word[i+2:] + word[:i+2] + "ay"
		}
		// Rule 4: consonant(s) + "y" (y not at position 0)
		if word[i] == 'y' && i > 0 {
			return word[i:] + word[:i] + "ay"
		}
		// Rule 2: found a vowel, split here
		if isVowel(word[i]) {
			return word[i:] + word[:i] + "ay"
		}
	}
	// Fallback: all consonants
	return word + "ay"
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
