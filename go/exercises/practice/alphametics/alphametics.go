package alphametics

import (
	"errors"
	"sort"
	"strings"
)

// Solve solves an alphametics puzzle and returns a mapping of letters to digits.
func Solve(puzzle string) (map[string]int, error) {
	tokens := strings.Fields(puzzle)
	var addends []string
	var result string
	seenEquals := false
	for _, tok := range tokens {
		switch tok {
		case "+":
			continue
		case "==":
			seenEquals = true
			continue
		}
		if seenEquals {
			result = tok
		} else {
			addends = append(addends, tok)
		}
	}

	// Compute coefficients for each letter.
	var coef [26]int
	var seen [26]bool
	for _, word := range addends {
		pow := 1
		for i := len(word) - 1; i >= 0; i-- {
			idx := word[i] - 'A'
			coef[idx] += pow
			seen[idx] = true
			pow *= 10
		}
	}
	pow := 1
	for i := len(result) - 1; i >= 0; i-- {
		idx := result[i] - 'A'
		coef[idx] -= pow
		seen[idx] = true
		pow *= 10
	}

	// Identify leading letters of multi-digit words.
	var leading [26]bool
	for _, word := range addends {
		if len(word) > 1 {
			leading[word[0]-'A'] = true
		}
	}
	if len(result) > 1 {
		leading[result[0]-'A'] = true
	}

	// Collect unique letters.
	var letters []int
	for i := 0; i < 26; i++ {
		if seen[i] {
			letters = append(letters, i)
		}
	}
	n := len(letters)

	// Sort by descending absolute coefficient for better pruning.
	sort.Slice(letters, func(i, j int) bool {
		ai := coef[letters[i]]
		if ai < 0 {
			ai = -ai
		}
		aj := coef[letters[j]]
		if aj < 0 {
			aj = -aj
		}
		return ai > aj
	})

	// Build arrays for the solver.
	c := make([]int, n)
	lead := make([]bool, n)
	for i, l := range letters {
		c[i] = coef[l]
		lead[i] = leading[l]
	}

	assignment := make([]int, n)
	var used [10]bool

	var solve func(idx, partialSum int) bool
	solve = func(idx, partialSum int) bool {
		if idx == n {
			return partialSum == 0
		}
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			if d == 0 && lead[idx] {
				continue
			}
			newSum := partialSum + c[idx]*d
			if idx < n-1 && !canReachZero(newSum, c[idx+1:]) {
				continue
			}
			used[d] = true
			assignment[idx] = d
			if solve(idx+1, newSum) {
				return true
			}
			used[d] = false
		}
		return false
	}

	if !solve(0, 0) {
		return nil, errors.New("no solution found")
	}

	res := make(map[string]int, n)
	for i, l := range letters {
		res[string(rune(l+'A'))] = assignment[i]
	}
	return res, nil
}

// canReachZero checks whether the remaining coefficients can potentially
// be assigned digits such that partialSum + remaining contribution = 0.
func canReachZero(partialSum int, remCoefs []int) bool {
	minSum, maxSum := 0, 0
	for _, c := range remCoefs {
		if c > 0 {
			maxSum += c * 9
		} else {
			minSum += c * 9
		}
	}
	return partialSum+minSum <= 0 && partialSum+maxSum >= 0
}
