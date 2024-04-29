package _594

import "math"

func repairCars(ranks []int, cars int) int64 {
	f := func(x int) int {
		sum := 0
		for _, r := range ranks {
			sum += int(math.Sqrt(float64(x / r)))
		}
		return sum
	}

	low, high := 1, math.MaxInt
	for low < high {
		mid := low + (high-low)/2
		if f(mid) < cars {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return int64(high)
}
