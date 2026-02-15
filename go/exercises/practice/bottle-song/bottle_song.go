package bottlesong

import "fmt"

var numberToWord = map[int]string{
	0: "no", 1: "one", 2: "two", 3: "three", 4: "four",
	5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
}

func bottleStr(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

func verse(n int) []string {
	return []string{
		fmt.Sprintf("%s green %s hanging on the wall,", Title(numberToWord[n]), bottleStr(n)),
		fmt.Sprintf("%s green %s hanging on the wall,", Title(numberToWord[n]), bottleStr(n)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %s green %s hanging on the wall.", numberToWord[n-1], bottleStr(n-1)),
	}
}

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := startBottles; i > startBottles-takeDown; i-- {
		result = append(result, verse(i)...)
		if i > startBottles-takeDown+1 {
			result = append(result, "")
		}
	}
	return result
}
