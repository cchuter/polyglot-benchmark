package bookstore

import "sort"

const bookPrice = 800

var discounts = [6]int{0, 0, 5, 10, 20, 25}

func Cost(books []int) int {
	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}
	counts := make([]int, 0, len(freq))
	for _, v := range freq {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return minCost(counts)
}

func minCost(counts []int) int {
	for len(counts) > 0 && counts[len(counts)-1] == 0 {
		counts = counts[:len(counts)-1]
	}
	if len(counts) == 0 {
		return 0
	}

	best := int(^uint(0) >> 1)
	for groupSize := 1; groupSize <= len(counts); groupSize++ {
		next := make([]int, len(counts))
		copy(next, counts)
		for i := 0; i < groupSize; i++ {
			next[i]--
		}
		sort.Sort(sort.Reverse(sort.IntSlice(next)))
		cost := groupCost(groupSize) + minCost(next)
		if cost < best {
			best = cost
		}
	}
	return best
}

func groupCost(size int) int {
	return bookPrice * size * (100 - discounts[size]) / 100
}
