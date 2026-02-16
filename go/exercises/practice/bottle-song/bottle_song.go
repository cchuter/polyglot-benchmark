package bottlesong

import (
	"fmt"
	"strings"
)

var numberToWord = map[int]string{
	0: "no", 1: "one", 2: "two", 3: "three", 4: "four",
	5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
}

func titleCase(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}

func bottleStr(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func verse(n int) []string {
	current := titleCase(numberToWord[n])
	next := numberToWord[n-1]
	return []string{
		fmt.Sprintf("%s green %s hanging on the wall,", current, bottleStr(n)),
		fmt.Sprintf("%s green %s hanging on the wall,", current, bottleStr(n)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", next, bottleStr(n-1)),
	}
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := startBottles; i > startBottles-takeDown; i-- {
		if len(result) > 0 {
			result = append(result, "")
		}
		result = append(result, verse(i)...)
	}
	return result
}
