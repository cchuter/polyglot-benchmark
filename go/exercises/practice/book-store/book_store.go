package bookstore

import "sort"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Cost calculates the total price (in cents) for a basket of books,
// applying the best possible discount by grouping different books.
func Cost(books []int) int {
	// Count frequency of each book (books are numbered 1-5)
	freq := make([]int, 5)
	for _, b := range books {
		freq[b-1]++
	}

	// Greedily form groups of distinct books
	groupCount := [6]int{} // groupCount[size] = number of groups of that size
	for {
		sort.Sort(sort.Reverse(sort.IntSlice(freq)))
		if freq[0] == 0 {
			break
		}
		size := 0
		for i := 0; i < 5; i++ {
			if freq[i] > 0 {
				freq[i]--
				size++
			}
		}
		groupCount[size]++
	}

	// Optimize: convert pairs of (5-group + 3-group) into (4-group + 4-group)
	// because 5*600 + 3*720 = 5160, while 4*640 + 4*640 = 5120 (cheaper)
	swaps := minInt(groupCount[5], groupCount[3])
	groupCount[5] -= swaps
	groupCount[3] -= swaps
	groupCount[4] += 2 * swaps

	// Price per book (in cents) for each group size
	price := [6]int{0, 800, 760, 720, 640, 600}

	total := 0
	for size := 1; size <= 5; size++ {
		total += groupCount[size] * size * price[size]
	}
	return total
}
