package alphametics

import (
	"errors"
	"sort"
	"strings"
)

type solver struct {
	letters    []byte
	weights    []int64
	isLeading  []bool
	nLetters   int
	assignment [26]int
	used       [10]bool
}

// Solve parses an alphametics puzzle and returns a mapping of letters to digits.
func Solve(puzzle string) (map[string]int, error) {
	sides := strings.Split(puzzle, "==")
	if len(sides) != 2 {
		return nil, errors.New("invalid puzzle: must contain exactly one ==")
	}

	addendStrs := strings.Split(sides[0], "+")
	if len(addendStrs) == 0 {
		return nil, errors.New("invalid puzzle: no addends")
	}

	var words []string
	for _, a := range addendStrs {
		w := strings.TrimSpace(a)
		if w == "" {
			return nil, errors.New("invalid puzzle: empty addend")
		}
		words = append(words, w)
	}

	result := strings.TrimSpace(sides[1])
	if result == "" {
		return nil, errors.New("invalid puzzle: empty result")
	}

	allWords := append(words, result)

	// Validate characters and collect unique letters
	var letterSet [26]bool
	for _, w := range allWords {
		for i := 0; i < len(w); i++ {
			c := w[i]
			if c < 'A' || c > 'Z' {
				return nil, errors.New("invalid puzzle: non-uppercase letter")
			}
			letterSet[c-'A'] = true
		}
	}

	var uniqueLetters []byte
	for i := 0; i < 26; i++ {
		if letterSet[i] {
			uniqueLetters = append(uniqueLetters, byte(i)+'A')
		}
	}
	if len(uniqueLetters) > 10 {
		return nil, errors.New("invalid puzzle: more than 10 unique letters")
	}

	// Compute weights
	var weightMap [26]int64
	for _, w := range words {
		var pow int64 = 1
		for i := len(w) - 1; i >= 0; i-- {
			weightMap[w[i]-'A'] += pow
			pow *= 10
		}
	}
	{
		var pow int64 = 1
		for i := len(result) - 1; i >= 0; i-- {
			weightMap[result[i]-'A'] -= pow
			pow *= 10
		}
	}

	// Identify leading letters
	var leadingSet [26]bool
	for _, w := range allWords {
		if len(w) > 1 {
			leadingSet[w[0]-'A'] = true
		}
	}

	// Build parallel slices
	n := len(uniqueLetters)
	letters := make([]byte, n)
	weights := make([]int64, n)
	isLeading := make([]bool, n)
	for i, c := range uniqueLetters {
		letters[i] = c
		weights[i] = weightMap[c-'A']
		isLeading[i] = leadingSet[c-'A']
	}

	// Sort by descending |weight|
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(a, b int) bool {
		wa := weights[indices[a]]
		wb := weights[indices[b]]
		if wa < 0 {
			wa = -wa
		}
		if wb < 0 {
			wb = -wb
		}
		return wa > wb
	})

	sortedLetters := make([]byte, n)
	sortedWeights := make([]int64, n)
	sortedLeading := make([]bool, n)
	for i, idx := range indices {
		sortedLetters[i] = letters[idx]
		sortedWeights[i] = weights[idx]
		sortedLeading[i] = isLeading[idx]
	}

	s := &solver{
		letters:   sortedLetters,
		weights:   sortedWeights,
		isLeading: sortedLeading,
		nLetters:  n,
	}
	for i := range s.assignment {
		s.assignment[i] = -1
	}

	if s.solve(0, 0) {
		result := make(map[string]int, n)
		for i := 0; i < n; i++ {
			result[string(s.letters[i])] = s.assignment[s.letters[i]-'A']
		}
		return result, nil
	}

	return nil, errors.New("no solution found")
}

func (s *solver) solve(depth int, currentSum int64) bool {
	if depth == s.nLetters {
		return currentSum == 0
	}

	// Compute admissible bounds for remaining letters (depth+1 onwards)
	var minDelta, maxDelta int64
	for i := depth + 1; i < s.nLetters; i++ {
		w := s.weights[i]
		lo := s.minAvail(s.isLeading[i])
		hi := s.maxAvail()
		var cMin, cMax int64
		if w >= 0 {
			cMin = w * int64(lo)
			cMax = w * int64(hi)
		} else {
			cMin = w * int64(hi)
			cMax = w * int64(lo)
		}
		minDelta += cMin
		maxDelta += cMax
	}

	// Include current letter's contribution bounds for pruning
	w := s.weights[depth]
	for d := 0; d <= 9; d++ {
		if s.used[d] {
			continue
		}
		if d == 0 && s.isLeading[depth] {
			continue
		}

		contrib := w * int64(d)
		newSum := currentSum + contrib

		// Prune: check if 0 is reachable
		if newSum+minDelta > 0 || newSum+maxDelta < 0 {
			continue
		}

		s.used[d] = true
		s.assignment[s.letters[depth]-'A'] = d
		if s.solve(depth+1, newSum) {
			return true
		}
		s.used[d] = false
		s.assignment[s.letters[depth]-'A'] = -1
	}

	return false
}

func (s *solver) minAvail(leading bool) int {
	start := 0
	if leading {
		start = 1
	}
	for d := start; d <= 9; d++ {
		if !s.used[d] {
			return d
		}
	}
	return 9
}

func (s *solver) maxAvail() int {
	for d := 9; d >= 0; d-- {
		if !s.used[d] {
			return d
		}
	}
	return 0
}
