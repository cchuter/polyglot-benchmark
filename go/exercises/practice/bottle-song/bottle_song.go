package bottlesong

import "strings"

var numberWords = []string{
	"no", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine", "ten",
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func bottlePlural(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := 0; i < takeDown; i++ {
		if i > 0 {
			result = append(result, "")
		}
		n := startBottles - i
		word := capitalize(numberWords[n])
		plural := bottlePlural(n)
		result = append(result,
			word+" green "+plural+" hanging on the wall,",
			word+" green "+plural+" hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be "+numberWords[n-1]+" green "+bottlePlural(n-1)+" hanging on the wall.",
		)
	}
	return result
}
