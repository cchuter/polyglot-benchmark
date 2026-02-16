package alphametics

import (
	"errors"
	"strings"
)

type colInfo struct {
	letterCoeffs map[byte]int // letter -> net coefficient in this column
	newLetters   []byte       // letters first encountered in this column
}

// Solve solves an alphametics puzzle and returns a mapping of letters to digits.
func Solve(puzzle string) (map[string]int, error) {
	parts := strings.SplitN(puzzle, "==", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid puzzle format")
	}

	lhsStr := strings.TrimSpace(parts[0])
	rhsWord := strings.TrimSpace(parts[1])

	var addends []string
	for _, w := range strings.Split(lhsStr, "+") {
		addends = append(addends, strings.TrimSpace(w))
	}

	// Collect all words (addends + result) and unique letters.
	allWords := make([]string, len(addends)+1)
	copy(allWords, addends)
	allWords[len(addends)] = rhsWord

	var leading [256]bool
	var seen [256]bool
	var letters []byte

	for _, word := range allWords {
		if len(word) > 1 {
			leading[word[0]] = true
		}
		for i := 0; i < len(word); i++ {
			c := word[i]
			if !seen[c] {
				seen[c] = true
				letters = append(letters, c)
			}
		}
	}

	// Build columns right-to-left. For each column position, compute a
	// coefficient for each letter: +1 for each LHS occurrence, -1 for RHS.
	maxLen := len(rhsWord)
	for _, w := range addends {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}
	// Add one extra column for potential carry overflow.
	numCols := maxLen + 1

	assignedSet := make(map[byte]bool)
	cols := make([]colInfo, numCols)

	for c := 0; c < numCols; c++ {
		cols[c].letterCoeffs = make(map[byte]int)
	}

	// Process columns right-to-left, tracking which letters are new per column.
	for col := 0; col < numCols; col++ {
		newInCol := make(map[byte]bool)
		// LHS addends.
		for _, w := range addends {
			idx := len(w) - 1 - col
			if idx >= 0 {
				ch := w[idx]
				cols[col].letterCoeffs[ch]++
				if !assignedSet[ch] {
					newInCol[ch] = true
				}
			}
		}
		// RHS (result) word.
		idx := len(rhsWord) - 1 - col
		if idx >= 0 {
			ch := rhsWord[idx]
			cols[col].letterCoeffs[ch]--
			if !assignedSet[ch] {
				newInCol[ch] = true
			}
		}

		for ch := range newInCol {
			cols[col].newLetters = append(cols[col].newLetters, ch)
			assignedSet[ch] = true
		}
	}

	// Solver state.
	var assignment [256]int
	for i := range assignment {
		assignment[i] = -1
	}
	var used [10]bool

	// Recursive solver: assign new letters per column, then check constraint.
	var solveCol func(col, carry int) bool
	var assign func(col, carry, idx int) bool

	solveCol = func(col, carry int) bool {
		if col == numCols {
			return carry == 0
		}
		return assign(col, carry, 0)
	}

	assign = func(col, carry, idx int) bool {
		if idx == len(cols[col].newLetters) {
			// All letters in this column assigned; check column constraint.
			colSum := carry
			for ch, coeff := range cols[col].letterCoeffs {
				colSum += coeff * assignment[ch]
			}
			if colSum < 0 || colSum%10 != 0 {
				return false
			}
			return solveCol(col+1, colSum/10)
		}

		ch := cols[col].newLetters[idx]
		start := 0
		if leading[ch] {
			start = 1
		}
		for d := start; d <= 9; d++ {
			if !used[d] {
				assignment[ch] = d
				used[d] = true
				if assign(col, carry, idx+1) {
					return true
				}
				used[d] = false
				assignment[ch] = -1
			}
		}
		return false
	}

	if !solveCol(0, 0) {
		return nil, errors.New("no solution found")
	}

	result := make(map[string]int)
	for _, ch := range letters {
		result[string(ch)] = assignment[ch]
	}
	return result, nil
}
