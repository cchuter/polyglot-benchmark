package piglatin

import "strings"

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
var specials = map[string]bool{"xr": true, "yt": true}
var vowelsY = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// Sentence translates a phrase from English to Pig Latin.
func Sentence(phrase string) string {
	words := strings.Fields(phrase)
	for i, word := range words {
		words[i] = translateWord(word)
	}
	return strings.Join(words, " ")
}

func translateWord(word string) string {
	// Rule 1: starts with vowel, "xr", or "yt"
	if vowels[word[0]] || (len(word) >= 2 && specials[word[:2]]) {
		return word + "ay"
	}

	// Scan consonant cluster to find split point
	// Handles Rule 2 (consonants), Rule 3 (qu), and Rule 4 (y as vowel)
	for pos := 1; pos < len(word); pos++ {
		letter := word[pos]
		if vowelsY[letter] {
			// Rule 3: if we hit 'u' preceded by 'q', include 'u' in consonant prefix
			if letter == 'u' && word[pos-1] == 'q' {
				pos++
			}
			return word[pos:] + word[:pos] + "ay"
		}
	}

	return word + "ay"
}
