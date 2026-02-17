package dominoes

// Domino represents a domino tile with two sides.
type Domino [2]int

// MakeChain attempts to arrange the input dominoes into a valid chain
// where adjacent dominoes match and the chain forms a loop.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	switch len(input) {
	case 0:
		return []Domino{}, true
	case 1:
		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}

	if !canChain(input) {
		return nil, false
	}

	chain = make([]Domino, len(input))
	used := make([]bool, len(input))

	// Try starting with the first domino in its original orientation.
	used[0] = true
	chain[0] = input[0]
	if search(input, used, chain, 1) {
		return chain, true
	}

	// Try starting with the first domino flipped.
	chain[0] = Domino{input[0][1], input[0][0]}
	if search(input, used, chain, 1) {
		return chain, true
	}

	return nil, false
}

// canChain checks necessary conditions: all vertices have even degree
// and the graph of domino values is connected.
func canChain(input []Domino) bool {
	degree := map[int]int{}
	for _, d := range input {
		degree[d[0]]++
		degree[d[1]]++
	}
	for _, deg := range degree {
		if deg%2 != 0 {
			return false
		}
	}

	// Connectivity check using union-find.
	parent := map[int]int{}
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

	// All values must share the same root.
	var root int
	first := true
	for v := range degree {
		if first {
			root = find(v)
			first = false
		} else if find(v) != root {
			return false
		}
	}
	return true
}

// search recursively tries to place unused dominoes into the chain.
func search(input []Domino, used []bool, chain []Domino, pos int) bool {
	if pos == len(input) {
		return chain[0][0] == chain[pos-1][1]
	}

	need := chain[pos-1][1]
	for i, d := range input {
		if used[i] {
			continue
		}
		used[i] = true
		if d[0] == need {
			chain[pos] = d
			if search(input, used, chain, pos+1) {
				return true
			}
		}
		if d[1] == need && d[0] != d[1] {
			chain[pos] = Domino{d[1], d[0]}
			if search(input, used, chain, pos+1) {
				return true
			}
		}
		used[i] = false
	}
	return false
}
