package alphametics

import (
	"errors"
	"strings"
)

// Solve solves an alphametics puzzle and returns a mapping of letters to digits.
func Solve(puzzle string) (map[string]int, error) {
	sides := strings.SplitN(puzzle, "==", 2)
	if len(sides) != 2 {
		return nil, errors.New("invalid puzzle format")
	}

	addendStrs := strings.Split(sides[0], "+")
	result := strings.TrimSpace(sides[1])

	words := make([]string, 0, len(addendStrs)+1)
	for _, s := range addendStrs {
		words = append(words, strings.TrimSpace(s))
	}
	words = append(words, result)

	// Collect unique letters and identify leading letters
	seen := make(map[byte]bool)
	leading := make(map[byte]bool)
	var letters []byte

	for _, w := range words {
		if len(w) > 1 {
			leading[w[0]] = true
		}
		for i := 0; i < len(w); i++ {
			if !seen[w[i]] {
				seen[w[i]] = true
				letters = append(letters, w[i])
			}
		}
	}

	// Compute coefficients: addends contribute positively, result negatively
	coeffs := make(map[byte]int)
	for _, w := range addendStrs {
		w = strings.TrimSpace(w)
		power := 1
		for i := len(w) - 1; i >= 0; i-- {
			coeffs[w[i]] += power
			power *= 10
		}
	}
	power := 1
	for i := len(result) - 1; i >= 0; i-- {
		coeffs[result[i]] -= power
		power *= 10
	}

	// Backtracking search
	assignment := make(map[byte]int)
	var used [10]bool

	if search(letters, coeffs, leading, &used, assignment, 0) {
		out := make(map[string]int, len(letters))
		for _, l := range letters {
			out[string(l)] = assignment[l]
		}
		return out, nil
	}

	return nil, errors.New("no solution found")
}

func search(letters []byte, coeffs map[byte]int, leading map[byte]bool, used *[10]bool, assignment map[byte]int, idx int) bool {
	if idx == len(letters) {
		sum := 0
		for l, c := range coeffs {
			sum += c * assignment[l]
		}
		return sum == 0
	}

	letter := letters[idx]
	start := 0
	if leading[letter] {
		start = 1
	}

	for d := start; d <= 9; d++ {
		if !used[d] {
			used[d] = true
			assignment[letter] = d
			if search(letters, coeffs, leading, used, assignment, idx+1) {
				return true
			}
			delete(assignment, letter)
			used[d] = false
		}
	}
	return false
}
