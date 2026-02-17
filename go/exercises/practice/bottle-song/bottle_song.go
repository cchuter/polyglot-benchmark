package bottlesong

import "fmt"

func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := 0; i < takeDown; i++ {
		n := startBottles - i
		if i > 0 {
			result = append(result, "")
		}
		word := Title(numberWord(n))
		nextWord := numberWord(n - 1)
		result = append(result,
			fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
			fmt.Sprintf("%s green %s hanging on the wall,", word, bottleStr(n)),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", nextWord, bottleStr(n-1)),
		)
	}
	return result
}

var numbers = []string{
	"no", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten",
}

func numberWord(n int) string {
	return numbers[n]
}

func bottleStr(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}
