package bookstore

import "sort"

var discount = [6]int{0, 0, 5, 10, 20, 25}

func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}

	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	memo := make(map[[5]int]int)
	return solve(toKey(counts), memo)
}

func toKey(counts []int) [5]int {
	var k [5]int
	for i := 0; i < len(counts) && i < 5; i++ {
		k[i] = counts[i]
	}
	return k
}

func solve(key [5]int, memo map[[5]int]int) int {
	if key == [5]int{} {
		return 0
	}
	if v, ok := memo[key]; ok {
		return v
	}

	// Count how many non-zero entries
	nonZero := 0
	for _, c := range key {
		if c > 0 {
			nonZero++
		}
	}

	best := 1<<63 - 1
	for g := 1; g <= nonZero; g++ {
		// Take one book from each of the top g slots
		var next [5]int
		next = key
		for i := 0; i < g; i++ {
			next[i]--
		}
		// Re-sort descending
		sorted := next[:]
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
		copy(next[:], sorted)

		groupCost := g * 800 * (100 - discount[g]) / 100
		total := groupCost + solve(next, memo)
		if total < best {
			best = total
		}
	}

	memo[key] = best
	return best
}
