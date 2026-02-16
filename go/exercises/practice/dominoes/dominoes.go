package dominoes

// Domino represents a domino stone with two sides.
type Domino [2]int

// MakeChain attempts to arrange the input dominoes into a valid chain
// where adjacent dominoes match and the chain is circular.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 {
		return []Domino{}, true
	}

	// Check necessary conditions: even degree and connectivity.
	if !hasEvenDegrees(input) || !isConnected(input) {
		return nil, false
	}

	// Backtrack to find a valid chain.
	used := make([]bool, len(input))
	chain = make([]Domino, 0, len(input))

	// Try starting with the first domino in both orientations.
	for _, d := range []Domino{input[0], {input[0][1], input[0][0]}} {
		used[0] = true
		chain = append(chain, d)
		if result := solve(input, chain, used); result != nil {
			return result, true
		}
		chain = chain[:0]
		used[0] = false
	}

	return nil, false
}

func solve(input []Domino, chain []Domino, used []bool) []Domino {
	if len(chain) == len(input) {
		// Check circularity: first left == last right.
		if chain[0][0] == chain[len(chain)-1][1] {
			return chain
		}
		return nil
	}

	need := chain[len(chain)-1][1]

	for i, d := range input {
		if used[i] {
			continue
		}
		used[i] = true
		if d[0] == need {
			if result := solve(input, append(chain, d), used); result != nil {
				return result
			}
		} else if d[1] == need {
			if result := solve(input, append(chain, Domino{d[1], d[0]}), used); result != nil {
				return result
			}
		}
		used[i] = false
	}
	return nil
}

func hasEvenDegrees(input []Domino) bool {
	degree := make(map[int]int)
	for _, d := range input {
		degree[d[0]]++
		degree[d[1]]++
	}
	for _, deg := range degree {
		if deg%2 != 0 {
			return false
		}
	}
	return true
}

func isConnected(input []Domino) bool {
	parent := make(map[int]int)

	var find func(int) int
	find = func(x int) int {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra, rb := find(a), find(b)
		if ra != rb {
			parent[ra] = rb
		}
	}

	for _, d := range input {
		union(d[0], d[1])
	}

	// All vertices must share the same root.
	var root int
	first := true
	for v := range parent {
		if first {
			root = find(v)
			first = false
		} else if find(v) != root {
			return false
		}
	}
	return true
}
