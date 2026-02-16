package dominoes

// Domino represents a domino with two numbered sides.
type Domino [2]int

// MakeChain attempts to arrange the given dominoes into a valid chain where
// adjacent dominoes share matching values and the chain forms a loop.
func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return []Domino{}, true
	}
	if len(input) == 1 {
		if input[0][0] == input[0][1] {
			return []Domino{input[0]}, true
		}
		return nil, false
	}
	if !feasible(input) {
		return nil, false
	}
	chain := make([]Domino, 0, len(input))
	used := make([]bool, len(input))
	start := input[0][0]
	if backtrack(input, used, &chain, start, start, 0) {
		return chain, true
	}
	return nil, false
}

// feasible checks if an Eulerian circuit is possible: all vertices must have
// even degree and the graph must be connected.
func feasible(input []Domino) bool {
	degree := map[int]int{}
	adj := map[int]map[int]bool{}
	for _, d := range input {
		degree[d[0]]++
		degree[d[1]]++
		if adj[d[0]] == nil {
			adj[d[0]] = map[int]bool{}
		}
		if adj[d[1]] == nil {
			adj[d[1]] = map[int]bool{}
		}
		adj[d[0]][d[1]] = true
		adj[d[1]][d[0]] = true
	}
	for _, deg := range degree {
		if deg%2 != 0 {
			return false
		}
	}
	// Check connectivity via BFS
	var startNode int
	for n := range adj {
		startNode = n
		break
	}
	visited := map[int]bool{startNode: true}
	queue := []int{startNode}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return len(visited) == len(adj)
}

// backtrack tries to build a chain using recursive DFS with backtracking.
func backtrack(input []Domino, used []bool, chain *[]Domino, start, current, placed int) bool {
	if placed == len(input) {
		return current == start
	}
	for i, d := range input {
		if used[i] {
			continue
		}
		// Try normal orientation
		if d[0] == current {
			used[i] = true
			*chain = append(*chain, d)
			if backtrack(input, used, chain, start, d[1], placed+1) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}
		// Try flipped orientation (only if different from normal)
		if d[1] == current && d[0] != d[1] {
			used[i] = true
			flipped := Domino{d[1], d[0]}
			*chain = append(*chain, flipped)
			if backtrack(input, used, chain, start, d[0], placed+1) {
				return true
			}
			*chain = (*chain)[:len(*chain)-1]
			used[i] = false
		}
	}
	return false
}
