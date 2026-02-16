package bookstore

import "sort"

var groupCost = [...]int{0, 800, 1520, 2160, 2560, 3000}

func Cost(books []int) int {
	// Count frequency of each book
	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}

	// Extract frequencies and sort descending
	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Build groups by peeling layers
	var groups []int
	for len(counts) > 0 {
		groups = append(groups, len(counts))
		// Decrease each frequency by 1, remove zeros
		j := 0
		for _, c := range counts {
			if c > 1 {
				counts[j] = c - 1
				j++
			}
		}
		counts = counts[:j]
	}

	// Optimize: convert 5+3 pairs into 4+4 pairs
	n5, n3 := 0, 0
	for _, g := range groups {
		if g == 5 {
			n5++
		} else if g == 3 {
			n3++
		}
	}
	swaps := n5
	if n3 < swaps {
		swaps = n3
	}
	if swaps > 0 {
		optimized := make([]int, 0, len(groups)+swaps)
		s5, s3 := swaps, swaps
		for _, g := range groups {
			if g == 5 && s5 > 0 {
				optimized = append(optimized, 4)
				s5--
			} else if g == 3 && s3 > 0 {
				optimized = append(optimized, 4)
				s3--
			} else {
				optimized = append(optimized, g)
			}
		}
		groups = optimized
	}

	// Sum costs
	total := 0
	for _, g := range groups {
		total += groupCost[g]
	}
	return total
}
