package _056

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	result = append(result, intervals[0])

	for _, interval := range intervals[1:] {
		last := result[len(result)-1]
		if interval[0] <= last[1] {
			last[1] = int(math.Max(float64(last[1]), float64(interval[1])))
		} else {
			result = append(result, interval)
		}
	}

	return result
}
