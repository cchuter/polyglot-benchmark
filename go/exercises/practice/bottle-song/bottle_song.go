package bottlesong

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var numberWord = []string{
	"no", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine", "ten",
}

func capitalize(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}

func bottleStr(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func Recite(startBottles int, takeDown int) []string {
	var result []string
	for i := startBottles; i > startBottles-takeDown; i-- {
		if len(result) > 0 {
			result = append(result, "")
		}
		cur := capitalize(numberWord[i])
		next := numberWord[i-1]
		result = append(result,
			fmt.Sprintf("%s green %s hanging on the wall,", cur, bottleStr(i)),
			fmt.Sprintf("%s green %s hanging on the wall,", cur, bottleStr(i)),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", next, bottleStr(i-1)),
		)
	}
	return result
}
