package bottlesong

import (
	"fmt"
	"strings"
)

var numberToWord = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func bottleWord(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func verse(n int) []string {
	word := numberToWord[n]
	nextWord := numberToWord[n-1]
	return []string{
		fmt.Sprintf("%s green %s hanging on the wall,", capitalize(word), bottleWord(n)),
		fmt.Sprintf("%s green %s hanging on the wall,", capitalize(word), bottleWord(n)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", nextWord, bottleWord(n-1)),
	}
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := startBottles; i > startBottles-takeDown; i-- {
		if i < startBottles {
			result = append(result, "")
		}
		result = append(result, verse(i)...)
	}
	return result
}
