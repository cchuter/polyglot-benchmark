package bookstore

import "sort"

const bookPrice = 800

var discounts = [6]int{0, 0, 5, 10, 20, 25}

// Cost calculates the minimum cost for a basket of books, applying the best discounts.
func Cost(books []int) int {
	if len(books) == 0 {
		return 0
	}

	// Count frequency of each book
	freq := make([]int, 5)
	for _, b := range books {
		freq[b-1]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))

	// Greedily form groups of distinct books
	groups := [6]int{}
	for freq[0] > 0 {
		distinct := 0
		for i := range freq {
			if freq[i] > 0 {
				distinct++
			}
		}
		for i := 0; i < distinct; i++ {
			freq[i]--
		}
		groups[distinct]++
		sort.Sort(sort.Reverse(sort.IntSlice(freq)))
	}

	// Optimize: two groups of 4 are cheaper than a group of 5 + group of 3
	swaps := groups[5]
	if groups[3] < swaps {
		swaps = groups[3]
	}
	groups[5] -= swaps
	groups[3] -= swaps
	groups[4] += 2 * swaps

	// Calculate total cost
	total := 0
	for size := 1; size <= 5; size++ {
		total += groups[size] * groupCost(size)
	}
	return total
}

func groupCost(size int) int {
	return bookPrice * size * (100 - discounts[size]) / 100
}
