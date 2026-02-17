package alphametics

import (
	"errors"
	"strings"
)

// Solve solves an alphametics puzzle and returns a map of letter-to-digit assignments.
func Solve(puzzle string) (map[string]int, error) {
	// Parse puzzle: split on "==" to get LHS and RHS
	parts := strings.SplitN(puzzle, "==", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid puzzle")
	}

	// Extract addend words from LHS (split on "+")
	lhsParts := strings.Split(parts[0], "+")
	var words []string
	for _, p := range lhsParts {
		w := strings.TrimSpace(p)
		if w != "" {
			words = append(words, w)
		}
	}
	result := strings.TrimSpace(parts[1])
	if result == "" || len(words) == 0 {
		return nil, errors.New("invalid puzzle")
	}

	// Collect unique letters and identify leading letters
	letterSet := make(map[byte]bool)
	leadingSet := make(map[byte]bool)

	for _, w := range words {
		for i := 0; i < len(w); i++ {
			letterSet[w[i]] = true
		}
		if len(w) > 1 {
			leadingSet[w[0]] = true
		}
	}
	for i := 0; i < len(result); i++ {
		letterSet[result[i]] = true
	}
	if len(result) > 1 {
		leadingSet[result[0]] = true
	}

	// Build ordered list of unique letters
	letters := make([]byte, 0, len(letterSet))
	for ch := range letterSet {
		letters = append(letters, ch)
	}

	if len(letters) > 10 {
		return nil, errors.New("too many unique letters")
	}

	// Compute coefficients for each letter
	coeffs := make(map[byte]int)
	for _, w := range words {
		pow := 1
		for i := len(w) - 1; i >= 0; i-- {
			coeffs[w[i]] += pow
			pow *= 10
		}
	}
	pow := 1
	for i := len(result) - 1; i >= 0; i-- {
		coeffs[result[i]] -= pow
		pow *= 10
	}

	// Sort letters by descending absolute coefficient for better pruning
	for i := 0; i < len(letters)-1; i++ {
		for j := i + 1; j < len(letters); j++ {
			ai := coeffs[letters[i]]
			if ai < 0 {
				ai = -ai
			}
			aj := coeffs[letters[j]]
			if aj < 0 {
				aj = -aj
			}
			if aj > ai {
				letters[i], letters[j] = letters[j], letters[i]
			}
		}
	}

	// Build arrays for the solver
	n := len(letters)
	coeffArr := make([]int, n)
	isLeading := make([]bool, n)
	for i, ch := range letters {
		coeffArr[i] = coeffs[ch]
		isLeading[i] = leadingSet[ch]
	}

	// Backtracking solver
	assignment := make([]int, n)
	used := uint16(0) // bitmask for digits 0-9

	// computeBounds returns the min and max possible sum contribution
	// from letters at indices [from, n) given the current used bitmask.
	computeBounds := func(from int, usedMask uint16) (int, int) {
		minRem, maxRem := 0, 0
		for i := from; i < n; i++ {
			c := coeffArr[i]
			lo := 0
			if isLeading[i] {
				lo = 1
			}
			minC, maxC := 0, 0
			first := true
			for d := lo; d <= 9; d++ {
				if usedMask&(1<<uint(d)) != 0 {
					continue
				}
				v := c * d
				if first {
					minC, maxC = v, v
					first = false
				}
				if v < minC {
					minC = v
				}
				if v > maxC {
					maxC = v
				}
			}
			minRem += minC
			maxRem += maxC
		}
		return minRem, maxRem
	}

	var solve func(idx int, partialSum int) bool
	solve = func(idx int, partialSum int) bool {
		if idx == n {
			return partialSum == 0
		}

		lo := 0
		if isLeading[idx] {
			lo = 1
		}
		for d := lo; d <= 9; d++ {
			if used&(1<<uint(d)) != 0 {
				continue
			}
			newSum := partialSum + coeffArr[idx]*d

			// Prune: check if remaining letters can bring sum to zero
			if idx < n-1 {
				used |= 1 << uint(d)
				minRem, maxRem := computeBounds(idx+1, used)
				used &^= 1 << uint(d)
				if newSum+minRem > 0 || newSum+maxRem < 0 {
					continue
				}
			}

			assignment[idx] = d
			used |= 1 << uint(d)
			if solve(idx+1, newSum) {
				return true
			}
			used &^= 1 << uint(d)
		}
		return false
	}

	if !solve(0, 0) {
		return nil, errors.New("no solution")
	}

	// Build result map
	resultMap := make(map[string]int, n)
	for i, ch := range letters {
		resultMap[string(ch)] = assignment[i]
	}
	return resultMap, nil
}
