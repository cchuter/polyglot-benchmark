package dominoes

// Domino represents a domino tile with two sides.
type Domino [2]int

// MakeChain attempts to arrange the input dominoes into a valid chain.
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
	// Place first domino
	chain[0] = input[0]
	used[0] = true
	if solve(chain, used, input, 1) {
		return chain, true
	}
	// Try first domino flipped
	chain[0] = Domino{input[0][1], input[0][0]}
	if solve(chain, used, input, 1) {
		return chain, true
	}
	return nil, false
}

// canChain checks Euler circuit preconditions: even degree + connectivity.
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

	// Connectivity check via union-find
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
	root := find(input[0][0])
	for v := range degree {
		if find(v) != root {
			return false
		}
	}
	return true
}

// solve uses backtracking to build the chain from position pos onward.
func solve(chain []Domino, used []bool, input []Domino, pos int) bool {
	if pos == len(input) {
		return chain[0][0] == chain[pos-1][1]
	}
	end := chain[pos-1][1]
	for i, d := range input {
		if used[i] {
			continue
		}
		used[i] = true
		if d[0] == end {
			chain[pos] = d
			if solve(chain, used, input, pos+1) {
				return true
			}
		}
		if d[1] == end && d[0] != d[1] {
			chain[pos] = Domino{d[1], d[0]}
			if solve(chain, used, input, pos+1) {
				return true
			}
		}
		used[i] = false
	}
	return false
}
