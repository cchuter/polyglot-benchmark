package bookstore

import (
	"math"
	"sort"
)

var discounts = [6]int{0, 0, 5, 10, 20, 25}

func Cost(books []int) int {
	var freq [5]int
	for _, b := range books {
		freq[b-1]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freq[:])))
	memo := make(map[[5]int]int)
	return minCost(freq, memo)
}

func minCost(freq [5]int, memo map[[5]int]int) int {
	if freq[0] == 0 {
		return 0
	}
	if v, ok := memo[freq]; ok {
		return v
	}

	distinct := 0
	for _, f := range freq {
		if f > 0 {
			distinct++
		}
	}

	best := math.MaxInt32
	for gs := 1; gs <= distinct; gs++ {
		var nf [5]int
		copy(nf[:], freq[:])
		for i := 0; i < gs; i++ {
			nf[i]--
		}
		sort.Sort(sort.Reverse(sort.IntSlice(nf[:])))
		candidate := groupCost(gs) + minCost(nf, memo)
		if candidate < best {
			best = candidate
		}
	}

	memo[freq] = best
	return best
}

func groupCost(n int) int {
	return n * 800 * (100 - discounts[n]) / 100
}
