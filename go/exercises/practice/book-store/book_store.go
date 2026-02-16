package bookstore

import "sort"

// Cost calculates the total cost in cents for a basket of books,
// applying the best possible group discounts.
func Cost(books []int) int {
	// Count frequency of each book
	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}

	// Extract counts and sort descending
	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Build groups greedily
	var groups []int
	for {
		size := 0
		for i := range counts {
			if counts[i] > 0 {
				counts[i]--
				size++
			}
		}
		if size == 0 {
			break
		}
		groups = append(groups, size)
	}

	// Count groups of 5 and 3; redistribute to 4+4
	fives, threes := 0, 0
	for _, g := range groups {
		if g == 5 {
			fives++
		}
		if g == 3 {
			threes++
		}
	}
	redistribute := min(fives, threes)

	// Calculate cost using group cost table
	groupCost := [6]int{0, 800, 1520, 2160, 2560, 3000}
	total := 0
	for _, g := range groups {
		total += groupCost[g]
	}
	// Apply redistribution: each 5+3 → 4+4 saves 40 cents
	total -= redistribute * 40

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
