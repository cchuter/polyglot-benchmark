package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	// Step 1: Normalize - keep only alphanumeric, lowercase
	var norm []byte
	for _, ch := range pt {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			norm = append(norm, byte(unicode.ToLower(ch)))
		}
	}
	n := len(norm)
	if n == 0 {
		return ""
	}

	// Step 2: Compute rectangle dimensions
	c := int(math.Ceil(math.Sqrt(float64(n))))
	r := int(math.Ceil(float64(n) / float64(c)))

	// Step 3: Read columns and build output
	var b strings.Builder
	for col := 0; col < c; col++ {
		if col > 0 {
			b.WriteByte(' ')
		}
		for row := 0; row < r; row++ {
			idx := row*c + col
			if idx < n {
				b.WriteByte(norm[idx])
			} else {
				b.WriteByte(' ')
			}
		}
	}
	return b.String()
}
