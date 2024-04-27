package _875

func minEatingSpeed(piles []int, h int) int {

	low, high := 1, 1
	for _, pile := range piles {
		high = max(high, pile)
	}

	minK := high
	for low <= high {
		mid := (low + high) / 2
		if f(piles, mid) > h {
			low = mid + 1
		} else {
			high = mid - 1
			minK = mid
		}
	}
	return minK
}

func f(piles []int, k int) int {
	cost := 0
	for _, pile := range piles {
		if pile%k == 0 {
			cost += pile / k
		} else {
			cost += pile/k + 1
		}
	}
	return cost
}
