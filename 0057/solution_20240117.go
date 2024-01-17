package _057

import "math"

func insert_(intervals [][]int, newInterval []int) [][]int {

	result, added := make([][]int, 0), false
	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			result, added = append(result, newInterval), true
			result = append(result, intervals[i:]...)
			break
		} else if newInterval[0] > interval[1] {
			result = append(result, interval)
		} else {
			newInterval[0] = int(math.Min(float64(newInterval[0]), float64(interval[0])))
			newInterval[1] = int(math.Max(float64(newInterval[1]), float64(interval[1])))
		}
	}
	if !added {
		result = append(result, newInterval)
	}
	return result

}
