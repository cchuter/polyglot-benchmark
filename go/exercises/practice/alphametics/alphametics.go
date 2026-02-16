package alphametics

import (
	"errors"
	"sort"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	sides := strings.Split(puzzle, "==")
	if len(sides) != 2 {
		return nil, errors.New("invalid puzzle")
	}

	rhsWord := strings.TrimSpace(sides[1])
	lhsParts := strings.Split(sides[0], "+")
	addends := make([]string, 0, len(lhsParts))
	for _, p := range lhsParts {
		p = strings.TrimSpace(p)
		if p != "" {
			addends = append(addends, p)
		}
	}

	allWords := make([]string, 0, len(addends)+1)
	allWords = append(allWords, addends...)
	allWords = append(allWords, rhsWord)

	// Leading letters of multi-char words cannot be zero
	leading := make(map[byte]bool)
	for _, w := range allWords {
		if len(w) > 1 {
			leading[w[0]] = true
		}
	}

	// Compute weight for each letter.
	// Addend letters get positive place values, result letters get negative.
	// A valid solution satisfies: sum(weight[ch] * digit[ch]) == 0
	weights := make(map[byte]int)
	for _, w := range addends {
		pv := 1
		for i := len(w) - 1; i >= 0; i-- {
			weights[w[i]] += pv
			pv *= 10
		}
	}
	pv := 1
	for i := len(rhsWord) - 1; i >= 0; i-- {
		weights[rhsWord[i]] -= pv
		pv *= 10
	}

	// Collect unique letters, sorted by absolute weight descending for better pruning
	letters := make([]byte, 0, len(weights))
	for ch := range weights {
		letters = append(letters, ch)
	}
	sort.Slice(letters, func(i, j int) bool {
		wi := weights[letters[i]]
		if wi < 0 {
			wi = -wi
		}
		wj := weights[letters[j]]
		if wj < 0 {
			wj = -wj
		}
		return wi > wj
	})

	n := len(letters)
	w := make([]int, n)
	isLeading := make([]bool, n)
	for i, ch := range letters {
		w[i] = weights[ch]
		isLeading[i] = leading[ch]
	}

	assignment := make([]int, n)
	used := [10]bool{}

	var solve func(idx, sum int) bool
	solve = func(idx, sum int) bool {
		if idx == n {
			return sum == 0
		}

		// Find min and max available digits for pruning bounds
		minAvail, maxAvail := -1, -1
		for d := 0; d <= 9; d++ {
			if !used[d] {
				if minAvail == -1 {
					minAvail = d
				}
				maxAvail = d
			}
		}

		// Compute loose bounds on remaining weighted sum
		remMin, remMax := 0, 0
		for j := idx; j < n; j++ {
			if w[j] >= 0 {
				remMin += w[j] * minAvail
				remMax += w[j] * maxAvail
			} else {
				remMin += w[j] * maxAvail
				remMax += w[j] * minAvail
			}
		}

		if sum+remMin > 0 || sum+remMax < 0 {
			return false
		}

		startD := 0
		if isLeading[idx] {
			startD = 1
		}

		for d := startD; d <= 9; d++ {
			if used[d] {
				continue
			}
			used[d] = true
			assignment[idx] = d
			if solve(idx+1, sum+w[idx]*d) {
				return true
			}
			used[d] = false
		}
		return false
	}

	if solve(0, 0) {
		result := make(map[string]int)
		for i, ch := range letters {
			result[string(ch)] = assignment[i]
		}
		return result, nil
	}

	return nil, errors.New("no solution found")
}
