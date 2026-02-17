package bottlesong

import "strings"

var numberWords = []string{
	"no", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine", "ten",
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := 0; i < takeDown; i++ {
		if i > 0 {
			result = append(result, "")
		}
		current := startBottles - i
		next := current - 1
		result = append(result,
			capitalize(numberWords[current])+" green "+plural(current)+" hanging on the wall,",
			capitalize(numberWords[current])+" green "+plural(current)+" hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be "+numberWords[next]+" green "+plural(next)+" hanging on the wall.",
		)
	}
	return result
}

func plural(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
