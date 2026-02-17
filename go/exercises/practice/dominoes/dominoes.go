package dominoes

// Domino represents a domino tile with two face values.
type Domino [2]int

// MakeChain attempts to arrange the input dominoes into a valid chain.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	n := len(input)
	if n == 0 {
		return []Domino{}, true
	}
	if n == 1 {
		if input[0][0] == input[0][1] {
			return []Domino{input[0]}, true
		}
		return nil, false
	}

	// Quick validation: even degree + connectivity is necessary and sufficient
	// for an Eulerian circuit to exist.
	if !canChain(input) {
		return nil, false
	}

	// Backtracking search
	used := make([]bool, n)
	chain = make([]Domino, 0, n)

	// Fix first domino (try both orientations)
	used[0] = true
	chain = append(chain, input[0])
	if solve(input, used, &chain) {
		return chain, true
	}
	chain = chain[:1]
	chain[0] = Domino{input[0][1], input[0][0]}
	if solve(input, used, &chain) {
		return chain, true
	}

	return nil, false
}

// canChain checks necessary and sufficient conditions for an Eulerian circuit:
// all vertices have even degree and the graph is connected.
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
	return connected(input)
}

// connected checks if all domino vertices form a single connected component
// using union-find.
func connected(input []Domino) bool {
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
	// All vertices should have the same root
	root := 0
	rootSet := false
	for v := range parent {
		r := find(v)
		if !rootSet {
			root = r
			rootSet = true
		} else if r != root {
			return false
		}
	}
	return true
}

// solve recursively tries to extend the chain using backtracking.
func solve(input []Domino, used []bool, chain *[]Domino) bool {
	if len(*chain) == len(input) {
		// Check that chain forms a loop
		return (*chain)[0][0] == (*chain)[len(*chain)-1][1]
	}
	last := (*chain)[len(*chain)-1][1]
	for i := range input {
		if used[i] {
			continue
		}
		// Try original orientation
		if input[i][0] == last {
			used[i] = true
			*chain = append(*chain, input[i])
			if solve(input, used, chain) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}
		// Try flipped orientation (skip if double, already tried above)
		if input[i][1] == last && input[i][0] != input[i][1] {
			used[i] = true
			*chain = append(*chain, Domino{input[i][1], input[i][0]})
			if solve(input, used, chain) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}
	}
	return false
}
