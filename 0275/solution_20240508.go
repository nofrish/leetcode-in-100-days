package _275

import "sort"

func hIndex(citations []int) int {

	sort.Ints(citations)

	lo, hi := 0, len(citations)
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if f(citations, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func f(citations []int, x int) bool {

	lo, hi := 0, len(citations)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if citations[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if hi != len(citations)-1 || citations[hi] >= x {
		return len(citations)-hi >= x
	}
	return false
}
