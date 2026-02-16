package alphametics

import (
	"errors"
	"math"
	"sort"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	sides := strings.SplitN(puzzle, "==", 2)
	if len(sides) != 2 {
		return nil, errors.New("invalid puzzle: missing ==")
	}

	addendStrs := strings.Split(sides[0], "+")
	resultStr := strings.TrimSpace(sides[1])

	var addends []string
	for _, a := range addendStrs {
		addends = append(addends, strings.TrimSpace(a))
	}

	// Compute weight for each letter and identify leading letters.
	weights := make(map[byte]int)
	leading := make(map[byte]bool)

	for _, word := range addends {
		if len(word) > 1 {
			leading[word[0]] = true
		}
		for i, ch := range word {
			pow := int(math.Pow(10, float64(len(word)-1-i)))
			weights[byte(ch)] += pow
		}
	}

	if len(resultStr) > 1 {
		leading[resultStr[0]] = true
	}
	for i, ch := range resultStr {
		pow := int(math.Pow(10, float64(len(resultStr)-1-i)))
		weights[byte(ch)] -= pow
	}

	// Collect unique letters sorted by descending absolute weight.
	letters := make([]byte, 0, len(weights))
	for ch := range weights {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool {
		ai := weights[letters[i]]
		if ai < 0 {
			ai = -ai
		}
		aj := weights[letters[j]]
		if aj < 0 {
			aj = -aj
		}
		return ai > aj
	})

	// Build ordered weight and leading arrays for fast access.
	n := len(letters)
	w := make([]int, n)
	isLeading := make([]bool, n)
	for i, ch := range letters {
		w[i] = weights[ch]
		isLeading[i] = leading[ch]
	}

	// Recursive backtracking.
	assignment := make([]int, n)
	var used [10]bool

	var solve func(idx, sum int) bool
	solve = func(idx, sum int) bool {
		if idx == n {
			return sum == 0
		}
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			if d == 0 && isLeading[idx] {
				continue
			}
			used[d] = true
			assignment[idx] = d
			if solve(idx+1, sum+d*w[idx]) {
				return true
			}
			used[d] = false
		}
		return false
	}

	if !solve(0, 0) {
		return nil, errors.New("no solution found")
	}

	result := make(map[string]int, n)
	for i, ch := range letters {
		result[string(ch)] = assignment[i]
	}
	return result, nil
}
