package _435

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	cur := intervals[0]
	for _, interval := range intervals[1:] {
		if cur[1] <= interval[0] {
			result = append(result, cur)
			cur = interval
		} else if cur[1] > interval[1] {
			cur = interval
		}
	}
	result = append(result, cur)

	return len(intervals) - len(result)
}
