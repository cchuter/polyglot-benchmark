package dominoes

// Domino represents a domino tile with two sides.
type Domino [2]int

// MakeChain attempts to arrange dominoes into a valid chain.
func MakeChain(input []Domino) (chain []Domino, ok bool) {
	if len(input) == 0 {
		return []Domino{}, true
	}

	// Check necessary conditions for Euler circuit:
	// 1. All vertices must have even degree
	// 2. The graph must be connected (considering only vertices that appear)

	degree := map[int]int{}
	for _, d := range input {
		degree[d[0]]++
		degree[d[1]]++
	}
	for _, deg := range degree {
		if deg%2 != 0 {
			return nil, false
		}
	}

	// Check connectivity using Union-Find
	parent := map[int]int{}
	var find func(int) int
	find = func(x int) int {
		if _, exists := parent[x]; !exists {
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

	// All vertices must be in the same component
	var root int
	first := true
	for v := range degree {
		if first {
			root = find(v)
			first = false
		} else if find(v) != root {
			return nil, false
		}
	}

	// Build chain via backtracking
	used := make([]bool, len(input))
	chain = make([]Domino, 0, len(input))

	var solve func() bool
	solve = func() bool {
		if len(chain) == len(input) {
			return chain[0][0] == chain[len(chain)-1][1]
		}
		for i, d := range input {
			if used[i] {
				continue
			}
			used[i] = true
			// Try normal orientation
			if len(chain) == 0 || chain[len(chain)-1][1] == d[0] {
				chain = append(chain, d)
				if solve() {
					return true
				}
				chain = chain[:len(chain)-1]
			}
			// Try flipped orientation (skip for self-loops)
			if d[0] != d[1] {
				flipped := Domino{d[1], d[0]}
				if len(chain) == 0 || chain[len(chain)-1][1] == flipped[0] {
					chain = append(chain, flipped)
					if solve() {
						return true
					}
					chain = chain[:len(chain)-1]
				}
			}
			used[i] = false
		}
		return false
	}

	if solve() {
		return chain, true
	}
	return nil, false
}
