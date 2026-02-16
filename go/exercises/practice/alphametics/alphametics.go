package alphametics

import (
	"errors"
	"strings"
)

// Solve solves an alphametics puzzle and returns a mapping of letters to digits.
func Solve(puzzle string) (map[string]int, error) {
	// Parse the puzzle into words and identify the result word.
	fields := strings.Fields(puzzle)
	var words []string
	for _, f := range fields {
		if f != "+" && f != "==" {
			words = append(words, f)
		}
	}
	if len(words) < 2 {
		return nil, errors.New("invalid puzzle")
	}

	// The last word is the result (right side of ==).
	addends := words[:len(words)-1]
	result := words[len(words)-1]

	// Compute coefficient for each letter.
	// Addend letters get positive place values, result letters get negative.
	coeffs := make(map[byte]int)
	for _, w := range addends {
		placeValue := 1
		for i := len(w) - 1; i >= 0; i-- {
			coeffs[w[i]] += placeValue
			placeValue *= 10
		}
	}
	placeValue := 1
	for i := len(result) - 1; i >= 0; i-- {
		coeffs[result[i]] -= placeValue
		placeValue *= 10
	}

	// Collect unique letters in a stable order.
	seen := make(map[byte]bool)
	var letters []byte
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if !seen[w[i]] {
				seen[w[i]] = true
				letters = append(letters, w[i])
			}
		}
	}

	// Identify leading letters of multi-digit words (cannot be zero).
	leading := make(map[byte]bool)
	for _, w := range words {
		if len(w) > 1 {
			leading[w[0]] = true
		}
	}

	n := len(letters)
	if n > 10 {
		return nil, errors.New("too many unique letters")
	}

	// Build coefficient array in letter order for fast lookup.
	coeffArr := make([]int, n)
	for i, ch := range letters {
		coeffArr[i] = coeffs[ch]
	}

	// Build leading-letter flags array.
	isLeading := make([]bool, n)
	for i, ch := range letters {
		isLeading[i] = leading[ch]
	}

	// Try all permutations of 10 digits taken n at a time.
	digits := make([]int, n)
	used := make([]bool, 10)

	var solve func(pos int) bool
	solve = func(pos int) bool {
		if pos == n {
			// Check if the equation is satisfied.
			sum := 0
			for i := 0; i < n; i++ {
				sum += coeffArr[i] * digits[i]
			}
			return sum == 0
		}
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			if d == 0 && isLeading[pos] {
				continue
			}
			digits[pos] = d
			used[d] = true
			if solve(pos + 1) {
				return true
			}
			used[d] = false
		}
		return false
	}

	if !solve(0) {
		return nil, errors.New("no solution")
	}

	// Build the result map.
	result_map := make(map[string]int, n)
	for i, ch := range letters {
		result_map[string(ch)] = digits[i]
	}
	return result_map, nil
}
