package _275

func hIndex(citations []int) int {

	low, high := 0, 0
	for _, c := range citations {
		high = max(high, c)
	}

	index := 0
	for low <= high {
		mid := (low + high) >> 1
		if f(citations, mid) {
			index = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return index
}

func f(citations []int, index int) bool {
	low, high := 0, len(citations)-1
	for low < high {
		mid := (low + high) >> 1
		if citations[mid] < index {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return len(citations)-high >= index
}
