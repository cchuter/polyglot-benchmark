package bookstore

import "sort"

// Cost calculates the total cost of a basket of books, applying the best possible discount.
func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	// Count frequency of each book
	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}

	// Collect and sort frequencies descending
	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Build group counts using histogram layer-peeling
	n := len(counts)
	groups := make([]int, 6) // groups[i] = number of groups of size i
	for w := n; w >= 1; w-- {
		next := 0
		if w < n {
			next = counts[w]
		}
		groups[w] = counts[w-1] - next
	}

	// Optimize: convert (5,3) pairs to (4,4) pairs
	pairs := groups[5]
	if groups[3] < pairs {
		pairs = groups[3]
	}
	groups[5] -= pairs
	groups[3] -= pairs
	groups[4] += 2 * pairs

	// Calculate total cost
	costTable := [6]int{0, 800, 1520, 2160, 2560, 3000}
	total := 0
	for i := 1; i <= 5; i++ {
		total += groups[i] * costTable[i]
	}
	return total
}
