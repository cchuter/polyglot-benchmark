package alphametics

import (
	"errors"
	"sort"
	"strings"
)

// Solve finds digit assignments for an alphametics puzzle such that
// the arithmetic equation holds, each letter maps to a unique digit,
// and no multi-digit number has a leading zero.
func Solve(puzzle string) (map[string]int, error) {
	// Parse: split on whitespace, discard "+" and "=="
	tokens := strings.Fields(puzzle)
	var words []string
	for _, t := range tokens {
		if t != "+" && t != "==" {
			words = append(words, t)
		}
	}
	if len(words) < 2 {
		return nil, errors.New("invalid puzzle")
	}

	// Compute letter coefficients.
	// Addends (all words except last) contribute positively,
	// result word (last) contributes negatively.
	// A correct assignment makes the total sum zero.
	coeffMap := make(map[byte]int)
	leading := make(map[byte]bool)
	for i, word := range words {
		sign := 1
		if i == len(words)-1 {
			sign = -1
		}
		place := 1
		for j := len(word) - 1; j >= 0; j-- {
			coeffMap[word[j]] += sign * place
			place *= 10
		}
		if len(word) > 1 {
			leading[word[0]] = true
		}
	}

	// Sort letters by |coefficient| descending for better pruning
	letters := make([]byte, 0, len(coeffMap))
	for ch := range coeffMap {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool {
		ai := coeffMap[letters[i]]
		if ai < 0 {
			ai = -ai
		}
		aj := coeffMap[letters[j]]
		if aj < 0 {
			aj = -aj
		}
		return ai > aj
	})

	n := len(letters)
	coeff := make([]int, n)
	isLeading := make([]bool, n)
	for i, ch := range letters {
		coeff[i] = coeffMap[ch]
		isLeading[i] = leading[ch]
	}

	// Backtracking solver with bounds pruning
	assignment := make([]int, n)

	var solve func(idx, sum, used int) bool
	solve = func(idx, sum, used int) bool {
		if idx == n {
			return sum == 0
		}
		for d := 0; d <= 9; d++ {
			if used&(1<<d) != 0 {
				continue
			}
			if d == 0 && isLeading[idx] {
				continue
			}
			newSum := sum + coeff[idx]*d
			newUsed := used | (1 << d)

			// Bounds pruning: check if remaining letters can bring sum to 0
			if idx < n-1 {
				lo, hi := 0, 0
				feasible := true
				for k := idx + 1; k < n; k++ {
					minD, maxD := availMinMax(newUsed, isLeading[k])
					if minD < 0 {
						feasible = false
						break
					}
					c := coeff[k]
					if c > 0 {
						lo += c * minD
						hi += c * maxD
					} else {
						lo += c * maxD
						hi += c * minD
					}
				}
				if !feasible || newSum+lo > 0 || newSum+hi < 0 {
					continue
				}
			}

			assignment[idx] = d
			if solve(idx+1, newSum, newUsed) {
				return true
			}
		}
		return false
	}

	if !solve(0, 0, 0) {
		return nil, errors.New("no solution found")
	}

	result := make(map[string]int, n)
	for i, ch := range letters {
		result[string(ch)] = assignment[i]
	}
	return result, nil
}

// availMinMax returns the smallest and largest available digits given a
// used-digit bitmask. If isLeading is true, digit 0 is excluded.
// Returns (-1, -1) if no digit is available.
func availMinMax(used int, isLeading bool) (int, int) {
	start := 0
	if isLeading {
		start = 1
	}
	minD := -1
	for d := start; d <= 9; d++ {
		if used&(1<<d) == 0 {
			minD = d
			break
		}
	}
	maxD := -1
	for d := 9; d >= start; d-- {
		if used&(1<<d) == 0 {
			maxD = d
			break
		}
	}
	return minD, maxD
}
