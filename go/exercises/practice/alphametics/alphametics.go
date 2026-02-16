package alphametics

import (
	"errors"
	"strings"
	"unicode"
)

var decDigits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

type problem struct {
	vDigits      [][]rune
	maxDigits    int
	letterValues [26]int
	lettersUsed  []rune
	nLetters     int
	leadingRunes []rune
}

func Solve(puzzle string) (map[string]int, error) {
	p := parsePuzzle(puzzle)
	if p == nil {
		return nil, errors.New("invalid puzzle")
	}
	return p.solvePuzzle()
}

func parsePuzzle(puzzle string) (p *problem) {
	var valueStrings []string
	p = new(problem)
	fields := strings.Fields(puzzle)
	for _, field := range fields {
		if field == "+" || field == "==" {
			continue
		}
		valueStrings = append(valueStrings, field)
		if len(field) > p.maxDigits {
			p.maxDigits = len(field)
		}
		for _, r := range field {
			if !unicode.IsUpper(r) {
				return nil
			}
			p.letterValues[r-'A'] = -1
		}
	}
	// Collect leading runes of multi-digit words.
	seen := [26]bool{}
	for _, s := range valueStrings {
		if len(s) > 1 {
			r := rune(s[0])
			idx := r - 'A'
			if !seen[idx] {
				seen[idx] = true
				p.leadingRunes = append(p.leadingRunes, idx)
			}
		}
	}
	// Count the letters used.
	for v := 0; v < len(p.letterValues); v++ {
		if p.letterValues[v] == -1 {
			p.nLetters++
		}
	}
	// Build column structure: vDigits[word][column] where column 0 = ones place.
	p.vDigits = make([][]rune, len(valueStrings))
	for i := range valueStrings {
		p.vDigits[i] = make([]rune, p.maxDigits)
		for d, r := range valueStrings[i] {
			j := len(valueStrings[i]) - 1 - d
			p.vDigits[i][j] = r - 'A' + 1
		}
	}
	// Build list of letters used (0 == 'A').
	p.lettersUsed = make([]rune, p.nLetters)
	for n, v := 0, 0; v < len(p.letterValues); v++ {
		if p.letterValues[v] == -1 {
			p.lettersUsed[n] = rune(v)
			n++
		}
	}
	return p
}

func (p *problem) solvePuzzle() (map[string]int, error) {
	for _, digValues := range permutations(decDigits, p.nLetters) {
		// Assign values to letters.
		for i, r := range p.lettersUsed {
			p.letterValues[r] = digValues[i]
		}
		// Check leading zeros for all multi-digit words.
		leadingZero := false
		for _, idx := range p.leadingRunes {
			if p.letterValues[idx] == 0 {
				leadingZero = true
				break
			}
		}
		if leadingZero {
			continue
		}
		if p.isPuzzleSolution() {
			return p.puzzleMap(), nil
		}
	}
	return nil, errors.New("no solution")
}

func (p *problem) isPuzzleSolution() bool {
	carry := 0
	for d := 0; d < p.maxDigits; d++ {
		sum := carry
		// Sum all addend rows (all except the last) for this column.
		for n := 0; n < len(p.vDigits)-1; n++ {
			r := p.vDigits[n][d]
			if r != 0 {
				sum += p.letterValues[r-1]
			}
		}
		carry = sum / 10
		sum %= 10
		// Check against the result row.
		r := p.vDigits[len(p.vDigits)-1][d]
		if r == 0 || sum != p.letterValues[r-1] {
			return false
		}
	}
	return carry == 0
}

func (p *problem) puzzleMap() map[string]int {
	pm := make(map[string]int, p.nLetters)
	for _, v := range p.lettersUsed {
		r := v + 'A'
		pm[string(r)] = p.letterValues[v]
	}
	return pm
}

func permutations(iterable []int, r int) (perms [][]int) {
	pool := iterable
	n := len(pool)
	if r > n {
		return
	}

	nperm := 1
	for i := n; i > 1; i-- {
		nperm *= i
	}
	if r < n {
		d := 1
		for i := n - r; i > 1; i-- {
			d *= i
		}
		nperm /= d
	}
	perms = make([][]int, 0, nperm)

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}
	p2 := make([]int, len(result))
	copy(p2, result)
	perms = append(perms, p2)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]
				for k := i; k < r; k++ {
					result[k] = pool[indices[k]]
				}
				p2 = make([]int, len(result))
				copy(p2, result)
				perms = append(perms, p2)
				break
			}
		}
		if i < 0 {
			return
		}
	}
	return
}
