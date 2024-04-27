package _274

func hIndex(citations []int) int {

	low, high := 0, 0
	for _, c := range citations {
		high = max(high, c)
	}

	index := 0
	for low <= high {
		mid := (low + high) / 2
		count := 0
		for _, c := range citations {
			if c >= mid {
				count++
			}
		}
		if count >= mid {
			index = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return index
}
