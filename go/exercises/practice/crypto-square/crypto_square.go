package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// Normalize: lowercase and keep only alphanumeric characters
	var normalized strings.Builder
	for _, ch := range pt {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			normalized.WriteRune(unicode.ToLower(ch))
		}
	}
	s := normalized.String()
	if len(s) == 0 {
		return ""
	}

	// Compute dimensions
	c := int(math.Ceil(math.Sqrt(float64(len(s)))))
	r := int(math.Ceil(float64(len(s)) / float64(c)))

	// Read down columns and build output chunks
	chunks := make([]string, c)
	for col := 0; col < c; col++ {
		var b strings.Builder
		for row := 0; row < r; row++ {
			pos := row*c + col
			if pos < len(s) {
				b.WriteByte(s[pos])
			} else {
				b.WriteByte(' ')
			}
		}
		chunks[col] = b.String()
	}

	return strings.Join(chunks, " ")
}
