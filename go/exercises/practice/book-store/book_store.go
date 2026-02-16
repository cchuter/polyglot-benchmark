package bookstore

import "sort"

// groupCost maps group size to its discounted price in cents.
var groupCost = [6]int{0, 800, 1520, 2160, 2560, 3000}

// Cost calculates the cheapest price for a basket of books.
func Cost(books []int) int {
	// Count frequency of each book (numbered 1-5).
	freq := make([]int, 5)
	for _, b := range books {
		freq[b-1]++
	}

	// Sort frequencies descending.
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))

	// Layer-peel: determine number of groups of each size.
	// With sorted desc frequencies, groups of size k are formed by
	// the difference between adjacent frequency levels.
	var groups [6]int
	for size := 1; size <= 5; size++ {
		next := 0
		if size <= 4 {
			next = freq[size]
		}
		groups[size] = freq[size-1] - next
	}

	// Adjustment: two groups of 4 (5120) are cheaper than a 5+3 pair (5160).
	pairs := groups[5]
	if groups[3] < pairs {
		pairs = groups[3]
	}
	groups[5] -= pairs
	groups[3] -= pairs
	groups[4] += 2 * pairs

	// Sum up the total cost.
	total := 0
	for size := 1; size <= 5; size++ {
		total += groups[size] * groupCost[size]
	}
	return total
}
