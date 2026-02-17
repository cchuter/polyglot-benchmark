package bookstore

import "sort"

// Cost calculates the total cost in cents for a basket of books,
// applying the best possible discount grouping.
func Cost(books []int) int {
	// Count frequency of each book (books are numbered 1-5)
	freq := make([]int, 5)
	for _, b := range books {
		freq[b-1]++
	}

	// Sort frequencies descending
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))

	// Greedily form groups (largest possible each time)
	var groups []int
	for freq[0] > 0 {
		size := 0
		for i := 0; i < 5; i++ {
			if freq[i] > 0 {
				freq[i]--
				size++
			}
		}
		groups = append(groups, size)
		// Re-sort after each extraction
		sort.Sort(sort.Reverse(sort.IntSlice(freq)))
	}

	// Calculate initial cost and count groups of 5 and 3
	cost := 0
	fives, threes := 0, 0
	for _, g := range groups {
		cost += groupCost(g)
		switch g {
		case 5:
			fives++
		case 3:
			threes++
		}
	}

	// Optimize: each (5+3) → (4+4) conversion saves 40 cents
	convert := fives
	if threes < convert {
		convert = threes
	}
	cost -= convert * 40

	return cost
}

func groupCost(n int) int {
	switch n {
	case 1:
		return 800
	case 2:
		return 1520
	case 3:
		return 2160
	case 4:
		return 2560
	case 5:
		return 3000
	default:
		return 0
	}
}
