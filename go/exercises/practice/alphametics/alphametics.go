package alphametics

import (
	"errors"
	"sort"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	// Step 1: Parse the puzzle
	sides := strings.SplitN(puzzle, "==", 2)
	if len(sides) != 2 {
		return nil, errors.New("invalid puzzle format")
	}
	lhs := strings.TrimSpace(sides[0])
	rhs := strings.TrimSpace(sides[1])

	addends := strings.Split(lhs, "+")
	for i := range addends {
		addends[i] = strings.TrimSpace(addends[i])
	}
	result := strings.TrimSpace(rhs)

	allWords := append(addends, result)

	// Step 2: Extract unique letters and identify leading letters
	seen := make(map[byte]bool)
	var letters []byte
	for _, word := range allWords {
		for i := 0; i < len(word); i++ {
			if !seen[word[i]] {
				seen[word[i]] = true
				letters = append(letters, word[i])
			}
		}
	}

	leading := make(map[byte]bool)
	for _, word := range allWords {
		if len(word) > 1 {
			leading[word[0]] = true
		}
	}

	n := len(letters)

	// Step 3: Compute letter weights
	weights := make(map[byte]int)
	for _, word := range addends {
		place := 1
		for i := len(word) - 1; i >= 0; i-- {
			weights[word[i]] += place
			place *= 10
		}
	}
	{
		place := 1
		for i := len(result) - 1; i >= 0; i-- {
			weights[result[i]] -= place
			place *= 10
		}
	}

	// Step 4: Sort letters by descending absolute weight
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

	// Build indexed arrays for fast access in recursion
	w := make([]int, n)
	isLeading := make([]bool, n)
	for i, ch := range letters {
		w[i] = weights[ch]
		isLeading[i] = leading[ch]
	}

	// Step 5: Recursive backtracking with pruning
	assignment := make([]int, n)
	var solve func(idx int, partialSum int, usedMask int) bool

	solve = func(idx int, partialSum int, usedMask int) bool {
		if idx == n {
			return partialSum == 0
		}

		for d := 0; d <= 9; d++ {
			if usedMask&(1<<d) != 0 {
				continue
			}
			if d == 0 && isLeading[idx] {
				continue
			}

			newSum := partialSum + w[idx]*d
			newMask := usedMask | (1 << d)

			// Compute loose bounds for remaining letters
			if idx+1 < n {
				minBound := 0
				maxBound := 0
				for j := idx + 1; j < n; j++ {
					wj := w[j]
					// Find min and max available digits for this letter
					var lo, hi int
					if isLeading[j] {
						// Leading letter: smallest available non-zero digit
						lo = -1
						for dd := 1; dd <= 9; dd++ {
							if newMask&(1<<dd) == 0 {
								if lo == -1 {
									lo = dd
								}
								hi = dd
							}
						}
					} else {
						lo = -1
						for dd := 0; dd <= 9; dd++ {
							if newMask&(1<<dd) == 0 {
								if lo == -1 {
									lo = dd
								}
								hi = dd
							}
						}
					}
					if lo == -1 {
						// No available digit — this path is impossible
						goto nextDigit
					}
					if wj > 0 {
						minBound += wj * lo
						maxBound += wj * hi
					} else {
						minBound += wj * hi
						maxBound += wj * lo
					}
				}
				// Check if zero is achievable
				if newSum+minBound > 0 || newSum+maxBound < 0 {
					goto nextDigit
				}
			}

			assignment[idx] = d
			if solve(idx+1, newSum, newMask) {
				return true
			}

		nextDigit:
		}
		return false
	}

	if !solve(0, 0, 0) {
		return nil, errors.New("no solution found")
	}

	// Step 6: Build result map
	result2 := make(map[string]int, n)
	for i, ch := range letters {
		result2[string(ch)] = assignment[i]
	}
	return result2, nil
}
