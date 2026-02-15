package bookstore

var groupCosts = [6]int{
	0,
	1 * 800,
	2 * 800 * 95 / 100,
	3 * 800 * 90 / 100,
	4 * 800 * 80 / 100,
	5 * 800 * 75 / 100,
}

func Cost(books []int) int {
	freq := make(map[int]int)
	for _, b := range books {
		freq[b]++
	}

	counts := make([]int, 0, len(freq))
	for _, c := range freq {
		counts = append(counts, c)
	}

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

	fives, threes := 0, 0
	for _, g := range groups {
		if g == 5 {
			fives++
		}
		if g == 3 {
			threes++
		}
	}
	swaps := fives
	if threes < swaps {
		swaps = threes
	}
	if swaps > 0 {
		newGroups := make([]int, 0, len(groups))
		fivesLeft, threesLeft := swaps, swaps
		for _, g := range groups {
			if g == 5 && fivesLeft > 0 {
				newGroups = append(newGroups, 4)
				fivesLeft--
			} else if g == 3 && threesLeft > 0 {
				newGroups = append(newGroups, 4)
				threesLeft--
			} else {
				newGroups = append(newGroups, g)
			}
		}
		groups = newGroups
	}

	total := 0
	for _, g := range groups {
		total += groupCosts[g]
	}
	return total
}
